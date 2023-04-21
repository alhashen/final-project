package route

import (
	"finalproj/handler"
	"finalproj/middleware"
	"finalproj/service"

	_ "finalproj/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			MyGram API
//	@version		1.0
//	@description	Tugas Akhir pelatihan DTS Kominfo Hacktiv8
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:2400
//	@BasePath	/

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				provide "Bearer" yourself
//	@scope.write				Grants write access
//	@scope.admin				Grants read and write access to administrative information

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func RegisterAPI(e *gin.Engine, svc service.ServiceInterface) {
	hdlr := handler.NewHttpServer(svc)

	userRouter := e.Group("/users")
	{
		userRouter.POST("/register", hdlr.RegisterUser)
		userRouter.POST("/login", hdlr.LogInUser)
	}

	photoRouter := e.Group("/photos")
	{
		photoRouter.Use(middleware.Authenticate())
		photoRouter.POST("/", hdlr.CreatePhoto)
		photoRouter.GET("/", hdlr.GetAllPhoto)
		photoRouter.GET("/:id", hdlr.GetPhoto)
		photoRouter.PUT("/:id", hdlr.UpdatePhoto)
		photoRouter.DELETE("/:id", hdlr.DeletePhoto)
	}

	commentRouter := e.Group("/comments")
	{
		commentRouter.Use(middleware.Authenticate())
		commentRouter.POST("/photo/:id", hdlr.CreateComment)
		commentRouter.GET("/photo/:id", hdlr.GetAllComment)
		commentRouter.GET("/:id", hdlr.GetComment)
		commentRouter.PUT("/:id", hdlr.UpdateComment)
		commentRouter.DELETE("/:id", hdlr.DeleteComment)
	}

	socialMediaRouter := e.Group("/social_medias")
	{
		socialMediaRouter.Use(middleware.Authenticate())
		socialMediaRouter.POST("/", hdlr.CreateSocialMedia)
		socialMediaRouter.GET("/", hdlr.GetAllSocialMedia)
		socialMediaRouter.GET("/:id", hdlr.GetSocialMedia)
		socialMediaRouter.PUT("/:id", hdlr.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", hdlr.DeleteSocialMedia)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))

}
