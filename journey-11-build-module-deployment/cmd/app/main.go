package main

import (
    "fmt"
    "log"

    "build-module-deployment/pkg/config"
    "build-module-deployment/internal/core"
)

func main() {
    cfg, err := config.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("error loading config: %v", err)
    }

    fmt.Println("App Name:", cfg.App.Name)
    fmt.Println(core.Hello("Gopher!"))
}
