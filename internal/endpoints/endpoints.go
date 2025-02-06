package endpoints

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrDecode = errors.New("Ошибка декодирования")
	ErrWrite  = errors.New("Ошибка записи файла")
)

type (
	Services interface {
		ImageWorker(ImageUploadRequest) (string, error)
	}
	Endpoints struct {
		services Services
	}
	ImageUploadRequest struct {
		Image string `json:"image"`
	}
)

func New(services Services) *Endpoints {
	return &Endpoints{
		services: services,
	}
}

func (e *Endpoints) Render(c echo.Context) error {

	return c.Render(http.StatusOK, "hello", nil)

}

func (e *Endpoints) ImageWork(c echo.Context) error {
	var request ImageUploadRequest
	if err := c.Bind(&request); err != nil {
		log.Printf("Ошибка привязки запроса: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Неверный формат запроса",
		})
	}

	if request.Image == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Изображение не найдено в запросе",
		})
	}

	filename, err := e.services.ImageWorker(request)
	switch {
	case errors.Is(err, ErrDecode):
		log.Printf("Ошибка декодирования base64: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Ошибка декодирования изображения",
		})
	case errors.Is(err, ErrWrite):
		log.Printf("Ошибка сохранения файла: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Ошибка сохранения файла",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message":  "Файл успешно загружен",
		"filename": filename,
	})

}
