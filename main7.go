package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var html7 = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/static/js/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.SetHTMLTemplate(html7)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/static/js/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务
	r.Run(":8080")
}
