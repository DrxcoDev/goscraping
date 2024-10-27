package main

import (
    "log"
    "github.com/gocolly/colly/v2"
    "goscraping/utils" // Asegúrate de usar el nombre correcto de tu módulo
)

func main() {
    // Inicializa el colector
    c := colly.NewCollector()

    // Configura el User-Agent
    c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36"

    // Define la lógica de respuesta
    c.OnResponse(func(r *colly.Response) {
        if r.StatusCode != 200 {
            log.Printf("Error: recibí el código de estado %d de %s\n", r.StatusCode, r.Request.URL)
            return // Salimos si el código no es 200
        }
        log.Printf("Código de estado %d: %s\n", r.StatusCode, r.Request.URL)
    })

    // Define la lógica de scraping
    var quotes []utils.Quote
    c.OnHTML(".quote", func(e *colly.HTMLElement) {
        quote := utils.Quote{
            Text:   e.ChildText(".text"),
            Author: e.ChildText(".author"),
        }
        quotes = append(quotes, quote)
    })

    // Visita la página web (reemplaza con la URL que desees)
    err := c.Visit("http://quotes.toscrape.com/") // Cambia el dominio por el que quieras
    if err != nil {
        log.Fatal(err)
    }

    // Guarda las citas en un archivo JSON
    err = utils.SaveQuote(quotes)
    if err != nil {
        log.Printf("Error al guardar citas: %s\n", err)
    }

    log.Println("Citas guardadas con éxito.")
}