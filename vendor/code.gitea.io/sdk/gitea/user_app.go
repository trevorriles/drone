// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// BasicAuthEncode generate base64 of basic auth head
func BasicAuthEncode(user, pass string) string {
	return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
}

// AccessToken represents a API access token.
// swagger:response AccessToken
type AccessToken struct {
	Name string `json:"name"`
	Sha1 string `json:"sha1"`
}

// AccessTokenList represents a list of API access token.
// swagger:response AccessTokenList
type AccessTokenList []*AccessToken 

// ListAccessTokens lista all the access tokens of user
func (c *Client) ListAccessTokens(user, pass string) ([]*AccessToken, error) {
	tokens := make([]*AccessToken, 0, 10)
	return tokens, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/tokens", user),
		http.Header{"Authorization": []string{"Basic " + BasicAuthEncode(user, pass)}}, nil, &tokens)
}

// CreateAccessTokenOption options when create access token
// swagger:parameters userCreateToken
type CreateAccessTokenOption struct {
	Name string `json:"name" binding:"Required"`
}

// CreateAccessToken create one access token with options
func (c *Client) CreateAccessToken(user, pass string, opt CreateAccessTokenOption) (*AccessToken, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	t := new(AccessToken)
	return t, c.getParsedResponse("POST", fmt.Sprintf("/users/%s/tokens", user),
		http.Header{
			"content-type":  []string{"application/json"},
			"Authorization": []string{"Basic " + BasicAuthEncode(user, pass)}},
		bytes.NewReader(body), t)
}
