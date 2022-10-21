package models

type Event struct {
	// IDS and timestamps
	UUID      string `json:"uuid" bson:"_id"`
	ID        string `json:"id"`
	Millisecs int64  `json:"pubMillis"`

	// Locations
	Country  string `json:"country"`
	NearBy   string `json:"nearby"`
	Location string `json:"location"`
	Street   string `json:"street"`

	// Fiability
	Reliability  int `json:"reliability"`
	ReportRating int `json:"reportRating"`

	// Subjective info
	AdditionalInfo string `json:"additionalInfo"`
	Comments       string `json:"comments"`

	// Types
	Type    string `json:"type"`
	Subtype string `json:"subType"`
}
