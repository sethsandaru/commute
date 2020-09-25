package models

import "time"

type Profile struct {
	tableName struct{} `pg:"profiles"`

	Id      	int    		`json:"id" pg:"id"`
	Name   		string 		`json:"name" pg:"name"`
	UserId 		int 		`json:"userId" pg:"user_id"`
	CreatedAt 	time.Time 	`json:"createdAt" pg:"created_at,default:now()"`
	UpdatedAt 	time.Time 	`json:"updatedAt" pg:"updated_at"`
}

func getProfileById(id int) *Profile {
	profile := new(Profile)

	err := db.Model(profile).
		Where("id = ?", id).
		Limit(1).
		Select()
	if err != nil {
		panic(err)
	}

	return profile
}

func createProfile(profile *Profile) bool {
	result, err := db.Model(profile).Insert()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected() > 0
}

func updateProfile(profile *Profile) bool {
	result, err := db.Model(profile).Update()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected() > 0
}

func deleteProfile(id int) bool {
	profile := new(Profile)
	profile.Id = id

	result, err := db.Model(profile).
		Where("id = ?", id).
		Limit(1).
		Delete()

	if err != nil {
		panic(err)
	}

	return result.RowsAffected() > 0
}