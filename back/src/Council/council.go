package council

type CouncilCreateDTO struct {
	UserID         uint   `json:"user_id" binding:"required"`
	LeaderName     string `json:"leader_name" binding:"required"`
	LeaderDocument string `json:"leader_document" binding:"required"`
}

type CouncilUpdateDTO struct {
	LeaderName     string `json:"leader_name" binding:"required"`
	LeaderDocument string `json:"leader_document" binding:"required"`
}

type UserInCouncilResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CouncilResponseDTO struct {
	ID             uint                     `json:"id"`
	LeaderName     string                   `json:"leader_name"`
	LeaderDocument string                   `json:"leader_document"`
	User           UserInCouncilResponseDTO `json:"user"`
}
