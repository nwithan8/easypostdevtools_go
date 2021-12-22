package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Webhooks struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}

func (w *Webhooks) GetMap() map[string]interface{} {
	_map := map[string]interface{}{
		"url": "https://www.example.com/webhooks",
	}
	return _map
}

func (w *Webhooks) Get() easypost.Webhook {
	// create webhook locally + POST to API
	_map := w.GetMap()
	webhook, err := w.client.CreateWebhook(_map["url"].(string))
	if err != nil {
		panic("Could not create webhook")
	}
	return *webhook
}
