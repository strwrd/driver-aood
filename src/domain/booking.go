package domain

type Booking struct {
	ID          string `json:"ID"`
	DriverID    string `json:"driverId"`
	IsCompleted bool   `json:"status"`
	Distance    int    `json:"distance"`
}
