package service

import (
	"fmt"
	"guruchan-back/app/models/db"
	"guruchan-back/app/models/entity"

	"github.com/labstack/echo/v4"
)

var result entity.GFoodSlice
var err error

// 查询附复数
func QueryAllFoodList(userID string, typeCode string) (entity.GFoodSlice, error) {
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
	// request := e.Request().Body.Read()
	// foodObj := entity.GFood{
	// 	UUserId:request.
	// 	}
	// foodObj.FoodId
	//存数据自动采番问题?
}
