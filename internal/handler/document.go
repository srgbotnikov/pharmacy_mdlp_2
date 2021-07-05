package handler

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) sendDocument(c *gin.Context) {
	//Получить входные данные
	file := c.PostForm("file")
	requestID := c.PostForm("requestID")

	//Проверить есть ли file и requestID sendLargeDocument
	if file == "" || requestID == "" {
		newErrorResponse(c, http.StatusBadRequest, "Expected file and requestID input data")
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	//Подписать файл
	dataSigned, sign, err := h.services.Sign.SignFile(file, requestID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//Отправить
	docID, statusCode, err := h.services.DocumentAPI.SendDocument(token, dataSigned, sign, requestID, false)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, map[string]interface{}{
			"document_id": docID,
		})
	case "xml":
		c.XML(http.StatusOK, map[string]interface{}{
			"document_id": docID,
		})
	default:
		c.XML(http.StatusOK, map[string]interface{}{
			"document_id": docID,
		})
	}
}

//Объединено 3 запроса в один. Отправка, загрузка и завершение.
func (h *Handler) sendLargeDocument(c *gin.Context) {
	//Получить входные данные
	file := c.PostForm("file")
	requestID := c.PostForm("requestID")

	//Проверить есть ли file и requestID sendLargeDocument
	if file == "" || requestID == "" {
		newErrorResponse(c, http.StatusBadRequest, "Expected file and requestID input data")
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	//Посчитать хеш сумму
	sum := sha256.Sum256([]byte(file))
	s := fmt.Sprintf("%x", sum)
	hashSum := string(s)

	//Подписать файл
	_, sign, err := h.services.Sign.SignFile(file, requestID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//Отправить
	docLarge, statusCode, err := h.services.DocumentAPI.SendDocumentLarge(token, sign, hashSum, requestID)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	//Загрузить xml с помощью curl
	err = h.services.Sign.CurlUploadWebDAVFile(token, file, requestID, docLarge.Link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//time.Sleep(50 * time.Millisecond)

	//Завершить
	outRequestID, statusCode, err := h.services.DocumentAPI.SendFinished(token, docLarge.DocumentID)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, map[string]interface{}{
			"request_id": outRequestID,
		})
	case "xml":
		c.XML(http.StatusOK, map[string]interface{}{
			"request_id": outRequestID,
		})
	default:
		c.XML(http.StatusOK, map[string]interface{}{
			"request_id": outRequestID,
		})
	}
}

func (h *Handler) cancelDocument(c *gin.Context) {
	var input model.CancelDocumentInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	statusCode, err = h.services.DocumentAPI.CancelSendDocument(token, input)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

}

func (h *Handler) getOutcomeDocuments(c *gin.Context) {
	var input model.DocFilterInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.DocumentAPI.GetOutcomeDocuments(token, input)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, output)
	case "xml":
		c.XML(http.StatusOK, output)
	default:
		c.XML(http.StatusOK, output)
	}
}

func (h *Handler) getIncomeDocuments(c *gin.Context) {
	var input model.DocFilterInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.DocumentAPI.GetIncomeDocuments(token, input)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, output)
	case "xml":
		c.XML(http.StatusOK, output)
	default:
		c.XML(http.StatusOK, output)
	}
}

func (h *Handler) setMarkReadDocuments(c *gin.Context) {
	var input model.DocumentMarkReadInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.DocumentAPI.SetMarkReadDocuments(token, input)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, output)
	case "xml":
		c.XML(http.StatusOK, output)
	default:
		c.XML(http.StatusOK, output)
	}
}

func (h *Handler) getMetadataDocument(c *gin.Context) {
	docId := c.Param("docId")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.DocumentAPI.GetMetadataDocument(token, docId)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, output)
	case "xml":
		c.XML(http.StatusOK, output)
	default:
		c.XML(http.StatusOK, output)
	}
}

func (h *Handler) getDocById(c *gin.Context) {
	docId := c.Param("docId")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	link, statusCode, err := h.services.DocumentAPI.GetDocByID(token, docId)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	//Загрузить xml с помощью curl
	docXML, err := h.services.Sign.CurlDownloadFile(token, link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, docXML)
}

func (h *Handler) getDocumentsByRequestId(c *gin.Context) {
	requestId := c.Param("request_id")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.DocumentAPI.GetDocumentsByRequestId(token, requestId)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	format := c.Query("format")
	switch format {
	case "json":
		c.JSON(http.StatusOK, output)
	case "xml":
		c.XML(http.StatusOK, output)
	default:
		c.XML(http.StatusOK, output)
	}
}

func (h *Handler) getTicketById(c *gin.Context) {
	docId := c.Param("docId")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	link, statusCode, err := h.services.DocumentAPI.GetTicketById(token, docId)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	//Загрузить xml с помощью curl
	docXML, err := h.services.Sign.CurlDownloadFile(token, link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, docXML)
}
