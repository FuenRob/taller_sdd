# Taller práctico: SDD + Go + IA

Repositorio didáctico para una sesión sobre **Spec-Driven Development**: convertir una petición vaga en especificación, plan, tareas, código y pruebas.

## Estructura

- `starter/`: API HTTP minimalista sin rate limiting. **Punto de partida** para la demo.
- `specs/001-rate-limit/`: especificación funcional, plan técnico y tareas.
- `solution/`: implementación de referencia separada (no mirar durante la demo).
- `docs/GUIA_FORMADOR.md`: guion, tiempos y prompts para la formación.

## Requisitos

Go 1.22+, sin dependencias externas. Spec Kit es opcional.

```bash
cd starter
go test ./...
go run .
# en otra terminal:
curl -i http://localhost:8080/health
curl -i http://localhost:8080/api/products
```

## Objetivo

Un usuario solicita: «Protege nuestra API para que un cliente no pueda saturarla». Antes de programar aclaramos rutas, identidad del cliente, límite, respuesta HTTP, persistencia, despliegue y criterios de aceptación.

Revisa `specs/001-rate-limit/{spec,plan,tasks}.md`. Implementa las tareas en una copia de `starter/` y contrasta con `solution/`:

```bash
cd solution
go test -race ./...
go run .
# repetir cuatro veces:
for i in 1 2 3 4; do curl -i http://localhost:8080/api/products; done
```

**Aviso:** usa contador en memoria, ventanas fijas, sin limpieza de IPs antiguas ni persistencia. No es una implementación apta para producción multiinstancia. Se identifica al cliente por `RemoteAddr`, sin confiar en cabeceras de proxy no verificadas.

## Referencias

- [GitHub Spec Kit](https://github.com/github/spec-kit)
- [OpenSpec](https://github.com/Fission-AI/OpenSpec)
