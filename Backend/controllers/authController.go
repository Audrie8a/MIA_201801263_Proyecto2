package controllers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"../database"

	"../models"
	"github.com/gofiber/fiber"

	"encoding/json"
)

//Peticiones Ejemplo
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

//Registro
func CrearUsuario(c *fiber.Ctx) error {
	var resultado string
	resultado = "Error al registrar!"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	//TierInt, _ := strconv.Atoi(data["Tier"])
	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)
	dia := time.Now().Day()
	mes := 4
	anio := time.Now().Year()
	hora := time.Now().Hour()
	min := time.Now().Minute()
	fmt.Println(dia, '/', mes, '/', anio, ' ', hora, ':', min)
	fecha := strconv.Itoa(dia) + "/" + strconv.Itoa(mes) + "/" + strconv.Itoa(anio) + " " + strconv.Itoa(hora) + ":" + strconv.Itoa(min)
	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      data["FechaNac"],
		FechaRegistro: fecha,
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	queryString := "call CrearUsuario("
	queryString += "'" + user.Username + "' , '" + string(user.Password) + "'"
	queryString += ", '" + user.Nombre + "' , '" + user.Apellido + "'"
	queryString += ", " + strconv.Itoa(user.Tier)
	queryString += "," + "'" + user.FechaNac + "', '" + user.FechaRegistro + "' "
	queryString += ", '" + user.Correo + "' , '" + user.Foto + "')"
	res, err := database.DB.Exec(queryString)
	msj := models.Mensaje{
		Mensaje: resultado,
	}
	if err != nil {
		resultado = "Error al realizar Query!"
		return err
	}
	resultado = "Cliente Registrado!"
	println("Cliente Registrado! ", res)
	msj = models.Mensaje{
		Mensaje: resultado,
	}
	//return c.JSON(err)
	return c.JSON(msj)
}

//Registro
func CrearMembresia(c *fiber.Ctx) error {
	var resultado string
	resultado = "Error al registrar!"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//Temporada, _ := strconv.Atoi(data["idTemporada"])
	EstadoMembresia, _ := strconv.Atoi(data["idEstadoMembresia"])
	TipoMembresia, _ := strconv.Atoi(data["idTipoMembresia"])
	membresia := models.Membresia{
		IdMembresia:       0,
		IdTemporada:       0, //Temporada,   //Falta Corregir Temporadas
		IdEstadoMembresia: EstadoMembresia,
		IdTipoMembresia:   TipoMembresia,
	}
	queryString := "insert into Membresia (idTemporada, idEstadoMembresia,idTipoMembresia) values ("
	queryString += strconv.Itoa(membresia.IdTemporada) + "," + strconv.Itoa(membresia.IdEstadoMembresia) + "," + strconv.Itoa(membresia.IdTipoMembresia) + ")"
	res, err := database.DB.Query(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		println(resultado)
		return err
	}
	resultado = "Membresia Registrada!"
	println(resultado, res)

	queryString = "select * from (select * from Membresia order by idMembresia desc) where rownum=1"
	res2, err2 := database.DB.Query(queryString)
	if err2 != nil {
		resultado = "Error al realizar Query Obtener id Membresia!"
		println(resultado)
		return err2
	}
	print(resultado)
	defer res2.Close()

	var idMemb, idTemp, idEst, idTip int
	for res2.Next() {
		res2.Scan(&idMemb, &idEst, &idTemp, &idTip)
		if idTemp == 0 {
			fmt.Println("Error Escanenado Datos!")
			return err
		}
	}
	Membre := models.Membresia{
		IdMembresia:       idMemb,
		IdTemporada:       idTemp,
		IdEstadoMembresia: idEst,
		IdTipoMembresia:   idTip,
	}
	//return c.JSON(err)
	return c.JSON(Membre)
}

