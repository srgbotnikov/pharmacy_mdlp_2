package model

//SgtinFilter фильтр для поиска по реестрам КИЗ
type SgtinFilter struct {
	Status                    []string `json:"status,omitempty" xml:"status,omitempty"`
	Gtin                      string   `json:"gtin,omitempty" xml:"gtin,omitempty"`
	Sgtin                     string   `json:"sgtin,omitempty" xml:"sgtin,omitempty"`
	Batch                     string   `json:"batch,omitempty" xml:"batch,omitempty"`
	SysId                     string   `json:"sys_id,omitempty" xml:"sys_id,omitempty"`
	EmissionOperationDateFrom string   `json:"emission_operation_date_from,omitempty" xml:"emission_operation_date_from,omitempty"`
	EmissionOperationDateTo   string   `json:"emission_operation_date_to,omitempty" xml:"emission_operation_date_to,omitempty"`
	LastTracingOpDateFrom     string   `json:"last_tracing_op_date_from,omitempty" xml:"last_tracing_op_date_from,omitempty"`
	LastTracingOpDateTo       string   `json:"last_tracing_op_date_to,omitempty" xml:"last_tracing_op_date_to,omitempty"`
	OMSOrderId                string   `json:"oms_order_id,omitempty" xml:"oms_order_id,omitempty"`
	Source                    string   `json:"source,omitempty" xml:"source,omitempty"`
}

type SgtinFilterInput struct {
	Filter    *SgtinFilter `json:"filter" binding:"required"`
	StartFrom *int         `json:"start_from" binding:"required"`
	Count     int          `json:"count" binding:"required"`
}

type SgtinFilterOutput struct {
	Entries []*Sgtin `json:"entries" xml:"entries"`
	Total   int      `json:"total" xml:"total"`
}

type SgtinListFilter struct {
	Sgtins []string `json:"sgtins" xml:"sgtins"`
}

type SgtinByListInput struct {
	Filter *SgtinListFilter `json:"filter" binding:"required"`
}

type SgtinByListOutput struct {
	Total         int            `json:"total" xml:"total"`
	Failed        int            `json:"failed" xml:"failed"`
	Entries       []*Sgtin       `json:"entries" xml:"entries"`
	FailedEntries []*FailedSgtin `json:"failed_entries" xml:"failed_entries"`
}

