package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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
	file, header, err := r.FormFile("imagedata")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])

	io.Copy(&buf, file)

	hash := ksuid.New().String()

	fmt.Println(hash)
	accessKey := os.Getenv("DO_ACCESS_KEY")
	secKey := os.Getenv("DO_SECRET_ACCESS_KEY")
	ssl := true

	client, err := minio.New("fra1.digitaloceanspaces.com", accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())
	fileName := fmt.Sprintf("%s.png", hash)
	contentType := "image/png"
	userMetaData := map[string]string{"x-amz-acl": "public-read"}

	fmt.Println(fileName)

	n, err := client.PutObject("hugoimg", fileName, reader, reader.Size(), minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: userMetaData,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
	buf.Reset()

	fmt.Fprintf(w, "%s/%s", "https://img.hugo.gg", fileName)
}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":80", nil)
}
