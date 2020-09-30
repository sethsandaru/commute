package controllers

import (
	"github.com/gin-gonic/gin"
	"mach/app/dto"
	"net/http"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

// Render type error for all controllers
func renderError(c *gin.Context, e error) {
	result := new(dto.Result)
	result.Status = false
	result.Msg = e.Error()

	c.JSON(http.StatusNotAcceptable, result)
}