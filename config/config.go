package config

import "os"

func ReadEnvString(config string) string {
	return os.Getenv(config)
}
