package main

import (
	"flag"
	"log"
	"pion-tests/icebreaker"
)

func main() {
	// Define flags without default values (not passing a value will cause an error)
	gameId := flag.Uint64("game-id", 0, "The ID of the game session")
	accessToken := flag.String("access-token", "", "The access token for authentication")
	apiRoot := flag.String("api-root", "https://api.faforever.com/ice", "The root uri of the icebreaker api")

	// Parse the command-line flags
	flag.Parse()

	// Validate that the required flags are provided
	if *gameId == 0 {
		log.Fatalf("Error: --game-id is required and must be a valid uint64.")
	}

	if *accessToken == "" {
		log.Fatalf("Error: --access-token is required and cannot be empty.")
	}

	icebreakerClient := icebreaker.NewClient(*apiRoot, *gameId, *accessToken)

	sessionGameResponse, err := icebreakerClient.GetGameSession()

	if err != nil {
		log.Fatal("Could not query session game: ", err)
	}

	// Use the parsed data
	log.Printf("Parsed API Response: %+v\n", *sessionGameResponse)

	go icebreakerClient.Listen()

	select {}
}
