package handler

import (
	"net/http"
)

//GetToken Получает токен из кеша. Если в кеше нет или протух, запрашивает новый и записывает в кеш
func (h *Handler) GetToken() (string, int, error) {
	statusCode := http.StatusOK

	thumbprint, _ := h.services.GetThumbprint()

	tokenCache, err := h.services.AuthoriztionCache.GetTokenFromCache(thumbprint)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if tokenCache == "" {
		code, statusCode, err := h.services.AuthorizationAPI.GetAuthCode()
		if err != nil {
			return "", statusCode, err
		}

		_, sign, err := h.services.Sign.SignFile(code, "code")
		if err != nil {
			return "", http.StatusInternalServerError, err
		}

		token, statusCode, err := h.services.AuthorizationAPI.GetToken(code, sign)
		if err != nil {
			return "", statusCode, err
		}

		//Записать токен в кеш
		err = h.services.SetTokenCache(token, thumbprint)
		if err != nil {
			return "", http.StatusInternalServerError, err
		}

		return token, statusCode, nil
	} else {
		return tokenCache, statusCode, nil
	}
}
