package services

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Save(r *gin.Context) {

	fmt.Println("Service acessado com sucesso ...")
	r.JSON(http.StatusOK, map[string]string{
		"Status": "Gravado ok",
	})
}
