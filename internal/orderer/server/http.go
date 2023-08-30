package server

import (
	"github.com/gin-gonic/gin"
	"nunu-fabric/internal/organization/handler"
	middleware2 "nunu-fabric/internal/organization/pkg/middleware"
	"nunu-fabric/pkg/helper/resp"
	"nunu-fabric/pkg/jwt"
	"nunu-fabric/pkg/log"
)

func NewServerHTTP(
	logger *log.Logger,
	jwt *jwt.JWT,
	userHandler handler.UserHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(
		middleware2.CORSMiddleware(),
		middleware2.ResponseLogMiddleware(logger),
		middleware2.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// No route group has permission
	noAuthRouter := r.Group("/")
	{

		noAuthRouter.GET("/", func(ctx *gin.Context) {
			logger.WithContext(ctx).Info("hello")
			resp.HandleSuccess(ctx, map[string]interface{}{
				"say": "Hi Nunu!",
			})
		})

		noAuthRouter.POST("/register", userHandler.Register)
		noAuthRouter.POST("/login", userHandler.Login)
	}
	// Non-strict permission routing group
	noStrictAuthRouter := r.Group("/").Use(middleware2.NoStrictAuth(jwt, logger))
	{
		noStrictAuthRouter.GET("/user", userHandler.GetProfile)
	}

	// Strict permission routing group
	strictAuthRouter := r.Group("/").Use(middleware2.StrictAuth(jwt, logger))
	{
		strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
	}

	return r
}