//Sgtin ...
type Sgtin struct {
	ID                    string                   `json:"id" xml:"id"`
	INN                   string                   `json:"inn" xml:"inn"`
	Gtin                  string                   `json:"gtin" xml:"gtin"`
	Sgtin                 string                   `json:"sgtin" xml:"sgtin"`
	Status                string                   `json:"status" xml:"status"`
	StatusDate            string                   `json:"status_date" xml:"status_date"`
	Batch                 string                   `json:"batch" xml:"batch"`
	Owner                 string                   `json:"owner" xml:"owner"`
	EmissionType          int                      `json:"emission_type,omitempty" xml:"emission_type,omitempty"`
	ReleaseDate           string                   `json:"release_date" xml:"release_date"`
	EmissionOperationDate string                   `json:"emission_operation_date" xml:"emission_operation_date"`
	FederalSubjectCode    string                   `json:"federal_subject_code,omitempty" xml:"federal_subject_code,omitempty"`
	FederalSubjectName    string                   `json:"federal_subject_name" xml:"federal_subject_name"`
	ExpirationDate        string                   `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
	ProdName              string                   `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	SellName              string                   `json:"sell_name,omitempty" xml:"sell_name,omitempty"`
	FullProdName          string                   `json:"full_prod_name,omitempty" xml:"full_prod_name,omitempty"`
	RegHolder             string                   `json:"reg_holder,omitempty" xml:"reg_holder,omitempty"`
	Pack1Desc             string                   `json:"pack1_desc,omitempty" xml:"pack1_desc,omitempty"`
	Pack3ID               string                   `json:"pack3_id,omitempty" xml:"pack3_id,omitempty"`
	LastTracingOpDate     string                   `json:"last_tracing_op_date,omitempty" xml:"last_tracing_op_date,omitempty"`
	SourceType            int                      `json:"source_type,omitempty" xml:"source_type,omitempty"`
	DrugCode              string                   `json:"drug_code,omitempty" xml:"drug_code,omitempty"`
	ProdFromName          string                   `json:"prod_from_name,omitempty" xml:"prod_from_name,omitempty"`
	ProdDName             string                   `json:"prod_d_name,omitempty" xml:"prod_d_name,omitempty"`
	CustomPointID         string                   `json:"custom_point_id,omitempty" xml:"custom_point_id,omitempty"`
	OmsOrderID            string                   `json:"oms_order_id,omitempty" xml:"oms_order_id,omitempty"`
	BillingInfo           *SgtinBillingInformation `json:"billing_info,omitempty" xml:"billing_info,omitempty"`
	BillingState          int                      `json:"billing_state,omitempty" xml:"billing_state,omitempty"`
	VznDrug               bool                     `json:"vzn_drug,omitempty" xml:"vzn_drug,omitempty"`
	Gnvlp                 bool                     `json:"gnvlp,omitempty" xml:"gnvlp,omitempty"`
	HaltDocDate           string                   `json:"halt_doc_date,omitempty" xml:"halt_doc_date,omitempty"`
	HaltDate              string                   `json:"halt_date,omitempty" xml:"halt_date,omitempty"`
	HaltDocNum            string                   `json:"halt_doc_num,omitempty" xml:"halt_doc_num,omitempty"`
	HaltID                string                   `json:"halt_id,omitempty" xml:"halt_id,omitempty"`
	SysID                 string                   `json:"sys_id,omitempty" xml:"sys_id,omitempty"`
	WithdrawalType        int                      `json:"withdrawal_type,omitempty" xml:"withdrawal_type,omitempty"`
	ViaDevice             int                      `json:"via_device,omitempty" xml:"via_device,omitempty"`
	PackingInn            string                   `json:"packing_inn,omitempty" xml:"packing_inn,omitempty"`
	PackingName           string                   `json:"packing_name,omitempty" xml:"packing_name,omitempty"`
	PackingId             string                   `json:"packing_id,omitempty" xml:"packing_id,omitempty"`
	ControlInn            string                   `json:"control,omitempty" xml:"control,omitempty"`
	ControlName           string                   `json:"control_name,omitempty" xml:"control_name,omitempty"`
	ControlId             string                   `json:"control_id,omitempty" xml:"control_id,omitempty"`
}

type SgtinBillingInformation struct {
	IsPrepaid   bool                       `json:"is_prepaid" xml:"is_prepaid"`
	FreeCode    bool                       `json:"free_code" xml:"free_code"`
	IsPaid      bool                       `json:"is_paid" xml:"is_paid"`
	ContainsVzn bool                       `json:"contains_vzn" xml:"contains_vzn"`
	Payments    []*SgtinPaymentInformation `json:"payments" xml:"payments"`
}

type SgtinPaymentInformation struct {
	CreatedDate string `json:"created_date" xml:"created_date"`
	PaymentDate string `json:"payment_date" xml:"payment_date"`
	Tariff      int    `json:"tariff" xml:"tariff"`
}

type FailedSgtin struct {
	Sgtin     string `json:"sgtin" xml:"sgtin"`
	ErrorCode int    `json:"error_code" xml:"error_code"`
	ErrorDesc string `json:"error_desc" xml:"error_desc"`
}

//PublicSgtin ...
type PublicSgtin struct {
	Sgtin          string `json:"sgtin" xml:"sgtin"`
	Batch          string `json:"batch" xml:"batch"`
	ExpirationDate string `json:"expiration_date" xml:"expiration_date"`
	ProdName       string `json:"prod_name" xml:"prod_name"`
	SellName       string `json:"sell_name" xml:"sell_name"`
	ProdDName      string `json:"prod_d_name" xml:"prod_d_name"`
	ProdFromName   string `json:"prod_from_name" xml:"prod_from_name"`
	RegDate        string `json:"reg_date" xml:"reg_date"`
	RegNumber      string `json:"reg_number" xml:"reg_number"`
	DrugCode       string `json:"drug_code" xml:"drug_code"`
	RegHolder      string `json:"reg_holder" xml:"reg_holder"`
	Status         string `json:"status" xml:"status"`
	EmissionType   int    `json:"emission_type" xml:"emission_type"`
	BranchId       string `json:"branch_id" xml:"branch_id"`
	Sscc           string `json:"sscc" xml:"sscc"`
	WithdrawalType int    `json:"withdrawal_type" xml:"withdrawal_type"`
	ViaDevice      int    `json:"via_device" xml:"via_device"`
}

