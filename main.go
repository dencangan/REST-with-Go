package main
 
import (
    "fmt"
    "log"
    "net/http"
    //"github.com/gorilla/mux"
    "encoding/json"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}


func mainPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
 
func pageOne(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is pageOne!")
}
 
func returnArticles(w http.ResponseWriter, r *http.Request) {
    
    articles := []Article{
        Article{Title: "BookOne", Desc: "BookOne summary", Content: "BookOne content"},
        Article{Title: "BookTwo", Desc: "BookTwo summary", Content: "BookTwo content"},
    }

    json.NewEncoder(w).Encode(articles)
}


func main() {

    // Serving static page
    fs := http.FileServer(http.Dir("../src/static"))
    http.Handle("/", fs)

    log.Println("Listening on :8080...")
    http.ListenAndServe(":8080", nil)
	// router := mux.NewRouter().StrictSlash(true)
    // // Home and other pages
    // // router.HandleFunc("/", mainPage)
    // // router.HandleFunc("/pageone", pageOne)

    // // // Custom page
    // // router.HandleFunc("/articles", returnArticles)
    
    // log.Fatal(http.ListenAndServe(":8080", router))
}
 