package commerce

type CommerceCreateDTO struct {
	UserID       uint   `json:"user_id" binding:"required"`
	Rif          string `json:"rif" binding:"required"`
	BossName     string `json:"boss_name" binding:"required"`
	BossDocument string `json:"boss_document" binding:"required"`
}

type CommerceUpdateDTO struct {
	Rif          string `json:"rif" binding:"required"`
	BossName     string `json:"boss_name" binding:"required"`
	BossDocument string `json:"boss_document" binding:"required"`
}

type UserInCommerceResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CommerceResponseDTO struct {
	ID           uint                      `json:"id"`
	Rif          string                    `json:"rif"`
	BossName     string                    `json:"boss_name"`
	BossDocument string                    `json:"boss_document"`
	User         UserInCommerceResponseDTO `json:"user"`
}
