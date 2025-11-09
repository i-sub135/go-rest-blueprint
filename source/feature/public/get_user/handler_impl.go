package getuser

import (
	"strconv"

	"github.com/gin-gonic/gin"
	httpresputils "github.com/i-sub135/go-rest-blueprint/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
	"github.com/i-sub135/go-rest-blueprint/source/service/constant"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		errMsg := "Invalid user ID"
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	requestID := c.GetString(constant.RequestIDKey)

	h.repo.LogUserAccess(ctx, uint(id), requestID)

	user, err := h.repo.GetByID(ctx, uint(id))
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	httpresputils.HttpRespOK(c, user, nil)
}
