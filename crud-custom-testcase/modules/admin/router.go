package admin

import (
	"crud/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterAdmin struct {
	AdminRequestHandler RequestHandlerAdmin
}

func NewRouter(db *gorm.DB) RouterAdmin {
	return RouterAdmin{AdminRequestHandler: NewAdminRequestHandler(db)}
}

func (r RouterAdmin) Handle(engine *gin.Engine) {
	basePath := "/admin"

	admin := engine.Group(basePath)
	admin.POST("/login", r.AdminRequestHandler.LoginAdmin)
	admin.POST("/register-admin", r.AdminRequestHandler.RegisterAdmin)
	admin.Use(middleware.Authentication())
	admin.GET("/:id", r.AdminRequestHandler.GetAdminById)
	admin.PUT("/:id", r.AdminRequestHandler.UpdateAdminById)
	admin.DELETE("/:id", r.AdminRequestHandler.DeleteAdminById)

	// About Customer
	admin.Use(middleware.Authentication())
	admin.POST("/create-customer", r.AdminRequestHandler.CreateCustomer)
	admin.DELETE("/delete-customer/:id", r.AdminRequestHandler.DeleteCustomerById)
	admin.GET("/customers", r.AdminRequestHandler.GetAllCustomers)
	admin.GET("/fetch-customers", r.AdminRequestHandler.SaveCustomersFromAPI)
}
