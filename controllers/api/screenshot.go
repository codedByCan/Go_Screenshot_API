package api

import (
	"encoding/base64"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

var (
	browser *rod.Browser
	once    sync.Once
)

func initBrowser() {
	once.Do(func() {
		u := launcher.New().
			Headless(true).
			Set("disable-gpu").
			Set("no-sandbox").
			Set("disable-dev-shm-usage").
			Set("disable-setuid-sandbox").
			Set("disable-web-security").
			Set("disable-background-networking").
			Set("disable-default-apps").
			Set("disable-extensions").
			Set("disable-sync").
			Set("disable-translate").
			Set("hide-scrollbars").
			Set("metrics-recording-only").
			Set("mute-audio").
			Set("no-first-run").
			MustLaunch()

		browser = rod.New().
			ControlURL(u).
			MustConnect()
	})
}

type ScreenshotRequest struct {
	Domain string `json:"domain" binding:"required,url"`
}

type ScreenshotResponse struct {
	Success bool   `json:"success"`
	Image   string `json:"image,omitempty"`
	Error   string `json:"error,omitempty"`
}

func HandleScreenshot(c *gin.Context) {
	var req ScreenshotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ScreenshotResponse{Success: false, Error: err.Error()})
		return
	}

	initBrowser()

	domain := req.Domain
	if !strings.HasPrefix(domain, "http") {
		domain = "https://" + domain
	}

	page := browser.MustPage(domain)
	defer page.MustClose()

	page.MustSetWindow(0, 0, 1280, 720)
	page.MustWaitLoad()

	screenshot, err := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  1280,
			Height: 720,
			Scale:  1,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, ScreenshotResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, ScreenshotResponse{
		Success: true,
		Image:   "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(screenshot),
	})
}