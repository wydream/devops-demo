package status

import (
	"github.com/gin-gonic/gin"
	"github.com/wydream/devops-demo/web/ctrls"
	"net/http"
)

func GetOk(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func init() {
	ctrls.Register("status", func(r gin.IRouter) {
		r.Any("/ok", GetOk)
		r.Any("/ok.htm", GetOk)
		r.Any("/ok.html", GetOk)
	})
}
