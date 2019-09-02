// main.go
package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
	conf := readConfig()
  listenAt := fmt.Sprintf(":%d", conf.port)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
  })

  log.Printf("Open the following URL in the browser: http://localhost:%d\n", conf.port)
  log.Fatal(http.ListenAndServe(listenAt, nil))
}