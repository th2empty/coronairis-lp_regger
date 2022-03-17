package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"coronairis-lp_web-app/pkg/utils/system"
	"coronairis-lp_web-app/pkg/utils/vk"
)

type Input struct {
	AccessToken string `json:"access_token" binding:"required"`
}

func (h *Handler) GetIndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}

func (h *Handler) Register(ctx *gin.Context) {
	var input Input

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := vk.GetUserId(input.AccessToken)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
		return
	}

	if err := h.services.Registration.RegisterUser(input.AccessToken, id); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "User already registered")
		return
	}

	if _, err := system.RunScript(restartServiceScript); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError,
			"Registration was successful, but the session could not be started. Contact administrator")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Вы успешно подключены. В течении 2-х минут сможете пользоваться ботом",
	})
}

func (h *Handler) UpdateToken(ctx *gin.Context) {
	var input Input

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := vk.GetUserId(input.AccessToken)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, "Invalid token")
		return
	}

	if err := h.services.Registration.UpdateUser(input.AccessToken, id); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := system.RunScript(restartServiceScript); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError,
			"Token successfully updated, but the session could not be started. Contact administrator")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Токен обновлен. В течении нескольких минут вы сможете начать пользоваться ботом.",
	})
}
