package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()

	engine.POST("/sonarQube/webhook", func(context *gin.Context) {

		data, err := context.GetRawData()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "error"})
		}
		var obj = make(map[string]interface{})
		err = json.Unmarshal(data, &obj)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "error"})
		}
		marshal, err := json.Marshal(obj)
		fmt.Println(string(marshal))
		context.JSON(http.StatusOK, obj)
	})

	engine.GET("/get", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"name": "zzh"})
	})

	err := engine.Run(":8080")
	if err != nil {
		return
	}
}
