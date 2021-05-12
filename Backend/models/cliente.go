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
type Deportes struct {
	Sports []Deporte
}
type Deporte struct {
	IdDeporte int
	Nombre    string
	Imagen    string
	Color     string
}