//Login
func Login(c *fiber.Ctx) error {
	var resultado2 string
	resultado2 = "Acceso Denegado"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      data["FechaNac"],
		FechaRegistro: data["FechaRegistro"],
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	stringQuery := "Select Username from Cliente where Username='"
	stringQuery += user.Username + "' and Password= '" + string(user.Password) + "'"

	println(stringQuery)
	res, err := database.DB.Query(stringQuery)
	msj := models.Mensaje{
		Mensaje: resultado2,
	}
	if err != nil {

		return err
	}
	println(res)

	defer res.Close()

	var nombre string
	for res.Next() {
		res.Scan(&nombre)
		if nombre != "" {
			resultado2 = "Acceso Concedido!"
		} else {
			resultado2 = "Acceso Denegado! No hay usuarios registrados con los datos ingresados"

		}
	}
	msj = models.Mensaje{
		Mensaje: resultado2,
	}
	println(resultado2)
	//return c.Response().Write([]byte("Hello"))
	return c.JSON(msj) //c.SendString(resultado2)
}
func UpdateDeporte(c *fiber.Ctx) error {
	var resultado2 string
	resultado2 = "Acceso Denegado"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return err
	}
	ID, _ := strconv.Atoi(data["idDeporte"])
	user := models.Deporte{
		IdDeporte: ID,
		Nombre:    data["Nombre"], //password,
		Imagen:    data["Imagen"],
		Color:     data["Color"],
	}
	stringQuery := "update Deporte set Color ='" + user.Color + "', Imagen='" + user.Imagen + "' where idDeporte= " + strconv.Itoa(user.IdDeporte)
	fmt.Println(stringQuery)
	res, err := database.DB.Query(stringQuery)

	if err != nil {

		return err
	}
	println(res)

	defer res.Close()
	resultado2 = "Deporte Actualizado!"
	msj := models.Mensaje{
		Mensaje: resultado2,
	}
	println(resultado2)
	return c.JSON(msj)
}

//Obtener Datos Usuario
func GetUsuario(c *fiber.Ctx) error {
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      data["FechaNac"],
		FechaRegistro: data["FechaRegistro"],
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	stringQuery := "Select Username, Password,Nombre, Apellido, Tier, FechaNac, FechaRegistro, Correo, Foto from Cliente where Username='"
	stringQuery += user.Username + "'"

	println(stringQuery)
	res, err := database.DB.Query(stringQuery)

	if err != nil {
		fmt.Println("Error durante Query!")
		return err
	}

	defer res.Close()

	var Username, Password, Nombre, Apellido, Correo, Foto string
	var Tier int
	var FechaNac, FechaRegistro string
	for res.Next() {
		res.Scan(&Username, &Password, &Nombre, &Apellido, &Tier, &FechaNac, &FechaRegistro, &Correo, &Foto)
		if Username == "" {
			fmt.Println("Error Escanenado Datos!")
			return err
		}
	}

	//Obtener datos Membresia
	stringQuery = "Select Membresia.idTipoMembresia from Membresia, Cliente where Cliente.Tier=Membresia.idMembresia and  Cliente.Username ='"
	stringQuery += user.Username + "'"

	println(stringQuery)
	res2, err2 := database.DB.Query(stringQuery)

	if err2 != nil {
		fmt.Println(err2)
		fmt.Println("Error durante Query!")
		return err
	}

	defer res.Close()
	var tipoTier int
	for res2.Next() {
		res2.Scan(&tipoTier)
	}
	Usuario := models.User{
		Username:      Username,
		Password:      Password, //password,
		Nombre:        Nombre,
		Apellido:      Apellido,
		Tier:          tipoTier,
		FechaNac:      FechaNac,
		FechaRegistro: FechaRegistro,
		Correo:        Correo,
		Foto:          Foto,
	}

	return c.JSON(Usuario)
}

