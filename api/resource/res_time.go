package resource

// ResGetTime returns current timestamp and date.
type ResGetTime struct {
	Timestamp string `json:"timestamp"`
	Date      string `json:"date"`
}
