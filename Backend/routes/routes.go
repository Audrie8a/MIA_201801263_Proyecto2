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

	app.Post("/LoginProc", controllers.LoginProc)

	app.Get("/Usuarios", controllers.GetUsuarios)

	app.Post("/datosUsuario", controllers.GetUsuario)

	app.Post("/updateUsuario", controllers.UpdateUsuario)

	app.Post("/membresia", controllers.UpdateUsuario)

	app.Post("/CargarDatos", controllers.CargaMasiva)

	app.Get("/Deportes", controllers.GetDeportes)

	app.Post("/updateDeporte", controllers.UpdateDeporte)

	app.Post("/QuinelasUsuario", controllers.GetQuinelasUsuario)

	app.Get("/Quinelas", controllers.GetQuinelas)

	app.Get("/Temporadas", controllers.GetDatosTemporadas)

	app.Get("/Eventos", controllers.GetEventos)

	app.Post("/ProcMembresia", controllers.ProcMembresia)
}
