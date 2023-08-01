package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PostData struct {
	Action     any `json:"action"`
	Membership any `json:"membership"`
}

func HandlePost(c *gin.Context) {

	var postData PostData
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Novo post recebido do Github: %s", postData)

	c.JSON(http.StatusOK, gin.H{"status": postData})
}
