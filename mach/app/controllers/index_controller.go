package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}