type SgtinPublicByListOutput struct {
	Total         int            `json:"total" xml:"total"`
	Failed        int            `json:"failed" xml:"failed"`
	Entries       []*PublicSgtin `json:"entries" xml:"entries"`
	FailedEntries []string       `json:"failed_entries" xml:"failed_entries"`
}

type GtinInfo struct {
	Id              string `json:"id" xml:"id"`
	Gtin            string `json:"gtin" xml:"gtin"`
	RegNumber       string `json:"reg_number" xml:"reg_number"`
	RegDate         string `json:"reg_date" xml:"reg_date"`
	TypeFrom        string `json:"type_from" xml:"type_from"`
	ProdPack1EdName string `json:"prod_pack1_ed_name" xml:"prod_pack1_ed_name"`
	PackerAddress   string `json:"packer_address" xml:"packer_address"`
	ProdName        string `json:"prod_name" xml:"prod_name"`
	ProdSellName    string `json:"prod_sell_name" xml:"prod_sell_name"`
	ProdContent     string `json:"prod_content" xml:"prod_content"`
	ProdDesc        string `json:"prod_desc" xml:"prod_desc"`
	ProdPack1Ed     string `json:"prod_pack_1_ed" xml:"prod_pack_1_ed"`
	RegEndDate      string `json:"reg_end_date" xml:"reg_end_date"`
	ProdDName       string `json:"prod_d_name" xml:"prod_d_name"`
	ProdPack1Name   string `json:"prod_pack_1_name" xml:"prod_pack_1_name"`
	ProdPack2Name   string `json:"prod_pack_2_name" xml:"prod_pack_2_name"`
	ProdPack12      string `json:"prod_pack_1_2" xml:"prod_pack_1_2"`
	TnVed           string `json:"tn_ved" xml:"tn_ved"`
	Gnvlp           bool   `json:"gnvlp" xml:"gnvlp"`
	MaxGnvlp        string `json:"max_gnvlp" xml:"max_gnvlp"`
	MaxGnvlpRegDate string `json:"max_gnvlp_reg_date" xml:"max_gnvlp_reg_date"`
	RegHolder       string `json:"reg_holder" xml:"reg_holder"`
	RegCountry      string `json:"reg_country" xml:"reg_country"`
	ProdStatus      string `json:"prod_status" xml:"prod_status"`
	MinZdrav        bool   `json:"min_zdrav" xml:"min_zdrav"`
	Gs1             bool   `json:"gs1" xml:"gs1"`
	CostLimit       string `json:"cost_limit" xml:"cost_limit"`
	RegInn          string `json:"reg_inn" xml:"reg_inn"`
	Completeness    string `json:"completeness" xml:"completeness"`
	ProdFromName    string `json:"prod_from_name" xml:"prod_from_name"`
	GlfName         string `json:"glf_name" xml:"glf_name"`
	GlfCountry      string `json:"glf_country" xml:"glf_country"`
	DrugCode        string `json:"drug_code" xml:"drug_code"`
	DrugCodeVersion int    `json:"drug_code_version" xml:"drug_code_version"`
}

type SgtinDetailInfoOutput struct {
	SgtinInfo *Sgtin    `json:"sgtin_info" xml:"sgtin_info"`
	GtinInfo  *GtinInfo `json:"gtin_info" xml:"gtin_info"`
}

type ValidSgtin struct {
	Sgtin string `json:"sgtin" xml:"sgtin"`
	Valid bool   `json:"valid" xml:"valid"`
	Found bool   `json:"found" xml:"found"`
}

type ValidSgtinListOutput struct {
	Entries []*ValidSgtin `json:"entries" xml:"entries"`
}

