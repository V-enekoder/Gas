package user

import (
	"errors"

	"github.com/V-enekoder/GasManager/src/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrEmailExists        = errors.New("a user with this email already exists")
	ErrUserHasRelations   = errors.New("cannot delete user because it has associated data (orders, commerce, etc.)")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserInactive       = errors.New("user account is inactive")
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

func RegisterUserService(dto RegisterRequestDTO) (UserResponseDTO, error) {
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

	// --- CAMBIO CLAVE ---
	// Después de crear el usuario, `newUser` tiene el ID asignado.
	// Ahora lo volvemos a buscar en la BD usando la función que carga las relaciones.
	createdUserWithRelations, err := GetUserByIDWithRelationsRepository(newUser.ID)
	if err != nil {
		// En un caso real, aquí se debería manejar la posibilidad de que el usuario no se encuentre,
		// aunque sería muy raro justo después de crearlo.
		return UserResponseDTO{}, err
	}

	// Mapeamos el usuario (que ahora tiene los datos del municipio) al DTO de respuesta.
	return mapToUserResponseDTO(createdUserWithRelations), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// LoginUserService maneja la lógica de negocio para la autenticación.
func LoginUserService(dto LoginRequestDTO) (LoginResponseDTO, error) {
	// 1. Buscar al usuario por email con todas sus relaciones de rol.
	user, err := GetUserByEmailWithRelationsRepository(dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return LoginResponseDTO{}, ErrInvalidCredentials
		}
		return LoginResponseDTO{}, err // Otro error de base de datos
	}

	// 2. Verificar si la contraseña es correcta.
	if !checkPasswordHash(dto.Password, user.Password) {
		return LoginResponseDTO{}, ErrInvalidCredentials
	}

	// 3. Verificar si el usuario está activo.
	if !user.Active {
		return LoginResponseDTO{}, ErrUserInactive
	}

	// 4. Construir la respuesta.
	response := LoginResponseDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	// 5. Determinar el rol y adjuntar los datos del objeto hijo.
	if user.Commerce != nil {
		response.Role = "commerce"
		response.RoleData = CommerceResponseDTO{
			ID:           user.Commerce.ID,
			Rif:          user.Commerce.Rif,
			BossName:     user.Commerce.BossName,
			BossDocument: user.Commerce.BossDocument,
		}
	} else if user.Council != nil {
		response.Role = "council"
		response.RoleData = CouncilResponseDTO{
			ID:             user.Council.ID,
			LeaderName:     user.Council.LeaderName,
			LeaderDocument: user.Council.LeaderDocument,
		}
	} else if user.Disabled != nil {
		response.Role = "disabled"
		response.RoleData = DisabledResponseDTO{
			ID:         user.Disabled.ID,
			Document:   user.Disabled.Document,
			Disability: user.Disabled.Disability,
		}
	} else {
		// Es un usuario estándar sin un rol específico.
		response.Role = "user"
		response.RoleData = nil
	}

	return response, nil
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
