package handler

import (
	"net/http"

	"ci.drugs.main/okit/pharmacy_mdlp_2/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getFilteredReestrKIZ(c *gin.Context) {
	var input model.SgtinFilterInput

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

	output, statusCode, err := h.services.ReestrAPI.GetFilteredReestrKIZ(token, input)
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

func (h *Handler) getReestrKizByList(c *gin.Context) {
	var input model.SgtinByListInput

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

	output, statusCode, err := h.services.ReestrAPI.GetReestrKizByList(token, input)
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

func (h *Handler) getPublicReestrKizByList(c *gin.Context) {
	var input model.SgtinByListInput

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

	output, statusCode, err := h.services.ReestrAPI.GetPublicReestrKizByList(token, input)
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

func (h *Handler) getPublicArchReestrKizByList(c *gin.Context) {
	var input model.SgtinByListInput

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

	output, statusCode, err := h.services.ReestrAPI.GetPublicArchReestrKizByList(token, input)
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

func (h *Handler) getDetailInfoSgtin(c *gin.Context) {
	sgtin := c.Param("sgtin")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.ReestrAPI.GetDetailInfoSgtin(token, sgtin)
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

func (h *Handler) getDetailArchInfoSgtin(c *gin.Context) {
	sgtin := c.Param("sgtin")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.ReestrAPI.GetDetailArchInfoSgtin(token, sgtin)
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

func (h *Handler) getDocumentsBySgtin(c *gin.Context) {
	sgtin := c.Query("sgtin")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.ReestrAPI.GetDocumentsBySgtin(token, sgtin)
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

func (h *Handler) validateSgtin(c *gin.Context) {
	var input model.SgtinListFilter

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

	output, statusCode, err := h.services.ReestrAPI.ValidateSgtin(token, input)
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

func (h *Handler) getSsccHierarchy(c *gin.Context) {
	sscc := c.Param("sscc")

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.ReestrAPI.GetSsccHierarchy(token, sscc)
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

func (h *Handler) getKIZInfoFromSscc(c *gin.Context) {
	sscc := c.Param("sscc")
	var input model.SgtinsInfoSsccInput

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

	output, statusCode, err := h.services.ReestrAPI.GetKIZInfoFromSscc(token, sscc, input)
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

func (h *Handler) getSsccFullHierarchy(c *gin.Context) {
	arrSSCC := c.QueryArray("sscc")
	SSCC := ""

	for _, element := range arrSSCC {
		SSCC = SSCC + "sscc=" + element + "&"
	}

	//Получить токен
	token, statusCode, err := h.GetToken()
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	output, statusCode, err := h.services.ReestrAPI.GetSsccFullHierarchy(token, SSCC)
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

func (h *Handler) addTrustedPartners(c *gin.Context) {
	var input model.TrustedPartnersInput

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

	output, statusCode, err := h.services.ReestrAPI.AddTrustedPartners(token, input)
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

func (h *Handler) deleteTrustedPartners(c *gin.Context) {
	var input model.TrustedPartnersInput

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

	statusCode, err = h.services.ReestrAPI.DeleteTrustedPartners(token, input)
	if err != nil {
		newErrorResponse(c, statusCode, err.Error())
		return
	}
}

func (h *Handler) getFilterTrustedPartners(c *gin.Context) {
	var input model.TrustedPartnerFilterInput

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

	output, statusCode, err := h.services.ReestrAPI.GetFilterTrustedPartners(token, input)
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

func (h *Handler) getFilterReestrPartners(c *gin.Context) {
	var input model.ReestrPartnerFilterInput

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

	output, statusCode, err := h.services.ReestrAPI.GetFilterReestrPartners(token, input)
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

func (h *Handler) getEsklp(c *gin.Context) {
	var input model.EsklpInput

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

	output, statusCode, err := h.services.ReestrAPI.GetEsklp(token, input)
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
