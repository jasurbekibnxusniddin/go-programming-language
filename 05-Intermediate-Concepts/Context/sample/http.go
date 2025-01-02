package sample

import (
	"context"
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			// wrap the context with stuff -- we'll see how soon!
			req = req.WithContext(ctx)
			handler.ServeHTTP(rw, req)
		},
	)
}
func logic(ctx context.Context, info string) (string, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-ctx.Done():
			return
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	data := req.FormValue("data")

	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	rw.WriteHeader(http.StatusOK)

	rw.Write([]byte(result))
}

func Main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
