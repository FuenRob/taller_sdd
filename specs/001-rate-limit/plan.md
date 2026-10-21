# Plan técnico

1. Crear middleware `rateLimit(next, limiter)` envolviendo `ServeMux`.
2. Implementar `map[string]bucket` protegido por `sync.Mutex` (solo una instancia).
3. Inyectar reloj `now func() time.Time` para pruebas deterministas.
4. Extraer IP desde `r.RemoteAddr` con `net.SplitHostPort`, sin confiar en `X-Forwarded-For`.
5. Eximir `/health` y devolver 429 JSON con `Retry-After` al superar el límite.
6. Ejecutar `go test -race ./...` y verificar con `curl`.

## Riesgos y decisiones
- Las ventanas fijas permiten ráfagas en los límites.
- El mapa crece sin eviction de clientes antiguos.
- Múltiples instancias no comparten contadores.
- Con reverse proxy todos los clientes podrían compartir IP.
- No usar un reloj mutable compartido en tests concurrentes.
