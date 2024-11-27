package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Obtener la IP y el agente del navegador
	clientIP := r.RemoteAddr
	userAgent := r.UserAgent()

	// Crear el mensaje de log
	logMessage := fmt.Sprintf("Cliente conectado: IP=%s, Agente=%s\n", clientIP, userAgent)

	// Registrar en la consola
	log.Print(logMessage)

	// Registrar en el archivo
	logToFile(logMessage)

	// Responder al cliente
	fmt.Fprintf(w, "¡Hola Mundo!")
}

func logToFile(message string) {
	// Abrir o crear el archivo de log
	file, err := os.OpenFile("accesos.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo de log: %v\n", err)
		return
	}
	defer file.Close()

	// Escribir el mensaje en el archivo
	logger := log.New(file, "", log.LstdFlags)
	logger.Print(message)
}

func main() {
	// Configurar el servidor HTTP
	http.HandleFunc("/", handler)
	fmt.Println("Servidor ejecutándose en http://localhost:8080")

	// Iniciar el servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
