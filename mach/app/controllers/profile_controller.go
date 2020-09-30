package controllers

import (
	"github.com/gin-gonic/gin"
	"mach/app/dto"
	"mach/app/models/profile"
	"strconv"
)

func GetProfileList(c *gin.Context) {
	err, profileLists := profile.GetProfiles("", 10, 0);

	if err != nil {
		renderError(c, err)
		return
	}

	resultBucket := new(dto.ProfileResultList)
	resultBucket.Items = profileLists
	render(c, gin.H{
		"payload": resultBucket,
	}, "")
}

func GetSingleProfile(c *gin.Context) {

}

func InsertProfile(c *gin.Context) {

}

func UpdateProfile(c *gin.Context) {

}

func DeleteProfile(c *gin.Context) {
	idParam := c.Param("id")
	id, parseErr := strconv.Atoi(idParam)

	if (parseErr != nil) {
		renderError(c, parseErr)
		return
	}

	err, status := profile.DeleteProfile(id)
	if err != nil {
		renderError(c, err)
		return
	}

	resultBucket := new(dto.Result)
	resultBucket.Status = status
	render(c, gin.H{
		"payload": resultBucket,
	}, "")
}