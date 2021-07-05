package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/mdlpapi"
	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
)

type DocumentApiService struct {
	api mdlpapi.DocumentAPI
}

func NewDocumentApiService(api mdlpapi.DocumentAPI) *DocumentApiService {
	return &DocumentApiService{api}
}

func (s *DocumentApiService) SendDocument(token, datasigned, sign, requestID string, bulkProcessing bool) (string, int, error) {
	output, statusCode, err := s.api.SendDocument(token, datasigned, sign, requestID, bulkProcessing)
	if err != nil {
		return "", statusCode, err
	}
	return output, statusCode, nil
}

func (s *DocumentApiService) SendDocumentLarge(token, sign, hashSum, requestID string) (*model.DocumentLarge, int, error) {
	output, statusCode, err := s.api.SendDocumentLarge(token, sign, hashSum, requestID)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) SendFinished(token, documentID string) (string, int, error) {
	output, statusCode, err := s.api.SendFinished(token, documentID)
	if err != nil {
		return "", statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) CancelSendDocument(token string, input model.CancelDocumentInput) (int, error) {
	statusCode, err := s.api.CancelSendDocument(token, input)
	if err != nil {
		return statusCode, err
	}
	return statusCode, err
}

func (s *DocumentApiService) GetOutcomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error) {
	output, statusCode, err := s.api.GetOutcomeDocuments(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) GetIncomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error) {
	output, statusCode, err := s.api.GetIncomeDocuments(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) SetMarkReadDocuments(token string, input model.DocumentMarkReadInput) (*model.DocumentMarkReadOutput, int, error) {
	output, statusCode, err := s.api.SetMarkReadDocuments(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) GetMetadataDocument(token, documentID string) (*model.Document, int, error) {
	output, statusCode, err := s.api.GetMetadataDocument(token, documentID)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) GetDocByID(token, documentID string) (string, int, error) {
	output, statusCode, err := s.api.GetDocByID(token, documentID)
	if err != nil {
		return "", statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) GetDocumentsByRequestId(token, requestID string) (*model.DocumentList, int, error) {
	output, statusCode, err := s.api.GetDocumentsByRequestId(token, requestID)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, err
}

func (s *DocumentApiService) GetTicketById(token, documentID string) (string, int, error) {
	output, statusCode, err := s.api.GetTicketById(token, documentID)
	if err != nil {
		return "", statusCode, err
	}
	return output, statusCode, err
}
