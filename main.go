package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"

	"time"

	"github.com/jamiebarnett/gobackground/response"
	_ "github.com/mattn/go-sqlite3"
)

const (
	clientID     = "b8e7c700488326b"
	clientSecret = "f429c3cacc1ac1affd7f25691cb0d6e477e3963c"
)

func main() {

	url := "https://api.imgur.com/3/gallery/r/wallpaper/top/day?mature=false"

	rq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rq.Header.Add("authorization", "Client-ID "+clientID)

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
			link = inst.Images[0].Link
		}
	}

	log.Println(link)

	imgrq, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal("error creating image request ", err)
	}

	imgrs, err := http.DefaultClient.Do(imgrq)
	if err != nil {
		log.Fatal("error executing image request ", err)
	}
	defer imgrs.Body.Close()

	if _, err := os.Stat("img"); os.IsNotExist(err) {
		os.Mkdir("img", 0777)
	}

	file, err := os.Create("./img/back.jpg")
	if err != nil {
		log.Fatal("error creating file ", err)
	}

	_, err = io.Copy(file, imgrs.Body)
	if err != nil {
		log.Fatal("error copying file ", err)
	}
	defer file.Close()

	//works for OSX sierra
	command := fmt.Sprintf(
		"sqlite3 ~/Library/Application\\ Support/Dock/desktoppicture.db \"update data set value = '%s/src/github.com/jamiebarnett/gobackground/img/back.jpg'\" && killall Dock",
		os.Getenv("GOPATH"),
	)

	log.Println(command)
	err = exec.Command("bash", "-c", command).Run()
	if err != nil {
		log.Fatal("error executing command ", err)
	}
}
