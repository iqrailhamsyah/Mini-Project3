package main

import (
	"crud/modules/admin"
	"crud/modules/customer"
	"crud/modules/superadmin"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	engine := gin.New()

	// membuka koneksi ke db
	dbCrud := db.GormMysql()

	//menyambungkan router customer dengan db dan engine
	customerRouter := customer.NewRouter(dbCrud)
	customerRouter.Handle(engine)

	//menyambungkan router admin dengan db dan engine
	adminRouter := admin.NewRouter(dbCrud)
	adminRouter.Handle(engine)

	//menyambungkan router superadmin dengan db dan engine
	superadminRouter := superadmin.NewRouter(dbCrud)
	superadminRouter.Handle(engine)

	//menjalankan engine
	errRouter := engine.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
