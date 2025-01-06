
package climate

import (
    "fmt"
    "strings"
)

type Analyzer struct {
    cities []*City
}

func NewAnalyzer() *Analyzer {
    return &Analyzer{
        cities: make([]*City, 0),
    }
}

func (a *Analyzer) AddCity(city *City) error {
    
    for _, c := range a.cities {
        if strings.EqualFold(c.name, city.name) {
            return fmt.Errorf("city %s already exists", city.name)
        }
    }

    if city.temperature < -273.15 {
        return fmt.Errorf("invalid temperature: cannot be below absolute zero")
    }

    if city.rainfall < 0 {
        return fmt.Errorf("invalid rainfall: cannot be negative")
    }

    a.cities = append(a.cities, city)
    return nil
}

func (a *Analyzer) GetHighestTemperatureCity() (*City, error) {
    if len(a.cities) == 0 {
        return nil, fmt.Errorf("no cities available")
    }

    highest := a.cities[0]
    for _, city := range a.cities {
        if city.temperature > highest.temperature {
            highest = city
        }
    }
    return highest, nil
}

func (a *Analyzer) GetLowestTemperatureCity() (*City, error) {
    if len(a.cities) == 0 {
        return nil, fmt.Errorf("no cities available")
    }

    lowest := a.cities[0]
    for _, city := range a.cities {
        if city.temperature < lowest.temperature {
            lowest = city
        }
    }
    return lowest, nil
}

func (a *Analyzer) GetAverageRainfall() (float64, error) {
    if len(a.cities) == 0 {
        return 0, fmt.Errorf("no cities available")
    }

    total := 0.0
    for _, city := range a.cities {
        total += city.rainfall
    }
    return total / float64(len(a.cities)), nil
}

func (a *Analyzer) GetCitiesAboveRainfall(threshold float64) []*City {
    var result []*City
    for _, city := range a.cities {
        if city.rainfall > threshold {
            result = append(result, city)
        }
    }
    return result
}

func (a *Analyzer) SearchByName(name string) (*City, error) {
    for _, city := range a.cities {
        if strings.EqualFold(city.name, name) {
            return city, nil
        }
    }
    return nil, fmt.Errorf("city %s not found", name)
}

func (a *Analyzer) GetAllCities() []*City {
    return a.cities
}