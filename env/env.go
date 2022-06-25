package env

import "os"

func GetEnvWithFallback(name string, fallback string) string {
	value, present := os.LookupEnv(name)
	if present {
		return value
	} else {
		return fallback
	}
}
