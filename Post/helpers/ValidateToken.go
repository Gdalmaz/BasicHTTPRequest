package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const authURL = "http://auth:9090/api/v1/user/token-control"

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type TokenResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *User  `json:"data"`
}

// CheckToken, verilen token'ı auth mikroservisine gönderir ve geçerli ise kullanıcı bilgilerini döndürür
func CheckToken(token string) (*TokenResponse, error) {
	req, err := http.NewRequest("POST", authURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unauthorized: %s", resp.Status)
	}

	// Yanıtın ham verilerini kontrol et
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	// Eğer bodyBytes boş ise bunu loglayın
	if len(bodyBytes) == 0 {
		log.Println("Response body is empty")
		return nil, fmt.Errorf("response body is empty")
	}

	// Ham yanıtı loglayın
	log.Printf("Raw response body: %s", string(bodyBytes))
	log.Println("asdlkmasdlmkasdlmkasdlmkasdlmk")
	if err = json.Unmarshal([]byte(bodyBytes), &TokenResponse{}); err != nil {

		log.Printf("Failed to unmarshal response: %v", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// JSON ayrıştırma başarılıysa, sonucu döndür
	return &TokenResponse{}, nil
}
