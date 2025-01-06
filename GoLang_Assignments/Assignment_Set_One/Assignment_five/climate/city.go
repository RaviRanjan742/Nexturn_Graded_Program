package climate

type City struct {
    name        string
    temperature float64
    rainfall    float64
}

func NewCity(name string, temperature, rainfall float64) *City {
    return &City{
        name:        name,
        temperature: temperature,
        rainfall:    rainfall,
    }
}

func (c *City) GetName() string {
    return c.name
}

func (c *City) GetTemperature() float64 {
    return c.temperature
}

func (c *City) GetRainfall() float64 {
    return c.rainfall
}