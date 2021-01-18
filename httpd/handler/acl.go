package handler

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        string
	LoginType string
}
type App struct {
	Id          string
	DbName      string
	Environment string
}
type JWT struct {
	Role string
	User User
	App  App
}

func TokenExtraction(jwt *string) (string, error) {
	just_the_token := strings.Split(*jwt, " ")[1]
	just_the_body := strings.Split(just_the_token, ".")[1]
	sDec, _ := b64.StdEncoding.DecodeString(just_the_body)

	return string(sDec), nil
	//auth_token := JWT{}
	//json.Unmarshal(sDec, &auth_token)
	//return &auth_token, nil
}

func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth_token := c.Request.Header["Authorization"][0]
		res, err := TokenExtraction(&auth_token)
		if err != nil {
			c.JSON(http.StatusForbidden, err.Error())
			return
		}
		//fmt.Print(res)
		//userid := "s"  //res.User.Id
		//role := "s"    //res.Role
		//db_name := "s" //res.App.DbName
		//c.JSON(http.StatusOK, gin.H{"userid": userid, "role": role, "db_name": db_name})
		c.JSON(401, res)
	}
}
