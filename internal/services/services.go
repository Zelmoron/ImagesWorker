package services

import (
	"ImageWorkr/internal/endpoints"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type (
	Services struct {
	}
)

func New() *Services {
	return &Services{}

}

func (s *Services) ImageWorker(request endpoints.ImageUploadRequest) (string, error) {
	base64Data := request.Image
	if strings.Contains(base64Data, ",") {
		base64Data = strings.Split(base64Data, ",")[1]
	}

	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", endpoints.ErrDecode
	}
	uploadDir := "./static"

	filename := fmt.Sprintf("%d.png", time.Now().UnixNano())
	filepath := filepath.Join(uploadDir, filename)

	err = os.WriteFile(filepath, imageData, 0666)
	if err != nil {
		return "", endpoints.ErrWrite
	}

	log.Printf("Файл успешно сохранен: %s", filename)
	return filename, nil

}
