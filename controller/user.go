package controller

import (
	"gin-project-base/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//JSON格式输出
func GetUser(c *gin.Context){
	id := c.Param("id")
	var user model.Users
	nid,_ := strconv.Atoi(id)
	arr := user.GetList(nid)

	//var res = map[string]interface{}{"name":"","age": "", "jiguan":""}
	//if getData == "1" {
	//	res["name"] = "bruce";
	//	res["age"] = 32
	//	res["jiguan"] = "luotian"
	//}

	c.JSON(http.StatusOK, gin.H{"params": arr})
}