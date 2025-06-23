package council

import (
	"errors"

	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyHasCouncilRecord = errors.New("this user already has a council record")
	ErrLeaderDocumentExists        = errors.New("a council with this leader document already exists")
)

func mapToCouncilResponseDTO(c schema.Council) CouncilResponseDTO {
	return CouncilResponseDTO{
		ID:             c.ID,
		LeaderName:     c.LeaderName,
		LeaderDocument: c.LeaderDocument,
		User: UserInCouncilResponseDTO{
			ID:    c.User.ID,
			Name:  c.User.Name,
			Email: c.User.Email,
		},
	}
}

func CreateCouncilService(dto CouncilCreateDTO) (CouncilResponseDTO, error) {
	userExists, err := CouncilExistsByUserIDRepository(dto.UserID)
	if err != nil {
		return CouncilResponseDTO{}, err
	}
	if userExists {
		return CouncilResponseDTO{}, ErrUserAlreadyHasCouncilRecord
	}

	docExists, err := CouncilExistsByLeaderDocumentRepository(dto.LeaderDocument, 0)
	if err != nil {
		return CouncilResponseDTO{}, err
	}
	if docExists {
		return CouncilResponseDTO{}, ErrLeaderDocumentExists
	}

	newCouncil := schema.Council{
		UserID:         dto.UserID,
		LeaderName:     dto.LeaderName,
		LeaderDocument: dto.LeaderDocument,
	}

	if err := CreateCouncilRepository(&newCouncil); err != nil {
		return CouncilResponseDTO{}, err
	}

	createdCouncil, err := GetCouncilByIDRepository(newCouncil.ID)
	if err != nil {
		return CouncilResponseDTO{}, err
	}

	return mapToCouncilResponseDTO(createdCouncil), nil
}

func GetAllCouncilsService() ([]CouncilResponseDTO, error) {
	councils, err := GetAllCouncilsRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []CouncilResponseDTO
	for _, c := range councils {
		responseDTOs = append(responseDTOs, mapToCouncilResponseDTO(c))
	}
	return responseDTOs, nil
}

func GetCouncilByIDService(id uint) (CouncilResponseDTO, error) {
	council, err := GetCouncilByIDRepository(id)
	if err != nil {
		return CouncilResponseDTO{}, err
	}
	return mapToCouncilResponseDTO(council), nil
}

func UpdateCouncilService(id uint, dto CouncilUpdateDTO) (CouncilResponseDTO, error) {
	docExists, err := CouncilExistsByLeaderDocumentRepository(dto.LeaderDocument, id)
	if err != nil {
		return CouncilResponseDTO{}, err
	}
	if docExists {
		return CouncilResponseDTO{}, ErrLeaderDocumentExists
	}

	updateData := map[string]interface{}{
		"leader_name":     dto.LeaderName,
		"leader_document": dto.LeaderDocument,
	}

	if err := UpdateCouncilRepository(id, updateData); err != nil {
		return CouncilResponseDTO{}, err
	}

	return GetCouncilByIDService(id)
}

func DeleteCouncilService(id uint) error {
	err := DeleteCouncilRepository(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
