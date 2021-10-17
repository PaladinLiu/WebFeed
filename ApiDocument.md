# API Document

（顺带把服务层对应的函数的大致思路写下）

### 用户系统 User

对象为用户，支持用户登录、创建、注销、修改、查看信息

[问题1] 关于key的问题，我前端在开始的时候是不知道uuid是什么的，所以查询如果按uuid去查在开始是行不通的，我目前的解决方法是登录过后返回uuid。但在之后处理其他推文之类的问题的时候也会出现问题，解决方法都是创建后返回，但这必然导致多次访问的问题。所以查询有其他更好的解决方法吗，比如用其他我能在前端本地就能获取的值去查（比如user_name）

[问题2] 关于转赞评uuid 如果我拿到uuid不能知道他的类别的话，岂不是需要每次都要遍历所有表

##### 1.创建用户

创建用户 同时分配uuid

Request:

```
URL: api/v1/user/create

POST 
{
	"user_name": xxx,
    "passwd": xxx,
    "nick_name": xxx,
}
```

Response:

```
status: 201
{"code": "201", "msg": "create success"}
status: 400
{"code": "400", "msg": "create failed"}
```



##### 2.用户登录

登陆状态需要设置？

Request:

```
URL: api/v1/user/login

POST 
{
	"user_name": xxx,
    "passwd": xxx,
}
```

Response:

```
status: 200
{"code": "200", "msg": "login success"}

status: 401
{"code": "401", "msg": "login failed"}
```



##### 3.修改用户信息

除id外全部覆盖

Request:

```
URL: api/v1/user/modify

PATCH
{
	"user_name": xxx,
	"passwd": xxx,
	"nickname": xxx,
}
```

Response:

```
status: 200
{"code": "200", "msg": "modify success"}
status: 401
{"code": "401", "msg": "modify failed"}
```



##### 4.注销用户

Request:

```
URL: api/v1/user/delete

DELETE
{
	"uuid": xxx,
}
```

Response:

```
status: 200
{"code": "200", "msg": "delete success"}
status: 401
{"code": "401", "msg": "delete failed"}
```



##### 5.查看用户信息

Request:

```
URL: api/v1/user/query

GET
{
	"uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":{
		"uuid": xxx,
		"user_name": xxx,
		"nickname": xxx,
	}
	"msg": "query success"
}
	
status: 401
{"code": "401", "msg": "query failed"}
```





### 互动系统 Interact

转赞评

##### 1.创建评论 

创建评论，生成uuid

Request:

```
URL: api/v1/interact/createComment

POST
{
	"user_uuid": xxx,
	"tweet_uuid": xxx,
	"content": xxx,
}
```

Response:

```
status: 201
{
	"code": "201", 
	"msg": "create success"，
}
	
status: 401
{"code": "401", "msg": "create failed"}
```



##### 2.修改评论 

内容覆盖

Request:

```
URL: api/v1/interact/updateComment

PATCH
{
	"comment_uuid": xxx,
	"content": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "modify success"，
}
	
status: 401
{"code": "401", "msg": "modify failed"}
```



##### 3.删除评论

Request:

```
URL: api/v1/interact/deleteComment

DELETE
{
	"comment_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "delete success"，
}
	
status: 401
{"code": "401", "msg": "delete failed"}
```



##### 4.创建点赞

Request:

```
URL: api/v1/interact/createLike

POST
{
	"user_uuid": xxx,
	"tweet_uuid": xxx,
}
```

Response:

```
status: 201
{
	"code": "201", 
	"msg": "create success"，
}
	
status: 401
{"code": "401", "msg": "create failed"}
```



##### 5.取消点赞

Request:

```
URL: api/v1/interact/deleteLike

DELETE
{
	"like_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "delete success",
}
	
status: 401
{"code": "401", "msg": "delete failed"}
```



##### 6.创建转发 new tweet 

Request:

```
URL: api/v1/interact/createRetweet

POST
{
	"user_uuid": xxx,
	"tweet_uuid": xxx,
}
```

Response:

```
status: 201
{
	"code": "201", 
	"msg": "create success"，
}
	
status: 401
{"code": "401", "msg": "create failed"}
```



##### 7.取消转发

Request:

```
URL: api/v1/interact/deleteRetweet

DELETE
{
	"retweet_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "delete success"，
}
	
status: 401
{"code": "401", "msg": "delete failed"}
```



##### 8.查看某条推文下的转发

Request:

```
URL: api/v1/interact/getRetweet

GET
{
	"tweet_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":
		[{"retweet_uuid": xxx,
		  "user_uuid": xxx,
		  "time": xxx,
		  },{},{}...],
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```



##### 9.查看某条推文下的赞

Request:

```
URL: api/v1/interact/getLike

GET
{
	"tweet_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":
		[{"like_uuid": xxx,
		  "user_uuid":xxx,
		  "time": xxx,
		  },{},{}...],
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```



##### 10.查看某条推文下的评论

Request:

```
URL: api/v1/interact/getComment

GET
{
	"tweet_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":
		[{"comment_uuid": xxx,
		  "user_uuid":xxx,
		  "time": xxx,
		  "content": xxx,
		  },{},{}...],
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```





### 查找系统 query

[问题3] 后期问题 查找做补全、近似、根据内容查找推文等功能在服务层完成？猜测： 接收请求-处理近似过后再查库，但这样还是需要在处理内容的时候遍历库？目前查找功能仅支持uuid和发推人

[问题4] 查找 现实中应该无论按id nickname username 都可以查到 是按这三者直接进库查 还是经过了转换呢（转换类似dns，把查到的对应表存在本地这种，本地没有就去查库）

##### 1.查找用户

这里需要加入关注列表？

三个参数有一个即可

Request:

```
URL: api/v1/query/user

GET
{
	"uuid": xxx,
	"nickname": xxx,
	"username": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":{
		"uuid": xxx,
		"user_name": xxx,
		"nickname": xxx,
	}
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```



##### 2.根据uuid查找推文

Request:

```
URL: api/v1/query/tweetByUUID

GET
{
	"uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":{
		"uuid": xxx,
		"time": xxx,
		"user_uuid": xxx,
		"content": xxx,
	}
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```



##### 3.根据发推人查找推文

Request:

```
URL: api/v1/query/tweetByUser

GET
{
	"user_uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"data":{
		"uuid": xxx,
		"time": xxx,
		"user_uuid": xxx,
		"content": xxx,
	}
	"msg": "query success",
}
	
status: 401
{"code": "401", "msg": "query failed"}
```





### 推荐系统 recommend

##### 1.关注

Request:

```
URL: api/v1/recommend/follow

GET
{
	"uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "follow success",
}
	
status: 401
{"code": "401", "msg": "follow failed"}
```



##### 2.取消关注

Request:

```
URL: api/v1/recommend/unfollow

GET
{
	"uuid": xxx,
}
```

Response:

```
status: 200
{
	"code": "200", 
	"msg": "unfollow success",
}
	
status: 401
{"code": "401", "msg": "unfollow failed"}
```



3.