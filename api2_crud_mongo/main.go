package main

import (
	"fmt"

	"api2_crud_mongo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Iniciando a aplicação ...")

	r := gin.Default()

	r.POST("/novo", services.Save)
	r.Run(":3031")
}
