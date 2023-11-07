package routes

import (
	"guruchan-back/app/handlers"

	"github.com/labstack/echo/v4"
)

// 注册所有路由
func RegisterRoutes(e *echo.Echo) {
	// 定义路由和处理函数映射
	var baseurl string = "food"
	if baseurl == "food" {
		RegisterRouteFood(e, baseurl)
	} else if baseurl == "user" {
		RegisterRouteUser(e, baseurl)
	} else if baseurl == "tag" {
		RegisterRouteTag(e, baseurl)
	}
}

// food
func RegisterRouteFood(e *echo.Echo, baseurl string) {
	e.GET(baseurl+"/getFoodList/:userId/:typeCode", handlers.GetFoodList)
	//typeCode 0:all 1:気に入り
	e.GET(baseurl+"/getOneFoodL/:userId/:foodId", handlers.GetOneFood)
	e.POST(baseurl+"/CreateFood", handlers.CreateFood)
	e.PUT(baseurl+"/UpdateFood", handlers.UpdateFood)
	//
	e.DELETE(baseurl+"/DeleteFood/:fodId", handlers.DeleteFood)

}

// user
func RegisterRouteUser(e *echo.Echo, baseurl string) {
	// e.GET(baseurl+"/", handlers.foodObj)
	// e.POST(baseurl+"/", handlers.CreateRestaurant)
	// e.PUT(baseurl+"/", handlers.CreateRestaurant)
	// e.DELETE(baseurl+"/", handlers.DeleteRestaurant)
}

// tag
func RegisterRouteTag(e *echo.Echo, baseurl string) {
	e.GET(baseurl+"/getTagList/:userId", handlers.GetTagList)
	e.POST(baseurl+"/CreateTag", handlers.CreateTag)
	e.PUT(baseurl+"/UpdateTag", handlers.UpdateTag)
	e.DELETE(baseurl+"/DeleteTag", handlers.DeleteTag)
}
