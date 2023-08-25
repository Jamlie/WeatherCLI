package main

import (
    "fmt"
    "context"
    "os"
    "log"

    "github.com/Jamlie/weather_cli/weather"
    "github.com/cristalhq/acmd"
)

func main() {
    cmd := []acmd.Command{
        {
            Name: "now",
            Description: "Run the weather command",
            ExecFunc: func(ctx context.Context, args []string) error {
                fmt.Print("Enter city name: ")
                var city string
                fmt.Scanln(&city)
                temp, err := weather.FetchWeatehr(city)
                if err != nil {
                    log.Fatal(err)
                    return err
                }

                fmt.Println(temp)
                return nil
            },
        },
    }

    r := acmd.RunnerOf(cmd, acmd.Config{
        AppName: "weather",
        AppDescription: "Weather is a command line tool to get the weather",
        Version: "1.0.0",
        Output: os.Stdout,
    })

    r.Run()
}
