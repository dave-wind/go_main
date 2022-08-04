package middleware

import (
	"context"
	"net/http"
	"time"
)

// TimeoutMiddleware...
type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 如果中间件是唯一中间件
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}
	// 从请求里拿到 当前的 context
	ctx := r.Context()
	// 修改当前Context,给其设置超时
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	// 替换Context
	r.WithContext(ctx)

	// 新建一个channel, 其意图: 如果这个请求能在3秒内完成的话,channel将会收到一个信号:空struct: struct{}
	ch := make(chan struct{})

	go func() {
		tm.Next.ServeHTTP(w, r)
		// 发送信号, struct{}是类型
		ch <- struct{}{}
	}()

	// 设置一个竞争·关系:
	select {
	// 如果先收到空的channel 就表示请求在3秒内完成
	case <-ch:
		return
	case <-ctx.Done(): // 表示超时, 会触发ctx.Done(),其会返回一个 channel
		w.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
