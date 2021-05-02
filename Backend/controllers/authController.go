package controllers

import (
	"fmt"
	"strconv"
	"time"

	"../database"

	"../models"
	"github.com/gofiber/fiber"
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
	layout := "01-01-2000 13:34"
	fecha, _ := time.Parse(layout, data["FechaNac"])
	fechaR, _ := time.Parse(layout, data["FechaRegistro"])
	//TierInt, _ := strconv.Atoi(data["Tier"])
	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          0,
		FechaNac:      fecha,
		FechaRegistro: fechaR,
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	queryString := "Insert into Cliente(Username, Password,Nombre, Apellido, Tier,Correo,Foto ) values ("
	queryString += "'" + user.Username + "' , '" + string(user.Password) + "'"
	queryString += ", '" + user.Nombre + "' , '" + user.Apellido + "'"
	queryString += ", " + strconv.Itoa(user.Tier)
	//queryString+= ","+"'"+ user.FechaNac+"')"
	queryString += ", '" + user.Correo + "' , '" + user.Foto + "')"
	res, err := database.DB.Query(queryString)
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
		FechaNac:      time.Now(),
		FechaRegistro: time.Now(),
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
		FechaNac:      time.Now(),
		FechaRegistro: time.Now(),
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	stringQuery := "Select * from Cliente where Username='"
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
	var FechaNac, FechaRegistro time.Time
	for res.Next() {
		res.Scan(&Username, &Password, &Nombre, &Apellido, &Tier, &FechaNac, &FechaRegistro, &Correo, &Foto)
		if Username == "" {
			fmt.Println("Error Escanenado Datos!")
			return err
		}
	}

	Usuario := models.User{
		Username:      Username,
		Password:      Password, //password,
		Nombre:        Nombre,
		Apellido:      Apellido,
		Tier:          Tier,
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

	stringQuery := "Select * from Cliente"
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
		FechaNac:      time.Now(),
		FechaRegistro: time.Now(),
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	stringQuery := "update cliente set password='"
	stringQuery += user.Password + "' , nombre='" + user.Nombre + "', "
	stringQuery += "apellido='" + user.Apellido + "' , Tier=" + strconv.Itoa(user.Tier)
	//stringQuery+= ", FechaNac='"+user.FechaNac+"' , FechaRegistro='"+user.FechaRegistro
	stringQuery += ", correo='" + user.Correo + "' , Foto='" + user.Foto + "'"
	stringQuery += "where username='" + user.Username + "'"

	res, err := database.DB.Query(stringQuery)

	msj := models.Mensaje{
		Mensaje: resultado,
	}
	if err != nil {
		fmt.Println("Error durante Query!")
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
		FechaNac:      time.Now(),
		FechaRegistro: time.Now(),
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