//Ejemplo Get
func GetUsuarios(c *fiber.Ctx) error {
	database.Connect()

	stringQuery := "Select Username, Password,Nombre, Apellido, Tier, FechaNac, FechaRegistro, Correo, Foto from Cliente"
	rows, err := database.DB.Query(stringQuery)

	if err != nil {

		fmt.Print("Error running Query!")
		return err
	}

	defer rows.Close()

	result := models.Users{}

	for rows.Next() {
		Usuario := models.User{}

		err := rows.Scan(&Usuario.Username, &Usuario.Password, &Usuario.Nombre, &Usuario.Apellido, &Usuario.Tier, &Usuario.FechaNac, &Usuario.FechaRegistro, &Usuario.Correo, &Usuario.Foto)

		if err != nil {
			fmt.Println("Error recorriendo Usuarios!")
			return err
		}
		result.Users = append(result.Users, Usuario)
	}

	return c.JSON(result)
}

//Actualizar datos Usuario
func UpdateUsuario(c *fiber.Ctx) error {
	var resultado string
	resultado = "Error al Editar!"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)
	tier, _ := strconv.Atoi(data["Tier"])

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          tier,
		FechaNac:      data["FechaNac"],
		FechaRegistro: data["FechaRegistro"],
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	stringQuery := "call UpdateUsuario('" + user.Username + "', '"
	stringQuery += user.Password + "' , '" + user.Nombre + "', "
	stringQuery += "'" + user.Apellido
	stringQuery += "', '" + user.Correo + "' , '" + user.Foto + "')"

	res, err := database.DB.Exec(stringQuery)

	msj := models.Mensaje{
		Mensaje: resultado,
	}
	if err != nil {
		fmt.Println("Error durante Query!", err)
		return err
	}
	println("Datos Actualizados! ", res)
	resultado = "Datos Actualizados!"
	msj = models.Mensaje{
		Mensaje: resultado,
	}

	return c.JSON(msj)
}

func LoginProc(c *fiber.Ctx) error {
	var resultado2 string
	resultado2 = "Acceso Denegado"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      data["FechaNac"],
		FechaRegistro: data["FechaRegistro"],
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	println(user.Apellido)
	stringQuery := "call login_usuario('"
	stringQuery += user.Username + "', '" + string(user.Password) + "')"
	//stringQuery := "Set serverout on"
	println(stringQuery)
	res, err := database.DB.Query(stringQuery)
	msj := models.Mensaje{
		Mensaje: resultado2,
	}
	if err != nil {

		return err
	}
	println(res)

	//defer res.Close()

	println(resultado2)
	//return c.Response().Write([]byte("Hello"))
	return c.JSON(msj) //c.SendString(resultado2)
}

//Ejemplo Get
func GetDeportes(c *fiber.Ctx) error {
	database.Connect()

	stringQuery := "Select * from Deporte "
	rows, err := database.DB.Query(stringQuery)

	if err != nil {

		fmt.Print("Error running Query!", err)
		return err
	}

	defer rows.Close()

	result := models.Deportes{}

	for rows.Next() {
		Deporte := models.Deporte{}

		err := rows.Scan(&Deporte.IdDeporte, &Deporte.Nombre, &Deporte.Imagen, &Deporte.Color)

		if err != nil {
			fmt.Println("Error recorriendo Usuarios!", err)
			return err
		}
		result.Sports = append(result.Sports, Deporte)
	}

	return c.JSON(result)
}

