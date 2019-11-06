package main

import ( 
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	)

type User struct {
	id int64
	name string
}

type Trade struct {
	id int64
	trade_type string
	user User
	symbol string
	shares int
	price float64
	timestamp string
}


func main() {
	var userdata User
	var tradedata Trade
	
	userdata = User{ id: 1, name: "Horacio" }

	tradedata = Trade{ id: 1,
			trade_type: "buy",
			user: userdata,
			symbol: "AC",
			shares: 28,
			price: 162.17,
			timestamp: "2014-06-14 13:13:13" } 


	r := mux.NewRouter()
        r.HandleFunc("/trades/id/:id", func... )

	r.HandleFunc("/trades", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode( tradedata )
	})



}
