# 📸 Go Screenshot API

Web sitelerinin ekran görüntüsünü alabileceğiniz hızlı bir REST API servisi.

## ✨ Özellikler

- 🌐 Herhangi bir web sitesinin ekran görüntüsünü alabilme
- 🖼️ Base64 formatında resim çıktısı
- ⚡ Yüksek performanslı ve hafif yapı
- 🎯 1280x720 çözünürlükte görüntü çıktısı

## 🛠️ Teknolojiler

- Go 1.23.4
- Gin Web Framework
- Rod (Web otomasyon kütüphanesi)

## 📦 Kurulum

1. Repoyu klonlayın
```bash
git clone https://github.com/fastuptime/Go_Screenshot_API.git
cd Go_Screenshot_API
```

2. Bağımlılıkları yükleyin
```bash
go mod download
```

3. Projeyi başlatın
```bash
go run main.go
```

## 🔥 API Kullanımı

### Ekran Görüntüsü Alma

**Endpoint:** `POST /api/v1/screenshot`

**İstek:**
```json
{
    "domain": "https://example.com"
}
```

**Başarılı Yanıt:**
```json
{
    "success": true,
    "image": "data:image/jpeg;base64,..."
}
```

**Hata Durumu:**
```json
{
    "success": false,
    "error": "Geçersiz domain"
}
```

## 📝 Lisans

MIT © [Fast Uptime](https://github.com/fastuptime)

## ⭐ Projeyi Destekle

Beğendiyseniz projeye yıldız vermeyi unutmayın!