type SSCCInfo struct {
	Sscc         string `json:"sscc" xml:"sscc"`
	ReleaseDate  string `json:"release_date" xml:"release_date"`
	SystemSubjID string `json:"system_subj_id" xml:"system_subj_id"`
}

type SSCCInfoListOutput struct {
	Up        []*SSCCInfo `json:"up,omitempty" xml:"up,omitempty"`
	Down      []*SSCCInfo `json:"down,omitempty" xml:"down,omitempty"`
	ErrorCode int         `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ErrorDesc string      `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
}

type SgtinsInfoSsccInput struct {
	StartFrom *int   `json:"start_from" binding:"required"`
	Count     int    `json:"count" binding:"required"`
	Source    string `json:"source,omitempty"`
}

type SgtinsInfoOutput struct {
	Entries   []*Sgtin `json:"entries" xml:"entries"`
	Total     int      `json:"total" xml:"total"`
	ErrorCode int      `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ErrorDesc string   `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
}

type HierarchySsccSgtinInfo struct {
	Sgtin                 string                    `json:"sgtin,omitempty" xml:"sgtin,omitempty"`
	Sscc                  string                    `json:"sscc" xml:"sscc"`
	Gtin                  string                    `json:"gtin,omitempty" xml:"gtin,omitempty"`
	Status                string                    `json:"status,omitempty" xml:"status,omitempty"`
	Batch                 string                    `json:"batch,omitempty" xml:"batch,omitempty"`
	PackingDate           string                    `json:"packing_date,omitempty" xml:"packing_date,omitempty"`
	ExpirationDate        string                    `json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
	OwnerID               string                    `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerOrganizationName string                    `json:"owner_organization_name,omitempty" xml:"owner_organization_name,omitempty"`
	Childs                []*HierarchySsccSgtinInfo `json:"childs,omitempty" xml:"childs,omitempty"`
}

type HierarchySsccOutput struct {
	Up   *HierarchySsccSgtinInfo `json:"up,ommitempty" xml:"up,ommitempty"`
	Down *HierarchySsccSgtinInfo `json:"down,ommitempty" xml:"down,ommitempty"`
	//	ErrorDescription string                  `json:"error_description,ommitempty" xml:"error_description,ommitempty"`
}

type TrustedPartnersInput struct {
	TrustedPartners []string `json:"trusted_partners" xml:"trusted_partners"`
}

type TrustedPartnersOutput struct {
	Code           int      `json:"code" xml:"code"`
	FailedPartners []string `json:"failed_partners,omitempty" xml:"failed_partners,omitempty"`
}

type TrustedPartnerFilterInput struct {
	Filter    *TrustedPartnersFilter `json:"filter" binding:"required"`
	StartFrom *int                   `json:"start_from" binding:"required"`
	Count     *int                   `json:"count" binding:"required"`
}

type TrustedPartnersFilter struct {
	TrustedInn   string `json:"trusted_partners,omitempty" xml:"trusted_partners,omitempty"`
	TrustedSysId string `json:"trusted_sys_id,omitempty" xml:"trusted_sys_id,omitempty"`
}

type TrustedPartnersFilterOutput struct {
	Entries []*TrustedPartnerEntry `json:"entries" xml:"entries"`
	Total   int                    `json:"total" xml:"total"`
}

type TrustedPartnerEntry struct {
	SysId   string `json:"sys_id" xml:"sys_id"`
	Inn     string `json:"inn" xml:"inn"`
	OrgName string `json:"org_name" xml:"org_name"`
}

type ReestrPartnerFilterInput struct {
	Filter    *PartnersFilter `json:"filter" binding:"required"`
	StartFrom *int            `json:"start_from" binding:"required"`
	Count     *int            `json:"count" binding:"required"`
}

