package config

import (
	"os"
	"reflect"
	"strconv"
)

func InitConfig() Config {
	return Config{
		DBConnectionString:        getEnv("DATABASE_URL", "postgresql://postgres@localhost:5432/gultix?sslmode=disable"),
		Port:                      getEnv("PORT", "8080"),
		ClientId:                  getEnv("CLIENT_ID", ""),
		ClientSecret:              getEnv("CLIENT_SECRET", ""),
		RedirectUrl:               getEnv("REDIRECT_URL", ""),
		JwtSecret:                 getEnv("JWT_SECRET", "supersecret"),
		JwtExpiredTime:            getEnv("JWT_EXPIRED_TIME", "2h"),
		JwtIssuer:                 getEnv("JWT_ISSUER", "Gultix.Backend.Dev"),
		MidtransServerKey:         getEnv("MIDTRANS_SERVER_KEY", ""),
		Env:                       getEnv("ENV", "dev"),
		ApiKey:                    getEnv("API_KEY", "supersecret"),
		BaseUrl:                   getEnv("BASE_URL", "http://localhost:8080"),
		DefaultBlasterAccessToken: getEnv("DEFAULT_BLASTER_ACCESS_TOKEN", ""),
		DefaultRefreshToken:       getEnv("DEFAULT_REFRESH_TOKEN", ""),
		DefaultTokenType:          getEnv("DEFAULT_TOKEN_TYPE", "Bearer"),
		DefaultExpiredTime:        getEnv("DEFAULT_EXPIRED_TIME", "2024-10-04T09:07:23.754532777Z"),
	}
}

func getEnv[T any](key string, defaultValue T) T {
	value, exists := os.LookupEnv(key)
	if !exists {
		// If defaultValue is Nil then it's required
		if reflect.ValueOf(defaultValue).IsZero() {
			panic("Environment variable " + key + " is required")
		}
		return defaultValue
	}

	// Parsing types to T
	var returnValue T
	reflectType := reflect.TypeOf(defaultValue).Kind()

	switch reflectType {
	case reflect.String:
		returnValue = reflect.ValueOf(value).Interface().(T)
	case reflect.Int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			returnValue = defaultValue
		}
		returnValue = reflect.ValueOf(intValue).Interface().(T)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			returnValue = defaultValue
		}
		returnValue = reflect.ValueOf(boolValue).Interface().(T)
	}

	return returnValue
}
