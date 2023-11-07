package handlers

import (
	"fmt"
	"guruchan-back/app/constants"
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
	//验证req
	// fmt.Println("ok")
	foodReq, err := service.CheckReqCreate(e)
	if err != nil {
		// 处理check请求验证
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusBadRequest, result)
		// return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	serviceErr := service.CreateFood(e, foodReq)
	if serviceErr != nil {
		// 处理查询错误
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusInternalServerError, result)
		// return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// 返回 HTTP 200 响应
	result := constants.Result{Code: 1, Message: "Success"}
	return e.JSON(http.StatusOK, result)
}

func UpdateFood(e echo.Context) error {
	return nil
}

func DeleteFood(e echo.Context) error {
	return nil
}
