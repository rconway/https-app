package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("STARTING...")

	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /")
		w.Write([]byte(`
			<body>
				<h2 style="text-align: center;">SUCCESS : Web App Go</h2>
				<p><a href="data">DATA</a></p>
				<p><a href="info">INFO</a></p>
			</body>
		`))
	})

	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /index")
		http.Redirect(w, r, "/data", 307)
	})

	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /data")
		//w.WriteHeader(204)
		w.Write([]byte(`
			<body>
				<h2 style="text-align: center;">You got some DATA</h2>
				<p><a href="./">HOME</a></p>
				<p><a href="info">INFO</a></p>
			</body>
		`))
	})

	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /info")
		//w.WriteHeader(204)
		w.Write([]byte(`
			<body>
				<h2 style="text-align: center;">You got some INFO</h2>
				<p><a href="./">HOME</a></p>
				<p><a href="data">DATA</a></p>
			</body>
		`))
	})

	fmt.Println("LISTEN on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))

	fmt.Println("EXITING...")
}
