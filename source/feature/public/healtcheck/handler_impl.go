package healtcheck

import (
	"github.com/gin-gonic/gin"
	httpresputils "github.com/i-sub135/go-rest-blueprint/source/service/share/glob_utils/http_resp_utils"
)

func (h *Handler) HealtCheck(c *gin.Context) {

	if err := checkConnection(h.db, c.Request.Context()); err != nil {
		errMsg := err.Error()
		httpresputils.HttpRespBadGateway(c, &errMsg, nil)
		return
	}

	msg := "db connect ok"
	httpresputils.HttpRespOK(c, nil, &msg)

}
