package zendesk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

const getAllTicketsForUserPath = "/api/v2/users/%d/tickets/requested.json?page=%d"
const getAllCommentsForTicketPath = "/api/v2/tickets/%d/comments.json?page=%d"

type ZendeskClient struct {
	baseURL  url.URL
	client   http.Client
	token    string
	username string
}

func (zendeskClient *ZendeskClient) DownloadCallRecording(requestUrl url.URL, outputPath string) {

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	req, err := http.NewRequest("GET", requestUrl.String(), nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(zendeskClient.username, zendeskClient.token)

	resp, err := zendeskClient.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}

func (zendeskClient *ZendeskClient) GetAllCommentsForTicket(ticketId int) []Comment {
	next := true
	page := 1
	ticketComments := make([]Comment, 0)

	for next {
		path, _ := url.Parse(fmt.Sprintf(getAllCommentsForTicketPath, ticketId, page))
		requestUrl := zendeskClient.baseURL.ResolveReference(path)

		req, err := createZendeskRequest(zendeskClient, "GET", requestUrl.String(), nil)
		if err != nil {
			panic(err)
		}

		resp, err := zendeskClient.client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var getTicketCommentsResponse GetTicketCommentsResponse
		json.NewDecoder(resp.Body).Decode(&getTicketCommentsResponse)

		pageTicketComments := getTicketCommentsResponse.Comments

		ticketComments = append(ticketComments, pageTicketComments...)
		page = page + 1

		if getTicketCommentsResponse.NextPage == nil {
			next = false
		} else {
			// Wait so we don't exceed rate limit
			time.Sleep(5 * time.Second)
		}
	}

	return ticketComments
}

func (zendeskClient *ZendeskClient) GetAllTicketsForUser(userId int64) []Ticket {

	next := true
	page := 1
	tickets := make([]Ticket, 0)

	for next {
		path, _ := url.Parse(fmt.Sprintf(getAllTicketsForUserPath, userId, page))
		requestUrl := zendeskClient.baseURL.ResolveReference(path)

		req, err := createZendeskRequest(zendeskClient, "GET", requestUrl.String(), nil)
		if err != nil {
			panic(err)
		}

		resp, err := zendeskClient.client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var getTicketsResponse GetTicketsResponse
		json.NewDecoder(resp.Body).Decode(&getTicketsResponse)

		pageTickets := getTicketsResponse.Tickets

		tickets = append(tickets, pageTickets...)
		page = page + 1

		if getTicketsResponse.NextPage == nil {
			next = false
		} else {
			// Wait so we don't exceed rate limit
			time.Sleep(5 * time.Second)
		}
	}

	return tickets
}

func createZendeskRequest(zendeskClient *ZendeskClient, method string, requestUrl string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("GET", requestUrl, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(zendeskClient.username, zendeskClient.token)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func NewClient(httpClient *http.Client, baseURL string, username string, token string) (*ZendeskClient, error) {

	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	parsedUsername := username + "/token"

	return &ZendeskClient{
		baseURL:  *parsedUrl,
		username: parsedUsername,
		token:    token,
		client:   *httpClient,
	}, nil
}
