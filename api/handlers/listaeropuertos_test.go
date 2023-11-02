package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListarAeropuertos(t *testing.T) {
	req, err := http.NewRequest("GET", "/aeropuertos", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Crear un ResponseRecorder (un implementación de http.ResponseWriter) para registrar la respuesta
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListarAeropuertos)

	// Llamar al handler con la solicitud falsa.
	handler.ServeHTTP(rr, req)

	// Verificar el código de estado de la respuesta
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler devolvió un código de estado incorrecto: esperado %v pero obtuvo %v", http.StatusOK, status)
	}

	// Verificar el tipo de contenido de la respuesta
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler devolvió un tipo de contenido incorrecto: esperado %s pero obtuvo %s", expectedContentType, contentType)
	}

	// Verificar el cuerpo de la respuesta (solo para ilustrar, puedes hacer pruebas más detalladas según tu necesidad)
	expectedResponseBody := `[{"ID": 1, "Aeropuerto": "Nombre del aeropuerto, Nombre de la ciudad, Nombre del país"}]`
	if rr.Body.String() != expectedResponseBody {
		t.Errorf("Handler devolvió un cuerpo de respuesta incorrecto: esperado %s pero obtuvo %s", expectedResponseBody, rr.Body.String())
	}
}
