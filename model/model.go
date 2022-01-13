package model

type IpData struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type JSONError struct {
	Err string `json:"error"`
}
