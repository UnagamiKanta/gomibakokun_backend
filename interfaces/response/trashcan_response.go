package response

type CreateTrashcanReq struct {
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	Image           string   `json:"image"`
	TrashType       []string `json:"trashType"`
	NearestBuilding string   `json:"nearestBuilding"`
}
