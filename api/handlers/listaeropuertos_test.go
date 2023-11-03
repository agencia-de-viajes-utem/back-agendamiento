package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func TestListarAeropuertos(t *testing.T) {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatal("Error al cargar el archivo .env:", err)
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

	// Comparar la estructura de datos real con la estructura de datos esperada
	expectedAeropuertos := []Aeropuerto{
		{ID: 1, Aeropuerto: "Aeropuerto Internacional Comodoro Arturo Merino Benítez, Santiago, Chile"},
		{ID: 2, Aeropuerto: "Aeropuerto Internacional de Pudahuel, Santiago, Chile"},
		{ID: 3, Aeropuerto: "Aeropuerto Carriel Sur, Concepción, Chile"},
		{ID: 4, Aeropuerto: "Aeropuerto Internacional Ministro Pistarini, Buenos Aires, Argentina"},
		{ID: 5, Aeropuerto: "Aeropuerto Jorge Newbery, Buenos Aires, Argentina"},
		{ID: 6, Aeropuerto: "Aeropuerto Internacional Cataratas del Iguazú, Puerto Iguazú, Argentina"},
		{ID: 7, Aeropuerto: "Aeropuerto Internacional Jorge Chávez, Lima, Perú"},
		{ID: 8, Aeropuerto: "Aeropuerto Internacional Alejandro Velasco Astete, Cusco, Perú"},
		{ID: 9, Aeropuerto: "Aeropuerto Internacional Rodríguez Ballón, Arequipa, Perú"},
		{ID: 10, Aeropuerto: "Aeropuerto Internacional Viru Viru, Santa Cruz de la Sierra, Bolivia"},
		{ID: 11, Aeropuerto: "Aeropuerto Internacional El Alto, La Paz, Bolivia"},
		{ID: 12, Aeropuerto: "Aeropuerto Internacional Alcantarí, Cobija, Bolivia"},
		{ID: 13, Aeropuerto: "Aeropuerto Internacional de São Paulo-Guarulhos, São Paulo, Brasil"},
		{ID: 14, Aeropuerto: "Aeropuerto Santos Dumont, Río de Janeiro, Brasil"},
		{ID: 15, Aeropuerto: "Aeropuerto Internacional de Brasília, Brasília, Brasil"},
		{ID: 16, Aeropuerto: "Aeropuerto Internacional de Carrasco, Montevideo, Uruguay"},
		{ID: 17, Aeropuerto: "Aeropuerto de Punta del Este, Punta del Este, Uruguay"},
		{ID: 18, Aeropuerto: "Aeropuerto de Laguna del Sauce, Punta del Este, Uruguay"},
		{ID: 19, Aeropuerto: "Aeropuerto Internacional Silvio Pettirossi, Asunción, Paraguay"},
		{ID: 20, Aeropuerto: "Aeropuerto Internacional Guarani, Ciudad del Este, Paraguay"},
		{ID: 21, Aeropuerto: "Aeropuerto Internacional Juan de Ayolas, Ayolas, Paraguay"},
		{ID: 22, Aeropuerto: "Aeropuerto Internacional Simón Bolívar, Maiquetía, Venezuela"},
		{ID: 23, Aeropuerto: "Aeropuerto Internacional La Chinita, Maracaibo, Venezuela"},
		{ID: 24, Aeropuerto: "Aeropuerto Internacional Arturo Michelena, Valencia, Venezuela"},
		{ID: 25, Aeropuerto: "Aeropuerto Internacional El Dorado, Bogotá, Colombia"},
		{ID: 26, Aeropuerto: "Aeropuerto Internacional José María Córdova, Medellín, Colombia"},
		{ID: 27, Aeropuerto: "Aeropuerto Internacional Rafael Núñez, Cartagena, Colombia"},
		{ID: 28, Aeropuerto: "wakanda, Santiago, Chile"},
		// Agrega más aeropuertos esperados según la respuesta real de la API
	}

	if !compareAeropuertos(expectedAeropuertos, aeropuertos) {
		t.Errorf("Handler devolvió una lista de aeropuertos incorrecta")
	}
}

type Aeropuerto struct {
	ID         int    `json:"id"`
	Aeropuerto string `json:"aeropuerto"`
}

func compareAeropuertos(expected []Aeropuerto, actual []Aeropuerto) bool {
	if len(expected) != len(actual) {
		return false
	}

	// Comparar cada aeropuerto en las listas
	for i := range expected {
		if expected[i].ID != actual[i].ID || expected[i].Aeropuerto != actual[i].Aeropuerto {
			return false
		}
	}

	return true
}
