package routes

import (
	"fmt"
	"net/http"
)

// Don't really need this - Just messing with composition/promotion with interfaces.
type FredbobHandler struct {
	http.ServeMux
}

func NewFredbobHandler() http.Handler {
	subHandler := &FredbobHandler{}
	subHandler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[fredbob] hit /")
		w.Write([]byte("zzz this is /fredbob/ zzz"))
	})
	subHandler.HandleFunc("/barry", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[fredbob] hit /barry")
		w.Write([]byte(`
			<body>
				<h2>[fredbob]</h2>
				this is /barry
			</body>
		`))
	})
	return subHandler
}
