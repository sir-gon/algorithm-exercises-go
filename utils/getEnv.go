package utils

import (
	"fmt"
	"os"
)

// GetEnv get key environment variable if exist otherwise return defalutValue
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Printf("[%s] from environment: [%s] with fallback: [%s] \n", key, value, defaultValue)

	if len(value) == 0 {
		return defaultValue
	}
	return value
}
