package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

//
//import "github.com/gin-gonic/gin"
//
//func main() {
//	// 禁用控制台颜色
//	// gin.DisableConsoleColor()
//
//	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
//	router := gin.Default()
//
//	router.GET("/someGet", getting)
//	router.POST("/somePost", posting)
//	router.PUT("/somePut", putting)
//	router.DELETE("/someDelete", deleting)
//	router.PATCH("/somePatch", patching)
//	router.HEAD("/someHead", head)
//	router.OPTIONS("/someOptions", options)
//
//	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
//	router.Run()
//	// router.Run(":3000") hardcode 端口号
//}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
