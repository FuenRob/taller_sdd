package main
import ("log";"net/http";"os";"time")
func main() {
 mux:=http.NewServeMux()
 mux.HandleFunc("GET /health",healthHandler)
 mux.HandleFunc("GET /api/products",productsHandler)
 server:=&http.Server{Addr:":8080",Handler:mux,ReadHeaderTimeout:5*time.Second}
 log.Printf("listening on http://localhost%s",server.Addr)
 if err:=server.ListenAndServe();err!=nil&&err!=http.ErrServerClosed {log.New(os.Stderr,"server: ",0).Fatal(err)}
}
