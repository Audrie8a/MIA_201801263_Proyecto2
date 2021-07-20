package models

type User struct {
	Username      string
	Password      string //[]byte
	Nombre        string
	Apellido      string
	Tier          int
	FechaNac      string
	FechaRegistro string
	Correo        string
	Foto          string
}

type Mensaje struct {
	Mensaje string
}

type Users struct {
	Users []User
}

type Membresia struct {
	IdMembresia       int
	IdTemporada       int
	IdEstadoMembresia int
	IdTipoMembresia   int
}

type MembresiaProc struct {
	IdTipoMembresia   int
	IdEstadoMembresia int
	Usuario           string
}

type Deportes struct {
	Sports []Deporte
}
type Deporte struct {
	IdDeporte int
	Nombre    string
	Imagen    string
	Color     string
}
type QuinelasUusairo struct {
	Quinelas []QuinelaUsuario
}
type QuinelaUsuario struct {
	Username           string
	IdQuinela          int
	Puntaje            int
	Nombre             string
	Local              int
	Visitante          int
	NombreVisitante    string
	NombreLocal        string
	ResultadoVisitante int
	ResultadoLocal     int
	Fecha              string
}
type TemporadasDatos struct {
	Datos []TemporadaDato
}

type TemporadaDato struct {
	Username string
	Nombre   string
	Total    int
}
type Eventos struct {
	Eventoss []Evento
}
type Evento struct {
	IdEvento        string
	IdJornada       string
	Nombre          string
	Fecha           string
	NombreLocal     string
	NombreVisitante string
}
