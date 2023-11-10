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

// bd初始化
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

////

type FoodDao struct {
}
type TagDao struct {
}

// get某一个对象
func (f FoodDao) GetObjById(ObjId string) (*models.Food, error) {
	ctx := context.Background()
	raws, err := models.Foods(qm.Where("foodId=?", ObjId)).One(ctx, db)
	return raws, err
}
func (t TagDao) GetObjById(ObjId string) (*models.Tag, error) {
	ctx := context.Background()
	raws, err := models.Tags(qm.Where("tagId=?", ObjId)).One(ctx, db)
	return raws, err
}

// 更新Obj
func (f FoodDao) UpdateObj(foodReq models.Food) error {
	ctx := context.Background()
	_, err := foodReq.Update(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
func (t TagDao) UpdateObj(tagReq models.Tag) error {
	ctx := context.Background()
	_, err := tagReq.Update(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// //////////
// 查询
func (f FoodDao) QueryAllList(userID string) (models.FoodSlice, error) {
	ctx := context.Background()
	raws, err := models.Foods(qm.Where("userId=?", userID)).All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return raws, nil
	// defer db.Close()
}
func (t TagDao) QueryAllList(userID string) (models.TagSlice, error) {
	ctx := context.Background()
	raws, err := models.Tags(qm.Where("userId=?", userID)).All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return raws, nil
	// defer db.Close()
}

// 查询收藏food
func QueryFavoriteFoodList(userID string) (models.FoodSlice, error) {
	ctx := context.Background()
	raws, err := models.Foods(qm.Where("isLikeflag=?", "1")).All(ctx, db)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return raws, nil
}

// 新增
func (f FoodDao) InsertObj(foodReq models.Food) error {
	ctx := context.Background()
	err := foodReq.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
func (t TagDao) InsertObj(tagReq models.Tag) error {
	ctx := context.Background()
	err := tagReq.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

// 查询是否存在
func (FoodDao) IsExists(foodId int) (bool, error) {
	ctx := context.Background()
	isExist, err := models.Foods(qm.Where("foodId=?", foodId)).Exists(ctx, db)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
func (TagDao) IsExists(tagId int) (bool, error) {
	ctx := context.Background()
	isExist, err := models.Tags(qm.Where("tagId=?", tagId)).Exists(ctx, db)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

// // 删除
// func DeleteFood(foodReq *models.Food) (int64, error) {
// 	ctx := context.Background()
// 	rows, err := foodReq.Delete(ctx, db)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rows, nil
// }
