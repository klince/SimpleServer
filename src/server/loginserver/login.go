package loginserver

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Println("loginserver inited.")

}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	switch req.URL.Path {
	case "/getlist":
		fmt.Fprintf(w, "%s: %s\n", "ip", "haha")
	case "/login":
		usr := req.URL.Query().Get("usr")
		pwd := req.URL.Query().Get("pwd")
		ok := pwd == "123"
		if !ok {
			//w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "pwd not accepted. usr=%s\n", usr)
			return
		}
		fmt.Fprintf(w, "ok\n")
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func StartListen() {
	th := http.HandlerFunc(httpHandler)

	// We use http.Handle instead of mux.Handle...
	http.Handle("/", th)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
