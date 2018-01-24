package maalfrid

import (
	"context"
	"strings"

	"github.com/kapsteur/franco"

	"github.com/nlnwa/maalfrid/api"
)

type maalfridApi struct{}

func NewApiServer() api.MaalfridServer {
	return new(maalfridApi)
}

func (m *maalfridApi) DetectLanguage(ctx context.Context, req *api.DetectLanguageRequest) (*api.DetectLanguageReply, error) {
	var languages []*api.Language

	res := franco.Detect(req.Text)

	for i := range res[:5] {
		code := api.Code(api.Code_value[strings.ToUpper(res[i].Code)])
		l := &api.Language{Code: code, Count: res[i].Count}
		languages = append(languages, l)
	}

	return &api.DetectLanguageReply{Languages: languages}, nil
}
