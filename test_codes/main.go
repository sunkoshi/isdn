package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

const ZIP_NAME = "isdn_fn_zip.zip"

func main() {
	filePath := os.Args[1]
	configPath := os.Args[2]

	if filePath == "" {
		filePath = "."
	}

	if configPath == "" {
		filePath = "./isdn_config.json"
	}

	fnMd := GetMetadata(configPath)
	log.Println("METADATA", fnMd)

	zipPath := zipAndGetPath(path.Join(filePath))
	log.Println("zipped files for uploading", zipPath)

	request, err := newfileUploadRequest(fmt.Sprintf("%v", fnMd["isdn_url"]), fnMd, zipPath, fmt.Sprintf("%v", fnMd["token"]))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode != 201 {
			log.Println("function uploaded failed")
			return
		}
		log.Println("function upload successful")
		var pb map[string]map[string]interface{}
		json.Unmarshal(body.Bytes(), &pb)
		log.Println("Function Call Url", fmt.Sprintf("%s/%v", fmt.Sprintf("%v", fnMd["isdn_url"]), pb["data"]["id"]))
	}
	clean(filePath)
}

func GetMetadata(path string) map[string]interface{} {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	var v map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(file, &v)
	if err != nil {
		log.Fatalln(err)
	}
	return v
}

func zipAndGetPath(source string) string {
	target := path.Join(source, ZIP_NAME)
	zipfile, err := os.Create(target)
	if err != nil {
		log.Fatalln(err)
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		log.Fatalln(err)
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatalln(err)
		}

		if info.Name() == baseDir && info.IsDir() {
			return nil
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			log.Fatalln(err)
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})
	return target
}

func clean(filePath string) {
	log.Println("removing temporary files")
	err := os.RemoveAll(path.Join(filePath, ZIP_NAME))
	if err != nil {
		log.Fatalln(err)
	}
}

func newfileUploadRequest(uri string, params map[string]interface{}, filePath string, token string) (*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("code", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, fmt.Sprintf("%v", val))
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req, err
}
