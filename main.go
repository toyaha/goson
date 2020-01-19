package goson

func NewClient(config *Config, metaList ...interface{}) *Client {
	client := &Client{}
	client.Init(config, metaList...)
	return client
}

func NewClientDefault(metaList ...interface{}) *Client {
	return NewClient(nil, metaList...)
}

func NewConfig() *Config {
	config := &Config{}
	config.Init()
	return config
}
