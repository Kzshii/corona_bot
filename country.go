package main

// Country struct
type Country struct {
	ID             string    `json:"id"`
	DisplayName    string    `json:"displayName"`
	TotalConfirmed int       `json:"totalConfirmed"`
	TotalDeaths    int       `json:"totalDeaths"`
	TotalRecovered int       `json:"totalRecovered"`
	LastUpdated    string    `json:"lastUpdated"`
	Areas          []Country `json:"areas"`
}
