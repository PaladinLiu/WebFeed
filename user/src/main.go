package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
    "strconv"
)

var db *sql.DB

type userInfo struct {
	id int
	uuid int
	username string
	passwd string
}

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

	// set addr&port
	err := router.Run(":8080")
	if err != nil {return}
}



//01.userCreate
//URL: api/v1/user/create
//POST x-www-form-urlencoded
func userCreate(context *gin.Context){
	isStatusOK := true //exit flag
	username := context.PostForm("user_name")
	passwd := context.PostForm("passwd")
	uuid := uuidGenerator(1)

	//db execution
	_, err := db.Exec("INSERT INTO tb_user (uuid, user_name, passwd) VALUES(?, ?, ?)", uuid, username, passwd)

	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
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


//02.userLogin
//URL: api/v1/user/login
//POST x-www-form-urlencoded
func userLogin(context *gin.Context){
	var stdUser userInfo
	isStatusOK := true //exit flag
	username := context.PostForm("user_name")
	passwd := context.PostForm("passwd")

	//db query
	row := db.QueryRow("select * from tb_user where user_name=?", username)
	//when scan completed, connection closes.
	err := row.Scan(&stdUser.id, &stdUser.uuid, &stdUser.username, &stdUser.passwd)
	if err != nil{
		fmt.Printf("scan failed, err:%v", err)
		isStatusOK = false
	}else if stdUser.passwd != passwd{
		isStatusOK = false
	}
	fmt.Printf("%v ",stdUser.passwd)

	//response
	if isStatusOK == true{
		context.JSON(http.StatusOK,gin.H{
			"code": 200,
			"msg": "login success",
		})
	}else{ //isStatusOK == false
		context.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg": "login failed",
		})
	}
}


//03.userInfoModify
//URL: api/v1/user/update
//PATCH x-www-form-urlencoded
func userInfoUpdate(context *gin.Context){
	isStatusOK := true //exit flag
	username := context.PostForm("user_name")
	newPasswd := context.PostForm("passwd")

	//db execution
	_, err := db.Exec("UPDATE tb_user SET passwd=? WHERE user_name=?", newPasswd, username)
	if err != nil{
		fmt.Printf("Update failed,err:%v",err)
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


//04.userDelete
//URL: api/v1/user/delete
//DELETE Query
func userDelete(context *gin.Context){
	isStatusOK := true //exit flag
	uuid := context.Query("uuid")

	//db execution
	_, err := db.Exec("delete from tb_user where uuid=?", uuid)
	if err != nil {
		fmt.Printf("user delete faied, error:[%v]", err.Error())
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

//05.getUserInfo
//URL: api/v1/user/query
//GET Query
func getUserInfo(context *gin.Context){
	var stdUser userInfo
	isStatusOK := true //exit flag
	uuid := context.Query("uuid")

	//db query
	row := db.QueryRow("select * from tb_user where uuid=?", uuid)
	//when scan completed, connection closes.
	err := row.Scan(&stdUser.id, &stdUser.uuid, &stdUser.username, &stdUser.passwd)
	if err != nil{
		fmt.Printf("scan failed, err:%v", err)
		isStatusOK = false
	}

	//response
	if isStatusOK == true{
		//create response data dict
		data := map[string]string{
			"id": strconv.Itoa(stdUser.id),
			"uuid": strconv.Itoa(stdUser.uuid),
			"user_name": stdUser.username,
			"passwd": stdUser.passwd,
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
