package dto

type CreateArtworkRequest struct {
	Name           string
	Description    string
	Image          string
	Starting_price float64
	Live_period    string
	Owner_id       string
	Category       string
}