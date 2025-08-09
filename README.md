# 🤪 Convertidor de Texto a Jerga de Internet de los 2000 en Go

Este proyecto es un script en Go que toma una frase de texto y la transforma en una versión que simula la jerga de internet de los años 2000. El objetivo es recrear el estilo de los chats y foros de la época, lleno de abreviaturas, emojis y errores ortográficos intencionales.

## 🚀 Características

* **Abreviaturas Comunes:** Reemplaza frases como "por favor" por "xfa" o "te quiero" por "tqm".
* **Errores de Ortografía:** Simula errores intencionales, como cambiar 's' por 'z', 'c' por 'k', etc.
* **Emojis Característicos:** Añade emojis y emoticons populares de la época (`:P`, `;)`, `xD`).
* **Aleatoriedad:** Cada conversión es única gracias al uso de funciones aleatorias para la selección de abreviaturas y la aplicación de errores.

## ⚙️ Cómo Ejecutar el Proyecto

1.  Asegúrate de tener **Go** instalado en tu sistema.
2.  Clona el repositorio:
    ```bash
    git clone https://github.com/santiagourdaneta/convertidor-jerga-2000-go/
    cd convertidor-jerga-2000-go
    ```
3.  Ejecuta el programa:
    ```bash
    go run main.go
    ```
    El script imprimirá ejemplos de conversión directamente en la consola.

## ✅ Pruebas

El proyecto incluye pruebas para asegurar que la lógica de conversión funcione como se espera.

* **Pruebas Unitarias:** Validan que las funciones individuales, como la de aplicar errores, se comporten correctamente.
* **Pruebas de Integración:** Verifican que el flujo completo de conversión produzca una salida transformada y no idéntica a la original.
* **Pruebas de Estrés:** Utilizan el sistema de *benchmarking* de Go para medir el rendimiento del convertidor al procesar grandes cantidades de texto.

Para ejecutar todas las pruebas y los benchmarks:
```bash
go test -bench=.

🛠️ Estructura del Código
main.go: Contiene la lógica principal del programa, incluyendo las reglas de reemplazo y las funciones de conversión.
main_test.go: Archivo que contiene todas las pruebas unitarias, de integración y de estrés para el proyecto.

✍️ Contribuir
¡Si tienes ideas para nuevas reglas de jerga, mejoras en los algoritmos o cualquier otra sugerencia, eres bienvenido a contribuir! Abre un issue o envía un pull request.

Labels & Tags

go golang script texto manipulacion-cadenas jerga 2000s

Hashtags

#go #golang #jerga #2000s #chat #programacion #stringmanipulation
