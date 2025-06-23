package disabled

import (
	"errors"

	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyHasDisabilityRecord = errors.New("this user already has a disability record")
)

func mapToDisabledResponseDTO(d schema.Disabled) DisabledResponseDTO {
	return DisabledResponseDTO{
		ID:         d.ID,
		Document:   d.Document,
		Disability: d.Disability,
		User: UserInDisabledResponseDTO{
			ID:    d.User.ID,
			Name:  d.User.Name,
			Email: d.User.Email,
		},
	}
}

func CreateDisabledService(dto DisabledCreateDTO) (DisabledResponseDTO, error) {
	exists, err := DisabledExistsByUserIDRepository(dto.UserID)
	if err != nil {
		return DisabledResponseDTO{}, err
	}
	if exists {
		return DisabledResponseDTO{}, ErrUserAlreadyHasDisabilityRecord
	}

	newDisabled := schema.Disabled{
		UserID:     dto.UserID,
		Document:   dto.Document,
		Disability: dto.Disability,
	}

	if err := CreateDisabledRepository(&newDisabled); err != nil {
		return DisabledResponseDTO{}, err
	}

	// Fetch again to preload the User data for the response
	createdDisabled, err := GetDisabledByIDRepository(newDisabled.ID)
	if err != nil {
		return DisabledResponseDTO{}, err
	}

	return mapToDisabledResponseDTO(createdDisabled), nil
}

func GetAllDisabledService() ([]DisabledResponseDTO, error) {
	disabledList, err := GetAllDisabledRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []DisabledResponseDTO
	for _, d := range disabledList {
		responseDTOs = append(responseDTOs, mapToDisabledResponseDTO(d))
	}
	return responseDTOs, nil
}

func GetDisabledByIDService(id uint) (DisabledResponseDTO, error) {
	disabled, err := GetDisabledByIDRepository(id)
	if err != nil {
		return DisabledResponseDTO{}, err
	}
	return mapToDisabledResponseDTO(disabled), nil
}

func UpdateDisabledService(id uint, dto DisabledUpdateDTO) (DisabledResponseDTO, error) {
	updateData := map[string]interface{}{
		"document":   dto.Document,
		"disability": dto.Disability,
	}

	if err := UpdateDisabledRepository(id, updateData); err != nil {
		return DisabledResponseDTO{}, err
	}

	return GetDisabledByIDService(id)
}

func DeleteDisabledService(id uint) error {
	err := DeleteDisabledRepository(id)
	// No need to check for other relations, just delete the record.
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