func GetQuinelasUsuario(c *fiber.Ctx) error {
	var resultado2 string
	resultado2 = "Acceso Denegado"
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		fmt.Println(resultado2, err)
		return err
	}

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      data["FechaNac"],
		FechaRegistro: data["FechaRegistro"],
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}

	stringQuery := "select Quinela.UsernameCliente, Quinela.idQuinela, Quinela.Puntaje,Temporada.Nombre, Quinela.Local, Quinela.Visitante, eventodeportivo.nombrevisitante, eventodeportivo.nombrelocal, eventodeportivo.resultadovisitante, eventodeportivo.resultadolocal, eventodeportivo.fecha "
	stringQuery += " from Cliente, Quinela, Temporada, EventoDeportivo "
	stringQuery += " where EventoDeportivo.idEventoDeportivo=Quinela.idEventoDeportivo "
	stringQuery += " and Temporada.idTemporada=Quinela.idTemporadaQ "
	stringQuery += " and Cliente.Username = Quinela.UsernameCliente "
	stringQuery += " and Cliente.Username='" + user.Username + "'"
	rows, err := database.DB.Query(stringQuery)
	//fmt.Println(stringQuery)

	if err != nil {

		fmt.Print("Error running Query!", err)
		return err
	}

	defer rows.Close()

	result := models.QuinelasUusairo{}

	for rows.Next() {
		Quinela := models.QuinelaUsuario{}

		err := rows.Scan(&Quinela.Username, &Quinela.IdQuinela, &Quinela.Puntaje, &Quinela.Nombre, &Quinela.Local, &Quinela.Visitante, &Quinela.NombreVisitante, &Quinela.NombreLocal, &Quinela.ResultadoVisitante, &Quinela.ResultadoLocal, &Quinela.Fecha)

		if err != nil {
			fmt.Println("Error recorriendo Quinelas Usuario!", err)
			return err
		}
		result.Quinelas = append(result.Quinelas, Quinela)
	}

	return c.JSON(result)
}
func GetQuinelas(c *fiber.Ctx) error {

	database.Connect()

	stringQuery := "select Quinela.UsernameCliente, Quinela.idQuinela, Quinela.Puntaje,Temporada.Nombre, Quinela.Local, Quinela.Visitante, eventodeportivo.nombrevisitante, eventodeportivo.nombrelocal, eventodeportivo.resultadovisitante, eventodeportivo.resultadolocal, eventodeportivo.fecha "
	stringQuery += " from Cliente, Quinela, Temporada, EventoDeportivo "
	stringQuery += " where EventoDeportivo.idEventoDeportivo=Quinela.idEventoDeportivo "
	stringQuery += " and Temporada.idTemporada=Quinela.idTemporadaQ "
	stringQuery += " and Cliente.Username = Quinela.UsernameCliente "
	rows, err := database.DB.Query(stringQuery)
	//fmt.Println(stringQuery)

	if err != nil {

		fmt.Print("Error running Query!", err)
		return err
	}

	defer rows.Close()

	result := models.QuinelasUusairo{}

	for rows.Next() {
		Quinela := models.QuinelaUsuario{}

		err := rows.Scan(&Quinela.Username, &Quinela.IdQuinela, &Quinela.Puntaje, &Quinela.Nombre, &Quinela.Local, &Quinela.Visitante, &Quinela.NombreVisitante, &Quinela.NombreLocal, &Quinela.ResultadoVisitante, &Quinela.ResultadoLocal, &Quinela.Fecha)

		if err != nil {
			fmt.Println("Error recorriendo Quinelas Usuario!", err)
			return err
		}

		result.Quinelas = append(result.Quinelas, Quinela)
	}
	fmt.Println("Quinelas Obtenida con exito!")
	return c.JSON(result)
}

func GetDatosTemporadas(c *fiber.Ctx) error {

	database.Connect()

	stringQuery := " select Cliente.Username, Temporada.Nombre, sum(Quinela.Puntaje) as Total "
	stringQuery += " from ClienteMembresia, Temporada, Membresia, Cliente, Quinela "
	stringQuery += " where  ClienteMembresia.Membresia=Membresia.idMembresia "
	stringQuery += " and quinela.idtemporadaq=temporada.idtemporada"
	stringQuery += " and Temporada.idTemporada=Membresia.idTemporada "
	stringQuery += " and ClienteMembresia.Usuario =Cliente.Username "
	//stringQuery += " and Temporada.Nombre=(select Nombre from (select * from Temporada order by Nombre desc) where rownum=1) "
	stringQuery += " group by Cliente.Username, Temporada.Nombre"
	rows, err := database.DB.Query(stringQuery)
	//fmt.Println(stringQuery, rows)

	if err != nil {

		fmt.Print("Error running Query!", err)
		return err
	}

	defer rows.Close()

	result := models.TemporadasDatos{}

	for rows.Next() {
		Temp := models.TemporadaDato{}

		err := rows.Scan(&Temp.Username, &Temp.Nombre, &Temp.Total)

		if err != nil {
			fmt.Println("Error recorriendo Datos Temporada!", err)
			return err
		}

		result.Datos = append(result.Datos, Temp)
	}
	fmt.Println("Datos Temporada Actual Obtenida con exito!")
	return c.JSON(result)
}

