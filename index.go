package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	if op == "" {
		sendHTML(w, r)
		return
	}

	switch op {
	case "text":
		sendText(w, r)
	case "json":
		sendJSON(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func sendHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, content)
}

func sendText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello from Zeit!"))
}

func sendJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := struct {
		Time    string `json:"time"`
		Message string `json:"message"`
		Error   string `json:"error,omitempty"`
	}{
		Time:    time.Now().Format(time.RFC3339),
		Message: "Hello from Zeit!",
		Error:   "",
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const content = `<DOCTYPE html>
<html>
  <head>
    <title>Hello, Zeit!</title>
	<style>
		html {
		   font-family: Arial, Helvetica, sans-serif;
		   background: #003333;
		}
		section {
			background: #226666;
			color: #669999;
			font-size: 3rem;
			border-radius: 1rem;
			padding: 1rem;
			position: absolute;
			top: 50%;
			left: 50%;
			margin-right: -50%;
			transform: translate(-50%, -50%);
		}
		a:link, a:visited {
			text-decoration: none;
			color: #D49A6A;
		}
		a:hover {
			color: #FFD1AA;
		}
	</style>
  </head>
  <body>
    <section>
		<p>Hello, <a href="https://zeit.co">Zeit</a>!</p>
	</section>
  </body>
</html>`
