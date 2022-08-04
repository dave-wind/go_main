package middleware

import "net/http"

type AuthMiddleware struct {
	Next http.Handler
}

// 返回不可以是其他类型
func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	// 简单的判断鉴权
	auth := r.Header.Get("Authorization")
	if auth != "" {
		// 相当于 callback 执行下一个 Hanlder
		am.Next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