type PartnersFilter struct {
	SystemSubjId        string `json:"system_subj_id,omitempty" xml:"system_subj_id,omitempty"`
	FederalSubjectCode  string `json:"federal_subject_code,omitempty" xml:"federal_subject_code,omitempty"`
	FederalDistrictCode string `json:"federal_district_code,omitempty" xml:"federal_district_code,omitempty"`
	Country             string `json:"country,omitempty" xml:"country,omitempty"`
	OrgName             string `json:"org_name,omitempty" xml:"org_name,omitempty"`
	Inn                 string `json:"inn,omitempty" xml:"inn,omitempty"`
	Kpp                 string `json:"kpp,omitempty" xml:"kpp,omitempty"`
	Ogrn                string `json:"ogrn,omitempty" xml:"ogrn,omitempty"`
	StartDate           string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	EndDate             string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	RegEntityType       int    `json:"reg_entity_type" xml:"reg_entity_type"`
	OpExecDateStart     string `json:"op_exec_date_start,omitempty" xml:"op_exec_date_start,omitempty"`
	OpExecDateEnd       string `json:"op_exec_date_end,omitempty" xml:"op_exec_date_end,omitempty"`
}

type ReestrPartnersOutput struct {
	//TODO добавить ForeignCounterparty
	FilteredRecords      []*RegistrationEntry `json:"filtered_records" xml:"filtered_records"`
	FilteredRecordsCount int                  `json:"filtered_records_count" xml:"filtered_records_count"`
}

type RegistrationEntry struct {
	SystemSubjID       string                 `json:"system_subj_id" xml:"system_subj_id"`
	Branches           []*ResolvedFiasAddress `json:"branches" xml:"branches"`
	SafeWarehouse      []*ResolvedFiasAddress `json:"safe_warehose" xml:"safe_warehose"`
	Inn                string                 `json:"inn" xml:"inn"`
	KPP                string                 `json:"KPP" xml:"KPP"`
	OrgName            string                 `json:"ORG_NAME" xml:"ORG_NAME"`
	OGRN               string                 `json:"OGRN" xml:"OGRN"`
	FirstName          string                 `json:"FIRST_NAME" xml:"FIRST_NAME"`
	MiddleName         string                 `json:"MIDDLE_NAME" xml:"MIDDLE_NAME"`
	LastName           string                 `json:"LAST_NAME" xml:"LAST_NAME"`
	EntityType         int                    `json:"entity_type" xml:"entity_type"`
	OpDate             *OperationDate         `json:"op_date" xml:"op_date"`
	OpExecDate         string                 `json:"op_exec_date" xml:"op_exec_date"`
	CountryCode        string                 `json:"country_code" xml:"country_code"`
	FederalSubjectCode string                 `json:"federal_subject_code" xml:"federal_subject_code"`
	Itin               string                 `json:"itin" xml:"itin"`
	RegNum             string                 `json:"regNum" xml:"regNum"`
	OrgAddress         string                 `json:"org_address" xml:"org_address"`
	Kpp2               string                 `json:"kpp" xml:"kpp"`
	Ogrn2              string                 `json:"ogrn" xml:"ogrn"`
	RegDate            string                 `json:"regDate" xml:"regDate"`
}

type OperationDate struct {
	ODate string `json:"$date" xml:"$date"`
}

type ResolvedFiasAddress struct {
	ID              string           `json:"id" xml:"id"`
	AddressFias     *AddressFias     `json:"address_fias" xml:"address_fias"`
	AddressResolved *AddressResolved `json:"address_resolved" xml:"address_resolved"`
	Status          int              `json:"status" xml:"status"`
}

type AddressFias struct {
	Aoguid    string `json:"aoguid" xml:"aoguid"`
	Houseguid string `json:"houseguid" xml:"houseguid"`
	Room      string `json:"room" xml:"room"`
}

type AddressResolved struct {
	Code    int    `json:"code" xml:"code"`
	Address string `json:"address" xml:"address"`
}

type EsklpInput struct {
	Filter    *EsklpFilter `json:"filter" binding:"required"`
	StartFrom *int         `json:"start_from" binding:"required"`
	Count     *int         `json:"count" binding:"required"`
}

