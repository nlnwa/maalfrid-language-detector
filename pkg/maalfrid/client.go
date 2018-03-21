// Copyright © 2017 Marius Elsfjordstrand Beck <marius.beck@nb.no>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maalfrid

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/nlnwa/maalfrid-language-detector/api"
	"github.com/pkg/errors"
)

type Client struct {
	api.MaalfridClient
	address string
	conn    *grpc.ClientConn
}

type clientOption func(*Client)

func WithAddress(host string, port int) clientOption {
	return func(c *Client) {
		c.address = fmt.Sprintf("%s:%d", host, port)
	}
}

func NewApiClient(opts ...clientOption) *Client {
	client := new(Client)

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) Hangup() error {
	return c.conn.Close()
}

func (c *Client) Dial() error {
	var err error
	if c.conn, err = grpc.Dial(c.address, grpc.WithInsecure()); err != nil {
		return errors.Wrapf(err, "failed to dial: %s", c.address)
	} else {
		c.MaalfridClient = api.NewMaalfridClient(c.conn)
		return nil
	}
}

func (c *Client) DetectLanguage(text string) ([]*api.Language, error) {
	req := &api.DetectLanguageRequest{Text: text}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if res, err := c.MaalfridClient.DetectLanguage(ctx, req); err != nil {
		return nil, err
	} else {
		return res.GetLanguages(), nil
	}
}
