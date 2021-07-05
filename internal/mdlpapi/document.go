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

type DocumentMdlp struct {
	Url     string
	UrlMdlp string
}

func NewDocumentMdlp(url, urlMdlp string) *DocumentMdlp {
	return &DocumentMdlp{
		Url:     url,
		UrlMdlp: urlMdlp,
	}
}

//DocumentInput структура для отправки документа в МДЛП
type DocumentInput struct {
	Document       string `json:"document" binding:"required"`
	Sign           string `json:"sign" binding:"required"`
	RequestID      string `json:"request_id" binding:"required"`
	BulkProcessing bool   `json:"bulk_processing,omitempty"`
}

//DocumentResp структура в которой хранится ответ после отправки документа в МДЛП
type DocumentResp struct {
	DocumentID string `json:"document_id"`
}

//SendDocument отправка документа в МДЛП
func (p *DocumentMdlp) SendDocument(token, datasigned, sign, requestID string, bulkProcessing bool) (string, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/send", p.Url)
	statusCode := http.StatusOK

	input := &DocumentInput{
		Document:       datasigned,
		Sign:           sign,
		RequestID:      requestID,
		BulkProcessing: bulkProcessing,
	}

	docJSON, err := json.Marshal(input)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	payload := strings.NewReader(string(docJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output *DocumentResp

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return output.DocumentID, statusCode, nil
}

type DocumentLargeInput struct {
	Sign           string `json:"sign" binding:"required"`
	HashSum        string `json:"hash_sum" binding:"required"`
	RequestID      string `json:"request_id" binding:"required"`
	BulkProcessing bool   `json:"bulk_processing,omitempty"`
}

//SendDocumentLarge отправка большого документа
func (p *DocumentMdlp) SendDocumentLarge(token, sign, hashSum, requestID string) (*model.DocumentLarge, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/send_large", p.Url)
	statusCode := http.StatusOK

	b := &DocumentLargeInput{
		Sign:      sign,
		HashSum:   hashSum,
		RequestID: requestID}

	docJSON, err := json.Marshal(b)

	payload := strings.NewReader(string(docJSON))

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

	//Здесь в старом коде был таймаут

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return nil, res.StatusCode, errors.New(string(body))
	}

	var output *model.DocumentLarge

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	//Заменить адрес сервиса МДЛП на адрес Stunnel
	output.Link = strings.Replace(output.Link, p.UrlMdlp, p.Url, -1)

	return output, statusCode, nil
}

type DocumentFinishedInput struct {
	DocumentID string `json:"document_id" binding:"required"`
}

type DocumentFinishedOutput struct {
	RequestID string `json:"request_id"`
}

//SendFinished завершение загрузки большого документа
func (p *DocumentMdlp) SendFinished(token, documentID string) (string, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/send_finished", p.Url)
	statusCode := http.StatusOK

	b := &DocumentFinishedInput{
		DocumentID: documentID,
	}

	docJSON, err := json.Marshal(b)

	payload := strings.NewReader(string(docJSON))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	//Здесь в старом коде был таймаут

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output *DocumentFinishedOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return output.RequestID, statusCode, nil
}

//CancelSendDocument отмена отправки документа
func (p *DocumentMdlp) CancelSendDocument(token string, input model.CancelDocumentInput) (int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/cancel", p.Url)
	statusCode := http.StatusOK

	docJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(docJSON))

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

//GetOutcomeDocuments получение списка исходящих документов
func (p *DocumentMdlp) GetOutcomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/outcome", p.Url)
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

	var output *model.DocumentList

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetIncomeDocuments получение списка входящих документов
func (p *DocumentMdlp) GetIncomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/income", p.Url)
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

	var output *model.DocumentList

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//SetMarkReadDocuments передача информации о прочтении документа
func (p *DocumentMdlp) SetMarkReadDocuments(token string, input model.DocumentMarkReadInput) (*model.DocumentMarkReadOutput, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/income/mark_read", p.Url)
	statusCode := http.StatusOK

	docJSON, err := json.Marshal(input)

	payload := strings.NewReader(string(docJSON))

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

	var output *model.DocumentMarkReadOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetMetadataDocument получить метаданные документа
func (p *DocumentMdlp) GetMetadataDocument(token, documentID string) (*model.Document, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/%s", p.Url, documentID)
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

	var output *model.Document

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

type DocumentDownloadOutput struct {
	Link string `json:"link"`
}

//GetDocByID получение документа по ИД
func (p *DocumentMdlp) GetDocByID(token, documentID string) (string, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/download/%s", p.Url, documentID)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output DocumentDownloadOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	//Заменить адрес сервиса МДЛП на адрес Stunnel
	output.Link = strings.Replace(output.Link, p.UrlMdlp, p.Url, -1)

	return output.Link, statusCode, nil
}

//GetDocumentsByRequestId Получение списка документов по ИД запроса
func (p *DocumentMdlp) GetDocumentsByRequestId(token, requestID string) (*model.DocumentList, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/request/%s", p.Url, requestID)
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

	var output *model.DocumentList

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return output, statusCode, nil
}

//GetTicketById получение квитанции по номеру исходящего документа
func (p *DocumentMdlp) GetTicketById(token, documentID string) (string, int, error) {
	url := fmt.Sprintf("%s/api/v1/documents/%s/ticket", p.Url, documentID)
	statusCode := http.StatusOK

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Authorization", "token "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode >= 300 {
		return "", res.StatusCode, errors.New(string(body))
	}

	var output DocumentDownloadOutput

	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	//Заменить адрес сервиса МДЛП на адрес Stunnel
	output.Link = strings.Replace(output.Link, p.UrlMdlp, p.Url, -1)

	return output.Link, statusCode, nil
}
