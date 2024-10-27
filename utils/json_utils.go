package utils

import (
    "encoding/json"
    "os"
)

// Quote representa una cita
type Quote struct {
    Text   string `json:"text"`
    Author string `json:"author"`
}

// SaveQuote guarda las citas en un archivo JSON
func SaveQuote(quotes []Quote) error {
    file, err := os.Create("build/quotes.json")
    if err != nil {
        return err // Devuelve el error si no se puede crear el archivo
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    return encoder.Encode(quotes) // Devuelve el error del encoder si ocurre
}
