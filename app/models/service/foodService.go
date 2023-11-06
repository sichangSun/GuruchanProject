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
func CreateFood(e echo.Context) error {
	var foodReq models.Food
	err := e.Bind(&foodReq)
	if err != nil {
		return err
	}
	// request := e.Request().Body.Read()
	// foodObj := models.GFood{
	// 	UUserId:request.
	// 	}
	// foodObj.FoodId
	//存数据自动采番问题?
	return nil
}
