package middleware

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vmware.com/migrator-poc/address-book-api/pkg/api"
	"vmware.com/migrator-poc/address-book-api/pkg/db"
	"vmware.com/migrator-poc/address-book-api/pkg/docs"
	"vmware.com/migrator-poc/address-book-api/pkg/healthz"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/address-book

func SetupRouter() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Address Book API"
	docs.SwaggerInfo.Description = "This is a Address book server."
	docs.SwaggerInfo.Version = "1.0"
	serviceHost := os.Getenv("SERVICE_HOST")

	if len(serviceHost) == 0 {
		serviceHost = "localhost:8080"
	}

	docs.SwaggerInfo.Host = serviceHost
	docs.SwaggerInfo.BasePath = "/api/address-book"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	docs.SwaggerInfo.ReadDoc()
	docs.SwaggerInfo.InstanceName()

	router := gin.Default()

	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	router.Use(func(c *gin.Context) {
		dbClient := db.DBClient{Db: dbConn}
		c.Set("dbConn", dbClient)
	})

	router.Use(cors.Default())
	v1 := router.Group("/api/address-book")
	{
		contacts := v1.Group("/contacts")
		{
			contacts.GET("", api.GetContacts)
			contacts.GET(":id", api.FindContact)
			contacts.POST("/create", api.CreateContact)
			contacts.PATCH(":id", api.UpdateContact)
			contacts.DELETE(":id", api.DeleteContact)
		}
	}
	// router.GET("/api/address-book/contacts", api.GetContacts)
	// router.GET("/api/address-book/contacts/:id", api.FindContact)
	// router.POST("/api/address-book/contact", api.CreateContact)
	// router.PATCH("/api/address-book/contacts/:id", api.UpdateContact)
	// router.DELETE("/api/address-book/contacts/:id", api.DeleteContact)

	router.GET("/api/address-book/healthz", healthz.GetStatus)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("0.0.0.0:8080")
}
