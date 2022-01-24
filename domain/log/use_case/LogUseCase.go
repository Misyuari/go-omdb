package use_case

import (
	"github.com/misyuari/go-omdb/domain/log/model"
	"github.com/misyuari/go-omdb/domain/log/repository"
	"time"
)

type LogUseCase interface {
	Create(serverType string, module string, payload interface{})
}

type LogUseCaseImpl struct {
	repository repository.LogRepository
}

func NewLogUseCase(repository repository.LogRepository) LogUseCase {
	return &LogUseCaseImpl{repository: repository}
}

func (useCase *LogUseCaseImpl) Create(serverType string, module string, payload interface{}) {
	useCase.repository.Create(&model.Log{
		Module:         module,
		RequestPayload: payload,
		ServerType:     serverType,
		Timestamp:      time.Now(),
	})
}
