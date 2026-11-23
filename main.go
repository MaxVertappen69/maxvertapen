package main

import "fmt"

func main() {
    var celsius float64
    fmt.Println("=== Program Konversi Suhu ===")
    fmt.Print("Masukkan suhu dalam Celsius (°C): ")
    fmt.Scanln(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32
    fmt.Printf("Hasil Fahrenheit: %.2f °F\n", fahrenheit)
}