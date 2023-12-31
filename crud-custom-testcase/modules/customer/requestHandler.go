package customer

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	ctrl ControllerCustomer
}

func NewCustomerRequestHandler(db *gorm.DB) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctrl: ControllerCustomer{
			uc: UsecaseCustomer{
				customerRepo: repositories.NewCustomer(db),
			},
		},
	}
}

func (rh RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}

	//cek untuk binding request json with gin context, jika request tidak bisa dibind maka "bad request"
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	//cek untuk melakukan perintah oleh controller untuk melakukan create customer berdasarkan request, jika terdapat error maka status "internal server error"
	res, err := rh.ctrl.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	//jika tak ada masalah maka status menjadi "ok"
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.GetCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerCustomer) UpdateCustomerById(c *gin.Context) {
	id := c.Param("id")
	request := CustomerParam{}

	// Bind JSON
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	response, err := rh.ctrl.UpdateCustomerById(uint(customerID), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, response)
}

func (rh RequestHandlerCustomer) DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
	}

	err = rh.ctrl.DeleteCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete Customer Data Successfully",
	})
}
