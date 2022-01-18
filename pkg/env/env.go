package env

import (
	"os"
	"strconv"
)

// PORT defines the server setup port
var PORT = port()

func port() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return 8000
	}

	return port
}
