package service

import (
	"errors"
	"fmt"
	"guruchan-back/app/common/db"
	models "guruchan-back/app/models/entity"
	// "github.com/volatiletech/null/v8"
)

var result models.FoodSlice
var err error

// 查询附复数
// food查询
func QueryAllFoodList(userID string, typeCode string) (models.FoodSlice, error) {
	//typeCode 0:all 1:気に入り
	if typeCode == "0" {
		//All
		foodDao := new(db.FoodDao)
		result, err = foodDao.QueryAllList(userID)
	}
	if typeCode == "1" {
		//favorite
		// result, err = db.QueryFavoriteFoodList(userID)
	}
	//TODO:isdel论理删除过滤
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return result, nil

}

// 添加新的
func CreateFoodObj(foodReq *models.Food) error {
	//处理req
	foodObj := models.Food{}
	foodObj = *foodReq
	//创建foodDao对象操作数据库
	foodDao := new(db.FoodDao)
	err := foodDao.InsertObj(foodObj)
	if err != nil {
		return err
	}
	return nil
}

// 更新
func UpdateFood(foodReq *models.Food) error {
	//处理req
	foodObj := models.Food{}
	foodObj = *foodReq
	//创建foodDao对象操作数据库
	foodDao := new(db.FoodDao)
	isHaveflg, err := foodDao.IsExists(foodObj.FoodId)
	if err != nil {
		return err
	}
	if !isHaveflg {
		//不存在,返回err
		errMsg := "faild to find this colum"
		err = errors.New(errMsg)
		return err
	}
	//更新
	err = foodDao.UpdateObj(foodObj)
	if err != nil {
		return err
	}
	return nil
}

// 论理删除
func LogicalDeleteFood(foodId string) error {
	foodDao := db.FoodDao{}
	raws, err := foodDao.GetObjById(foodId)
	if err != nil {
		return err
	}
	if raws == nil {
		//不存在,返回err
		errMsg := "faild to find this colum"
		err := errors.New(errMsg)
		return err
	}
	//更新
	raws.IsDel = 1
	err = UpdateFood(raws)
	if err != nil {
		return err
	}
	return nil

}

// // 删除,实际删除,不用
// func DeleteFood(e echo.Context) (int64, error) {
// 	foodReq := models.Food{}
// 	if err := e.Bind(&foodReq); err != nil {
// 		return 0, err
// 	}
// 	rows, err := db.DeleteFood(&foodReq)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rows, nil
// }

// 验证Req
/*
   echo上下文,中包含req的请求数据等等,所以不应该传到service层
*/
// func CheckFoodReqCreate(e echo.Context) (*models.Food, error) {
// 	foodReq := models.Food{}
// 	err := e.Bind(&foodReq)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &foodReq, nil
// }
