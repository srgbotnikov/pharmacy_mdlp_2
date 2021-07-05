package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/mdlpapi"
)

type AuthApiService struct {
	api mdlpapi.AuthorizationAPI
}

func NewAuthApiService(api mdlpapi.AuthorizationAPI) *AuthApiService {
	return &AuthApiService{api}
}

func (s *AuthApiService) GetAuthCode() (string, int, error) {
	code, statusCode, err := s.api.GetAuthCode()
	if err != nil {
		return "", statusCode, err
	}

	return code, statusCode, nil
}

func (s *AuthApiService) GetToken(code, sign string) (string, int, error) {

	token, statusCode, err := s.api.GetToken(code, sign)

	if err != nil {
		return "", statusCode, err
	}
	return token, statusCode, nil
}
