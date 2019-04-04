package corll

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myProject/db"
	"net/http"
)

func GetUserInfo(g *gin.Context) {
	rsp := new(Rsp)
	db := db.InitMysqlDb()
	Rows, err := db.Query("select  * from  userinfo;")
	fmt.Println(Rows, err)
	var user = new(UserInfo)
	var users []UserInfo
	for Rows.Next() {
		//Rows.Scan(&user.Id,&user.UserName)
		Rows.Scan(user)
		users = append(users, *user)
	}
	rsp.Code = 200
	rsp.Data = users
	g.JSON(http.StatusOK, rsp)
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}
