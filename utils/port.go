package utils

import (
	"os"
	"strconv"
)

func GetPort() string {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8080 // Usa 8080 si no se puede obtener el puerto
	}

	return ":" + strconv.Itoa(port)
}
