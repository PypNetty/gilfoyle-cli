// cmd/cli/main.go
package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "gilfoyle-cli/internal/cli"  // Utilise le nom du module local
)

func main() {
    cli, err := cli.New()
    if err != nil {
        log.Fatalf("Erreur d'initialisation (comme c'était prévisible): %v", err)
    }

    // Gestion des signaux
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-sigChan
        cli.Cleanup()
        os.Exit(0)
    }()

    cli.Run()
}