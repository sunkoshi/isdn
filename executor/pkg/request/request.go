package request

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Download(url string, at string) error {
	out, err := os.Create(at)
	if err != nil {
		return err
	}
	defer out.Close()
	log.Println("downloading", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
