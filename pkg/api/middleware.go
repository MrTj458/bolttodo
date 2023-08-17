package api

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (api *Api) recoverPanic(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				if api.Debug {
					api.Logger.Error("internal server error", "err", err, "stack", string(stack))
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "<h1>Panic: %v</h1><pre>%s</pre>", err, string(stack))
					return
				}

				api.internalErrorResponse(w, r, "err", err, "stack", stack)
			}
		}()

		next.ServeHTTP(w, r)
	}
}
