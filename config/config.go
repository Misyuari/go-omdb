package config

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Env struct {
		RestHost   string `json:"REST_HOST"`
		GrpcHost   string `json:"GRPC_HOST"`
		DbUri      string `json:"DB_URI"`
		DbDriver   string `json:"DB_DRIVER"`
		OmdbApiKey string `json:"OMDB_API_KEY"`
	}
	DBConnection *sql.DB
	Validator    *validator.Validate
}

var GlobalConfig = Config{}

func LoadConfig() {
	GlobalConfig.Validator = validator.New()
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	jsonByte, errMarshal := json.Marshal(env)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}
	errUnmarshal := json.Unmarshal(jsonByte, &GlobalConfig.Env)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}
}
