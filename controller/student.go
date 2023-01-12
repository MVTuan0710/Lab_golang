package controller

import (
	"myapp/model"
	db "myapp/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
	//students, err := GetRepoStudents()
	//if err != nil {
	//	return c.JSON(http.StatusNotFound, utils.NewError(err))
	//}
	//return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := db.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
func Create() ([]model.Students, error) {

}
