package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"net/http"
)


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}


func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Choose the folder to serve
	staticDir := "/static/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	return router
}
//
//func main() {
//	port := "8080"
//
//	if port == "" {
//		log.Fatal("$PORT must be set")
//	}
//
//	router := gin.New()
//	router.Use(gin.Logger())
//	router.LoadHTMLGlob("templates/*.tmpl.html")
//	router.Static("/static", "static")
//
//	router.GET("/", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.tmpl.html", nil)
//	})
//	router.GET("/mark", func(c *gin.Context) {
//		c.String(http.StatusOK, "testMark")
//	})
//
//
//	router.Run(":" + port)
//	//handleRequests()
//}
func main() {
	router := NewRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
