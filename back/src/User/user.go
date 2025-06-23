package user

type UserCreateDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=8"`
	MunicipalityID uint   `json:"municipality_id" binding:"required"`
}

type UserUpdateDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	MunicipalityID uint   `json:"municipality_id" binding:"required"`
	Active         *bool  `json:"active" binding:"required"`
}

type MunicipalityInUserResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserResponseDTO struct {
	ID           uint                          `json:"id"`
	Name         string                        `json:"name"`
	Email        string                        `json:"email"`
	Active       bool                          `json:"active"`
	Municipality MunicipalityInUserResponseDTO `json:"municipality"`
}
