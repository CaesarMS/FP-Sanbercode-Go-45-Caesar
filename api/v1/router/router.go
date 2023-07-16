package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/v1/middleware"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// admin
	r.POST("/admin/login", AuthenticateAdmin)

	adminKeyMiddleware := r.Group("/admin/register")
	adminKeyMiddleware.Use(middleware.AdminKeyMiddleware())
	adminKeyMiddleware.POST("/", InsertAdmin)

	adminTokenMiddleware := r.Group("/admin")
	adminTokenMiddleware.Use(middleware.AdminTokenMiddleware())
	adminTokenMiddleware.GET("/", GetAdmins)
	adminTokenMiddleware.GET("/user", GetUser)
	adminTokenMiddleware.GET("/seller", GetSeller)
	adminTokenMiddleware.PATCH("/password", UpdateAdminPassword)
	adminTokenMiddleware.DELETE("/", DeleteAdmin)

	// category
	r.GET("/category", GetCategories)
	categoryMiddleware := r.Group("/category")
	categoryMiddleware.Use(middleware.CategoryMiddleware())
	categoryMiddleware.POST("/", InsertCategory)
	categoryMiddleware.PATCH("/:id", UpdateCategory)
	categoryMiddleware.DELETE("/:id", DeleteCategory)

	// user
	r.POST("/user/register", InsertUser)
	r.POST("/user/login", AuthenticateUser)
	userMiddleware := r.Group("/user")
	userMiddleware.Use(middleware.UserMiddleware())
	userMiddleware.PATCH("/password", UpdateUserPassword)
	userMiddleware.PATCH("/data", UpdateUserData)
	userMiddleware.PATCH("/seller", UpdateUserToSeller)
	userMiddleware.DELETE("/", DeleteUser)

	// product
	r.GET("/product", GetProducts)
	productMiddleware := r.Group("/product")
	productMiddleware.Use(middleware.ProductMiddleware())
	productMiddleware.POST("/", InsertProduct)
	productMiddleware.PATCH("/:id", UpdateProduct)
	productMiddleware.DELETE("/:id", DeleteProduct)

	// invoices
	invoiceMiddleware := r.Group("/invoice")
	invoiceMiddleware.Use(middleware.InvoiceMiddleware())
	invoiceMiddleware.GET("/", GetInvoices)
	invoiceMiddleware.POST("/", InsertInvoice)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
