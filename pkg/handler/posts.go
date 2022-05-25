package handler

import (
	"net/http"
	"twittie"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return 
	}

	var input twittie.Post
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Post.Create(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllPosts(c *gin.Context) {
	
}

func (h *Handler) getPostById(c *gin.Context) {

}

func (h *Handler) deletePost(c *gin.Context) {

}