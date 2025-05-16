package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api", middlewareApiKey(apiHandler()))
	log.Fatal(http.ListenAndServe(":1234", mux))
}

func middlewareApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("api-key")
		h := md5.New()
		h.Write([]byte(apiKey))
		hStr := hex.EncodeToString(h.Sum(nil))
		if hStr == "5f4dcc3b5aa765d61d8327deb882cf99" {
			ctx := context.WithValue(r.Context(), "api-user-id", "BloomFilter")
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte("403 Forbidden"))
	})
}

func apiHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := r.Context().Value("api-user-id").(string)
		_, _ = w.Write([]byte(rid))
	}
}
