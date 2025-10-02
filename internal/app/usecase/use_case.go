package usecase

type UseCase struct {
	CreateCategoryUseCase
	CreateNoteUseCase
	CreateTaskUseCase
	FindAllCategoryUseCase
	FindAllTaskUseCase
	FindOneTaskUseCase
	UpdateTaskDoneUseCase
	UpdateTaskUndoneUseCase
	DeleteTaskUseCase
	HealthCheckUseCase
	CreateUserUseCase
	ValidateUserPasswordUseCase
	GenerateUserTokenUseCase
	FindUserUUIDUseCase
	UpdateUserNameUseCase
	FindAllNoteUseCase
	FindByIDNoteUseCase
}
