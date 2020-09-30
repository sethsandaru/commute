package profile

import (
	"time"
	"mach/app/models"
)

type Profile struct {
	tableName struct{} `pg:"profiles"`

	Id      	int    		`json:"id" pg:"id"`
	Name   		string 		`json:"name" pg:"name"`
	UserId 		int 		`json:"userId" pg:"user_id"`
	CreatedAt 	time.Time 	`json:"createdAt" pg:"created_at,default:now()"`
	UpdatedAt 	time.Time 	`json:"updatedAt" pg:"updated_at"`
}

func GetProfiles(keyword string, limit int, start int) (error, []Profile) {
	var profiles []Profile

	query := models.DB.Model(&profiles)

	// keyword filtering
	if keyword != "" {
		query.Where("name LIKE ?", "%" + keyword + "%")
	}

	// limit & start
	if limit == 0 {
		limit = 10
	}

	err := query.Limit(limit).Offset(start).Select()
	if err != nil {
		return err, nil
	}

	return nil, profiles
}

func GetProfileById(id int) *Profile {
	profile := new(Profile)

	err := models.DB.Model(profile).
		Where("id = ?", id).
		Limit(1).
		Select()

	if err != nil {
		panic(err)
	}

	return profile
}

func CreateProfile(profile *Profile) bool {
	result, err := models.DB.Model(profile).Insert()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected() > 0
}

func UpdateProfile(profile *Profile) bool {
	result, err := models.DB.Model(profile).Update()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected() > 0
}

func DeleteProfile(id int) (error, bool) {
	profile := new(Profile)
	profile.Id = id

	result, err := models.DB.Model(profile).
		Where("id = ?", id).
		Limit(1).
		Delete()

	if err != nil {
		return err, false
	}

	return err, result.RowsAffected() > 0
}