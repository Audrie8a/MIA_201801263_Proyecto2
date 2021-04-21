package controllers

import (
	"log"
	"strconv"
	"time"

	"../database"

	"../models"
	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
)

//Peticiones Ejemplo
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World !")
}

func CrearUsuario(c *fiber.Ctx) error {
	var data map[string]string
	database.Connect()
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	layout := "01-01-2000 13:34"
	fecha, _ := time.Parse(layout, data["FechaNac"])
	fechaR, _ := time.Parse(layout, data["FechaRegistro"])
	TierInt, _ := strconv.Atoi(data["Tier"])
	password, _ := bcrypt.GenerateFromPassword([]byte(data["Password"]), 14)

	user := models.User{
		Username:      data["Username"],
		Password:      password, //data["Password"],
		Nombre:        data["Nombre"],
		Apellido:      data["Apellido"],
		Tier:          TierInt,
		FechaNac:      fecha,
		FechaRegistro: fechaR,
		Correo:        data["Correo"],
		Foto:          data["Foto"],
	}
	queryString := "Insert into Cliente(Username, Nombre, Apellido, Tier, Correo, Foto ) values ("
	res, err := database.DB.Query("Insert into Cliente(Username, Nombre, Apellido, Tier, Correo, Foto ) values (?, ?,?,?,?,?)", user.Username, user.Nombre, user.Apellido, user.Tier, user.Correo, user.Foto)

	//res, err := database.DB.Query("insert into Cliente (Username, Password, Nombre, Apellido, Tier,Correo, Foto)values ('Audrie8a4', 'Contraseña', 'Audrie', 'Annelisse', 1, 'ann.audrie8a@gmail.com', 'foto')")

	if err != nil {
		return err
	}
	log.Print(res)
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"Usuario": user,
	}); err != nil {
		return err
	}
	return c.JSON(user)
}
