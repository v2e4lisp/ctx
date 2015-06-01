package main

import (
        "fmt"
        "log"
        "net/http"

        "github.com/v2e4lisp/ctx"
)

func SetURIHandler(h http.Handler) http.Handler {
        f := func(w http.ResponseWriter, r *http.Request) {
                ctx.For(r).Set("uri", r.URL.RequestURI())
                h.ServeHTTP(w, r)
        }
        return http.HandlerFunc(f)
}

func handler(w http.ResponseWriter, r *http.Request) {
        uri, _ := ctx.For(r).Get("uri")
        fmt.Fprintln(w, uri.(string))
}

func main() {
        h := ctx.Handler(SetURIHandler(http.HandlerFunc(handler)))
        http.Handle("/", h)
        log.Fatal(http.ListenAndServe(":8080", nil))
}
