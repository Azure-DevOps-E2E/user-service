package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"polyglot-shop/user-service/internal/httpx"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repository *Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, ListResponse{Items: h.repository.FindAll()})
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, found := h.repository.FindByID(id)
	if !found {
		httpx.WriteError(c, http.StatusNotFound, "USER_NOT_FOUND", fmt.Sprintf("User %s does not exist", id))
		return
	}

	c.JSON(http.StatusOK, item)
}
