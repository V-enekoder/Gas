package disabled

type DisabledCreateDTO struct {
	UserID     uint   `json:"user_id" binding:"required"`
	Document   string `json:"document"`
	Disability string `json:"disability"`
}

type DisabledUpdateDTO struct {
	Document   string `json:"document"`
	Disability string `json:"disability"`
}

type UserInDisabledResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DisabledResponseDTO struct {
	ID         uint                      `json:"id"`
	Document   string                    `json:"document"`
	Disability string                    `json:"disability"`
	User       UserInDisabledResponseDTO `json:"user"`
}
