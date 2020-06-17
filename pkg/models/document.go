package models

// Document repretents a document to be revised
type Document struct {
	Base
	URL        string   `json:"url"`
	Comment    string   `json:"comment"`
	Rectified  bool     `json:"rectified"`
	AccessCode string   `json:"accessCode"`
	Project    *Project `json:"project"`
	ProjectID  int      `json:"projectID"`
}
