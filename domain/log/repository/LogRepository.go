package repository

import (
	"encoding/json"
	"github.com/misyuari/go-omdb/config"
	"github.com/misyuari/go-omdb/domain/log/model"
)

type LogRepository interface {
	Create(log *model.Log)
}

type LogRepositoryImpl struct{}

func NewLogRepository() LogRepository {
	return &LogRepositoryImpl{}
}

func (*LogRepositoryImpl) Create(log *model.Log) {
	SQL := "INSERT INTO log(module, request_payload, server_type, timestamp) values (?, ?, ?, ?)"
	jsonBytePayload, _ := json.Marshal(log.RequestPayload)
	_, _ = config.GlobalConfig.DBConnection.Exec(SQL, log.Module, jsonBytePayload, log.ServerType, log.Timestamp)
}
