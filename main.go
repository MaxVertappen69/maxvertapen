package main

import "fmt"

func main() {
    var celsius float64
    fmt.Println("=== Program Konversi Suhu ===")
    fmt.Print("Masukkan suhu dalam Celsius (°C): ")
    fmt.Scanln(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32
    kelvin := celsius + 273.15

    fmt.Printf("Hasil Fahrenheit : %.2f °F\n", fahrenheit)
    fmt.Printf("Hasil Kelvin     : %.2f K\n", kelvin)
}