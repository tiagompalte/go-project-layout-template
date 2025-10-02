package usecase

import (
	"github.com/google/wire"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/cache"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/healthcheck"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

var ProviderSet = wire.NewSet(
	NewCreateCategoryUseCaseImpl,
	NewCreateNoteUseCaseImpl,
	NewCreateTaskUseCaseImpl,
	NewFindAllCategoryUseCaseImpl,
	NewFindAllTaskUseCaseImpl,
	NewFindOneTaskUseCaseImpl,
	NewUpdateTaskDoneUseCaseImpl,
	NewUpdateTaskUndoneUseCaseImpl,
	NewDeleteTaskUseCaseImpl,
	NewCreateUserUseCaseImpl,
	NewValidateUserPasswordUseCaseImpl,
	NewGenerateUserTokenUseCaseImpl,
	NewFindUserUUIDUseCaseImpl,
	ProviderHealthCheckUseCase,
	NewUpdateUserNameUseCaseImpl,
	NewFindAllNoteUseCaseImpl,
	NewFindByIDNoteUseCaseImpl,
)

func ProviderHealthCheckUseCase(cache cache.Cache, dataSqlManager repository.DataSqlManager, dataMongoManager repository.DataMongoManager) HealthCheckUseCase {
	return NewHealthCheckUseCaseImpl([]healthcheck.HealthCheck{
		cache, dataSqlManager, dataMongoManager,
	})
}
