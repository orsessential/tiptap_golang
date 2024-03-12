package routers

import (
	"gin-framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:bookID", controllers.GetBook)
	router.PUT("/book/:bookID", controllers.UpdateBook)
	router.DELETE("/book/:bookID", controllers.DeleteBook)
	router.GET("/book/allBooks", controllers.GetAllBooks)

	return router
}
