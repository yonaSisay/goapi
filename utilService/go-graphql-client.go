package utilService

import (
	"net/http"
	"os"
	// "fmt"
	"github.com/hasura/go-graphql-client"
)

type headersTransport struct {
	headers http.Header
	base    http.RoundTripper
}

func (t *headersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Set(k, v[0])
	}
	return t.base.RoundTrip(req)
}

func Client() *graphql.Client {	
	// Set up the HTTP client with the request headers
	headers := http.Header{}
	headers.Add("X-Hasura-Admin-Secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))
	// An HTTP transport that adds headers to requests
	httpClient := &http.Client{Transport: &headersTransport{headers, http.DefaultTransport}}
	
	// Set up the GraphQL client
	newClient :=  graphql.NewClient( os.Getenv("HASURA_GRAPHQL_ENDPOINT"), httpClient)
    
	// fmt.Println("newClient object", os.Getenv("HASURA_GRAPHQL_ENDPOINT"))
	// fmt.Println("admin_secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))

	return newClient
}