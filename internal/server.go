package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eric-sison/pulse/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Setup an HTTP server with graceful shutdown handling based on received signals from the operating system.
func (a *App) StartServer() {

	router := gin.Default()

	s := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", getPort()),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// This channel is used to receive signals from the operating system, such
	// as `os.Interrupt`, `syscall.SIGTERM`, `syscall.SIGINT`, and `syscall.SIGHUP`, which are then
	// handled to gracefully shut down the server when one of these signals is received.
	stop := make(chan os.Signal, 1)

	// Setup up a notification mechanism to listen for specific signals sent to the program.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	// A helper function that starts the HTTP server in a goroutine and then
	// executes a callback function. In this case, the callback function logs a message indicating that
	// the server is running on a specific address and port.
	startServer(s)

	//This line is blocking until a signal is received on the `stop` channel,
	// at which point it will proceed with the received signal value.
	sig := <-stop

	log.Println("Shutting down the server...", sig)

	// Create a context with a 5-second timeout to ensure the server has
	// enough time to complete ongoing requests before shutting down.
	// If the server doesn't shut down within this period, it forces a shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown %v.", err)
	}

	log.Println("Server exited gracefully.")
}

// Starts an HTTP server in a goroutine.
func startServer(s *http.Server) {
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server is running on http://" + s.Addr + ".")
}

// Retrieve the application port from the environment variable or defaults to a predefined port if not found.
func getPort() int {
	var defaultPort = 3214

	port := os.Getenv("APP_PORT")

	if port == "" {
		log.Printf("Unable to load application port. Defaulting to %d.", defaultPort)

		return defaultPort
	}

	return utils.ToInt(port)
}
