package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, "hello, world!\n")
}

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	ctx, cancel := context.WithCancel(ctx)
	svr := &http.Server{Addr: ":8080"}
	g.Go(func() error {
		http.HandleFunc("/hello", Handle)
		return svr.ListenAndServe()
	})
	g.Go(func() error {
		<-ctx.Done()
		fmt.Println("http服务停止")
		return svr.Shutdown(ctx)
	})
	g.Go(func() error {
		exitSignals := []os.Signal{os.Interrupt}

		sig := make(chan os.Signal, len(exitSignals))
		signal.Notify(sig, exitSignals...)
		for {
			fmt.Println("监听停止信号")
			select {
			case <-ctx.Done():
				fmt.Println("上下文信号")
				return ctx.Err()
			case <-sig:
				cancel()

			}

		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		fmt.Println("wait 错误: ", err)
	}
}
