package mdlpapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type AuthMdlp struct {
	Url          string
	ClientID     string
	ClientSecret string
	UserID       string
	AuthType     string
}

func NewAuthMdlp(url, clientID, clientSecret, userID, authType string) *AuthMdlp {
	return &AuthMdlp{
		Url:          url,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		UserID:       userID,
		AuthType:     authType,
	}
}

//AuthInput структура для запроса кода аутентификации
type AuthInput struct {
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	UserID       string `json:"user_id" binding:"required"`
	AuthType     string `json:"auth_type" binding:"required"`
}

//AuthResp структура в которой хранится ответ на запрос кода аутентификации
type AuthResp struct {
	Code string `json:"code"`
}

func (p *AuthMdlp) GetAuthCode() (string, int, error) {
	url := p.Url + "/api/v1/auth"
	statusCode := http.StatusOK

	input := &AuthInput{
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
		UserID:       p.UserID,
		AuthType:     p.AuthType}

	authJSON, err := json.Marshal(input)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	payload := strings.NewReader(string(authJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output *AuthResp

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return output.Code, statusCode, nil
}

//TokenInput структура для запроса токена
type TokenInput struct {
	Code      string `json:"code" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

//TokenResp ответ на запрос токена
type TokenResp struct {
	Token    string `json:"token"`
	Lifetime int    `json:"life_time"`
}

func (p *AuthMdlp) GetToken(code, sign string) (string, int, error) {
	statusCode := http.StatusOK

	url := p.Url + "/api/v1/token"

	input := &TokenInput{
		Code:      code,
		Signature: sign}

	authVar, err := json.Marshal(input)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	payload := strings.NewReader(string(authVar))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output *TokenResp

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return output.Token, statusCode, nil
}
