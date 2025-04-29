package domain

type Trashcan struct {
	ID              string
	Latitude        float64
	Longitude       float64
	Image           string
	TrashType       []string
	NearestBuilding string
}
