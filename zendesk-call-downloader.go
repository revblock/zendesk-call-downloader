package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cinch-labs/zendesk-call-downloader/zendesk"
)

func main() {

	zendeskUrl := flag.String("url", "", "Your Zendesk workspace URL")
	zendeskUsername := flag.String("username", "", "Your Zendesk username")
	zendeskToken := flag.String("api-token", "", "Your Zendesk API token")
	customerId := flag.Int64("customer-id", 0, "The ID of the customer you want to download calls for")

	flag.Parse()

	flagError := false

	if *zendeskUrl == "" {
		fmt.Println("Missing Zendesk URL")
		flagError = true
	}

	if *zendeskUsername == "" {
		fmt.Println("Missing Zendesk username")
		flagError = true
	}

	if *zendeskToken == "" {
		fmt.Println("Missing Zendesk API Token")
		flagError = true
	}

	if *customerId == 0 {
		fmt.Println("Missing Customer ID")
		flagError = true
	}

	if flagError {
		return
	}

	zendeskClient, err := zendesk.NewClient(&http.Client{}, *zendeskUrl, *zendeskUsername, *zendeskToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Getting all tickets for customer: %d\n", *customerId)

	tickets := zendeskClient.GetAllTicketsForUser(*customerId)

	fmt.Printf("Found %d tickets\n", len(tickets))

	var comments []zendesk.Comment

	for _, ticket := range tickets {

		time.Sleep(5 * time.Second)

		fmt.Printf("Getting all ticket comments for ticket: %d\n", ticket.ID)

		ticketComments := zendeskClient.GetAllCommentsForTicket(ticket.ID)

		fmt.Printf("Found %d ticket comments for ticket: %d\n", len(ticketComments), ticket.ID)

		for _, comment := range ticketComments {
			if comment.Data.RecordingURL != nil {
				comments = append(comments, comment)
				fmt.Printf("Recording found: %v\n", *comment.Data.RecordingURL)
			}
		}
	}

	fmt.Printf("\n\nFound %d ticket comments with recordings", len(comments))

	for _, comment := range comments {
		recordingUrl, _ := url.Parse(*comment.Data.RecordingURL)

		date := comment.Data.StartedAt.Format("2006-01-02-150405")
		filename := date + ".mp3"

		zendeskClient.DownloadCallRecording(*recordingUrl, filename)

		fmt.Printf("Recording saved: %v\n", filename)
	}

	fmt.Println("done")
}
