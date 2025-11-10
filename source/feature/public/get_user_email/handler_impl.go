package get_user_email

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	httpresputils "github.com/i-sub135/go-rest-blueprint/source/common/glob_utils/http_resp_utils"
	"github.com/i-sub135/go-rest-blueprint/source/pkg/logger"
)

func (h *Handler) Impl(c *gin.Context) {

	email := c.Query("email")
	if email == "" {
		errMsg := "user email can`t be empty"
		logger.Error().Err(errors.New(errMsg)).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	ctx := c.Request.Context()
	user, err := h.repo.GetByEmail(ctx, email)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	// Extract first name from email (before first dot)
	// Example: "James.Martinez762@outlook.com" -> "James"
	emailParts := strings.Split(email, "@")
	if len(emailParts) == 0 {
		errMsg := "invalid email format"
		logger.Error().Err(errors.New(errMsg)).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	localPart := emailParts[0] // "James.Martinez762"
	nameParts := strings.Split(localPart, ".")
	firstName := nameParts[0] // "James"

	custmer, err := h.repo.GetCustomerFirstName(ctx, firstName)
	if err != nil {
		errMsg := err.Error()
		logger.Error().Err(err).Caller().Msg(errMsg)
		httpresputils.HttpRespBadRequest(c, &errMsg)
		return
	}

	httpresputils.HttpRespOK(c,
		gin.H{
			"user":     user,
			"customer": custmer,
		},
		nil,
	)
}
