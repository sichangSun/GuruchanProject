package common

// import "github.com/labstack/echo"

type Inter interface {
	// 返回某一个food/tag对象
	GetObjById(ObjId string) (interface{}, error)
	UpdateObj(interface{}) error
	QueryAllList(userID string) (interface{}, error)
	IsExists(foodId int) (bool, error)
	InsertObj(interface{}) error
}
