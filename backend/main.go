package main

import (
	"context"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"goLearn/gateway"
	"goLearn/utils"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	// 将gin的log添加到zap日志中管理
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(utils.Logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(utils.Logger, true))
	if !utils.CONFIG.GetBool("DebugMode") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	gateway.InitRoutines(r)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(utils.CONFIG.GetInt("port")),
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			utils.Infof("TerminalMgr Service listen: %s\n", err)
		}
	}()

	// 安全的结束进行
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	utils.Infof("Shutdown TerminalMgr Service ...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Infof("TerminalMgr Service Shutdown:\n", err)
	}
}
