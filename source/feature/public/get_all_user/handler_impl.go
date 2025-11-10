package get_all_user

import (
	"github.com/gin-gonic/gin"
	httpresputils "github.com/i-sub135/go-rest-blueprint/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
)

func (h *Handler) Impl(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.repo.GetAll(ctx)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	httpresputils.HttpRespOK(c, users, nil)
}
