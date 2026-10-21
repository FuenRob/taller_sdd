package main
import ("net/http";"net/http/httptest";"strconv";"sync";"testing";"time")
func fixture()(*limiter,*time.Time){
 now:=time.Date(2026,1,1,0,0,0,0,time.UTC)
 return newLimiter(3,time.Minute,func()time.Time{return now}),&now
}
func request(h http.Handler,path,address string)*httptest.ResponseRecorder{
 req:=httptest.NewRequest(http.MethodGet,path,nil)
 req.RemoteAddr=address
 rr:=httptest.NewRecorder()
 h.ServeHTTP(rr,req)
 return rr
}
func TestThreeThen429WithRetryAfter(t *testing.T){
 l,_:=fixture();h:=rateLimit(http.HandlerFunc(productsHandler),l)
 for i:=0;i<3;i++ {if rr:=request(h,"/api/products","192.0.2.1:1234");rr.Code!=200{t.Fatalf("request %d: %d",i+1,rr.Code)}}
 rr:=request(h,"/api/products","192.0.2.1:4567")
 if rr.Code!=429{t.Fatalf("want 429, got %d",rr.Code)}
 if rr.Header().Get("Retry-After")!="60"{t.Fatalf("retry-after: %q",rr.Header().Get("Retry-After"))}
}
func TestDifferentClientAndHealthExemption(t *testing.T){
 l,_:=fixture();mux:=http.NewServeMux()
 mux.HandleFunc("GET /health",healthHandler)
 mux.HandleFunc("GET /api/products",productsHandler)
 h:=rateLimit(mux,l)
 for i:=0;i<4;i++{request(h,"/api/products","192.0.2.1:1234")}
 if rr:=request(h,"/api/products","192.0.2.2:1234");rr.Code!=200{t.Fatalf("other IP: %d",rr.Code)}
 if rr:=request(h,"/health","192.0.2.1:1234");rr.Code!=200{t.Fatalf("health: %d",rr.Code)}
}
func TestWindowResets(t *testing.T){
 now:=time.Date(2026,1,1,0,0,0,0,time.UTC)
 l:=newLimiter(1,time.Minute,func()time.Time{return now})
 h:=rateLimit(http.HandlerFunc(productsHandler),l)
 if rr:=request(h,"/api/products","192.0.2.1:10");rr.Code!=200{t.Fatal(rr.Code)}
 if rr:=request(h,"/api/products","192.0.2.1:10");rr.Code!=429{t.Fatal(rr.Code)}
 now=now.Add(time.Minute)
 if rr:=request(h,"/api/products","192.0.2.1:10");rr.Code!=200{t.Fatal(rr.Code)}
}
func TestConcurrentLimit(t *testing.T){
 l:=newLimiter(3,time.Minute,time.Now)
 var wg sync.WaitGroup
 results:=make(chan int,40)
 for i:=0;i<40;i++{wg.Add(1);go func(n int){defer wg.Done();results<-request(rateLimit(http.HandlerFunc(productsHandler),l),"/api/products","192.0.2.1:"+strconv.Itoa(1000+n)).Code}(i)}
 wg.Wait();close(results)
 allowed:=0
 for code:=range results{if code==200{allowed++}else if code!=429{t.Errorf("unexpected %d",code)}}
 if allowed!=3{t.Fatalf("allowed %d, want 3",allowed)}
}
