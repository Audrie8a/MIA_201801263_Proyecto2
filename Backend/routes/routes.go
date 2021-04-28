package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

//Peticiones Ejemplo
func Setup(app *fiber.App) {

	//Plantilla Peticiones
	app.Get("/", controllers.Hello)

	app.Post("/Registro", controllers.CrearUsuario)

	app.Post("/Login", controllers.Login)

	app.Get("/Usuarios", controllers.GetUsuarios)

	app.Post("/datosUsuario", controllers.GetUsuario)
}
