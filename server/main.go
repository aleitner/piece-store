package main
import (
  "github.com/aleitner/piece-store/pkg"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
  "html/template"
)

const dataDir = "./piece-store-data"

func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

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
    return
  }

  message := fmt.Sprintf("Successfully uploaded file!\nName: %s\nHash: %s\nSize: %v\n", header.Filename, hash, header.Size)
  w.Write([]byte(message))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  if err := renderByPath(w, "./server/templates/index.html"); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
}

func ShowUploadForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  if err := renderByPath(w, "./server/templates/uploadform.html"); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
  }
}

func renderByPath(w http.ResponseWriter, path string) error {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")

  fp := http.Dir(path)
  fmt.Println(fp)
  tmpl, err := template.ParseFiles(path)
  if err != nil {
      return err
  }

  tmpl.Execute(w, "")
  return nil
}

func main() {
  fmt.Println("Starting server...")
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/upload", ShowUploadForm)
  router.ServeFiles("/files/*filepath", http.Dir(dataDir))
  router.POST("/upload", UploadFile)
  log.Fatal(http.ListenAndServe(":8080", router))
}
