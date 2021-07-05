package model

type DocumentLarge struct {
	Link       string `json:"link" xml:"link"`
	DocumentID string `json:"document_id" xml:"document_id"`
}

type CancelDocumentInput struct {
	DocumentID string `json:"document_id" binding:"required"`
	RequestID  string `json:"request_id" binding:"required"`
}

//DocFilter Содержит информацию для фильтрации списка документов
type DocFilter struct {
	//Дата начала периода фильтрации
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	//Дата окончания периода фильтрации
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	//Уникальный идентификатор документа
	DocumentID string `json:"document_id,omitempty" xml:"document_id,omitempty"`
	//Массив идентификаторов документов
	DocumentIds []string `json:"document_ids,omitempty" xml:"document_ids,omitempty"`
	//Уникальный идентификатор запроса
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//Тип документа
	DocType int `json:"doc_type,omitempty" xml:"doc_type,omitempty"`
	//Статус документа
	DocStatus string `json:"doc_status,omitempty" xml:"doc_status,omitempty"`
	//Тип загрузки в систему
	FileUploadType int `json:"file_uploadtype,omitempty" xml:"file_uploadtype,omitempty"`
	//Дата обработки документа: начало периода
	ProcessedDateFrom string `json:"processed_date_from,omitempty" xml:"processed_date_from,omitempty"`
	//Дата обработки документа: окончание периода
	ProcessedDateTo string `json:"processed_date_to,omitempty" xml:"processed_date_to,omitempty"`
	//Уникальный идентификатор отправителя
	SenderID string `json:"sender_id,omitempty" xml:"sender_id,omitempty"`
	//Уникальный идентификатор получателя
	ReceiverID string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
	//Идентификатор отчета СУЗ
	SkzkmReportID string `json:"skzkm_report_id,omitempty" xml:"skzkm_report_id,omitempty"`
	//Номер документа основания
	DocNum string `json:"doc_num,omitempty" xml:"doc_num,omitempty"`
	//Статус обработки документа. Только для DocOutcomeFilter
	ProcessingDocumentStatus string `json:"processing_document_status,omitempty" xml:"processing_document_status,omitempty"`
	//Признак, регулирующий возможность получать только новые (непрочитанные) документы. Только для DocIncomeFilter
	OnlyNew bool `json:"only_new,omitempty" xml:"only_new,omitempty"`
}

type DocFilterInput struct {
	Filter *DocFilter `json:"filter" binding:"required"`
	//Индекс первой записи в списке возвращаемых документов
	StartFrom *int `json:"start_from" binding:"required"`
	//Количество записей в списке возвращаемых документов
	Count int `json:"count" binding:"required"`
}

type Document struct {
	//Уникальный идентификатор запроса
	RequestID string `json:"request_id" xml:"request_id"`
	//Уникальный идентификатор документа
	DocumentID string `json:"document_id" xml:"document_id"`
	//Дата получения документа
	Date string `json:"date" xml:"date"`
	//Дата обработки документа
	ProcessedDate string `json:"processed_date" xml:"processed_date"`
	//Отправитель документа
	Sender string `json:"sender" xml:"sender"`
	//Получатель документа
	Receiver string `json:"receiver" xml:"receiver"`
	//Идентификатор субъекта обращения в МДЛП
	SysID string `json:"sys_id" xml:"sys_id"`
	//Тип документа
	DocType int `json:"doc_type" xml:"doc_type"`
	//Статус документа
	DocStatus string `json:"doc_status" xml:"doc_status"`
	//Тип загрузки в систему
	FileUploadType int `json:"file_uploadtype" xml:"file_uploadtype"`
	//Версия документа
	Version string `json:"version" xml:"version"`
	//Номер документа основания. OutcomeDocument
	DocNum string `json:"doc_num" xml:"doc_num"`
	//Уникальный идентификатор регистратора событий. OutcomeDocument
	DeviceID string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	//Уникальный идентификатор системы, сформировавшей сообщение. OutcomeDocument
	SkzkmOriginMsgID string `json:"skzkm_origin_msg_id,omitempty" xml:"skzkm_origin_msg_id,omitempty"`
	//Идентификатор отчета СУЗ. OutcomeDocument
	SkzkmReportID string `json:"skzkm_report_id,omitempty" xml:"skzkm_report_id,omitempty"`
	//Статус обработки документа. OutcomeDocument
	ProcessingDocumentStatus string `json:"processing_document_status,omitempty" xml:"processing_document_status,omitempty"`
	//Идентификатор отправителя документа в МДЛП. IncomeDocument
	SenderSysID string `json:"sender_sys_id,omitempty" xml:"sender_sys_id,omitempty"`
	//Признак, регулирующий возможность получать только новые(непрочитанные) документы. IncomeDocument
	IsNew bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
}

// //OutcomeDocumentList список исходящих документов
// type OutcomeDocumentList struct {
// 	Documents []*Document `json:"documents"`
// 	Total     int         `json:"total"`
// }

//IncomeDocumentList список входящих документов
// type IncomeDocumentList struct {
// 	Documents []*Document `json:"documents"`
// 	Total     int         `json:"total"`
// }

//DocumentList список документов
type DocumentList struct {
	Documents []*Document `json:"documents" xml:"documents"`
	Total     int         `json:"total" xml:"total"`
}

//DocumentMarkReadInput список ИД документов для постановки статуса "Прочитано". Максимум 100
type DocumentMarkReadInput struct {
	//Массив идентификаторов документов
	DocumentIds []string `json:"document_ids" binding:"required"`
}

//DocumentMarkReadOutput выходные данные после постановки документам статуса "Прочитано"
type DocumentMarkReadOutput struct {
	//Массив идентификаторов документов, которым проставлен статус "прочитано"
	MarkedDocumentIds []string `json:"marked_document_ids" xml:"marked_document_ids"`
	//Количество документов, для которых проставлен статус "прочитано"
	Total int `json:"total" xml:"total"`
}

//DocumentInfo выходные данные инофрмация по документу
type DocumentInfo struct {
	RequestID      string `json:"request_id" xml:"request_id"`
	DocumentID     string `json:"document_id" json:"document_id"`
	Date           string `json:"date" xml:"date"`
	ProcessedDate  string `json:"processed_date" xml:"processed_date"`
	Sender         string `json:"sender" xml:"sender"`
	Receiver       string `json:"receiver" xml:"receiver"`
	SysID          string `json:"sys_id" xml:"sys_id"`
	DocType        int    `json:"doc_type" xml:"doc_type"`
	DocStatus      string `json:"doc_status" xml:"doc_status"`
	FileUploadType int    `json:"file_uploadtype" xml:"file_uploadtype"`
	Version        string `json:"version" xml:"version"`
	DocNum         string `json:"doc_num" xml:"doc_num"`
	Sscc           string `json:"sscc" xml:"sscc"`
}

//DocumentInfoList список DocumentInfo
type DocumentInfoList struct {
	Entries []*DocumentInfo `json:"entries" xml:"entries"`
	Total   int             `json:"total" xml:"total"`
}
