package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type pageData struct {
	Files []string
	Name  string
}

var templates = template.Must(template.ParseFiles("tmpl/gallery.html", "tmpl/sub_gallery.html"))

func getDirectoryContents(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	galleryNames, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}
	return galleryNames, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data pageData) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// IndexHandler - home page, greeting
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// GalleryHandler - show links to all subgalleries
func GalleryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	galleryNames, err := getDirectoryContents("public/img/gallery")
	if err != nil {
	}
	data := &pageData{
		Files: galleryNames,
		Name:  "Gallery"}
	renderTemplate(w, "gallery", *data)
}

// SubgalleryHandler - show all images in particular gallery
func SubgalleryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	galleryContents, err := getDirectoryContents("public/img/gallery/" + name)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	data := &pageData{
		Files: galleryContents,
		Name:  name}
	renderTemplate(w, "sub_gallery", *data)
}
