package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/0xEg0x/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", api.getStudents)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) Start() error {
	// Start server
	if err := api.Echo.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
		return err // 🔥 retorna o erro para o caller
	}

	return nil // servidor foi encerrado corretamente
}

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("GET %s student", id)
	return c.String(http.StatusOK, getStud)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudente(student); err != nil {
		return c.String(http.StatusInternalServerError, "error to create student")
	}

	return c.String(http.StatusOK, "create student")
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("UPDATE %s student", id)
	return c.String(http.StatusOK, updateStud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("DELETE %s student", id)
	return c.String(http.StatusOK, deleteStud)
}
