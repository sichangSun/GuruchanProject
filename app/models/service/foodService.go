package service

import (
	"errors"
	"fmt"
	"guruchan-back/app/models/db"
	models "guruchan-back/app/models/entity"

	"github.com/labstack/echo/v4"
	// "github.com/volatiletech/null/v8"
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

	//TODO:isdel论理删除过滤
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

// 论理删除
func LogicalDeleteFood(foodId string) error {
	// isExist, err := db.IsExists(foodId)
	// if err != nil {
	// 	return err
	// }
	// if !isExist {
	// 	//不存在,返回err
	// 	errMsg := "faild to find this colum"
	// 	err := errors.New(errMsg)
	// 	return err
	// }
	foodObj, err := db.GetFoodByFoodId(foodId)
	if err != nil {
		return err
	}
	if foodObj == nil {
		//不存在,返回err
		errMsg := "faild to find this colum"
		err := errors.New(errMsg)
		return err
	}
	//存在,做更新isdel处理
	foodObj.IsDel = 1
	err = db.UpdateFood(*foodObj)
	if err != nil {
		return err
	}
	return nil

}

// 删除,实际删除,不用
func DeleteFood(e echo.Context) (int64, error) {
	foodReq := models.Food{}
	if err := e.Bind(&foodReq); err != nil {
		return 0, err
	}
	rows, err := db.DeleteFood(&foodReq)
	if err != nil {
		return 0, err
	}
	return rows, nil
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
