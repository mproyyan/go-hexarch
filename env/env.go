package env

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Environment struct {
	// Application Envs
	APP_PORT string `env:"APP_PORT"`

	// Database Envs
	DB_CONN     string `env:"DB_CONN"`
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     string `env:"DB_PORT"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
}

func Load(envfile string, env *Environment) {
	err := godotenv.Load(envfile)
	if err != nil {
		log.Fatal(err)
	}

	v := reflect.ValueOf(env).Elem()
	if !v.CanAddr() {
		panic("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := v.Type().Field(i)
		tag := fieldType.Tag.Get("env")
		env := os.Getenv(tag)

		if fieldVal.CanSet() {
			fieldVal.Set(reflect.ValueOf(env))
		}
	}
}
