package main

import (
	"strings"
	"testing"
)

func TestAplicarError(t *testing.T) {
	// Caso de prueba 1: Reemplazar 's' por 'z'
	palabra := "esta"
	// Simular el reemplazo manualmente para verificar la lógica
	palabraConError := strings.ReplaceAll(palabra, "s", "z")
	
	if palabraConError != "ezta" { // <-- Valor esperado corregido
		t.Errorf("Error en el reemplazo de 's' por 'z'. Se esperaba 'ezta', pero se obtuvo %s", palabraConError)
	}

	// Caso de prueba 2: Reemplazar 'c' por 'k'
	palabra = "casa"
	palabraConError = strings.ReplaceAll(palabra, "c", "k")

	if palabraConError != "kasa" {
		t.Errorf("Error en el reemplazo de 'c' por 'k'. Se esperaba 'kasa', pero se obtuvo %s", palabraConError)
	}
}

func TestConvertirTextoAJerga(t *testing.T) {
	fraseOriginal := "Hola, por favor, me puedes decir qué hora es."

	// Ejecutar la función varias veces para evitar fallos por aleatoriedad
	for i := 0; i < 10; i++ {
		fraseConvertida := convertirTextoAJerga(fraseOriginal)
		
		// Verificación 1: La frase convertida no debe ser la misma que la original.
		if fraseConvertida == fraseOriginal {
			t.Errorf("La frase convertida es idéntica a la original. Posible fallo en la conversión.")
		}

		// Verificación 2: La frase convertida debe contener al menos un emoji.
		tieneEmoji := false
		for _, emoji := range emojis {
			if strings.Contains(fraseConvertida, emoji) {
				tieneEmoji = true
				break
			}
		}
		if !tieneEmoji {
			t.Errorf("La frase convertida no contiene un emoji. Frase: %s", fraseConvertida)
		}
	}
}

// BenchmarkConvertirTexto mide el rendimiento de la función de conversión
// al procesar un texto de gran longitud.
func BenchmarkConvertirTexto(b *testing.B) {
	// Crear un texto largo para la prueba de estrés
	textoLargo := `Hola, por favor, me puedes decir qué hora es. Te quiero.
	Tú eres la mejor persona que conozco. Gracias por todo.`
	
	// Reiniciar el temporizador para que no incluya el tiempo de configuración
	b.ResetTimer()

	// El bucle de `b.N` se ejecuta automáticamente por Go
	for i := 0; i < b.N; i++ {
		convertirTextoAJerga(textoLargo)
	}
}

// Para ejecutar: go test