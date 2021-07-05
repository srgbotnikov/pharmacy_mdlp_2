package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/mdlpapi"
	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
)

type ReestrApiService struct {
	api mdlpapi.ReestrAPI
}

func NewReestrApiService(api mdlpapi.ReestrAPI) *ReestrApiService {
	return &ReestrApiService{api}
}

func (s *ReestrApiService) GetFilteredReestrKIZ(token string, input model.SgtinFilterInput) (*model.SgtinFilterOutput, int, error) {
	output, statusCode, err := s.api.GetFilteredReestrKIZ(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinByListOutput, int, error) {
	output, statusCode, err := s.api.GetReestrKizByList(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetPublicReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error) {
	output, statusCode, err := s.api.GetPublicReestrKizByList(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetPublicArchReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error) {
	output, statusCode, err := s.api.GetPublicArchReestrKizByList(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetDetailInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error) {
	output, statusCode, err := s.api.GetDetailInfoSgtin(token, sgtin)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetDetailArchInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error) {
	output, statusCode, err := s.api.GetDetailArchInfoSgtin(token, sgtin)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetDocumentsBySgtin(token, sgtin string) (*model.DocumentInfoList, int, error) {
	output, statusCode, err := s.api.GetDocumentsBySgtin(token, sgtin)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) ValidateSgtin(token string, input model.SgtinListFilter) (*model.ValidSgtinListOutput, int, error) {
	output, statusCode, err := s.api.ValidateSgtin(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetSsccHierarchy(token, sscc string) (*model.SSCCInfoListOutput, int, error) {
	output, statusCode, err := s.api.GetSsccHierarchy(token, sscc)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetKIZInfoFromSscc(token, sscc string, input model.SgtinsInfoSsccInput) (*model.SgtinsInfoOutput, int, error) {
	output, statusCode, err := s.api.GetKIZInfoFromSscc(token, sscc, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetSsccFullHierarchy(token, sscc string) ([]*model.HierarchySsccOutput, int, error) {
	output, statusCode, err := s.api.GetSsccFullHierarchy(token, sscc)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) AddTrustedPartners(token string, input model.TrustedPartnersInput) (*model.TrustedPartnersOutput, int, error) {
	output, statusCode, err := s.api.AddTrustedPartners(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) DeleteTrustedPartners(token string, input model.TrustedPartnersInput) (int, error) {
	statusCode, err := s.api.DeleteTrustedPartners(token, input)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (s *ReestrApiService) GetFilterTrustedPartners(token string, input model.TrustedPartnerFilterInput) (*model.TrustedPartnersFilterOutput, int, error) {
	output, statusCode, err := s.api.GetFilterTrustedPartners(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetFilterReestrPartners(token string, input model.ReestrPartnerFilterInput) (*model.ReestrPartnersOutput, int, error) {
	output, statusCode, err := s.api.GetFilterReestrPartners(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}

func (s *ReestrApiService) GetEsklp(token string, input model.EsklpInput) (*model.EsklpOutput, int, error) {
	output, statusCode, err := s.api.GetEsklp(token, input)
	if err != nil {
		return nil, statusCode, err
	}
	return output, statusCode, nil
}