type EsklpFilter struct {
	RegDate       string `json:"REG_DATE,omitempty" xml:"REG_DATE"`
	RegEndDate    string `json:"REG_END_DATE,omitempty" xml:"REG_END_DATE"`
	RegID         string `json:"reg_id,omitempty" xml:"reg_id"`
	RegHolder     string `json:"REG_HOLDER,omitempty" xml:"REG_HOLDER"`
	ProdSellName  string `json:"PROD_SELL_NAME,omitempty" xml:"PROD_SELL_NAME"`
	ProdName      string `json:"PROD_NAME,omitempty" xml:"PROD_NAME"`
	RegHolderCode string `json:"REG_HOLDER_CODE,omitempty" xml:"REG_HOLDER_CODE"`
	RegStatus     string `json:"REG_STATUS,omitempty" xml:"REG_STATUS"`
}

type EsklpOutput struct {
	Entries []*InfoEsklp `json:"entries" xml:"entries"`
	Total   int          `json:"total" xml:"total"`
}

type InfoEsklp struct {
	ID               string `json:"id" xml:"id"`
	RegId            string `json:"reg_id" xml:"reg_id"`
	ProdName         string `json:"PROD_NAME" xml:"PROD_NAME"`
	RegHolderCode    string `json:"REG_HOLDER_CODE" xml:"REG_HOLDER_CODE"`
	ProdPack1Id      string `json:"PROD_PACK_1_ID" xml:"PROD_PACK_1_ID"`
	ProdPack1Name    string `json:"PROD_PACK_1_NAME" xml:"PROD_PACK_1_NAME"`
	ProdPack1Ed      string `json:"PROD_PACK_1_ED" xml:"PROD_PACK_1_ED"`
	ProdPack1EdName  string `json:"PROD_PACK_1_ED_NAME" xml:"PROD_PACK_1_ED_NAME"`
	ProdPack2Id      string `json:"PROD_PACK_2_ID" xml:"PROD_PACK_2_ID"`
	ProdPack2Name    string `json:"PROD_PACK_2_NAME" xml:"PROD_PACK_2_NAME"`
	RegCountry       string `json:"REG_COUNTRY" xml:"REG_COUNTRY"`
	ProdPack1Size    string `json:"PROD_PACK_1_SIZE" xml:"PROD_PACK_1_SIZE"`
	ProdD            string `json:"PROD_D" xml:"PROD_D"`
	RegHolderCodeF   string `json:"REG_HOLDER_CODE_F" xml:"REG_HOLDER_CODE_F"`
	TnVed            string `json:"TN_VED" xml:"TN_VED"`
	ProdDName        string `json:"PROD_D_NAME" xml:"PROD_D_NAME"`
	ProdFromName     string `json:"PROD_FORM_NAME" xml:"PROD_FORM_NAME"`
	ProdId           string `json:"PROD_ID" xml:"PROD_ID"`
	ProdPack1        string `json:"PROD_PACK_1" xml:"PROD_PACK_1"`
	ProdSellName     string `json:"PROD_SELL_NAME" xml:"PROD_SELL_NAME"`
	MaxGnvlp         string `json:"MAX_GNVLP" xml:"MAX_GNVLP"`
	ProdPack12       string `json:"PROD_PACK_1_2" xml:"PROD_PACK_1_2"`
	RegDate          string `json:"REG_DATE" xml:"REG_DATE"`
	RegHolder        string `json:"REG_HOLDER" xml:"REG_HOLDER"`
	Gnvlp            string `json:"GNVLP" xml:"GNVLP"`
	DrugCode         string `json:"DRUG_CODE" xml:"DRUG_CODE"`
	DrugCodeVersion  int    `json:"DRUG_CODE_VERSION" xml:"DRUG_CODE_VERSION"`
	Completeness     string `json:"COMPLETENESS" xml:"COMPLETENESS"`
	GlfName          string `json:"GLF_NAME" xml:"GLF_NAME"`
	GlfCountry       string `json:"GLF_COUNTRY" xml:"GLF_COUNTRY"`
	ProdDNormName    string `json:"PROD_D_NORM_NAME" xml:"PROD_D_NORM_NAME"`
	ProdNormName     string `json:"PROD_NORM_NAME" xml:"PROD_NORM_NAME"`
	ProdFormNormName string `json:"PROD_FORM_NORM_NAME" xml:"PROD_FORM_NORM_NAME"`
}
