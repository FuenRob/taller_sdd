# Especificación funcional: limitar tráfico HTTP

## Problema
Un consumidor puede enviar demasiadas peticiones a `/api/products` y afectar a otros. Queremos una protección observable y verificable para una sola instancia.

## Requisitos y criterios de aceptación

1. Por IP de `RemoteAddr`, como máximo 3 peticiones a `/api/products` en una ventana fija de 60 segundos. La cuarta devuelve 429.
2. Al comenzar una nueva ventana se admiten otras 3 peticiones.
3. Al superar el límite, devolver HTTP 429, `Content-Type: application/json`, cuerpo `{"error":"rate_limit_exceeded"}` y `Retry-After` en segundos redondeados al alza.
4. `/health` no está limitado.
5. Cada IP dispone de cupo independiente.
6. Concurrencia: no admitir más de 3 peticiones simultáneas por IP en la misma ventana.

## Fuera de alcance
Redis, rate limit distribuido, autenticación, limpieza de claves antiguas, límites configurables, encabezados forwarded de proxies. No apto para producción multiinstancia.

## Verificación
Pruebas de límite, ventana, aislamiento por IP, excepción de salud y `go test -race ./...`.

## Preguntas en directo
- ¿IP, usuario o API key?
- ¿Ventana fija, deslizante o token bucket?
- ¿Cuántas réplicas?
- ¿Hay reverse proxy y qué encabezados consideramos fiables?
