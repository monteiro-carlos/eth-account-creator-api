package routes

import (
	"eth-account-creator-api/core/domains/account"
	"eth-account-creator-api/internal/container"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	docs "eth-account-creator-api/internal/swagger/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const corsMaxAge = 300

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Handler(dep *container.Dependency) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Origin", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           corsMaxAge,
	}))

	accountHandler := &account.Handler{
		Service: dep.Services.Account,
	}
	g := router.Group("/account")
	{
		g.GET("/create", accountHandler.CreateNewAccount)
		g.GET("/:publicKey", accountHandler.GetAccountFromPubKey)
	}

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/metrics", prometheusHandler())

	err := router.Run(":5000")
	if err != nil {
		return
	}
}
