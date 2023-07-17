package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

func (t Tag) List(c *gin.Context) {

}

func (t Tag) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "123",
	})
}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
