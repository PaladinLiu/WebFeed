package main

import (
	_"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"time"
)


type dictInfo struct {
	id int
	uuid int
	time time.Time
	tweetUuid int
	creatorUuid int
	contextText string
}

//01.createComment
//URL: api/v1/interact/createComment
//POST x-www-form-urlencoded
//{
//	"creator_uuid": xxx,
//	"tweet_uuid": xxx,
//	"content": xxx,
//}
func createComment(context *gin.Context){
	isStatusOK := true //exit flag
	creatorUuid := context.PostForm("user_uuid")
	tweetUuid := context.PostForm("tweet_uuid")
	contentText := context.PostForm("content")
	uuid := uuidGenerator(3)
	datetime := time.Now().Format("YYYY-MM-DD hh:mm:ss")

	//db execution
	_, err := db.Exec("INSERT INTO tb_comment (uuid, time_created, tweet_uuid, creator_uuid, content_text) VALUES(?, ?, ?, ?, ?)", uuid, datetime, creatorUuid, tweetUuid, contentText)

	if err != nil {
		fmt.Printf("comment insert failed, error:[%v]", err.Error())
		isStatusOK = false //exit flag
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusCreated,gin.H{
			"code": 201,
			"data": uuid,
			"msg": "create success",
		})
	}else{ //isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "create failed",
		})
	}
}

//02.updateComment
//URL: api/v1/interact/updateComment
//PATCH x-www-form-urlencoded
//{
//	"comment_uuid": xxx,
//	"content": xxx,
//}
func updateComment(context *gin.Context){
	isStatusOK := true //exit flag
	uuid := context.PostForm("comment_uuid")
	contentText := context.PostForm("content")

	//db update
	_, err := db.Exec("UPDATE tb_comment SET content_text=? WHERE uuid=?", contentText, uuid)
	if err != nil{
		fmt.Printf("update failed,err:%v",err)
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "update success",
		})
	}else{ //isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "update failed",
		})
	}
}

//03.deleteComment
//URL: api/v1/interact/deleteComment
//DELETE Query
//{
//	"comment_uuid": xxx,
//}
func deleteComment(context *gin.Context){
	isStatusOK := true //exit flag
	uuid := context.Query("comment_uuid")

	//db execution
	_, err := db.Exec("delete from tb_comment where uuid=?", uuid)
	if err != nil {
		fmt.Printf("comment delete faied, error:[%v]", err.Error())
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "delete success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "delete failed",
		})
	}
}


//04.createLike
//URL: api/v1/interact/createLike
//POST Query
//{
//"creator_uuid": xxx,
//"tweet_uuid": xxx,
//}
func createLike(context *gin.Context){
	isStatusOK := true //exit flag
	creatorUuid := context.Query("creator_uuid")
	tweetUuid := context.Query("tweet_uuid")
	uuid := uuidGenerator(4)
	datetime := time.Now().Format("YYYY-MM-DD hh:mm:ss")

	//db execution
	_, err := db.Exec("INSERT INTO tb_like (uuid, time_created, tweet_uuid, creator_uuid) VALUES(?, ?, ?, ?)", uuid, datetime, creatorUuid, tweetUuid)

	if err != nil {
		fmt.Printf("like insert failed, error:[%v]", err.Error())
		isStatusOK = false //exit flag
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusCreated,gin.H{
			"code": 201,
			"data": uuid,
			"msg": "create success",
		})
	}else{ //isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "create failed",
		})
	}
}

//05.deleteLike
//URL: api/v1/interact/deleteLike
//DELETE Query
//{
//"like_uuid": xxx,
//}
func deleteLike(context *gin.Context){
	isStatusOK := true //exit flag
	uuid := context.Query("like_uuid")

	//db execution
	_, err := db.Exec("delete from tb_like where uuid=?", uuid)
	if err != nil {
		fmt.Printf("like delete faied, error:[%v]", err.Error())
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "delete success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "delete failed",
		})
	}
}


