package mdlpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
)

type ReestrMdlp struct {
	Url     string
	UrlMdlp string
}

func NewReestrMdlp(url, urlMdlp string) *ReestrMdlp {
	return &ReestrMdlp{
		Url:     url,
		UrlMdlp: urlMdlp,
	}
}

//GetFilteredReestrKIZ Метод поиска по реестрам КИЗ
func (p *ReestrMdlp) GetFilteredReestrKIZ(token string, input model.SgtinFilterInput) (*model.SgtinFilterOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/filter", p.Url)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinFilterOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetReestrKizByList Метод поиска по реестрам КИЗ по списку значений
func (p *ReestrMdlp) GetReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinByListOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/sgtins-by-list", p.Url)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinByListOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetPublicReestrKizByList Метод поиска по общедоступным реестрам КИЗ по списку значений
func (p *ReestrMdlp) GetPublicReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/public/sgtins-by-list", p.Url)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinPublicByListOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetPublicArchReestrKizByList Метод поиска по общедоступным реестрам КИЗ в архивном хранилище по списку значений
func (p *ReestrMdlp) GetPublicArchReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/public/archive/sgtins-by-list", p.Url)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinPublicByListOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetDetailInfoSgtin Метод для получения из реестров КИЗ детальной информации о КИЗ и связанным с ним ЛП
func (p *ReestrMdlp) GetDetailInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/%s", p.Url, sgtin)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinDetailInfoOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetDetailArchInfoSgtin  Метод для получения детальной информации о КИЗ в архивном хранилище и связанным с ним ЛП
func (p *ReestrMdlp) GetDetailArchInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/archive/%s", p.Url, sgtin)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinDetailInfoOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetDocumentsBySgtin Метод для получения перечня документов по идентификатору КИЗ
func (p *ReestrMdlp) GetDocumentsBySgtin(token, sgtin string) (*model.DocumentInfoList, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sgtin/documents?sgtin=%s", p.Url, sgtin)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.DocumentInfoList

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//ValidateSgtin Метод валидации кода макрировки
func (p *ReestrMdlp) ValidateSgtin(token string, input model.SgtinListFilter) (*model.ValidSgtinListOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/sgtin/validate", p.Url)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.ValidSgtinListOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetSsccHierarchy Метод для получения информации об иерархии вложенности третичной упаковки
func (p *ReestrMdlp) GetSsccHierarchy(token, sscc string) (*model.SSCCInfoListOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sscc/%s/hierarchy", p.Url, sscc)
	//url := p.Url + "/api/v1/reestr/sscc/" + sscc + "/hierarchy"
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SSCCInfoListOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetKIZInfoFromSscc Метод для получения информации о КИЗ в третичной упаковке
func (p *ReestrMdlp) GetKIZInfoFromSscc(token, sscc string, input model.SgtinsInfoSsccInput) (*model.SgtinsInfoOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sscc/%s/sgtins", p.Url, sscc)
	statusCode := http.StatusOK

	filterJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(filterJSON))

	// log.Println(payload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.SgtinsInfoOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetSsccFullHierarchy Метод для получения информации о полной иерархии вложенности третичной упаковки для нескольких SSCC
func (p *ReestrMdlp) GetSsccFullHierarchy(token, sscc string) ([]*model.HierarchySsccOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/sscc/full-hierarchy?%s", p.Url, sscc)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output []*model.HierarchySsccOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//AddTrustedPartners Метод добавления доверенного контрагента
func (p *ReestrMdlp) AddTrustedPartners(token string, input model.TrustedPartnersInput) (*model.TrustedPartnersOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/trusted_partners/add", p.Url)
	statusCode := http.StatusOK

	inputJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(inputJSON))

	// log.Println(payload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.TrustedPartnersOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//DeleteTrustedPartners Метод удаления доверенного контрагента
func (p *ReestrMdlp) DeleteTrustedPartners(token string, input model.TrustedPartnersInput) (int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/trusted_partners/delete", p.Url)
	statusCode := http.StatusOK

	inputJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(inputJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return res.StatusCode, errors.New(string(body))
	}

	return statusCode, nil
}

//GetFilterTrustedPartners Метод фильтрации доверенных контрагентов
func (p *ReestrMdlp) GetFilterTrustedPartners(token string, input model.TrustedPartnerFilterInput) (*model.TrustedPartnersFilterOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/trusted_partners/filter", p.Url)
	statusCode := http.StatusOK

	inputJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(inputJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.TrustedPartnersFilterOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetFilterReestrPartners Метод фильтрации по субъектам обращения
func (p *ReestrMdlp) GetFilterReestrPartners(token string, input model.ReestrPartnerFilterInput) (*model.ReestrPartnersOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr_partners/filter", p.Url)
	statusCode := http.StatusOK

	inputJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(inputJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.ReestrPartnersOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetEsklp Фильтрация по реестру ЕСКЛП
func (p *ReestrMdlp) GetEsklp(token string, input model.EsklpInput) (*model.EsklpOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/reestr/esklp/filter", p.Url)
	statusCode := http.StatusOK

	inputJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(inputJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.EsklpOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}
