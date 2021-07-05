package service

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/cache"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/mdlpapi"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/repository"
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/sign"
	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
)

type Authorization interface {
}

type AuthorizationAPI interface {
	//GetAuthCode получение кода аутентификации
	GetAuthCode() (string, int, error)
	//GetToken получение токена.
	GetToken(code, sign string) (string, int, error)
}

type DocumentAPI interface {
	//SendDocument отправить документ
	SendDocument(token, datasigned, sign, requestID string, bulkProcessing bool) (string, int, error)
	//SendDocumentLarge отправка большого документа
	SendDocumentLarge(token, sign, hashSum, requestID string) (*model.DocumentLarge, int, error)
	//SendFinished завершение загрузки большого документа
	SendFinished(token, documentID string) (string, int, error)
	//CancelSendDocument отмена отправки документа
	CancelSendDocument(token string, input model.CancelDocumentInput) (int, error)
	//GetOutcomeDocuments получение списка исходящих документов
	GetOutcomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error)
	//GetIncomeDocuments получение списка входящих документов
	GetIncomeDocuments(token string, input model.DocFilterInput) (*model.DocumentList, int, error)
	//SetMarkReadDocuments передача информации о прочтении документа
	SetMarkReadDocuments(token string, input model.DocumentMarkReadInput) (*model.DocumentMarkReadOutput, int, error)
	//GetMetadataDocument получение метаданных документа
	GetMetadataDocument(token, documentID string) (*model.Document, int, error)
	//GetDocByID получение документа по ИД
	GetDocByID(token, documentID string) (string, int, error)
	//GetDocumentsByRequestId Получение списка документов по ИД запроса
	GetDocumentsByRequestId(token, requestID string) (*model.DocumentList, int, error)
	//GetTicketById получение квитанции по номеру исходящего документа
	GetTicketById(token, documentID string) (string, int, error)
}

type ReestrAPI interface {
	//GetFilteredReestrKIZ Метод поиска по реестрам КИЗ
	GetFilteredReestrKIZ(token string, input model.SgtinFilterInput) (*model.SgtinFilterOutput, int, error)
	//GetReestrKizByList Метод поиска по реестрам КИЗ по списку значений
	GetReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinByListOutput, int, error)
	//GetPublicReestrKizByList Метод поиска по общедоступным реестрам КИЗ по списку значений
	GetPublicReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error)
	//GetPublicArchReestrKizByList Метод поиска по общедоступным реестрам КИЗ в архивном хранилище по списку значений
	GetPublicArchReestrKizByList(token string, input model.SgtinByListInput) (*model.SgtinPublicByListOutput, int, error)
	//GetDetailInfoSgtin Метод для получения из реестров КИЗ детальной информации о КИЗ и связанным с ним ЛП
	GetDetailInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error)
	//GetDetailArchInfoSgtin  Метод для получения детальной информации о КИЗ в архивном хранилище и связанным с ним ЛП
	GetDetailArchInfoSgtin(token, sgtin string) (*model.SgtinDetailInfoOutput, int, error)
	//GetDocumentsBySgtin Метод для получения перечня документов по идентификатору КИЗ
	GetDocumentsBySgtin(token, sgtin string) (*model.DocumentInfoList, int, error)
	//ValidateSgtin Метод валидации кода макрировки
	ValidateSgtin(token string, input model.SgtinListFilter) (*model.ValidSgtinListOutput, int, error)
	//GetSsccHierarchy Метод для получения информации об иерархии вложенности третичной упаковки
	GetSsccHierarchy(token, sscc string) (*model.SSCCInfoListOutput, int, error)
	//GetKIZInfoFromSscc Метод для получения информации о КИЗ в третичной упаковке
	GetKIZInfoFromSscc(token, sscc string, input model.SgtinsInfoSsccInput) (*model.SgtinsInfoOutput, int, error)
	//GetSsccFullHierarchy Метод для получения информации о полной иерархии вложенности третичной упаковки для нескольких SSCC
	GetSsccFullHierarchy(token, sscc string) ([]*model.HierarchySsccOutput, int, error)
	//AddTrustedPartners Метод добавления доверенного контрагента
	AddTrustedPartners(token string, input model.TrustedPartnersInput) (*model.TrustedPartnersOutput, int, error)
	//DeleteTrustedPartners Метод удаления доверенного контрагента
	DeleteTrustedPartners(token string, input model.TrustedPartnersInput) (int, error)
	//GetFilterTrustedPartners Метод фильтрации доверенных контрагентов
	GetFilterTrustedPartners(token string, input model.TrustedPartnerFilterInput) (*model.TrustedPartnersFilterOutput, int, error)
	//GetFilterReestrPartners Метод фильтрации по субъектам обращения
	GetFilterReestrPartners(token string, input model.ReestrPartnerFilterInput) (*model.ReestrPartnersOutput, int, error)
	//GetEsklp Фильтрация по реестру ЕСКЛП
	GetEsklp(token string, input model.EsklpInput) (*model.EsklpOutput, int, error)
}

type AuthoriztionCache interface {
	//GetTokenFromCache получить токен из кеша
	GetTokenFromCache(thumbprint string) (string, error)
	//SetTokenCache записать токен в кеш
	SetTokenCache(token, thumbprint string) error
}

type Cdb interface {
}

type Sign interface {
	//SignFile подписать файл
	SignFile(data, requestID string) (string, string, error)
	//GetThumbprint получить отпечаток текущей подписи
	GetThumbprint() (string, error)
	//CurlUploadWebDAVFile загрузить документ большого объема с помощью curl
	CurlUploadWebDAVFile(token, data, requestID, link string) error
	//CurlDownloadFile скачать документ с помощью curl
	CurlDownloadFile(token, link string) (string, error)
}

type Service struct {
	Authorization
	AuthorizationAPI
	DocumentAPI
	ReestrAPI
	AuthoriztionCache
	Cdb
	Sign
}

func NewService(repos *repository.Repository, api *mdlpapi.MdlpAPI, sign *sign.Sign, cache *cache.Cache) *Service {
	return &Service{
		AuthorizationAPI:  NewAuthApiService(api),
		Sign:              NewCryptoproCSPService(sign),
		AuthoriztionCache: NewAuthCacheService(cache.Authorization),
		DocumentAPI:       NewDocumentApiService(api),
		ReestrAPI:         NewReestrApiService(api),
	}
}
