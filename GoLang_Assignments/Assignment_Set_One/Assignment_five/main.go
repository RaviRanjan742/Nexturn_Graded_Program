
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "Assignment_five/climate"
)

func readString(reader *bufio.Reader, prompt string) string {
    fmt.Print(prompt)
    str, _ := reader.ReadString('\n')
    return strings.TrimSpace(str)
}

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
    str := readString(reader, prompt)
    return strconv.ParseFloat(str, 64)
}

func showMenu() {
    fmt.Println("\n=== Climate Data Analysis System ===")
    fmt.Println("1. Add New City Data")
    fmt.Println("2. Show All Cities")
    fmt.Println("3. Find City with Highest Temperature")
    fmt.Println("4. Find City with Lowest Temperature")
    fmt.Println("5. Show Average Rainfall")
    fmt.Println("6. Filter Cities by Rainfall Threshold")
    fmt.Println("7. Search City by Name")
    fmt.Println("8. Exit")
    fmt.Print("Enter your choice: ")
}

func addCityData(reader *bufio.Reader, analyzer *climate.Analyzer) {
    fmt.Println("\n--- Add New City Data ---")
    
    name := readString(reader, "Enter City Name: ")
    if name == "" {
        fmt.Println("City name cannot be empty")
        return
    }

    temp, err := readFloat(reader, "Enter Average Temperature (°C): ")
    if err != nil {
        fmt.Println("Invalid temperature format")
        return
    }

    rainfall, err := readFloat(reader, "Enter Rainfall (mm): ")
    if err != nil {
        fmt.Println("Invalid rainfall format")
        return
    }

    city := climate.NewCity(name, temp, rainfall)
    err = analyzer.AddCity(city)
    if err != nil {
        fmt.Printf("Error adding city: %v\n", err)
        return
    }

    fmt.Println("City data added successfully!")
}

func displayCities(cities []*climate.City) {
    if len(cities) == 0 {
        fmt.Println("No cities found")
        return
    }

    fmt.Println("\nCity Data:")
    fmt.Printf("%-20s | %-15s | %-15s\n", "City", "Temperature (°C)", "Rainfall (mm)")
    fmt.Println(strings.Repeat("-", 56))
    
    for _, city := range cities {
        fmt.Printf("%-20s | %15.2f | %15.2f\n", 
            city.GetName(), city.GetTemperature(), city.GetRainfall())
    }
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    analyzer := climate.NewAnalyzer()

    

    for {
        showMenu()
        choice, err := readInt(reader, "")
        if err != nil {
            fmt.Println("Invalid input. Please try again.")
            continue
        }

        switch choice {
        case 1:
            addCityData(reader, analyzer)

        case 2:
            displayCities(analyzer.GetAllCities())

        case 3:
            city, err := analyzer.GetHighestTemperatureCity()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nCity with highest temperature:\n")
                displayCities([]*climate.City{city})
            }

        case 4:
            city, err := analyzer.GetLowestTemperatureCity()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nCity with lowest temperature:\n")
                displayCities([]*climate.City{city})
            }

        case 5:
            avg, err := analyzer.GetAverageRainfall()
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                fmt.Printf("\nAverage rainfall across all cities: %.2f mm\n", avg)
            }

        case 6:
            threshold, err := readFloat(reader, "Enter rainfall threshold (mm): ")
            if err != nil {
                fmt.Println("Invalid threshold format")
                continue
            }

            cities := analyzer.GetCitiesAboveRainfall(threshold)
            fmt.Printf("\nCities with rainfall above %.2f mm:\n", threshold)
            displayCities(cities)

        case 7:
            name := readString(reader, "Enter city name to search: ")
            city, err := analyzer.SearchByName(name)
            if err != nil {
                fmt.Printf("Error: %v\n", err)
            } else {
                displayCities([]*climate.City{city})
            }

        case 8:
            fmt.Println("Thank you for using the Climate Data Analysis System. Goodbye!")
            return

        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
    str := readString(reader, prompt)
    return strconv.Atoi(str)
}
