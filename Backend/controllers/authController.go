package controllers

import (
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

func CrearUsuario(c *fiber.Ctx) error {
	var resultado string
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	layout := "01-01-2000 13:34"
	fecha, _ := time.Parse(layout, data["FechaNac"])
	fechaR, _ := time.Parse(layout, data["FechaRegistro"])
	TierInt, _ := strconv.Atoi(data["Tier"])
	//password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      data["Password"], //password,
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          TierInt,
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

	if err != nil {
		resultado = "Error al realizar Query!"
		return err
	}
	resultado = "Cliente Registrado!"
	println("Cliente Registrado! ", res)

	//return c.JSON(err)
	return c.SendString(resultado)
}

func Login(c *fiber.Ctx) error {
	var resultado2 string
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
	stringQuery := "Select Nombre from Cliente where Username='"
	stringQuery += user.Username + "' and Password= '" + string(user.Password) + "'"

	println(stringQuery)
	res, err := database.DB.Query(stringQuery)

	if err != nil {
		return err
	}
	println(res)

	defer res.Close()

	var nombre string
	for res.Next() {
		res.Scan(&nombre)
		if nombre != "" {
			resultado2 = "Acceso Concedido! " + nombre
		} else {
			resultado2 = "Acceso Denegado! No hay usuarios registrados con los datos ingresados"
		}
	}
	return c.SendString(resultado2)
}
