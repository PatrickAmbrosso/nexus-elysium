package models

type Document struct {
	Text string
	URL  string
}

// Struct of an SOP Document
type SOPDocument struct {
	SOPTitle       string `json:"SOPTitle"`
	SOPNumber      string `json:"SOPNumber"`
	SOPVersion     string `json:"SOPVersion"`
	SOPDept        string `json:"SOPDept"`
	EffectiveDate  string `json:"EffectiveDate"`
	NextReviewDate string `json:"NextReviewDate"`
	DocumentURL    string `json:"DocumentURL"`
	SearchIndex    string `json:"SearchIndex"`
}
