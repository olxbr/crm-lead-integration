package models

type Lead struct {
	LeadOrigin      string `json:"leadOrigin"`
	Timestamp       string `json:"timestamp"`
	OriginLeadId    string `json:"originLead"`
	OriginListingId string `json:"originListingId"`
	ClientListingId string `json:"clientListingId"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	DDD             string `json:"ddd"`
	Phone           string `json:"phone"`
	Message         string `json:"message"`
}
