package user

import (
	"errors"

	"github.com/V-enekoder/GasManager/src/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrEmailExists      = errors.New("a user with this email already exists")
	ErrUserHasRelations = errors.New("cannot delete user because it has associated data (orders, commerce, etc.)")
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func mapToUserResponseDTO(u schema.User) UserResponseDTO {
	return UserResponseDTO{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Active: u.Active,
		Municipality: MunicipalityInUserResponseDTO{
			ID:   u.Municipality.ID,
			Name: u.Municipality.Name,
		},
	}
}

func CreateUserService(dto UserCreateDTO) (UserResponseDTO, error) {
	exists, err := UserExistsByEmailRepository(dto.Email)
	if err != nil {
		return UserResponseDTO{}, err
	}
	if exists {
		return UserResponseDTO{}, ErrEmailExists
	}

	hashedPassword, err := hashPassword(dto.Password)
	if err != nil {
		return UserResponseDTO{}, err
	}

	newUser := schema.User{
		Name:           dto.Name,
		Email:          dto.Email,
		Password:       hashedPassword,
		MunicipalityID: dto.MunicipalityID,
		Active:         true,
	}

	if err := CreateUserRepository(&newUser); err != nil {
		return UserResponseDTO{}, err
	}

	createdUser, err := GetUserByIDRepository(newUser.ID)
	if err != nil {
		return UserResponseDTO{}, err
	}

	return mapToUserResponseDTO(createdUser), nil
}

func GetAllUsersService() ([]UserResponseDTO, error) {
	users, err := GetAllUsersRepository()
	if err != nil {
		return nil, err
	}

	var responseDTOs []UserResponseDTO
	for _, u := range users {
		responseDTOs = append(responseDTOs, mapToUserResponseDTO(u))
	}
	return responseDTOs, nil
}

func GetUserByIDService(id uint) (UserResponseDTO, error) {
	user, err := GetUserByIDRepository(id)
	if err != nil {
		return UserResponseDTO{}, err
	}
	return mapToUserResponseDTO(user), nil
}

func UpdateUserService(id uint, dto UserUpdateDTO) (UserResponseDTO, error) {
	// Opcional: Verificar si el nuevo email ya está en uso por otro usuario
	// Para ello se necesitaría un repositorio UserExistsByEmailAndNotIDRepository(email, id)

	updateData := map[string]interface{}{
		"name":            dto.Name,
		"email":           dto.Email,
		"municipality_id": dto.MunicipalityID,
		"active":          *dto.Active,
	}

	if err := UpdateUserRepository(id, updateData); err != nil {
		return UserResponseDTO{}, err
	}

	return GetUserByIDService(id)
}

func DeleteUserService(id uint) error {
	err := DeleteUserRepository(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
