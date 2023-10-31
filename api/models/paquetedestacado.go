package models

type PaquetesDestacados struct {
	Vistas              int                  `json:"total_vistas"`
	ID                  int                  `json:"id"`
	Nombre              string               `json:"nombre"`
	IdOrigen            int                  `json:"id_ciudad_origen"`
	IdDestino           int                  `json:"id_ciudad_destino"`
	Descripcion         string               `json:"descripcion"`
	Detalles            string               `json:"detalles"`
	PrecioVuelo         float64              `json:"precio_vuelo"`
	TotalPersonas       int                  `json:"total_personas"`
	FechaInit           string               `json:"fechainit"`
	FechaFin            string               `json:"fechafin"`
	NombreCiudadOrigen  string               `json:"nombre_ciudad_origen"`
	NombreCiudadDestino string               `json:"nombre_ciudad_destino"`
	PrecioOfertaVuelo   float64              `json:"oferta_vuelo"`
	PrecioNoche         float64              `json:"precio_noche"`
	Imagenes            string               `json:"imagenes"`
	InfoPaquete         PaqueteInfoAdicional `json:"info_paquete"`
}

type PaqueteInfoAdicionalHotel struct {
	NombreOpcionHotel     string    `json:"nombre_opcion_hotel"`
	DescripcionHabitacion string    `json:"descripcion_habitacion"`
	ServiciosHabitacion   string    `json:"servicios_habitacion"`
	HotelInfo             HotelInfo `json:"hotel_info"`
	RowNum                int       `json:"row_num"`
}

type InfoHotel struct {
	NombreHotel       string  `json:"nombre_hotel"`
	DireccionHotel    string  `json:"direccion_hotel"`
	ValoracionHotel   float64 `json:"valoracion_hotel"`
	DescripcionHotel  string  `json:"descripcion_hotel"`
	ServiciosHotel    string  `json:"servicios_hotel"`
	TelefonoHotel     string  `json:"telefono_hotel"`
	CorreoElectronico string  `json:"correo_electronico_hotel"`
	SitioWeb          string  `json:"sitio_web_hotel"`
}
