package handlers

import (
	"github.com/No1ball/biocadTask/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetPageData(c *gin.Context) {
	pageStr, isNotFalse := c.GetQuery("page")
	if isNotFalse == false {

	}
	limitStr, isNotFalse := c.GetQuery("limit")
	if isNotFalse == false {

	}
	page, errPage := strconv.Atoi(pageStr)
	limit, errLimit := strconv.Atoi(limitStr)
	if errPage != nil || errLimit != nil {

	}
	offset := (page - 1) * limit
	data, err := h.service.GetDataFromPage(offset, limit)
	if err != nil {

	}
	c.JSON(http.StatusOK, data)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/:unit_guid", h.GetPageData)
	return router
}
