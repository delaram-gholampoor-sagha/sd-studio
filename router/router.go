package router

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
)

func NewRouter() (*message.Router, error) {
	config := message.RouterConfig{}

	r, err := message.NewRouter(config, nil)
	if err != nil {
		return nil, err
	}

	r.AddPlugin(plugin.SignalsHandler)

	return r, nil
}
