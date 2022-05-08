package model

type Group struct {
	ID         int    `json:"id"`
	GroupOwner int    `json:"group_owner"`
	GroupName  string `json:"group_name"`
}
