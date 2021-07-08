package main

import (
    "fmt"
	"log"
	"meta/internal/app/engine"
	"meta/internal/app/model"
    "meta/internal/app/service"
    "github.com/douyu/jupiter"
)

func main() {
	eng := engine.NewEngine()
	eng.RegisterHooks(jupiter.StageAfterStop, func() error {
        fmt.Println("exit jupiter app ...")
        return nil
      })

    model.Init()
    service.Init()
    if err := eng.Run(); err != nil {
    	log.Fatal(err)
    }
}

