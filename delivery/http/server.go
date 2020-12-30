package http

import (
	"email-messages/app"
	"email-messages/delivery/commands"
	"email-messages/domain"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponsePercentResult struct {
	PercentResult float32 `json:"percentResult"`
}

type server struct {
	*app.App
}

func NewHandler(e *echo.Echo, app *app.App) {
	handler := &server{
		app,
	}
	e.POST("/api/message", handler.AddMessage)
	e.POST("/api/send", handler.SendMessages)
	e.GET("/api/messages/:email", handler.GetMessagesByEmail)
}

func (s *server) AddMessage(c echo.Context) error {
	var cmd commands.AddMessageCmd

	err := c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.MessagesService.AddMessage(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func (s *server) GetMessagesByEmail(c echo.Context) error {
	var (
		err   error
		email = c.Param("email")
	)

	resp, err := s.MessagesService.GetMessages(&email)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) SendMessages(c echo.Context) error {
	var (
		err error
		cmd commands.SendMessagesCmd
	)

	err = c.Bind(&cmd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ResponseMessage{Message: err.Error()})
	}

	if err = c.Validate(cmd); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = s.MessagesService.SendMessages(&cmd)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseMessage{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseMessage{Message: "ok"})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error(err)
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
