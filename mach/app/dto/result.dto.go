package dto

import "mach/app/models/profile"

type Result struct {
	Status bool
	Msg string
}

type ProfileResultList struct {
	Items []profile.Profile
}