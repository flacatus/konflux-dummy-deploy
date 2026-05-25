package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Dummy Deployment</title>
  <style>
    body {
      font-family: system-ui, sans-serif;
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      margin: 0;
      background: #0f172a;
      color: #e2e8f0;
    }
    .card {
      background: #1e293b;
      border-radius: 12px;
      padding: 2rem 3rem;
      box-shadow: 0 4px 24px rgba(0,0,0,0.3);
      text-align: center;
    }
    h1 { font-size: 1rem; color: #94a3b8; margin: 0 0 1rem; }
    p  { font-size: 2rem; margin: 0; }
  </style>
</head>
<body>
  <div class="card">
    <h1>DISPLAY_TEXT</h1>
    <p>{{.}}</p>
  </div>
</body>
</html>`))

func main() {
	text := os.Getenv("DISPLAY_TEXT")
	if text == "" {
		text = "Hello from Dummy Deployment!!!!x!!!!!"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tmpl.Execute(w, text)
	})

	fmt.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
