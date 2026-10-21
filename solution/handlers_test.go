package main
import ("net/http";"net/http/httptest";"testing")
func TestHealth(t *testing.T) {
 rr:=httptest.NewRecorder()
 healthHandler(rr,httptest.NewRequest(http.MethodGet,"/health",nil))
 if rr.Code!=http.StatusOK {t.Fatalf("want 200, got %d",rr.Code)}
}
func TestProducts(t *testing.T) {
 rr:=httptest.NewRecorder()
 productsHandler(rr,httptest.NewRequest(http.MethodGet,"/api/products",nil))
 if rr.Code!=http.StatusOK {t.Fatalf("want 200, got %d",rr.Code)}
}
