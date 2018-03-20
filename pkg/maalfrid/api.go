package maalfrid

import (
	"context"
	"strings"

	"github.com/kapsteur/franco"

	"github.com/nlnwa/maalfrid-language-detector/api"
)

type maalfridApi struct {
	// limit number of suggested languages in response
	limit int
}

type serverOption func(*maalfridApi)

func WithLimit(n int) serverOption {
	return func(srv *maalfridApi) {
		srv.limit = n
	}
}

func NewApiServer(opts ...serverOption) api.MaalfridServer {
	srv := new(maalfridApi)
	// apply functional options
	for _, opt := range opts {
		opt(srv)
	}
	return srv
}

func (m *maalfridApi) DetectLanguage(ctx context.Context, req *api.DetectLanguageRequest) (*api.DetectLanguageReply, error) {
	var languages []*api.Language

	res := franco.Detect(req.Text)

	limit := m.limit
	if len(res) < m.limit {
		limit = len(res)
	}
	for i := range res[:limit] {
		code := api.Code(api.Code_value[strings.ToUpper(res[i].Code)])
		l := &api.Language{Code: code, Count: res[i].Count}
		languages = append(languages, l)
	}

	return &api.DetectLanguageReply{Languages: languages}, nil
}
