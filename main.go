package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Estructura para almacenar las reglas de reemplazo
type Regla struct {
	palabra string
	reemplazo []string
}

// Conjunto de reglas para convertir el texto
var reglas = []Regla{
	{"por favor", []string{"xfa", "plz", "xfavor"}},
	{"gracias", []string{"grx", "ty", "graciaz"}},
	{"hola", []string{"q onda", "ola", "q pex"}},
	{"qué", []string{"q", "ke"}},
	{"que", []string{"q", "ke"}},
	{"tú", []string{"tu", "u"}},
	{"yo", []string{"io", "iio"}},
	{"adiós", []string{"bye", "chau", "cya"}},
	{"te quiero", []string{"tqm", "tk"}},
	{"para", []string{"pa", "pA"}},
	{"estoy", []string{"estoi", "stoi"}},
	{"es", []string{"ez", "eh"}},
	{"así", []string{"azi", "asi"}},
	{"hace", []string{"aze", "ace"}},
}

// Conjunto de emojis aleatorios para añadir al final de la frase
var emojis = []string{
	"😜", "😉", "😅", "😎", "😂", "🤘", "👍", "🤙", ":P", ":)", ":D", ";)",
}

// Conjunto de errores ortográficos intencionales
var errores = map[string]string{
	"c": "k",
	"s": "z",
	"qu": "k",
	"v": "b",
	"y": "i",
}

// Convierte una frase a jerga de internet
func convertirTextoAJerga(frase string) string {
	// Poner todo en minúsculas para que coincida con las reglas
	frase = strings.ToLower(frase)

	// Reemplazar palabras completas según las reglas
	for _, regla := range reglas {
		reemplazo := regla.reemplazo[rand.Intn(len(regla.reemplazo))]
		frase = strings.ReplaceAll(frase, regla.palabra, reemplazo)
	}

	// Simular errores ortográficos aleatorios
	palabras := strings.Split(frase, " ")
	for i, palabra := range palabras {
		// Con una probabilidad del 30%, aplicar un error ortográfico
		if rand.Float64() < 0.3 {
			palabras[i] = aplicarError(palabra)
		}
	}
	frase = strings.Join(palabras, " ")

	// Añadir un emoji al final de la frase
	emoji := emojis[rand.Intn(len(emojis))]
	return fmt.Sprintf("%s %s", frase, emoji)
}

// Aplica un error ortográfico a una palabra
func aplicarError(palabra string) string {
	// Reemplazar aleatoriamente una letra
	letras := strings.Split(palabra, "")
	if len(letras) > 1 {
		indice := rand.Intn(len(letras))
		letra := letras[indice]
		if reemplazo, ok := errores[letra]; ok {
			letras[indice] = reemplazo
		}
	}
	return strings.Join(letras, "")
}

func main() {
	// Inicializar la semilla aleatoria
	rand.Seed(time.Now().UnixNano())

	fraseOriginal := "Hola, por favor, me puedes decir qué hora es. Te quiero."
	fmt.Printf("Frase original: %s\n", fraseOriginal)

	fraseConvertida := convertirTextoAJerga(fraseOriginal)
	fmt.Printf("Frase convertida: %s\n", fraseConvertida)

	// Ejemplo 2
	fraseOriginal2 := "Qué tal. Yo estoy para ayudarte."
	fmt.Printf("\nFrase original: %s\n", fraseOriginal2)
	
	fraseConvertida2 := convertirTextoAJerga(fraseOriginal2)
	fmt.Printf("Frase convertida: %s\n", fraseConvertida2)
}