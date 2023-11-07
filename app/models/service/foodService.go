package service

import (
	"fmt"
	"guruchan-back/app/models/db"
	models "guruchan-back/app/models/entity"

	"github.com/labstack/echo/v4"
)

var result models.FoodSlice
var err error

// 查询附复数
func QueryAllFoodList(userID string, typeCode string) (models.FoodSlice, error) {
	//typeCode 0:all 1:気に入り
	if typeCode == "0" {
		//All
		result, err = db.QueryAllFoodList(userID)
	}
	if typeCode == "1" {
		//favorite
		result, err = db.QueryFavoriteFoodList(userID, typeCode)
	}

	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return result, nil

}

// 添加新的
func CreateFood(e echo.Context, foodReq *models.Food) error {
	foodObj := models.Food{}
	foodObj = *foodReq
	err := db.InsertFood(foodObj)
	if err != nil {
		return err
	}
	return nil
}

// 更新
func UpdateFood(e echo.Context, foodReq *models.Food) error {
	foodObj := models.Food{}
	foodObj = *foodReq
	err := db.UpdateFood(foodObj)
	if err != nil {
		return err
	}
	return nil
}

// 增加和修改验证Req
func CheckReqCreate(e echo.Context) (*models.Food, error) {
	foodReq := models.Food{}
	err := e.Bind(&foodReq)
	if err != nil {
		return nil, err
	}
	return &foodReq, nil
}
