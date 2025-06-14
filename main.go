package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Загружаем HTML шаблоны
	r.LoadHTMLGlob("web/templates/*")

	// Настраиваем статические файлы
	r.Static("/static", "./web/static")

	// Маршруты
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Запускаем сервер
	r.Run(":8080")
}
