package db

import (
	"context"
	"database/sql"
	"fmt"
	models "guruchan-back/app/models/entity"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var db *sql.DB

func init() {
	confDB := "root:root1234@tcp(localhost:3306)/guruchanDB?parseTime=true"
	var err error
	db, err = sql.Open("mysql", confDB)
	if err != nil {
		fmt.Println("db连接失败")
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DBerr:", err)
		fmt.Println("dbPing失败")
		panic(err.Error())
		//panic(err) qubie
	}
}
func QueryAllFoodList(userID string) (models.FoodSlice, error) {
	// 原始打开数据库查询方法
	// rows, err := db.Query("SELECT * FROM g_food")
	// if err != nil {
	// 	fmt.Println("seccuse")

	// }
	// fmt.PrintFoodln("rows", rows)
	ctx := context.Background()
	raws, err := models.Foods().All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return raws, nil

	// defer db.Close()

}
func QueryFavoriteFoodList(userID string, typeCode string) (models.FoodSlice, error) {
	// 初始数据库查询方法
	// rows, err := db.Query("SELECT * FROM g_food WHERE isLikeflag = 1")
	// if err != nil {
	// 	fmt.Println("seccuse")
	// }
	// fmt.Println("rows", rows)
	ctx := context.Background()
	raws, err := models.Foods(qm.Where("isLikeflag=?", typeCode)).All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return raws, nil
}

// 返回某一个food对象
func GetFoodByFoodId(foodId string) (*models.Food, error) {
	ctx := context.Background()
	raws, err := models.Foods(qm.Where("foodId=?", foodId)).One(ctx, db)
	return raws, err
}

// 新增food
func InsertFood(foodReq models.Food) error {
	ctx := context.Background()
	err := foodReq.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// 更新
func UpdateFood(foodReq models.Food) error {
	ctx := context.Background()
	_, err := foodReq.Update(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// 论理删除(更新isdel)
// func LogicalDeleteFood(foodId string) error {
// 	ctx := context.Background()
// 	_, err := foodReq.Update(ctx, db, boil.Infer())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// 删除
func DeleteFood(foodReq *models.Food) (int64, error) {
	ctx := context.Background()
	rows, err := foodReq.Delete(ctx, db)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// 查询是否存在
func IsExists(foodId string) (bool, error) {
	ctx := context.Background()
	isExist, err := models.Foods(qm.Where("foodId=?", foodId)).Exists(ctx, db)
	if err != nil {
		return false, err
	}
	return isExist, nil

}
