package models

import "time"

type Profile struct {
	ID      	int    		`json:"id"`
	Name   		string 		`json:"name"`
	UserID 		int 		`json:"userId"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	int 		`json:"updatedAt"`
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