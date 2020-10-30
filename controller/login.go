package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	name := c.PostForm("name")
	password := c.PostForm("password")
	fmt.Println(name, password)
	if (name=="cs" && password=="123") {
		c.JSON(http.StatusOK, gin.H{
			"code":0,
			"msg": http.StatusText(http.StatusOK),
		})
		name, err := c.Cookie("user_name")
		if err != nil {
			c.SetCookie("user_name", name, 3600, "/", "localhost", false, true)
		}
	} else{
		c.Redirect(http.StatusFound, "index")
	}

}
