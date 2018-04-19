package main
import (
  "github.com/aleitner/piece-store/pkg"
  "fmt"
  "net/http"
  "strings"
)

const dataDir = "./piece-store-data"

const AddForm = `
<form enctype="multipart/form-data" action="http://127.0.0.1:8080/upload" method="post">
{{/* 1. File input */}}
<input type="file" name="uploadfile" />

{{/* 2. Submit button */}}
<input type="submit" value="upload file" />
</form>
`

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
  if r.Method == "GET" {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprint(w, AddForm)

  } else if r.Method == "POST" {
    // in your case file would be fileupload
    file, header, err := r.FormFile("uploadfile")
    if err != nil {
      fmt.Printf("Error: ", err.Error())
      return
    }

    defer file.Close()
    fmt.Printf("Uploading file %s...", header.Filename)

    hash := String(20)
    err = pstore.Store(hash, file, header.Size, 0, dataDir)

    if err != nil {
      fmt.Printf("Error: ", err.Error())
    }

    message := fmt.Sprintf("Successfully uploaded file!\nName: %s\nHash: %s\nSize: %v\n", header.Filename, hash, header.Size)
    w.Write([]byte(message))


  } else {
      fmt.Println("Unknown HTTP " + r.Method + "  Method")
  }
}

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}


func main() {
  fmt.Println("Starting server...")
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/upload", ReceiveFile)

  fs := http.FileServer(http.Dir(dataDir))
  http.Handle("/files/", http.StripPrefix("/files", fs))

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
