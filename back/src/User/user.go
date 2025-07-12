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

type RegisterRequestDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=8"`
	MunicipalityID uint   `json:"municipality_id" binding:"required"`
}

// MunicipalityInUserResponseDTO es un DTO anidado para la información del municipio.
type MunicipalityInUserResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// UserResponseDTO ahora incluye el objeto del municipio en lugar del ID.
type UserResponseDTO struct {
	ID           uint                          `json:"id"`
	Name         string                        `json:"name"`
	Email        string                        `json:"email"`
	Active       bool                          `json:"active"`
	Municipality MunicipalityInUserResponseDTO `json:"municipality"` // Campo modificado
}

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponseDTO es la respuesta al hacer login.
// Incluye datos básicos del usuario, su rol y los datos específicos del rol.
type LoginResponseDTO struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Role     string      `json:"role"`                // "commerce", "council", "disabled", o "user"
	RoleData interface{} `json:"role_data,omitempty"` // Contendrá el objeto hijo
}

// DTOs para cada tipo de rol, para no exponer el modelo completo
type CommerceResponseDTO struct {
	ID           uint   `json:"id"`
	Rif          string `json:"rif"`
	BossName     string `json:"boss_name"`
	BossDocument string `json:"boss_document"`
}
type CouncilResponseDTO struct {
	ID             uint   `json:"id"`
	LeaderName     string `json:"leader_name"`
	LeaderDocument string `json:"leader_document"`
}
type DisabledResponseDTO struct {
	ID         uint   `json:"id"`
	Document   string `json:"document"`
	Disability string `json:"disability"`
}
