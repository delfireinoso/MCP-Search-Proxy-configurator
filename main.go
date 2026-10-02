package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile   = "UPSTREAMS.txt"
	envExample  = ".env.example"
	outputFile  = ".env"
)

func main() {
	// Leemos el archivo de entrada
	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("❌ Archivo %s no encontrado. Crea este archivo con tus servicios MCP.\n", inputFile)
		}
		log.Fatalf("❌ Error al abrir archivo: %v\n", err)
	}
	defer file.Close()

	// Variables para construir la lista
	var validServices []string
	var comments []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Separar comentarios de servicios
		if line == "" || line[0] == '#' {
			comments = append(comments, line)
			continue
		}

		// Añadir línea válida a la lista
		validServices = append(validServices, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("❌ Error al leer archivo: %v\n", err)
	}

	// Generar el archivo .env
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("❌ Error al crear archivo %s: %v\n", outputFile, err)
	}
	defer f.Close()

	// Escribir encabezado
	header := `# MCP Converter - Generated file
# UPSTREAMS variable overwritten by converter
`
	if _, err := f.WriteString(header); err != nil {
		log.Fatalf("❌ Error al escribir encabezado: %v\n", err)
	}

	// Copiar contenido de .env.example EXCEPTO la línea UPSTREAMS=del ejemplo
	const exampleUpstreams = "UPSTREAMS=https://opendata.cat/api/mcp,https://foo.bar/mcp|sk-your-token-here"

	exampleFile, err := os.Open(envExample)
	if err != nil {
		log.Fatalf("❌ Error al abrir %s: %v\n", envExample, err)
	}
	defer exampleFile.Close()

	exampleScanner := bufio.NewScanner(exampleFile)
	for exampleScanner.Scan() {
		line := exampleScanner.Text()

		// Omitir la línea del ejemplo
		if line == exampleUpstreams {
			continue
		}

		// Escribir línea del .env.example
		if _, err := f.WriteString(line + "\n"); err != nil {
			log.Fatalf("❌ Error al escribir línea de .env.example: %v\n", err)
		}
	}

	if err := exampleScanner.Err(); err != nil {
		log.Fatalf("❌ Error al leer %s: %v\n", envExample, err)
	}

	// Escribir la variable UPSTREAMS (en el lugar donde estaba el ejemplo)
	if _, err := f.WriteString("\n"); err != nil {
		log.Fatalf("❌ Error al escribir nueva línea: %v\n", err)
	}

	// Construir la variable UPSTREAMS
	upstreamsStr := ""
	for i, us := range validServices {
		if i > 0 {
			upstreamsStr += ","
		}
		upstreamsStr += us
	}

	if _, err := f.WriteString(fmt.Sprintf("UPSTREAMS=%s\n", upstreamsStr)); err != nil {
		log.Fatalf("❌ Error al escribir UPSTREAMS: %v\n", err)
	}

	// Escribir comentarios al final
	if len(comments) > 0 {
		if _, err := f.WriteString("\n# --- Comments from UPSTREAMS.txt ---\n"); err != nil {
			log.Fatalf("❌ Error al escribir línea de comentarios: %v\n", err)
		}
		for _, c := range comments {
			if _, err := f.WriteString(c + "\n"); err != nil {
				log.Fatalf("❌ Error al escribir comentario: %v\n", err)
			}
		}
	}

	// Mostrar resultado en consola
	fmt.Println("✅ Conversión exitosa")
	fmt.Println()
	fmt.Println("📄 Archivo generado:", outputFile)
	fmt.Println()
	fmt.Println("📋 Lista de servicios válidos:")
	for _, us := range validServices {
		fmt.Println("   •", us)
	}
	fmt.Println()
	fmt.Println("📋 Comentarios del archivo:")
	for _, c := range comments {
		fmt.Println("   •", c)
	}
	fmt.Println()
	fmt.Println("🎯 Variable UPSTREAMS:")
	fmt.Println("   ", upstreamsStr)
}