func GetEventos(c *fiber.Ctx) error {

	database.Connect()

	stringQuery := " select EventoDeportivo.idEventoDeportivo, eventodeportivo.idjoranadaed, Deporte.Nombre,EventoDeportivo.Fecha, eventodeportivo.nombrelocal, eventodeportivo.nombrevisitante	"
	stringQuery += " from EventoDeportivo,Deporte "
	stringQuery += " where EventoDeportivo.idDeporte=deporte.iddeporte"
	rows, err := database.DB.Query(stringQuery)
	//fmt.Println(stringQuery, rows)

	if err != nil {

		fmt.Print("Error running Query!", err)
		return err
	}

	defer rows.Close()

	result := models.Eventos{}

	for rows.Next() {
		Temp := models.Evento{}

		err := rows.Scan(&Temp.IdEvento, &Temp.IdJornada, &Temp.Nombre, &Temp.Fecha, &Temp.NombreLocal, &Temp.NombreVisitante)

		if err != nil {
			fmt.Println("Error recorriendo Datos Evento!", err)
			return err
		}

		result.Eventoss = append(result.Eventoss, Temp)
	}
	fmt.Println("Datos Eventos Actual Obtenida con exito!")
	return c.JSON(result)
}
func ProcMembresia(c *fiber.Ctx) error {
	var resultado string
	resultado = ""
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Ocurrio un error", err)
		return err
	}

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)
	Tipo, _ := strconv.Atoi(data["IdTipoMembresia"])
	Estado, _ := strconv.Atoi(data["IdEstadoMembresia"])

	Memb := models.MembresiaProc{
		IdTipoMembresia:   Tipo,
		IdEstadoMembresia: Estado, //password,
		Usuario:           data["Username"],
	}

	queryString := "call Membresia_Usuario("
	queryString += strconv.Itoa(Memb.IdTipoMembresia) + "," + strconv.Itoa(Memb.IdEstadoMembresia) + ", '" + Memb.Usuario + "')"
	res, err := database.DB.Exec(queryString)
	fmt.Println(queryString)
	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return err
	}
	defer database.DB.Close()
	resultado = "Insert Membresia Correcto!"
	println(resultado, res)
	Commit()
	return err
}

//CARGA MASIVA --------------------------------------------------------------------------------
var colorArray []string

