package typecylinder

type TypeCylinderResponseDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Disponible  bool    `json:"disponible"`
}
