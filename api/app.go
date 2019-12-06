package api

import (
	"MVC/controller"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApp() {

	fmt.Println("Servar Start : 8080")
	route := mux.NewRouter()

	route.HandleFunc("/user", controller.GetUser).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", route); err != nil {
		panic(err)
	}
}
