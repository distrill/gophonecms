package controllers

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

func only(arr []string) string {
	if len(arr) != 1 {
		panic("Invalid array length")
	}
	return arr[0]
}

func downloadImage(directory string, url string) {
	// create directory if it doesn't already exist
	newpath := filepath.Join(".", "public", "img", "gallery", directory)
	os.MkdirAll(newpath, os.ModePerm)
	fmt.Println("Created directory " + newpath)

	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	// all kinds of number string fuckery
	now := time.Now().Unix()
	rand.Seed(now)
	nowString := strconv.FormatInt(now, 10)
	randString := strconv.FormatInt(rand.Int63(), 10)
	fileString := nowString + "_" + randString + ".jpg"

	//open a file for writing
	newFile := filepath.Join(newpath, fileString)
	file, err := os.Create(newFile)
	if err != nil {
		log.Fatal(err)
	}

	// Use io.Copy to dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Successfully downloaded image " + newFile)
}

// MessageHandler - handle incoming messages
func MessageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("message handler")

	r.ParseForm()
	body := r.Form

	// image urls and directory
	var mediaUrls []string
	directory := "misc"
	for key, value := range body {
		if strings.Contains(key, "MediaUrl") {
			mediaUrls = append(mediaUrls, only(value))
		} else if key == "Body" && only(value) != "" {
			directory = strings.Replace(strings.ToLower(only(value)), " ", "_", -1)
		}
	}

	// debug. remove me pls
	fmt.Print("Media urls: ")
	fmt.Println(mediaUrls)
	fmt.Println("directory: " + directory)

	// put it in specified directory
	for _, url := range mediaUrls {
		downloadImage(directory, url)
	}

}
