package routes

import (
	"github.com/gin-gonic/gin"
	"week4/intenal/routes/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	user := api.NewUser()

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{

		apiv1.GET("/getNameById", user.GetNameById)

	}

	return r
}
