package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"summershare/internal/entities"
	"summershare/internal/entities/web"
	userRepo "summershare/internal/repository/user"
	"summershare/pkg/utils"
	"time"
)

type authService struct {
	userRepository userRepo.UserRepository
}

func (a authService) Login(username string, password string) web.Response {
	user, err := a.userRepository.GetByUsername(username)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to get user",
			Data:       nil,
			Error:      "failed to get user",
		}
	}

	if err == gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "user not found",
			Data:       nil,
			Error:      "user not found",
		}
	}

	if !utils.ComparePassword(user.Password, password) {
		return web.Response{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "invalid password",
			Data:       nil,
			Error:      "invalid password",
		}
	}

	jwtToken, exp, err := utils.GenerateJWTToken(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate jwt token")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to generate jwt token",
			Data:       nil,
			Error:      "failed to generate jwt token",
		}
	}

	if user.RefreshToken == "" {
		refreshToken, err := utils.GenerateRefreshToken()
		if err != nil {
			log.Error().Err(err).Msg("failed to generate refresh token")
			return web.Response{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "failed to generate refresh token",
				Data:       nil,
				Error:      "failed to generate refresh token",
			}
		}
		user.RefreshToken = refreshToken
		user.UpdatedAt = time.Now()

		if err = a.userRepository.Update(user); err != nil {
			log.Error().Err(err).Msg("failed to update user")
			return web.Response{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "failed to update user",
				Data:       nil,
				Error:      "failed to update user",
			}
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "login successful",
		Data: fiber.Map{
			"token":         jwtToken,
			"exp":           exp,
			"refresh_token": user.RefreshToken,
		},
		Error: nil,
	}
}

func (a authService) Register(username string, email string, password string) web.Response {
	userbyusername, err := a.userRepository.CountByUsername(username)
	userbyemail, err := a.userRepository.CountByEmail(email)

	if userbyusername > 0 || userbyemail > 0 {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "username or email already exists",
			Data:       nil,
			Error:      "username or email already exists",
		}
	}

	hashedPassword, err := utils.GeneratePassword(password)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate password")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to generate password",
			Data:       nil,
			Error:      "failed to generate password",
		}
	}

	user := entities.User{
		ID:           uuid.Must(uuid.NewRandom()),
		Email:        email,
		Username:     username,
		Password:     hashedPassword,
		RefreshToken: "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = a.userRepository.Create(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to create user",
			Data:       nil,
			Error:      "failed to create user",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "user created",
		Data: fiber.Map{
			"username":   user.Username,
			"email":      user.Email,
			"id":         user.ID,
			"created_at": user.CreatedAt,
		},
		Error: nil,
	}
}

func (a authService) RefreshToken(id string, refreshToken string) web.Response {
	user, err := a.userRepository.GetByID(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to get user",
			Data:       nil,
			Error:      "failed to get user",
		}
	}

	if err == gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusNotFound,
			Message:    "user not found",
			Data:       nil,
			Error:      "user not found",
		}
	}

	if user.RefreshToken != refreshToken {
		return web.Response{
			StatusCode: fiber.StatusUnauthorized,
			Message:    "invalid refresh token",
			Data:       nil,
			Error:      "invalid refresh token",
		}
	}

	jwtToken, exp, err := utils.GenerateJWTToken(user)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate jwt token")
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "failed to generate jwt token",
			Data:       nil,
			Error:      "failed to generate jwt token",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "refresh token successful",
		Data: fiber.Map{
			"token":         jwtToken,
			"exp":           exp,
			"refresh_token": user.RefreshToken,
		},
		Error: nil,
	}
}

func NewAuthService(userRepo userRepo.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}
