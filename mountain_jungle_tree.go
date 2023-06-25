//Sustainable Living Blog
package main

import (
    "fmt"
    "time"
    "strings"
    "net/http"
    "html/template"
)

//Global Variables
var Layout string

//Structs
type Post struct {
    Title string
    Content string
    PublishedAt time.Time
    Author string
}

//Functions
func init() {
    Layout = "02-January-2006"
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    t, err := template.ParseFiles("home.html")
    if err != nil {
        fmt.Fprintf(w, "Error parsing files: %v", err)
        return
    }

    err = t.Execute(w, nil)
    if err != nil {
        fmt.Fprintf(w, "Error executing template: %v", err)
        return
    }
}

func postHandler(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) != 3 {
        http.NotFound(w, r)
        return
    }

    post, err := readPost(parts[2])
    if err != nil {
        http.NotFound(w, r)
        return
    }

    t, err := template.ParseFiles("post.html")
    if err != nil {
        fmt.Fprintf(w, "Error parsing files: %v", err)
        return
    }

    err = t.Execute(w, post)
    if err != nil {
        fmt.Fprintf(w, "Error executing template: %v", err)
        return
    }
}

func readPost(postName string) (*Post, error) {
    post := &Post{
        Title: postName,
        Content: "This is some sample content for " + postName,
        PublishedAt: time.Now(),
        Author: "John Doe",
    }
    return post, nil
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/post/", postHandler)

    http.ListenAndServe(":8080", nil)
}