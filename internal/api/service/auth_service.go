package service

import (
	"context"
	"errors"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	myerror "film-management-api-golang/internal/pkg/error"
	myjwt "film-management-api-golang/internal/pkg/jwt"
	"film-management-api-golang/internal/utils"
	"net/http"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
		GetMe(ctx context.Context, userId string) (dto.GetMe, error)
	}

	authService struct {
		userRepository repository.UserRepository
		db             *gorm.DB
	}
)

func NewAuth(userRepository repository.UserRepository,
	db *gorm.DB) AuthService {
	return &authService{
		userRepository: userRepository,
		db:             db,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	_, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err == nil {
		return dto.RegisterResponse{}, myerror.New("user with this email already exist", http.StatusConflict)
	}

	// hashPassword, err := utils.HashPassword(req.Password)
	// if err != nil {
	// 	return dto.RegisterResponse{}, err
	// }

	createResult, err := s.userRepository.Create(ctx, nil, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.DisplayName,
		Bio:         req.Bio,
		Role:        entity.RoleUser,
	})
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:          createResult.ID.String(),
		Username:    createResult.Username,
		Email:       createResult.Email,
		DisplayName: createResult.DisplayName,
		Bio:         createResult.Bio,
	}, nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.LoginResponse{}, myerror.New("email or password invalid", http.StatusBadRequest)
		}
		return dto.LoginResponse{}, err
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
	if !checkPassword || err != nil {
		return dto.LoginResponse{}, myerror.New("email or password invalid", http.StatusBadRequest)
	}

	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    string(user.Role),
	}, 24)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
		Role:  string(user.Role),
	}, nil
}

func (s *authService) GetMe(ctx context.Context, userId string) (dto.GetMe, error) {
	user, err := s.userRepository.GetById(ctx, nil, userId)
	if err != nil {
		return dto.GetMe{}, err
	}

	return dto.GetMe{
		PersonalInfo: dto.PersonalInfo{
			ID:          userId,
			Username:    user.Username,
			Email:       user.Email,
			DisplayName: user.DisplayName,
			Bio:         user.Bio,
			Role:        string(user.Role),
		},
	}, nil
}
