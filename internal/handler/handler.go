package handler

import (
	"ci.drugs.main/okit/pharmacy_mdlp_2/internal/service"
	logstash "ci.drugs.main/okit/pharmacy_mdlp_2/pkg/gin-logrus-logstash"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// type Handler struct {
// 	services *service.Service
// }

// func NewHandler(services *service.Service) *Handler {
// 	return &Handler{services: services}
// }

type Handler struct {
	services    *service.Service
	logstashURL string
}

func NewHandler(services *service.Service, logstashURL string) *Handler {
	return &Handler{
		services:    services,
		logstashURL: logstashURL,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	log := logrus.New()

	router := gin.New()

	gin.SetMode(gin.ReleaseMode)

	router.Use(logstash.Logger(log, h.logstashURL, "mdlp"), gin.Recovery())

	v2 := router.Group("/v2")
	{
		//v2.POST("/auth", h.authAPI)
		documents := v2.Group("/documents")
		{
			documents.POST("/send", h.sendDocument)
			documents.POST("/send_large", h.sendLargeDocument)
			documents.POST("/cancel", h.cancelDocument)
			documents.POST("/outcome", h.getOutcomeDocuments)
			documents.POST("/income", h.getIncomeDocuments)
			documents.POST("/income/mark_read", h.setMarkReadDocuments)
			documents.GET("/:docId", h.getMetadataDocument)
			documents.GET("/download/:docId", h.getDocById)
			documents.GET("/request/:request_id", h.getDocumentsByRequestId)
			documents.GET("/:docId/ticket", h.getTicketById)
		}

		v2.POST("/sgtin/validate", h.validateSgtin)

		reestr := v2.Group("/reestr")
		{
			sgtin := reestr.Group("/sgtin")
			{
				sgtin.POST("/filter", h.getFilteredReestrKIZ)
				sgtin.POST("/sgtins-by-list", h.getReestrKizByList)
				sgtin.POST("/public/sgtins-by-list", h.getPublicReestrKizByList)
				sgtin.POST("/public/archive/sgtins-by-list", h.getPublicArchReestrKizByList)
				sgtin.GET("/:sgtin", h.getDetailInfoSgtin)
				sgtin.GET("/archive/:sgtin", h.getDetailArchInfoSgtin)
				sgtin.GET("/documents", h.getDocumentsBySgtin)
			}
			sscc := reestr.Group("/sscc")
			{
				sscc.GET("/:sscc/hierarchy", h.getSsccHierarchy)
				sscc.POST("/:sscc/sgtins", h.getKIZInfoFromSscc)
				sscc.GET("/full-hierarchy", h.getSsccFullHierarchy)
			}
			partnersTrusted := reestr.Group("/trusted_partners")
			{
				partnersTrusted.POST("/add", h.addTrustedPartners)
				partnersTrusted.POST("/delete", h.deleteTrustedPartners)
				partnersTrusted.POST("/filter", h.getFilterTrustedPartners)
			}

			//TODO Добавить метод 7.10.1. Фильтрация по реестру ЕСКЛП
			reestr.POST("/esklp/filter", h.getEsklp)
		}
		partners := v2.Group("/reestr_partners")
		{
			partners.POST("/filter", h.getFilterReestrPartners)
		}
	}
	return router
}
