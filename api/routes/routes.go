package routes

import (
	"backend/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router) {
	//allowedOrigins := []string{"http://www.lumonidy.studio", "http://localhost:3000", "http://lumonidy.studio"}
	//c := middleware.CorsMiddleware(allowedOrigins)
	//r.Use(c)

	//Ruta estandar
	r.Handle("/api/agendamiento", http.HandlerFunc(handlers.HomeHandler))

	//Ruta para los aeropuertos
	r.Handle("/api/agendamiento/aeropuertos", http.HandlerFunc(handlers.ListarAeropuertos)).Methods("GET")

	//Ruta para los paquetes
	r.Handle("/api/agendamiento/paquetes", http.HandlerFunc(handlers.ObtenerPaquetes)).Methods("POST")

	//Ruta para los paquetes por mes
	r.Handle("/api/agendamiento/paquetes-mes", http.HandlerFunc(handlers.ObtenerPaquetesMes)).Methods("POST")

	//Ruta para los paquetes destacados
	r.Handle("/api/agendamiento/paquetes-destacados", http.HandlerFunc(handlers.ObtenerPaquetesDestacados)).Methods("GET")

	//Ruta para los paquetes en oferta
	r.Handle("/api/agendamiento/paquetes-ofertas", http.HandlerFunc(handlers.ObtenerPaquetesOfertas)).Methods("GET")

	//Ruta para los aeropuertos
	r.Handle("/api/agendamiento/aeropuerto", http.HandlerFunc(handlers.ObtenerAeropuertos))
}
