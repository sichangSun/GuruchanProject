package service

import (
	"guruchan-back/app/common/db"
	models "guruchan-back/app/models/entity"

	"github.com/labstack/echo/v4"
)

// tag查询
func QueryTagAllList(userID string, typeCode string) (models.TagSlice, error) {
	//tag查询
	tagDao := new(db.TagDao)
	result, err := tagDao.QueryAllList(userID)

	//TODO:isdel论理删除过滤
	if err != nil {
		return nil, err
	}
	return result, nil

}
func CheckTagReqCreate(e echo.Context) error {
	var obj models.Tag
	err := e.Bind(&obj)
	if err != nil {
		return err
	}
	return nil
}
