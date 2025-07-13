package main

import (
	"net/http"

	"github.com/V-enekoder/GasManager/config"
	commerce "github.com/V-enekoder/GasManager/src/Commerce"
	council "github.com/V-enekoder/GasManager/src/Council"
	delivery "github.com/V-enekoder/GasManager/src/Delivery"
	disabled "github.com/V-enekoder/GasManager/src/Disabled"
	municipality "github.com/V-enekoder/GasManager/src/Municipality"
	order "github.com/V-enekoder/GasManager/src/Order"
	orderstate "github.com/V-enekoder/GasManager/src/OrderState"
	payment "github.com/V-enekoder/GasManager/src/Payment"
	paymentstate "github.com/V-enekoder/GasManager/src/PaymentState"
	report "github.com/V-enekoder/GasManager/src/Report"
	reportstate "github.com/V-enekoder/GasManager/src/ReportState"
	reporttype "github.com/V-enekoder/GasManager/src/ReportType"
	typecylinder "github.com/V-enekoder/GasManager/src/TypeCylinder"
	user "github.com/V-enekoder/GasManager/src/User"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
	config.SyncDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.SetTrustedProxies(nil)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	commerce.RegisterRoutes(r)
	council.RegisterRoutes(r)
	delivery.RegisterRoutes(r)
	disabled.RegisterRoutes(r)
	municipality.RegisterRoutes(r)
	order.RegisterRoutes(r)
	orderstate.RegisterRoutes(r)
	payment.RegisterRoutes(r)
	paymentstate.RegisterRoutes(r)
	report.RegisterRoutes(r)
	reporttype.RegisterRoutes(r)
	reportstate.RegisterRoutes(r)
	typecylinder.RegisterRoutes(r)
	user.RegisterRoutes(r)

	r.Run()
}
