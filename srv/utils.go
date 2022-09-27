package srv

import (
	"os"
)

func EnvVal(key string) string {
	val := os.Getenv(key)
	return val
}
