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

func GetSingleProfile(c *gin.Context)  {
	idParam := c.Param("id")
	id, parseErr := strconv.Atoi(idParam)

	if (parseErr != nil) {
		renderError(c, parseErr)
		return
	}

	profileInfo := profile.GetProfileById(id)
	render(c, gin.H{
		"payload": profileInfo,
	}, "")
}

func InsertProfile(c *gin.Context) {
	// pick data
	name, _ := c.GetPostForm("name")
	userId, _ := c.GetPostForm("user_id")
	userIdInt, _ := strconv.Atoi(userId)

	// prepare statement
	sProfile := new(profile.Profile)
	sProfile.Name = name
	sProfile.UserId = userIdInt

	// save
	status, err := profile.CreateProfile(sProfile)
	if err != nil || !status {
		render(c, gin.H{
			"payload": gin.H{
				"success": false,
				"error": "Error while adding new Profile",
			},
		}, "")
	}

	render(c, gin.H{
		"payload": gin.H{
			"success": true,
		},
	}, "")
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