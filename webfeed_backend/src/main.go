package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main(){
	db = initDB()

	router := gin.Default()

	user := router.Group("/v1/user")
	{
		user.POST("/create", userCreate)
		user.POST("/login", userLogin)
		user.PATCH("/update", userInfoUpdate)
		user.DELETE("/delete", userDelete)
		user.GET("/query", getUserInfo)
	}

	interact := router.Group("/v1/interact")
	{
		interact.POST("/createComment", createComment)
		interact.PATCH("/updateComment", updateComment)
		interact.DELETE("/deleteComment", deleteComment)
		interact.POST("/createLike", createLike)
		interact.DELETE("/deleteLike", deleteLike)
		interact.POST("/createRetweet", createRetweet)
		interact.DELETE("/deleteRetweet", deleteRetweet)
		interact.GET("/getRetweet", getRetweet)
		interact.GET("/getLike", getLike)
		interact.GET("/getComment", getComment)
	}

	// set addr&port
	err := router.Run(":8080")
	if err != nil {return}
}


