package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("====== PROGRAM KONVERSI SUHU (GO) ======")

    for {
        fmt.Print("\nMasukkan suhu Celsius (°C) [ketik 'q' untuk keluar]: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        if strings.ToLower(input) == "q" {
            fmt.Println("Program selesai. Terima kasih!")
            break
        }

        celsius, err := strconv.ParseFloat(input, 64)
        if err != nil {
            fmt.Println("[Error] Input tidak valid, masukkan angka!")
            continue
        }

        if celsius < -273.15 {
            fmt.Println("[Peringatan] Suhu tidak boleh di bawah nol mutlak (-273.15 °C)!")
            continue
        }

        fmt.Printf("- Fahrenheit : %.2f °F\n", (celsius*9/5)+32)
        fmt.Printf("- Kelvin     : %.2f K\n", celsius+273.15)
        fmt.Printf("- Reamur     : %.2f °R\n", celsius*4/5)
    }
}