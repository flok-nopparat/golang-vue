package main

import (
	"fmt"
	"line/interview/utils"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

// Setup routes
func setupRoutes() {
	// Set up static directory
	http.Handle("/", http.FileServer(http.Dir("./client/line-interview/dist")))

	// Set up routes
	http.HandleFunc("/upload", utils.UploadFile)
	http.HandleFunc("/uploadMulticore", utils.UploadFileMulticore)

	// Start the server
	http.ListenAndServe(":"+os.Getenv("RUN_PORT"), nil)

}

// Load environment variables from .env file
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Main function
func main() {
	fmt.Println("Server is running at port " + os.Getenv("RUN_PORT"))
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println("CPU run on ", runtime.NumCPU())
	setupRoutes()
}
