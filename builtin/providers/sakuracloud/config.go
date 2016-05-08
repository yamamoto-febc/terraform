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
	return API.NewClient(c.AccessToken, c.AccessTokenSecret, c.Zone)
}
