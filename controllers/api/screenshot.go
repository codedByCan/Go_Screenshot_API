package api

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type BrowserPool struct {
	browser      *rod.Browser
	launcher     *launcher.Launcher
	requestCount int
	mu           sync.Mutex
	maxRequests  int
}

var (
	browserPool *BrowserPool
	once        sync.Once
)

func initBrowserPool() {
	once.Do(func() {
		browserPool = &BrowserPool{
			maxRequests: 100, 
		}
		browserPool.createBrowser()
	})
}

func (bp *BrowserPool) createBrowser() error {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	if bp.browser != nil {
		bp.browser.MustClose()
	}
	if bp.launcher != nil {
		bp.launcher.Cleanup()
	}

	bp.launcher = launcher.New().
		Headless(true).
		NoSandbox(true).
		Set("disable-gpu").
		Set("disable-dev-shm-usage").
		Set("disable-setuid-sandbox").
		Set("no-first-run").
		Set("no-default-browser-check")

	controlURL, err := bp.launcher.Launch()
	if err != nil {
		return err
	}

	bp.browser = rod.New().ControlURL(controlURL).MustConnect()
	bp.requestCount = 0

	return nil
}

func (bp *BrowserPool) getBrowser() (*rod.Browser, error) {
	bp.mu.Lock()
	defer bp.mu.Unlock()

	bp.requestCount++

	if bp.requestCount >= bp.maxRequests {
		if err := bp.createBrowser(); err != nil {
			return nil, err
		}
	}

	return bp.browser, nil
}

type ScreenshotRequest struct {
	Domain string `json:"domain" binding:"required"`
}

type ScreenshotResponse struct {
	Success bool   `json:"success"`
	Image   string `json:"image,omitempty"`
	Error   string `json:"error,omitempty"`
}

func HandleScreenshot(c *gin.Context) {

	initBrowserPool()

	var req ScreenshotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ScreenshotResponse{
			Success: false,
			Error:   "Geçersiz istek formatı",
		})
		return
	}

	domain := strings.TrimSpace(req.Domain)
	if domain == "" {
		c.JSON(http.StatusBadRequest, ScreenshotResponse{
			Success: false,
			Error:   "Domain boş olamaz",
		})
		return
	}

	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		domain = "https://" + domain
	}

	parsedURL, err := url.Parse(domain)
	if err != nil || parsedURL.Host == "" {
		c.JSON(http.StatusBadRequest, ScreenshotResponse{
			Success: false,
			Error:   "Geçersiz domain",
		})
		return
	}

	browser, err := browserPool.getBrowser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Browser başlatılamadı",
		})
		return
	}

	page, err := browser.Page(proto.TargetCreateTarget{URL: ""})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Sayfa oluşturulamadı",
		})
		return
	}

	defer func() {
		if page != nil {
			page.Close()
		}
	}()

	err = page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{
		Width:  1280,
		Height: 720,
		DeviceScaleFactor: 1,
		Mobile: false,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Viewport ayarlanamadı",
		})
		return
	}

	page = page.Timeout(30 * time.Second)

	err = page.Navigate(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Sayfa yüklenemedi: " + err.Error(),
		})
		return
	}

	err = page.WaitLoad()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Sayfa yüklenemedi (timeout)",
		})
		return
	}

	time.Sleep(2 * time.Second)

	img, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: 90,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{
			Success: false,
			Error:   "Screenshot alınamadı",
		})
		return
	}

	base64Image := base64.StdEncoding.EncodeToString(img)
	imageData := "data:image/jpeg;base64," + base64Image

	c.JSON(http.StatusOK, ScreenshotResponse{
		Success: true,
		Image:   imageData,
	})
}
