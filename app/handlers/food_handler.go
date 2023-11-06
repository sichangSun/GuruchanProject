package handlers

import (
	"fmt"
	"guruchan-back/app/models/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFoodList(e echo.Context) error {
	userID := e.Param("userId")
	typeCode := e.Param("typeCode")
	fmt.Println(userID, typeCode)
	//查询
	results, err := service.QueryAllFoodList(userID, typeCode)
	if err != nil {
		// 处理查询错误
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 返回 HTTP 200 响应
	return e.JSON(http.StatusOK, results)
}

func GetOneFood(e echo.Context) error {
	return nil
}

func CreateFood(e echo.Context) error {
	//添加
	err := service.CreateFood(e)
	if err != nil {
		// 处理查询错误
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// 返回 HTTP 200 响应
	return e.JSON(http.StatusOK, nil)
}

func UpdateFood(e echo.Context) error {
	return nil
}

func DeleteFood(e echo.Context) error {
	return nil
}
