package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/sign"
)

type CryptoproCSPService struct {
	sign sign.CryptoproCSP
}

func NewCryptoproCSPService(sign sign.CryptoproCSP) *CryptoproCSPService {
	return &CryptoproCSPService{sign}
}

func (s *CryptoproCSPService) SignFile(data, requestID string) (string, string, error) {
	dataSigned, sign, err := s.sign.SignFile(data, requestID)
	return dataSigned, sign, err
}

func (s *CryptoproCSPService) GetThumbprint() (string, error) {
	thumbprint, err := s.sign.GetThumbprint()
	return thumbprint, err
}

func (s *CryptoproCSPService) CurlUploadWebDAVFile(token, data, requestID, link string) error {
	err := s.sign.CurlUploadWebDAVFile(token, data, requestID, link)
	return err
}

func (s *CryptoproCSPService) CurlDownloadFile(token, link string) (string, error) {
	output, err := s.sign.CurlDownloadFile(token, link)
	return output, err
}
