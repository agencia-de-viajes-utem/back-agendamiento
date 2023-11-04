package handlers

import (
	"backend/api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestObtenerPaquetesOfertas(t *testing.T) {
	// Prepara una solicitud falsa
	payload := ciudadPaquete{
		Ciudad: "CiudadEjemplo",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("GET", "/paquetes-ofertas", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ObtenerPaquetesOfertas)

	// Llama al handler con la solicitud falsa.
	handler.ServeHTTP(rr, req)

	// Verifica el código de estado de la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}

	// Intenta decodificar el cuerpo de la respuesta en una lista de paquetes de oferta
	var paquetesOfertas []models.PaqueteOferta
	if err := json.NewDecoder(rr.Body).Decode(&paquetesOfertas); err != nil {
		t.Errorf("Error al decodificar la respuesta JSON: %v", err)
	}
}
