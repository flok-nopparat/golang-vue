package models

// Model Returns the result
type ReturnData struct {
	Total   int     `json:"total"`
	Sucess  int     `json:"success"`
	Fail    int     `json:"fail"`
	TimeUse float64 `json:"time_use"`
}

// Model to store the data
type Total struct {
	CountSuccess int
	CountFail    int
	CountSite    int
}
