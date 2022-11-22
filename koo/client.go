package koo

import (
	"net/http"

	"github.com/dghubble/sling"
)

const kooAPI = "https://www.kooapp.com/apiV1"

type Client struct {
	sling       *sling.Sling
	FeedService *FeedService
}

func NewClient(httpClient *http.Client) *Client {
	clientBase := sling.New().Client(httpClient).Base(kooAPI)

	return &Client{
		sling:       clientBase,
		FeedService: newFeedService(clientBase.New()),
	}
}
