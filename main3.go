package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件 ,  统一处理
func totalFun() gin.HandlerFunc {
	return func(context *gin.Context) {

		now := time.Now()
		context.Next()  // 调用后续的handler
		context.Abort() // 组织调用后续的handler

		since := time.Since(now)
		fmt.Println("接口花费时间", since)
		//context.Get()
	}
}

func main() {

	engine := gin.Default()

	engine.Use(totalFun()) // 全局处理中间件

	engine.Group("user")
	{
		engine.GET("/name", func(context *gin.Context) {
			context.JSON(200, gin.H{"name": "zzh"})
		})
		engine.GET("/age", func(context *gin.Context) {
			context.JSON(200, gin.H{"age": 18})
		})
	}

	engine.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("f1")

		dst := "./files" + file.Filename
		context.SaveUploadedFile(file, dst)

	})

	engine.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})
	engine.GET("/forward", func(context *gin.Context) {
		context.Request.URL.Path = "/redirect"
		engine.HandleContext(context)
	})
	engine.NoRoute(func(context *gin.Context) {
		context.Request.URL.Path = "/404"
		engine.HandleContext(context)
	})

	// 统一处理，中间件
	engine.GET("/totalFun", totalFun(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	err := engine.Run(":8080")
	if err != nil {
		fmt.Println("application fail run,", err)
		return
	}
}
