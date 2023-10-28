package main

import (
	"fmt"
	"net/http"

	"github.com/faiz-gh/go-lang/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go Api Service")

	fmt.Println(`
 #####    #####              ###    ######    ######  
##   ##  ### ###            ## ##    ##  ##     ##    
##       ##   ##           ##   ##   ##  ##     ##    
## ####  ##   ##           ##   ##   #####      ##    
##   ##  ##   ##           #######   ##         ##    
##   ##  ### ###           ##   ##   ##         ##    
 #####    #####            ##   ##  ####      ###### `)

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
	}
}
