package main

import (
    "log"
    "net/http"
    "itc/config"
    "itc/router"
)

// Middleware to handle CORS
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func main() {
    cfg := config.LoadConfig()
    r := router.SetupRouter()

    // Apply CORS middleware
    handler := corsMiddleware(r)

    log.Println("Server running on port", cfg.ServerPort)
    log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, handler))
}
