package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanexu/logn"
	"github.com/spf13/viper"
	"github.com/wydream/devops-demo/web/ctrls"
	_ "github.com/wydream/devops-demo/web/ctrls/health"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 启动http server, 这个方法会block.
func Run() {
	// viper 会从环境变量中读取，viper无视大小写
	appPort := viper.GetInt("app_port")
	//env := viper.GetString("ENV")
	if appPort == 0 {
		appPort = 8088
	}

	r := gin.New()
	//if env == "production" {
	//	gin.SetMode(gin.ReleaseMode)
	//}
	log := logn.GetLogger("web")
	ctrls.Init(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appPort),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("server shutdown: %s", err)
	}
}
