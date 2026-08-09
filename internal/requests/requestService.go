package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/ThinCrunch/crunchy-streams/internal/config"
)

type Service struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
}

func NewService(secrets config.TwitchSecrets) *Service {
	return &Service{
		ClientId:     secrets.ClientId,
		ClientSecret: secrets.ClientSecret,
		AccessToken:  secrets.AccessToken,
	}
}

func (s *Service) GetNewToken() {
	requestBody := url.Values{
		"client_id":     {s.ClientId},
		"client_secret": {s.ClientSecret},
		"grant_type":    {"client_credentials"},
	}
	resp, err := http.PostForm("https://id.twitch.tv/oauth2/token", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var tokenResponse TwitchTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("token: %s\n", tokenResponse.AccessToken)
	fmt.Printf("expires: %d\n", tokenResponse.ExpiresIn)
	fmt.Printf("type: %s\n", tokenResponse.TokenType)
	fmt.Printf("Don't forget to overwrite your token in .env!")
}

func (s *Service) GetLiveUsers(follows []string) LiveUsersResponse {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams", nil)
	if err != nil {
		log.Fatal(err)
	}
	query := req.URL.Query()
	for _, user := range follows {
		query.Add("user_login", user)
	}
	req.URL.RawQuery = query.Encode()

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.AccessToken))
	req.Header.Add("Client-Id", s.ClientId)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var liveUsersResponse LiveUsersResponse
	err = json.Unmarshal(body, &liveUsersResponse)
	if err != nil {
		log.Fatal(err)
	}

	return liveUsersResponse
}
