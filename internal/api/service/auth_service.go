package service

import (
	"context"
	"errors"
	"fmt"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	mailer "frs-planning-backend/internal/pkg/email"
	myerror "frs-planning-backend/internal/pkg/error"
	myjwt "frs-planning-backend/internal/pkg/jwt"
	"frs-planning-backend/internal/utils"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
)

type (
	AuthService interface {
		Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
		Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
		Verify(ctx context.Context, token string) error
		// ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) error
		// ChangePassword(ctx context.Context, req dto.ChangePasswordRequest) error
		GetMe(ctx context.Context, userId string) (dto.GetMe, error)
	}

	authService struct {
		userRepository repository.UserRepository
		mailService    mailer.Mailer
		db             *gorm.DB
	}
)

func NewAuthService(userRepository repository.UserRepository,
	mailService mailer.Mailer,
	db *gorm.DB) AuthService {
	return &authService{
		userRepository: userRepository,
		mailService:    mailService,
		db:             db,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	_, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err == nil {
		return dto.RegisterResponse{}, myerror.New("user with this email already exist", http.StatusConflict)
	}

	createResult, err := s.userRepository.Create(ctx, nil, entity.User{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		IsVerified:  true,
	})
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": createResult.ID.String(),
		"email":   createResult.Email,
	}, 24*time.Hour)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	token = fmt.Sprintf("%s/token=%s", os.Getenv("APP_URL"), token)
	if err := s.mailService.MakeMail("./internal/pkg/email/template/verification_email.html", map[string]any{
		"Username": createResult.Username,
		"Verify":   token,
	}).Send(createResult.Email, "Verify Your Account").Error; err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{
		ID:          createResult.ID.String(),
		Username:    createResult.Username,
		Email:       createResult.Email,
		PhoneNumber: createResult.PhoneNumber,
	}, nil
}

func (s *authService) Verify(ctx context.Context, token string) error {
	payloadToken, err := myjwt.GetPayloadInsideToken(token)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetByEmail(ctx, nil, payloadToken["email"])
	if err != nil {
		return err
	}

	user.IsVerified = true

	_, err = s.userRepository.Update(ctx, nil, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.LoginResponse{}, myerror.New("email or password invalid", http.StatusBadRequest)
		}
		return dto.LoginResponse{}, err
	}

	if !user.IsVerified {
		return dto.LoginResponse{}, myerror.New("user is not verify", http.StatusUnauthorized)
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
	if !checkPassword || err != nil {
		return dto.LoginResponse{}, myerror.New("email or password invalid", http.StatusBadRequest)
	}

	token, err := myjwt.GenerateToken(map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
	}, 24)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
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
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
