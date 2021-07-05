package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/cache"
)

type AuthCacheService struct {
	cache cache.Authorization
}

func NewAuthCacheService(cache cache.Authorization) *AuthCacheService {
	return &AuthCacheService{cache}
}

func (s *AuthCacheService) GetTokenFromCache(thumbprint string) (string, error) {
	token, err := s.cache.GetTokenFromCache(thumbprint)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthCacheService) SetTokenCache(token, thumbprint string) error {
	err := s.cache.SetTokenCache(token, thumbprint)
	return err
}
