package main
import ("encoding/json";"fmt";"net";"net/http";"strings";"sync";"time")
type bucket struct {count int;windowStart time.Time}
type limiter struct {mu sync.Mutex;clients map[string]bucket;maxRequests int;window time.Duration;now func()time.Time}
func newLimiter(max int,window time.Duration,now func()time.Time)*limiter {
 return &limiter{clients:make(map[string]bucket),maxRequests:max,window:window,now:now}
}
func (l *limiter) allow(key string)(bool,time.Duration){
 l.mu.Lock();defer l.mu.Unlock()
 now:=l.now()
 b,ok:=l.clients[key]
 if !ok||now.Sub(b.windowStart)>=l.window {b=bucket{windowStart:now}}
 if b.count>=l.maxRequests{return false,l.window-now.Sub(b.windowStart)}
 b.count++;l.clients[key]=b
 return true,0
}
func rateLimit(next http.Handler,l *limiter)http.Handler{
 return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
  if r.URL.Path=="/health" {next.ServeHTTP(w,r);return}
  // Never trust unconfigured forwarded-IP headers.
  ip,_,err:=net.SplitHostPort(r.RemoteAddr)
  if err!=nil {ip=r.RemoteAddr}
  ip=strings.Trim(ip,"[]")
  allowed,retry:=l.allow(ip)
  if !allowed {
   seconds:=int(retry.Seconds())
   if retry>time.Duration(seconds)*time.Second {seconds++}
   if seconds<1 {seconds=1}
   w.Header().Set("Retry-After",fmt.Sprintf("%d",seconds))
   w.Header().Set("Content-Type","application/json")
   w.WriteHeader(http.StatusTooManyRequests)
   _=json.NewEncoder(w).Encode(map[string]string{"error":"rate_limit_exceeded"})
   return
  }
  next.ServeHTTP(w,r)
 })
}
