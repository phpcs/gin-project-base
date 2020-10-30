package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
)

func Upload(c *gin.Context)  {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	c.SaveUploadedFile(file, "./")
	c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
}
