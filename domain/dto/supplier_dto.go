package dto

type CreateStoreRequest struct {
	FullAddress string `json:"full_address"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type CreateStoreResponse struct {
	ID          uint   `json:"id"`
	FullAddress string `json:"full_address"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
