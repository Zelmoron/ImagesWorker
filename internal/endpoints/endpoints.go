package endpoints

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Services interface {
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

	base64Data := request.Image
	if strings.Contains(base64Data, ",") {
		base64Data = strings.Split(base64Data, ",")[1]
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		log.Printf("Ошибка декодирования base64: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Ошибка декодирования изображения",
		})
	}
	uploadDir := "./static"

	filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
	filepath := filepath.Join(uploadDir, filename)

	err = os.WriteFile(filepath, imageData, 0666)
	if err != nil {
		log.Printf("Ошибка сохранения файла: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Ошибка сохранения файла",
		})
	}

	log.Printf("Файл успешно сохранен: %s", filename)

	return c.JSON(http.StatusOK, map[string]string{
		"message":  "Файл успешно загружен",
		"filename": filename,
	})
}
