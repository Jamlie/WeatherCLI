package weather

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

    "github.com/Jamlie/weather_cli/globals"
)

func FetchWeatehr(city string) (string, error) {
    var data struct {
        Main struct {
            Temp float64 `json:"temp"`
        } `json:"main"`
    }

    url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, globals.API_KEY)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
        return "", err
    }
    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        log.Fatal(err)
        return "", err
    }

    return fmt.Sprintf("%s: %.2f°C", city, data.Main.Temp), nil
}
