package main

import (
  "io/ioutil"
  "net/http"
  "os"
  "fmt"
  "bufio"
)

func buildFilePath(r *http.Request) (path string) {
  return "/tmp/go/" + r.URL.Path[3:]
}

func readHandler(w http.ResponseWriter, r *http.Request) {
  content, err := ioutil.ReadFile(buildFilePath(r))
  if err != nil {
    http.Error(w, err.Error(), http.StatusNotFound)
    return
  }
  w.Write(content)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func writeBodyToFile(completeFilePath string, body string) {
  file, err := os.Create(completeFilePath)
  check(err)
  defer file.Close()

  writer := bufio.NewWriter(file)
  defer writer.Flush()

  fmt.Fprintln(writer, body)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  if err != nil {
    http.Error(w, http.StatusText(http.StatusNotFound), 404)
    return
  }
  writeBodyToFile(buildFilePath(r), string(body))
}

func main() {
  http.HandleFunc("/r/", readHandler)
  http.HandleFunc("/w/", writeHandler)
  http.ListenAndServe(":8080", nil)
}

