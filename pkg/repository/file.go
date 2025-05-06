package repository

import (
	"bytes"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type BlobStorage struct {
}

func (storage *BlobStorage) UploadFile(file []byte, contentType string) (string, error) {
	body := bytes.NewReader(file)
	localStorageUrl, isEnvExist := os.LookupEnv("LOCAL_BLOB_STORAGE_URL")
	if !isEnvExist {
		localStorageUrl = "127.0.0.1:3001"
	}
	id := uuid.New().String()
	localStorageUrl = localStorageUrl + "/" + id
	req, err := http.NewRequest(http.MethodPut, localStorageUrl, body)
	if err != nil {
		return "", err
	}

	req.Header["Content-Type"] = []string{contentType}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	return localStorageUrl, nil
}
