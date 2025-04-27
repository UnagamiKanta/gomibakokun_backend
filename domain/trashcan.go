package domain

type Trashcan struct {
	ID              string
	Latitude        float64
	Longitude       float64
	NearestBuilding string
	TrashType       []string
}
