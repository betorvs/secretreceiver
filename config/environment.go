package config

import "os"

var (
	// Port to be listened by application
	Port string
	// EncodingRequest string
	EncodingRequest string
)

// GetEnv func return a default value if dont find a environment variable
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	Port = GetEnv("PORT", "9090")
	EncodingRequest = GetEnv("ENCODING_REQUEST", "disabled")
}
