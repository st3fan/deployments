package main

import (
  "embed"
  "log"
	"net/http"
  "os"
	"text/template"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    data := map[string]string {
      "Region": os.Getenv("FLY_REGION"),
      "Environment": os.Getenv("ENVIRONMENT"),
    }
    t.ExecuteTemplate(w, "index.html.tmpl", data)
  })

  log.Println("Listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, nil))
}
