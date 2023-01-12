package routers

import (
	"github.com/labstack/echo/v4"
	"myapp/controller"
)

func Api() {

	e := echo.New()

	e.GET("/student", controller.GetStudents)
	e.POST("/create", controller.Create)
	//e.GET("/index", controller)
	//e.POST("/store", Controllers.Store)
	//e.PUT("/update/:id", Controllers.Update)
	//e.DELETE("/delete/:id", Controllers.Delete)

	e.Logger.Fatal(e.Start(":1323"))
}
