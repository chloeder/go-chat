package env

import (
	"log"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := ".env"
	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		log.Println("Error reading .env file:", err)
		panic(err)
	}
}
