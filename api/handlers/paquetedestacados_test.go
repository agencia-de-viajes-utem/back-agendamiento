package handlers

import (
	"backend/api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObtenerPaquetesDestacados(t *testing.T) {
	req, err := http.NewRequest("GET", "/paquetes-destacados", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ObtenerPaquetesDestacados)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}

	var paquetesDestacados []models.PaquetesDestacados
	if err := json.NewDecoder(rr.Body).Decode(&paquetesDestacados); err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}
}
