package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	cli     *minio.Client
	baseURL string
}

func NewClient(
	endpoint string,
	accessKey string,
	secretKey string,
	useSSL bool,
	baseURL string,
) (*Client, error) {

	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		cli:     cli,
		baseURL: baseURL,
	}, nil
}

func (c *Client) Raw() *minio.Client {
	return c.cli
}
