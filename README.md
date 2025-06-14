# Магазин сувениров и духов на разлив

Одностраничный сайт для магазина сувениров и духов на разлив в Первоуральске.

## Технологии

- Go (Gin framework)
- HTML5
- CSS3
- Bootstrap 5
- JavaScript

## Локальная разработка

1. Установите Go (версия 1.16 или выше)
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/your-username/perfume-shop.git
   cd perfume-shop
   ```
3. Установите зависимости:
   ```bash
   go mod download
   ```
4. Запустите сервер:
   ```bash
   go run main.go
   ```
5. Откройте http://localhost:8080 в браузере

## Структура проекта

```
perfume-shop/
├── main.go           # Основной файл приложения
├── go.mod           # Файл зависимостей Go
├── go.sum           # Файл контрольных сумм зависимостей
└── web/             # Веб-ресурсы
    ├── static/      # Статические файлы (изображения, CSS, JS)
    └── templates/   # HTML шаблоны
        └── index.html  # Основной шаблон сайта
```

## Лицензия

MIT 