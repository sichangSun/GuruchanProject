package routes

import (
	"guruchan-back/app/handlers"

	"github.com/labstack/echo/v4"
)

// 注册所有路由
func RegisterRoutes(e *echo.Echo) {
	//food
	//查询
	e.GET("/api/food/getList/:userId/:typeCode", handlers.GetFoodList)
	//新增
	e.POST("/api/food/CreateFood", handlers.CreateFood)
	//更新
	e.PUT("/api/food/UpdateTag", handlers.UpdateFood)
	//論理削除
	e.GET("/LogicalDeleteFood/:foodId", handlers.LogicalDeleteFood)
	// e.GET("/api/food/getOneFoodL/:userId/:foodId", handlers.GetOneFood)

	//tag
	e.GET("/getTagList/:userId", handlers.GetTagList)
	e.POST("/CreateTag", handlers.CreateTag)
	e.PUT("/UpdateTag", handlers.UpdateTag)
	e.DELETE("/DeleteTag", handlers.DeleteTag)
}
