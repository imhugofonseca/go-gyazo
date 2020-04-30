package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/minio/minio-go"
	"github.com/segmentio/ksuid"
)

func upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // limit your max input length!
	if err != nil {
		fmt.Fprintf(w, "")
		return
	}
	var buf bytes.Buffer
	file, _, err := r.FormFile("imagedata")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(&buf, file)
	hash := ksuid.New().String()

	host := os.Getenv("BKT_HOST")
	accessKey := os.Getenv("BKT_ACCESS_KEY")
	secKey := os.Getenv("BKT_SECRET_ACCESS_KEY")
	spaceDomain := os.Getenv("BKT_SPACE_DOMAIN")
	bktName := os.Getenv("BKT_NAME")

	client, err := minio.New(host, accessKey, secKey, true)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())
	fileName := fmt.Sprintf("%s.png", hash)
	contentType := "image/png"
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = client.PutObject(bktName, fileName, reader, reader.Size(), minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: userMetaData,
	})
	if err != nil {
		log.Fatal(err)
	}
	buf.Reset()
	fmt.Fprintf(w, "https://%s/%s", spaceDomain, fileName)
}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":80", nil)
}
