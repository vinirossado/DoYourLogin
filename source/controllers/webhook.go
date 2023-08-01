package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GithubWebhookUserManagePayload struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`

	Member struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		Type  string `json:"type"`
	}

	Sender struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		Type  string `json:"type"`
	}

	Team struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	Organization struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	}
}

func HandlePost(c *gin.Context) {
	var postData GithubWebhookUserManagePayload
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Novo post recebido do Github: %s", postData)

	c.JSON(http.StatusOK, gin.H{"status": postData})
}
