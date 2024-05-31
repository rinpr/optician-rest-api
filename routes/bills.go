package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"optician-rest-api/controllers"
)

var RegisterImageDataRoutes = func(router *mux.Router) {
	router.HandleFunc("/bills-data", controllers.CreateBillData).Methods("POST")
	router.HandleFunc("/bills-data", controllers.GetBillsData).Methods("GET")
	router.HandleFunc("/bills-data/{billId}", controllers.GetBillData).Methods("GET")
	router.HandleFunc("/bills-data/{billId}", controllers.UpdateBillData).Methods("PUT")
	router.HandleFunc("/bills-data/{billId}", controllers.DeleteBillData).Methods("DELETE")
}

func ListenAndServe() {
	fmt.Println("Listen and serve on port 8080!")
	r := mux.NewRouter()
	RegisterImageDataRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
