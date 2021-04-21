package models

import "time"

type User struct {
	Username      string
	Password      []byte
	Nombre        string
	Apellido      string
	Tier          int
	FechaNac      time.Time
	FechaRegistro time.Time
	Correo        string
	Foto          string
}