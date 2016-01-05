package main

import (
	"fmt"
	"net/http"

	"github.com/wkirschbaum/myip"
	"github.com/wkirschbaum/yawn"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, myip.ParseIPFromRemoteAddr(r.RemoteAddr))
	})
	yawn.ListenAndServe(nil)
}
