package v2

import (
	sdk "github.com/tdak47/baidu-ai-go-sdk"
)

type FaceClient struct {
	*sdk.Client
}

func NewFaceClient(key, secret string) *FaceClient {
	return &FaceClient{
		Client: sdk.NewClient(key, secret),
	}
}
