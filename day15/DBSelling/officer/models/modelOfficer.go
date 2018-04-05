package models

type Officer struct {
	tblOfficerID    int    `json:"tblOfficerID"`
	OfficerCode     string `json:"OfficerCode"`
	OfficerName     string `json:"OfficerName"`
	OfficerPassword string `json:"OfficerPassword"`
	OfficerStatus   string `json:"OfficerStatus"`
}
