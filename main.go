package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"math/rand"

	"github.com/jamiebarnett/gobackground/response"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

const (
	clientID     = "b8e7c700488326b"
	clientSecret = "f429c3cacc1ac1affd7f25691cb0d6e477e3963c"
)

func main() {

	url := "https://api.imgur.com/3/gallery/r/wallpaper/top/day"

	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("authorization", "Client-ID " + clientID)

	rs, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Body.Close()

	body := new(response.Images)
	err = json.NewDecoder(rs.Body).Decode(body)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	ran := rand.Intn(len(body.Data))

	var link string

	for i, inst := range body.Data {
		if i == ran {
			link = inst.Link
			log.Println(link)
		}
	}

	imgrq, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal("error creating image request ",err)
	}

	imgrs, err := http.DefaultClient.Do(imgrq)
	if err != nil {
		log.Fatal("error executing image request ",err)
	}
	defer imgrs.Body.Close()

	file, err := os.Create("img/back.jpg")
	if err != nil {
		log.Fatal("error creating file ", err)
	}

	_, err = io.Copy(file, imgrs.Body)
	if err != nil {
		log.Fatal("error copying file ", err)
	}
	file.Close()

	//works for OSX sierra
	err = exec.Command("bash", "-c",
		"sqlite3 ~/Library/Application\\ Support/Dock/desktoppicture.db \"update data set value = '~/go/src/github.com/jamiebarnett/gobackground/img/back.jpg'\" && killall Dock").Run()
	if err != nil {
		log.Fatal("error executing command ",err)
	}
}