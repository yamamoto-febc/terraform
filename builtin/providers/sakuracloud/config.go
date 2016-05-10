package sakuracloud

import (
	API "github.com/yamamoto-febc/libsacloud/api"
)

type Config struct {
	AccessToken       string
	AccessTokenSecret string
	Zone              string
}

func (c *Config) NewClient() *API.Client {
	client := API.NewClient(c.AccessToken, c.AccessTokenSecret, c.Zone)
	//client.TraceMode = true
	return client
}
