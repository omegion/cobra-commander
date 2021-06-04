package client

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/client_mock.go -package=mocks github.com/omegion/go-cli/internal/client Interface
// Interface is an interface entrypoint for the application.
type Interface interface {
	ArithmeticInterface
}

// Client is an entrypoint to controllers.
type Client struct{}

// NewClient is a factory for Client.
func NewClient() *Client {
	return &Client{}
}
