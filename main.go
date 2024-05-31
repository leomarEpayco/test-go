package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const consultaFrecuencia = 15

// Función para consultar transacciones
func consultarTransacciones(comercio int) (*http.Response, error) {
	url := fmt.Sprintf("https://apify.epayco.io/ping?id=%d", comercio)

	// Crea un cliente HTTP con Timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Establece el timeout en 10 segundos
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Función para registrar transacciones
func registrarTransacciones(transacciones *http.Response, body []byte) error {
	// ... (Código para enviar la petición HTTPS) ...
	return nil // Reemplaza con la lógica de tu petición HTTPS
}

// Función para ejecutar las consultas
func ejecutarConsultas() {
	// Define el número de requests
	numRequests := 5000
	fmt.Printf("ENTRO")

	// Crea un WaitGroup para sincronizar las goroutines
	var wg sync.WaitGroup

	// Define el tiempo de inicio
	start := time.Now()
	wg.Add(numRequests)

	// Crea un canal para enviar las respuestas
	responses := make(chan http.Response, numRequests)

	fechaActual := time.Now()
	hora := fechaActual.Hour()
	minutos := fechaActual.Minute()
	segundos := fechaActual.Second()

	fmt.Printf("Inicio de consultas a las %d:%d:%d\n", hora, minutos, segundos)

	for i := 0; i < numRequests; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := consultarTransacciones(9999)
			if err != nil {
				fmt.Println("Error al consultar transacciones:", err)
				return
			}
			responses <- *resp // Envía la respuesta al canal
		}(i)
	}

	// Espera que todas las goroutines terminen
	wg.Wait()

	// Imprime el tiempo total
	elapsed := time.Since(start)
	fmt.Printf("Tiempo total: %v\n", elapsed)

	fechaActual = time.Now()
	hora = fechaActual.Hour()
	minutos = fechaActual.Minute()
	segundos = fechaActual.Second()
	fmt.Printf("Terminó de consultar a las %d:%d:%d\n", hora, minutos, segundos)

}

func main() {
	fmt.Printf("INICIO")
	var wg sync.WaitGroup
	wg.Add(1) // Incrementa el contador de WaitGroup
	go func() {
		defer wg.Done() // Decrementa el contador al finalizar
		ejecutarConsultas()
	}()
	wg.Wait()
}