package cargamasiva

var path = ""

type Todo struct {
	Datos MyData
}
type MyData struct {
	Info []Usuario
}

type Datos map[string]MyData

//Usuario
type Usuario struct {
	Nombre   string
	Apellido string
	Password string
	Username string
	Res      []Resultados
}
type Resultados struct {
	Temporada string
	Tier      string
	Jornadas  []Jornada
}

type Jornada struct {
	Jornada string
	Evento  []Predicciones
}

type Predicciones struct {
	Deporte   string
	Fecha     string
	Visitante string
	Local     string
	Pred      Prediccion
	Res       Resultado
}

type Prediccion struct {
	Visitante int
	Local     int
}

type Resultado struct {
	Visitante int
	Local     int
}
