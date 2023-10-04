package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	// Abrir el archivo CSV de entrada (ticket)
	ticketFile, err := os.Open("ticket.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo CSV de entrada:", err)
		return
	}
	defer ticketFile.Close()

	// Crear un lector CSV para el archivo de entrada
	ticketReader := csv.NewReader(ticketFile)

	// Inicializar variables para almacenar los totales
	importeTotal := 0.0
	impuestosTotales := 0.0

	// Procesar cada línea del archivo CSV de entrada
	for {
		row, err := ticketReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error al leer una línea del archivo CSV:", err)
			return
		}

		// Ignorar la primera línea (encabezados)
		if row[0] == "CodigoProducto" {
			continue
		}

		// Convertir los valores de la línea a tipos numéricos
		cantidad, _ := strconv.ParseFloat(row[1], 64)
		precioUnitario, _ := strconv.ParseFloat(row[2], 64)
		impuestosProducto, _ := strconv.ParseFloat(row[4], 64)

		//TODO Calcular el subtotal de cada fila
		//subtotal...

		//TODO Calcular los impuestos de cada fila (IVA)
		//iva := ...

		//TODO Calcular la sumatoria de los subtotales..
		//importeTotal...

		//TODO Calcular la sumatoria del IVA
		//impuestosTotales = ...
	}

	// Calcular el total (importe + IVA total)
	total := importeTotal + impuestosTotales

	// Crear un archivo CSV de salida para los totales
	resultFile, err := os.Create("result.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo CSV de salida:", err)
		return
	}
	defer resultFile.Close()

	// Crear un escritor CSV para el archivo de salida
	resultWriter := csv.NewWriter(resultFile)

	// Escribir los totales en el archivo CSV de salida
	resultWriter.Write([]string{"importe", "impuestos", "total"})
	resultWriter.Write([]string{fmt.Sprintf("%.2f", importeTotal), fmt.Sprintf("%.2f", impuestosTotales), fmt.Sprintf("%.2f", total)})

	// Finalizar la escritura en el archivo CSV de salida
	resultWriter.Flush()

	fmt.Println("Totales calculados y guardados en result.csv")
}
