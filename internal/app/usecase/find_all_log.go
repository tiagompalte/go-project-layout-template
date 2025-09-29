package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllLogUseCase interface {
	Execute(ctx context.Context, input FindAllLogInput) ([]entity.Log, error)
}

type FindAllLogInput struct {
	Limit int64
}

type FindAllLogUseCaseImpl struct {
	logRepo protocols.LogRepository
}

func NewFindAllLogUseCaseImpl(logRepo protocols.LogRepository) FindAllLogUseCase {
	return FindAllLogUseCaseImpl{
		logRepo: logRepo,
	}
}

func (uc FindAllLogUseCaseImpl) Execute(ctx context.Context, input FindAllLogInput) ([]entity.Log, error) {
	logs, err := uc.logRepo.FindAll(ctx, input.Limit)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return logs, nil
}
