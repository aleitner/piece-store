package main
import (
  "github.com/aleitner/piece-store/pkg"
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
  "log"
  "html/template"
  "strings"
  "strconv"
)

const dataDir = "./piece-store-data"

func UploadFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

  // in your case file would be fileupload
  file, header, err := r.FormFile("uploadfile")
  if err != nil {
    fmt.Printf("Error: ", err.Error())
    return
  }

  // We have to do stupid conversions
  // Is there a better way to convert []string into ints?
  r.ParseForm()
  var dataSize int64
  dataSizeStr := strings.Join(r.Form["size"], "")
  dataSizeInt64, _ := strconv.ParseInt(dataSizeStr, 10, 64);
  if dataSizeStr == "" || dataSizeInt64 <= 0 {
    dataSize = header.Size
  } else {
    dataSize = dataSizeInt64
  }

  var dataOffset int64
  dataOffsetStr := strings.Join(r.Form["offset"], "")
  dataOffsetInt64, _ := strconv.ParseInt(dataOffsetStr, 10, 64);
  if dataOffsetStr == "" && dataOffsetInt64 <= 0 {
    dataOffset = 0
  } else {
    dataOffset = dataOffsetInt64
  }

  defer file.Close()
  fmt.Printf("Uploading file (%s), Offset: (%v), Size: (%v)...\n", header.Filename, dataOffset, dataSize)

  hash := String(20)
  err = pstore.Store(hash, file, dataSize, dataOffset, dataDir)

  if err != nil {
    fmt.Printf("Error: ", err.Error())
    return
  }

  fmt.Printf("Successfully uploaded file %s...\n", header.Filename)


  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  message := fmt.Sprintf("Successfully uploaded file!\nName: %s\nHash: %s\nSize: %v\n", header.Filename, hash, header.Size)
  message = fmt.Sprintf("%s\n<a href=\"/files/\">List files</a>", message)
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

  tmpl, err := template.ParseFiles(path)
  if err != nil {
      return err
  }

  tmpl.Execute(w, "")
  return nil
}

func main() {
  router := httprouter.New()
  router.GET("/", Index)
  router.GET("/upload", ShowUploadForm)
  router.ServeFiles("/files/*filepath", http.Dir(dataDir))
  router.POST("/upload", UploadFile)
  log.Fatal(http.ListenAndServe(":8080", router))
}