//06.createRetweet
//URL: api/v1/interact/createRetweet
//POST
//{
//"user_uuid": xxx,
//"tweet_uuid": xxx,
//}
func createRetweet(context *gin.Context){
	isStatusOK := true //exit flag
	creatorUuid := context.PostForm("user_uuid")
	tweetUuid := context.PostForm("tweet_uuid")
	uuid := uuidGenerator(5)
	datetime := time.Now().Format("YYYY-MM-DD hh:mm:ss")

	//db execution
	_, err := db.Exec("INSERT INTO tb_retweet (uuid, time_created, creator_uuid, origin_tweet_uuid) VALUES(?, ?, ?, ?)", uuid, datetime, creatorUuid, tweetUuid)

	if err != nil {
		fmt.Printf("retweet insert failed, error:[%v]", err.Error())
		isStatusOK = false //exit flag
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusCreated,gin.H{
			"code": 201,
			"data": uuid,
			"msg": "create success",
		})
	}else{ //isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "create failed",
		})
	}
}


//07.deleteRetweet
//URL: api/v1/interact/deleteRetweet
//DELETE
//{
//"retweet_uuid": xxx,
//}
func deleteRetweet(context *gin.Context){
	isStatusOK := true //exit flag
	uuid := context.Query("retweet_uuid")

	//db execution
	_, err := db.Exec("delete from tb_retweet where uuid=?", uuid)
	if err != nil {
		fmt.Printf("retweet delete faied, error:[%v]", err.Error())
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "delete success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "delete failed",
		})
	}
}


//08.queryRetweet
//URL: api/v1/interact/getRetweet
//GET
//{
//"tweet_uuid": xxx,
//}
func getRetweet(context *gin.Context){
	var info dictInfo
	isStatusOK := true //exit flag
	uuid := context.Query("tweet_uuid")

	//db query
	row := db.QueryRow("select * from tb_retweet where uuid=?", uuid)
	//when scan completed, connection closes.
	err := row.Scan(&info.uuid, &info.time,  &info.creatorUuid, &info.tweetUuid)
	if err != nil{
		fmt.Printf("scan failed, err:%v", err)
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		//create response data dict
		data := map[string]string{
			"uuid": strconv.Itoa(info.uuid),
			"time_created": info.time.Format("YYYY-MM-DD hh:mm:ss"),
			"origin_tweet_uuid": strconv.Itoa(info.tweetUuid),
			"creator_uuid": strconv.Itoa(info.creatorUuid),
		}
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"data": data,
			"msg": "query success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"data": nil,
			"msg": "query failed",
		})
	}
}


//09.getLike
//URL: api/v1/interact/getLike
//GET
//{
//"like_uuid": xxx,
//}
func getLike(context *gin.Context){
	var info dictInfo
	isStatusOK := true //exit flag
	uuid := context.Query("like_uuid")

	//db query
	row := db.QueryRow("select * from tb_like where uuid=?", uuid)
	//when scan completed, connection closes.
	err := row.Scan(&info.uuid, &info.time, &info.tweetUuid, &info.creatorUuid)
	if err != nil{
		fmt.Printf("scan failed, err:%v", err)
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		//create response data dict
		data := map[string]string{
			"uuid": strconv.Itoa(info.uuid),
			"time_created": info.time.Format("YYYY-MM-DD hh:mm:ss"),
			"tweet_uuid": strconv.Itoa(info.tweetUuid),
			"creator_uuid": strconv.Itoa(info.creatorUuid),
		}
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"data": data,
			"msg": "query success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"data": nil,
			"msg": "query failed",
		})
	}
}


//10.getComment
//URL: api/v1/interact/getComment
//GET
//{
//	"comment_uuid": xxx,
//}
func getComment(context *gin.Context){
	var info dictInfo
	isStatusOK := true //exit flag
	uuid := context.Query("comment_uuid")

	//db query
	row := db.QueryRow("select * from tb_comment where uuid=?", uuid)
	//when scan completed, connection closes.
	err := row.Scan(&info.uuid, &info.time, &info.tweetUuid, &info.creatorUuid, &info.contextText)
	if err != nil{
		fmt.Printf("scan failed, err:%v", err)
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		//create response data dict
		data := map[string]string{
			"uuid": strconv.Itoa(info.uuid),
			"time_created": info.time.Format("YYYY-MM-DD hh:mm:ss"),
			"tweet_uuid": strconv.Itoa(info.tweetUuid),
			"creator_uuid": strconv.Itoa(info.creatorUuid),
			"context_text": info.contextText,
		}
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"data": data,
			"msg": "query success",
		})
	}else{ ////isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"data": nil,
			"msg": "query failed",
		})
	}
}


