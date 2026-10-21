package main

import ("encoding/json";"net/http")

func healthHandler(w http.ResponseWriter,_ *http.Request) {
 w.Header().Set("Content-Type","application/json")
 _ = json.NewEncoder(w).Encode(map[string]string{"status":"ok"})
}
func productsHandler(w http.ResponseWriter,_ *http.Request) {
 w.Header().Set("Content-Type","application/json")
 _ = json.NewEncoder(w).Encode([]map[string]any{{"id":1,"name":"Keyboard"},{"id":2,"name":"Mouse"}})
}