func CargaMasiva(c *fiber.Ctx) error {
	var resultado string
	resultado = "Error al Cargar Datos!"
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println(err)
		return err
	}
	//aux := strings.ReplaceAll(data["Info"], "[", " ")
	//aux = strings.ReplaceAll(aux, "]", " ")
	//fmt.Println(aux)
	in := []byte(data["Info"])
	var raw map[string]Usuario
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	for item, element := range raw {
		var Username string
		var nombre string
		var password string
		var correo string
		var apellido string

		Username = item
		nombre = element.Nombre
		password = element.Password
		correo = element.Username
		apellido = element.Apellido

		resultado = insertUsuario(Username, password, nombre, apellido, correo)

		for _, element2 := range element.Resultados {
			var nombreTemporada string
			var tipoTier string
			nombreTemporada = element2.Temporada
			tipoTier = element2.Tier

			resultado = insertResultados(nombreTemporada, tipoTier)
			var fechaIniTemporada string
			var fechaFinTemporada string
			var contadorTemp int = 0
			for _, element3 := range element2.Jornadas {
				var nombreJornada string = element3.Jornada
				resultado = insertJornadas(nombreJornada)
				var fechaIniJornada string
				var fechaFinJornada string
				var contador int = 0
				colorLst := [11]string{"red", "pink", "aqua", "blue", "brown", "olive", "green", "teal", "yellow", "fuchsia", "lime"}
				for _, element4 := range element3.Predicciones {
					var deporte string = element4.Deporte
					var fecha string = element4.Fecha
					var visitante string = element4.Visitante
					var local string = element4.Local
					var preVisitante int = element4.Prediccion.Visitante
					var preLocal int = element4.Prediccion.Local
					var resVisitante int = element4.Resultado.Visitante
					var resLocal int = element4.Resultado.Local
					var color string = colorLst[rand.Intn(10)]
					//if len(colorArray) == 0 {
					//	colorArray = append(colorArray, color)
					//} else {
					//	color = ComprobarColor(color, colorLst, 0)
					//}
					fechaFinJornada = fecha
					resultado = insert_EventoDeportivoDeportePrediccion(deporte, color, fecha, visitante, local, strconv.Itoa(preVisitante), strconv.Itoa(preLocal), strconv.Itoa(resVisitante), strconv.Itoa(resLocal))
					if contador == 0 {
						fechaIniJornada = fecha
					}
				}

				//Aquí va procedimiento para actualizar Datos Jornada
				resultado = updateFechasJornada(fechaIniJornada, fechaFinJornada)
				if contadorTemp == 0 {
					fechaIniTemporada = fechaIniJornada
				}

				fechaFinTemporada = fechaFinJornada

			}

			//Aquí va el proceso para actualizar Temporada
			resultado = updateFechasTemporada(fechaIniTemporada, fechaFinTemporada)

			//Aquí Termina de recorrer 1 Temporada

		}

	}
	msj := models.Mensaje{
		Mensaje: resultado,
	}
	return c.JSON(msj)
}

type Todo struct {
	Datos MyData
}
type MyData struct {
	Info []Usuario
}

//Usuario
type Usuario struct {
	Nombre     string       `json: "nombre"`
	Apellido   string       `json: "apellido"`
	Password   string       `json: "password"`
	Username   string       `json: "username"`
	Resultados []Resultados `json: "resultados"`
}
type Resultados struct {
	Temporada string    `json: "temporada"`
	Tier      string    `json: "tier"`
	Jornadas  []Jornada `json: "jornadas"`
}

type Jornada struct {
	Jornada      string          `json: "jornada"`
	Predicciones []Prediccioness `json: "predicciones"`
}

type Prediccioness struct {
	Deporte    string      `json: deporte`
	Fecha      string      `json: fecha`
	Visitante  string      `json: visitante`
	Local      string      `json: local`
	Prediccion Prediccions `json: prediccion`
	Resultado  Resultadoss `json: resultado`
}

type Prediccions struct {
	Visitante int `json: visitante`
	Local     int `json: local`
}

type Resultadoss struct {
	Visitante int `json: visitante`
	Local     int `json: local`
}

