package easypostdevtools

import "github.com/EasyPost/easypost-go"

type Batch struct {
	client *easypost.Client
	tools  *EasyPostDevTools
}
