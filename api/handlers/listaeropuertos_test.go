package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

type Aeropuerto struct {
	ID         int    `json:"id"`
	Aeropuerto string `json:"aeropuerto"`
}

func TestListarAeropuertos(t *testing.T) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	req, err := http.NewRequest("GET", "/aeropuertos", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un ResponseRecorder (una implementación de http.ResponseWriter) para registrar la respuesta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListarAeropuertos)

	// Llamar al handler con la solicitud falsa.
	handler.ServeHTTP(rr, req)

	// Verificar el código de estado de la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}

	// Convertir la respuesta real de la API a una estructura de datos
	var aeropuertos []Aeropuerto
	if err := json.NewDecoder(rr.Body).Decode(&aeropuertos); err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}
}
