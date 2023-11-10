package handlers

import (
	"fmt"
	"guruchan-back/app/constants"
	"guruchan-back/app/models/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Todo 新收藏某一个food
// 查询
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

// 添加
func CreateFood(e echo.Context) error {
	//验证req
	foodReq, err := service.CheckFoodReqCreate(e)
	if err != nil {
		// 处理check请求验证
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusBadRequest, result)
		// return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	serviceErr := service.CreateFoodObj(foodReq)
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

// 更新
func UpdateFood(e echo.Context) error {
	//验证格式
	foodReq, err := service.CheckFoodReqCreate(e)
	if err != nil {
		// 处理check请求验证
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusBadRequest, result)
		// return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	err = service.UpdateFood(foodReq)
	if err != nil {
		// 处理查询错误
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusInternalServerError, result)
		// return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	// 返回 HTTP 200 响应
	result := constants.Result{Code: 1, Message: "Success"}
	return e.JSON(http.StatusOK, result)
}

// 論理削除
func LogicalDeleteFood(e echo.Context) error {
	foodId := e.Param("foodId")
	err := service.LogicalDeleteFood(foodId)
	if err != nil {
		result := constants.Result{Code: 0, Message: err.Error()}
		return e.JSON(http.StatusInternalServerError, result)
	}
	result := constants.Result{Code: 1, Message: "Success"}
	return e.JSON(http.StatusOK, result)
}

// 删除,实际删除,不用
// func DeleteFood(e echo.Context) error {
// 	rows, err := service.DeleteFood(e)
// 	if err != nil {
// 		// 处理错误
// 		result := constants.Result{Code: 0, Message: err.Error()}
// 		return e.JSON(http.StatusInternalServerError, result)
// 		// return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}
// 	// 返回 HTTP 200 响应
// 	fmt.Println(rows)
// 	result := constants.Result{Code: 1, Message: "Success:deleted"}
// 	return e.JSON(http.StatusOK, result)
// }
