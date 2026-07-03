package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv               string
	AppPort              string
	DatabaseURL          string
	SupabaseURL          string
	SupabaseJWTSecret    string
	MidtransServerKey    string
	MidtransClientKey    string
	MidtransIsProduction bool
	CORSAllowedOrigins   string
	BynaraBaseURL        string
	BynaraAPIKey         string
	BynaraModel          string
}

func Load() Config {
	_ = godotenv.Load()
	return Config{
		AppEnv:               getenv("APP_ENV", "development"),
		AppPort:              getenv("APP_PORT", "8080"),
		DatabaseURL:          os.Getenv("DATABASE_URL"),
		SupabaseURL:          os.Getenv("SUPABASE_URL"),
		SupabaseJWTSecret:    os.Getenv("SUPABASE_JWT_SECRET"),
		MidtransServerKey:    os.Getenv("MIDTRANS_SERVER_KEY"),
		MidtransClientKey:    os.Getenv("MIDTRANS_CLIENT_KEY"),
		MidtransIsProduction: os.Getenv("MIDTRANS_IS_PRODUCTION") == "true",
		CORSAllowedOrigins:   getenv("CORS_ALLOWED_ORIGINS", "http://localhost:5173"),
		BynaraBaseURL:        getenv("BYNARA_BASE_URL", "https://router.bynara.id/v1"),
		BynaraAPIKey:         os.Getenv("BYNARA_API_KEY"),
		BynaraModel:          getenv("BYNARA_MODEL", "mistral-medium-3-5"),
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
