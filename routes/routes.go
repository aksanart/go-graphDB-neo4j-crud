package routes

import (
	"github.com/aksanart/go-graphDB-neo4j-crud/controller"
	"github.com/aksanart/go-graphDB-neo4j-crud/model"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/get_person", controller.GetPerson(model.PersonDB{}))
	r.POST("/add_person", controller.AddPerson(model.PersonDB{}))
	r.POST("/follow", controller.Follow(model.PersonDB{}))
	r.POST("/unfollow", controller.UnFollow(model.PersonDB{}))
	r.GET("/view_graph", controller.ViewGraph)
	return r
}
