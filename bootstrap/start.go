package bootstrap

import (
	"codetube.cn/interface-web/components"
	"codetube.cn/interface-web/routes/v1"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var BootErrChan chan error

func Start() {
	BootErrChan = make(chan error)
	go func() {
		//加载各版本的 API 路由
		v1.ApiRouter.Load(v1.LoadRoutes...)

		components.RouterEngine.Run(":8080")
	}()
	//监听事件
	go func() {
		sigC := make(chan os.Signal)
		signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)
		BootErrChan <- fmt.Errorf("%", <-sigC)
	}()
	getErr := <-BootErrChan
	log.Println(getErr)
}