func insertUsuario(usuario string, contra string, nombre string, apellido string, correo string) string {
	var resultado string
	resultado = ""
	database.Connect()

	queryString := "call Insert_Cliente("
	queryString += "'" + nombre + "', '" + usuario + "', '" + contra + "', '" + apellido + "', '" + correo + "' )"

	res, err := database.DB.Exec(queryString)
	fmt.Println(queryString)
	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	defer database.DB.Close()
	resultado = "Insert Usuario Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}
func insertResultados(nombreTemp string, tipoMemb string) string {
	var resultado string
	resultado = ""
	database.Connect()
	tipoMemb = strings.ToLower(tipoMemb)
	if tipoMemb == "gold" {
		tipoMemb = "Gold"
	} else if tipoMemb == "silver" {
		tipoMemb = "Silver"
	} else if tipoMemb == "bronze" {
		tipoMemb = "Bronze"
	}
	queryString := "call Insert_Resultados("
	queryString += "'" + nombreTemp + "', '" + tipoMemb + "')"

	res, err := database.DB.Exec(queryString)
	defer database.DB.Close()
	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	resultado = "Insert Resultados Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}
func insertJornadas(nombreJornada string) string {
	var resultado string
	resultado = ""
	database.Connect()

	queryString := "call Insert_Jornadas("
	queryString += "'" + nombreJornada + "')"

	res, err := database.DB.Exec(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	defer database.DB.Close()
	resultado = "Insert Jornadas Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}

func updateFechasTemporada(fechaIni string, fechaFin string) string {
	var resultado string
	resultado = ""
	database.Connect()
	//queryString := "ALTER SESSION SET nls_date_format = 'DD/MM/YYYY HH24:MI'"
	//res2, err2 := database.DB.Query(queryString)
	//
	//if err2 != nil {
	//	resultado = "Error al realizar Query!"
	//	fmt.Println(err2)
	//	return resultado
	//}
	//fmt.Println(res2)
	queryString := "call Update_Temporada("
	queryString += "'" + fechaIni + "', '" + fechaFin + "' )"

	res, err := database.DB.Exec(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	defer database.DB.Close()
	resultado = "Insert Fechas Temporada Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}

func updateFechasJornada(fechaIni string, fechaFin string) string {
	var resultado string
	resultado = ""
	database.Connect()
	//queryString := "ALTER SESSION SET nls_date_format = 'DD/MM/YYYY HH24:MI'"
	//res2, err2 := database.DB.Query(queryString)
	//
	//if err2 != nil {
	//	resultado = "Error al realizar Query!"
	//	fmt.Println(err2)
	//	return resultado
	//}
	//fmt.Println(res2)
	queryString := "call Update_Jornada("
	queryString += "'" + fechaIni + "', '" + fechaFin + "' )"

	res, err := database.DB.Exec(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	defer database.DB.Close()
	resultado = "Insert Fechas Jornada Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}
func insert_EventoDeportivoDeportePrediccion(deporte string, color string, fecha string, visitante string, local string, preVisitante string, preLocal string, resVisitante string, resLocal string) string {
	var resultado string
	resultado = ""
	database.Connect()
	//queryString := "ALTER SESSION SET nls_date_format = 'DD/MM/YYYY HH24:MI'"
	//res2, err2 := database.DB.Query(queryString)
	//
	//if err2 != nil {
	//	resultado = "Error al realizar Query!"
	//	fmt.Println(err2)
	//	return resultado
	//}
	//fmt.Println(res2)
	queryString := "call Insert_EventoDeportivoDeporte("
	queryString += "'" + deporte + "', '" + color + "', '" + fecha + "', '" + visitante + "', '" + local + "', " + preVisitante + "," + preLocal + "," + resVisitante + "," + resLocal + " )"

	res, err := database.DB.Exec(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}
	defer database.DB.Close()
	resultado = "Insert EventoDeportivo Correcto!"
	println(resultado, res)
	Commit()
	return resultado
}

func Commit() string {
	var resultado string
	database.Connect()
	resultado = "Guardado!"
	queryString := "commit"
	res, err := database.DB.Exec(queryString)

	if err != nil {
		resultado = "Error al realizar Query!"
		fmt.Println(err)
		return resultado
	}

	defer database.DB.Close()
	resultado = "Progreso Guardado!"
	println(resultado, res)

	return resultado
}

//
//func ComprobarColor(color string, colorLst [11]string, contador int) string {
//
//	if contador != 10 {
//		for _, clr := range colorArray {
//			if color == clr {
//				color = colorLst[contador]
//				color = ComprobarColor(color, colorLst, contador)
//			}
//		}
//		colorArray = append(colorArray, color)
//	}
//	contador++
//	return color
//}
