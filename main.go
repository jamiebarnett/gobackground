package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"math/rand"

	"github.com/jamiebarnett/gobackground/response"
	_ "github.com/mattn/go-sqlite3"
)

const (
	clientID     = "b8e7c700488326b"
	clientSecret = "f429c3cacc1ac1affd7f25691cb0d6e477e3963c"
)

func main() {

	url := "https://api.imgur.com/3/gallery/r/wallpaper/top/day"

	galleryReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	galleryReq.Header.Add("authorization", "Client-ID "+clientID)

	galleryRes, err := http.DefaultClient.Do(galleryReq)
	if err != nil {
		log.Fatal(err)
	}
	defer galleryRes.Body.Close()

	body := new(response.Images)
	err = json.NewDecoder(galleryRes.Body).Decode(body)
	if err != nil {
		log.Fatal(err)
	}

	ran := rand.Intn(len(body.Data))
	log.Println(ran)
	var link string

	for i, inst := range body.Data {
		if i == ran {
			link = inst.Link
			log.Println(link)
		}
	}

	imgReq, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	imgRes, err := http.DefaultClient.Do(imgReq)
	if err != nil {
		log.Fatal(err)
	}
	defer imgRes.Body.Close()

	file, err := os.Create("img/back.jpg")
	if err != nil {
		log.Fatal("error creating file", err)
	}

	_, err = io.Copy(file, imgRes.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	db, err := sql.Open("sqlite3", "/Users/jamiebarnett/Library/Application Support/Dock/desktoppicture.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE data SET value = './img/1.jpg';")
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("bash", "-c", "killall Dock").Run()
	if err != nil {
		log.Fatal(err)
	}

}
