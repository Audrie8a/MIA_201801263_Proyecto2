package models

import "time"

type User struct {
	Username      string
	Password      string //[]byte
	Nombre        string
	Apellido      string
	Tier          int
	FechaNac      time.Time
	FechaRegistro time.Time
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
