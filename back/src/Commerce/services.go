package commerce

import (
	"errors"

	"github.com/V-enekoder/GasManager/src/schema"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyHasCommerceRecord = errors.New("this user already has a commerce record")
	ErrRifExists                    = errors.New("a commerce with this RIF already exists")
	ErrBossDocumentExists           = errors.New("a commerce with this boss document already exists")
)

func mapToCommerceResponseDTO(c schema.Commerce) CommerceResponseDTO {
	return CommerceResponseDTO{
		ID:           c.ID,
		Rif:          c.Rif,
		BossName:     c.BossName,
		BossDocument: c.BossDocument,
		User: UserInCommerceResponseDTO{
			ID:    c.User.ID,
			Name:  c.User.Name,
			Email: c.User.Email,
		},
	}
}

func CreateCommerceService(dto CommerceCreateDTO) (CommerceResponseDTO, error) {
	userExists, err := CommerceExistsByUserIDRepository(dto.UserID)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	if userExists {
		return CommerceResponseDTO{}, ErrUserAlreadyHasCommerceRecord
	}

	rifExists, err := CommerceExistsByRifRepository(dto.Rif, 0)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	if rifExists {
		return CommerceResponseDTO{}, ErrRifExists
	}

	docExists, err := CommerceExistsByBossDocumentRepository(dto.BossDocument, 0)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	if docExists {
		return CommerceResponseDTO{}, ErrBossDocumentExists
	}

	newCommerce := schema.Commerce{
		UserID:       dto.UserID,
		Rif:          dto.Rif,
		BossName:     dto.BossName,
		BossDocument: dto.BossDocument,
	}

	if err := CreateCommerceRepository(&newCommerce); err != nil {
		return CommerceResponseDTO{}, err
	}

	createdCommerce, err := GetCommerceByIDRepository(newCommerce.ID)
	if err != nil {
		return CommerceResponseDTO{}, err
	}

	return mapToCommerceResponseDTO(createdCommerce), nil
}

func GetAllCommercesService() ([]CommerceResponseDTO, error) {
	commerces, err := GetAllCommercesRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []CommerceResponseDTO
	for _, c := range commerces {
		responseDTOs = append(responseDTOs, mapToCommerceResponseDTO(c))
	}
	return responseDTOs, nil
}

func GetCommerceByIDService(id uint) (CommerceResponseDTO, error) {
	commerce, err := GetCommerceByIDRepository(id)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	return mapToCommerceResponseDTO(commerce), nil
}

func UpdateCommerceService(id uint, dto CommerceUpdateDTO) (CommerceResponseDTO, error) {
	rifExists, err := CommerceExistsByRifRepository(dto.Rif, id)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	if rifExists {
		return CommerceResponseDTO{}, ErrRifExists
	}

	docExists, err := CommerceExistsByBossDocumentRepository(dto.BossDocument, id)
	if err != nil {
		return CommerceResponseDTO{}, err
	}
	if docExists {
		return CommerceResponseDTO{}, ErrBossDocumentExists
	}

	updateData := map[string]interface{}{
		"rif":           dto.Rif,
		"boss_name":     dto.BossName,
		"boss_document": dto.BossDocument,
	}

	if err := UpdateCommerceRepository(id, updateData); err != nil {
		return CommerceResponseDTO{}, err
	}

	return GetCommerceByIDService(id)
}

func DeleteCommerceService(id uint) error {
	err := DeleteCommerceRepository(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
