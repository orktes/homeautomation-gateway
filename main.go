package main

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.ListenAndServe(":"+os.Getenv("PORT"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, _ := httputil.DumpRequest(r, true)
		println(d)
		io.WriteString(w, "TO BE DONE")
	}))
}
