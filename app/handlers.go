package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anil1226/go-banking/service"
)

func getTime(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query().Get("tz")
	response := make(map[string]string, 0)
	timezones := strings.Split(vars, ",")

	if len(timezones) <= 1 {
		tz, error := time.LoadLocation(vars)
		if error != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", vars)))
		} else {
			response["current-time"] = time.Now().In(tz).String()
		}
	} else {
		for _, v := range timezones {
			tz, error := time.LoadLocation(v)
			if error != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s", v)))
			} else {
				response[v] = time.Now().In(tz).String()
			}
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Anil", "Durham", "27500"},
		{"Raj", "Cary", "27526"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	//custs := getAllCustomers(w,r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "Post request recieved")
// }

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip"`
}

type CustomerHandlers struct {
	service service.CustomerService
}
