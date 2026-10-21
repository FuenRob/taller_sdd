# Guía del formador — 25 minutos

**0–3 min · Antes de la IA.** Enseña `starter/main.go` y ejecuta `go test ./...`. Pregunta al grupo qué significa «proteger la API».

**3–8 min · Especificación.** Pide tres aclaraciones, muestra `specs/001-rate-limit/spec.md` e identifica criterios verificables y no objetivos.

**8–12 min · Plan y tareas.** Revisa `plan.md` y `tasks.md`. Señala los riesgos: multiinstancia, proxy, consumo de memoria y ráfagas en ventanas fijas.

**12–19 min · Agente de programación.** Copia `starter/` a un directorio de trabajo. Solicita T01–T06, revisa los cambios y luego pide T07–T09.

**19–22 min · Validación.** Ejecuta `go test -race ./...` y `curl -i`. Si falla algo: ¿era un requisito ambiguo, un bug de implementación o un test incorrecto?

**22–25 min · Debate.** ¿Qué haría falta para producción? Comparar con `solution/` al final.

## Con GitHub Spec Kit
Instala la CLI siguiendo [la documentación oficial](https://github.com/github/spec-kit). En una copia del proyecto configura la integración con tu agente. Recorre constitución → especificación → aclaraciones → plan → tareas → implementación → convergencia. Comprueba los comandos exactos de la versión instalada.

## Prompts para copiar

**Requisitos:** «Actúa como analista funcional. No escribas código. Haz preguntas para convertir “protege la API” en criterios de aceptación medibles. Indica decisiones fuera de alcance».

**Diseño:** «A partir de spec.md, produce un plan mínimo para Go 1.22+ con net/http, sin librerías externas. Identifica riesgos de concurrencia y del uso de IP».

**Implementación:** «Lee spec.md, plan.md y tasks.md. Implementa incrementos pequeños. Por cada tarea, explica el test que demuestra su cumplimiento».

**Revisión:** «Comprueba uno por uno los criterios de aceptación; busca requisitos sin tests, divergencias y supuestos no documentados; propón reparaciones».
