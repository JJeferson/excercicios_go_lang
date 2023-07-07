package main

import (
	"api1/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando a aplicação ...")

	r := gin.Default()
	r.GET("/ping", controllers.PingGet)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
