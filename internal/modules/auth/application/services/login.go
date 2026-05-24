package services

import (
	"context"
	"errors"
	"tradex/internal/modules/auth/adapter/inbound/rest/dto"
	"tradex/internal/modules/auth/application/inbound"
	"tradex/internal/modules/auth/application/outbound"
)

type loginService struct {
	repository outbound.Repository
	passwordHasher outbound.PasswordHasher
	tokenGenerator outbound.TokenGenerator
}

func NewLogin(
	repository outbound.Repository,
	passwordHasher outbound.PasswordHasher,
	tokenGenerator outbound.TokenGenerator,
) inbound.Login {
	return &loginService{
		repository: repository,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
	}
}

func (s *loginService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repository.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := s.passwordHasher.Compare(user.Password, request.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.tokenGenerator.Generate(user)
	if err != nil {
		return  nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}
