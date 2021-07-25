package core

import (
	"fmt"
	"net/http"
	"self-discipline/global"
	"self-discipline/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() *http.Server {

	router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	srv := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	SELF-DISCIPLNE服务已启动
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)

	return srv
}
