package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAgregarVista(t *testing.T) {
	// Prepara una solicitud POST falsa
	payload := SolicitudVista{
		Fk_fechaPaquete: 1,
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/agregar-vista", bytes.NewReader(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AgregarVista)

	// Llama al handler con la solicitud falsa.
	handler.ServeHTTP(rr, req)

	// Verifica el código de estado de la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}
}
