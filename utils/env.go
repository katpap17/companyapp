package utils

import "os"

func GetEnv(name string, fallback string) string {
	res := os.Getenv(name)
	if res == "" {
		return fallback
	}
	return res
}
