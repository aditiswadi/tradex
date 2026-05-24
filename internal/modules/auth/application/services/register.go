package services

import (
	"context"
	"errors"

	"tradex/internal/modules/auth/adapter/inbound/rest/dto"
	"tradex/internal/modules/auth/application/inbound"
	"tradex/internal/modules/auth/application/outbound"
	"tradex/internal/modules/user/domain"
)

type registerService struct {
	repository     outbound.Repository
	passwordHasher outbound.PasswordHasher
}

func NewRegister(
	repository outbound.Repository,
	passwordHasher outbound.PasswordHasher,
) inbound.Register {
	return &registerService{
		repository:     repository,
		passwordHasher: passwordHasher,
	}
}

func (s *registerService) Register(ctx context.Context, request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	existingUser, err := s.repository.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := s.passwordHasher.Hash(request.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     "USER",
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
