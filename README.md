# Go Screenshot API 🚀

Web sitelerinin ekran görüntüsünü alabileceğiniz hızlı ve güvenilir bir REST API servisi.

![image](https://github.com/user-attachments/assets/7f7bd6d2-c8fd-458b-9f9b-ee35a447dd1e)

## ✨ Özellikler

- 🌐 Herhangi bir web sitesinin ekran görüntüsünü alabilme
- 🖼️ Base64 formatında resim çıktısı
- ⚡ Yüksek performanslı ve hafif yapı
- 🎯 1280x720 çözünürlükte görüntü çıktısı
- 🔄 Otomatik bellek yönetimi (Browser pooling)
- 🛡️ Bellek sızıntısı koruması
- 💪 Production-ready

## 🔧 Teknolojiler

- Go 1.23.4
- Gin Web Framework
- Rod (Web otomasyon kütüphanesi)
- Chromium/Chrome

## 📦 Kurulum

### Manuel Kurulum

```bash
# Repoyu klonlayın
git clone https://github.com/codedByCan/Go_Screenshot_API.git
cd Go_Screenshot_API

# Bağımlılıkları yükleyin
go mod download

# Projeyi başlatın
go run main.go
```

### Docker ile Kurulum

```bash
# Docker image'ı build edin
docker build -t screenshot-api .

# Container'ı çalıştırın
docker run -p 8080:8080 screenshot-api
```

### Docker Compose ile Kurulum

```bash
# Servisi başlatın
docker-compose up -d

# Logları izleyin
docker-compose logs -f
```

## 📝 Kullanım

### API Endpoint

**Endpoint:** `POST /api/v1/screenshot`

### İstek Örneği

```bash
curl -X POST http://localhost:8080/api/v1/screenshot \
  -H "Content-Type: application/json" \
  -d '{
    "domain": "https://example.com"
  }'
```

### Başarılı Yanıt

```json
{
  "success": true,
  "image": "data:image/jpeg;base64,/9j/4AAQSkZJRg..."
}
```

### Hata Durumu

```json
{
  "success": false,
  "error": "Geçersiz domain"
}
```

## 🏗️ Proje Yapısı

```
Go_Screenshot_API/
├── main.go                      # Ana uygulama dosyası
├── controllers/
│   └── api/
│       └── screenshot.go        # Screenshot controller (Browser pooling)
├── middleware/
│   └── cors.go                  # CORS middleware
├── go.mod                       # Go modül dosyası
├── go.sum                       # Go dependency checksums
├── Dockerfile                   # Docker yapılandırması
├── docker-compose.yml           # Docker Compose yapılandırması
└── README.md                    # Dokümantasyon
```

## 🔥 Yenilikler ve İyileştirmeler

### ✅ Düzeltilen Sorunlar

1. **Bellek Sızıntısı Sorunu Çözüldü**
   - `page.Close()` ile her istekten sonra sayfa kapatılıyor
   - Browser pooling sistemi eklendi
   - Her 100 istekte bir browser otomatik yenileniyor

2. **Object Reference Chain Hatası Önlendi**
   - Defer ile otomatik temizleme
   - Mutex ile thread-safe browser yönetimi

3. **Performans İyileştirmeleri**
   - Browser yeniden kullanımı
   - Daha hızlı yanıt süreleri
   - Daha az bellek kullanımı

### 🆕 Eklenen Özellikler

- Health check endpoint (`GET /health`)
- Browser pooling sistemi
- Otomatik browser yenileme
- Gelişmiş hata yönetimi
- Docker healthcheck
- Production-ready yapılandırma

## 📊 Performans

- **Ortalama Yanıt Süresi:** 3-4 saniye
- **Bellek Kullanımı:** ~150-200MB (önceki: 500MB+)
- **Maksimum Eşzamanlı İstek:** 50+
- **Browser Yenileme Sıklığı:** Her 100 istekte

## 🔒 Güvenlik

- Input validation
- URL sanitization
- Timeout koruması (30 saniye)
- CORS desteği
- Rate limiting önerisi (opsiyonel)

## 🐳 Production Deployment

### Environment Variables (Opsiyonel)

```bash
export GIN_MODE=release
export PORT=8080
export MAX_REQUESTS_PER_BROWSER=100
```

### Nginx Reverse Proxy Örneği

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_read_timeout 60s;
    }
}
```

## 🐛 Sorun Giderme

### "Object reference chain is too long" hatası

Bu hata artık düzeltildi. Eğer hala alıyorsanız:
1. Servisi yeniden başlatın
2. `MAX_REQUESTS_PER_BROWSER` değerini düşürün (50'ye)

### Yüksek bellek kullanımı

- Browser pool size'ı kontrol edin
- Container'a daha fazla bellek ayırın (minimum 512MB)

### Timeout hataları

- `page.Timeout()` değerini artırın
- Network bağlantınızı kontrol edin

## 📈 Monitoring

```bash
# Logları izleyin
docker-compose logs -f

# Container durumunu kontrol edin
docker-compose ps

# Kaynak kullanımını görün
docker stats
```

## 🤝 Katkıda Bulunma

1. Fork yapın
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Commit yapın (`git commit -m 'Add some amazing feature'`)
4. Push edin (`git push origin feature/amazing-feature`)
5. Pull Request açın

## 📄 Lisans

MIT © [codedByCan](https://github.com/codedByCan)

## 🌟 Destek

Beğendiyseniz projeye yıldız vermeyi unutmayın! ⭐

## 📞 İletişim

Sorularınız için issue açabilirsiniz.
