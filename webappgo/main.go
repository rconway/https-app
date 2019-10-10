package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/hashicorp/go-uuid"

	"github.com/rconway/webappgo/routes"
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

	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /time")

		type TimeResult struct {
			Time time.Time `json:"time"`
		}
		timeResult := TimeResult{
			Time: time.Now(),
		}

		var jsonData []byte
		jsonData, err := json.Marshal(timeResult)
		if err != nil {
			log.Println(err)
		}

		w.Write(jsonData)
	})

	mux.HandleFunc("/uuid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /uuid")

		type UUIDResult struct {
			UUID string `json:"uuid"`
		}

		uuidResult := UUIDResult{}
		uuidResult.UUID, _ = uuid.GenerateUUID()

		var jsonData []byte
		jsonData, err := json.Marshal(uuidResult)
		if err != nil {
			log.Println(err)
		}

		w.Write(jsonData)
	})

	// MESSING with subroutes
	mux.Handle("/fredbob/", http.StripPrefix("/fredbob", routes.NewFredbobHandler()))
	// MESSING with subroutes

	fmt.Println("LISTEN on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))

	fmt.Println("EXITING...")
}
