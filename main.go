package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jamiebarnett/gobackground/response"

	"fmt"
	"os/exec"
	"strconv"
)

const (
	clientID     = "b8e7c700488326b"
	clientSecret = "f429c3cacc1ac1affd7f25691cb0d6e477e3963c"
)

func main() {

	url := "https://api.imgur.com/3/gallery/r/wallpaper/top/day"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("authorization", "Client-ID "+clientID)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body := new(response.Images)
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		log.Fatal(err)
	}

	for i, img := range body.Data {

		if img.Nsfw != true {

			req, err := http.NewRequest("GET", img.Link, nil)
			if err != nil {
				log.Fatal(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			file, err := os.Create("img/" + strconv.Itoa(i) + ".jpg")
			if err != nil {
				log.Fatal("error creating file", err)
			}

			_, err = io.Copy(file, res.Body)
			if err != nil {
				log.Fatal(err)
			}

			res.Body.Close()
			file.Close()
		}
	}

	out, err := exec.Command("switchbg.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output is %s\n", out)
}
