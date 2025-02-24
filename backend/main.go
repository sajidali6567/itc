package main

import (
    "log"
    "net/http"
    "itc/config"
    "itc/router"
)

func main() {
    cfg := config.LoadConfig()
    r := router.SetupRouter()

    log.Println("Server running on port", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, r))
}
