/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// credentials.go - the credentials data structure definition

// Package auth implements the authorization functionality for BCE.
// It use the BCE access key ID and secret access key with the specific sign algorithm to generate
// the authorization string. It also supports the temporary authorization by the STS token.
package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// AuthType indicates which authentication mechanism is used.
type AuthType int

const (
	AuthTypeAKSK AuthType = iota
	AuthTypeAccessToken
	AuthTypeApiKey
)

// Credentials is the common interface implemented by all credential types.
type Credentials interface {
	GetAuthType() AuthType
	// GetSigner returns the signer used to authenticate requests for this credential.
	GetSigner() Signer
}

// BceCredentials define the data structure for authorization
type BceCredentials struct {
	AccessKeyId     string // access key id to the service
	SecretAccessKey string // secret access key to the service
	SessionToken    string // session token generate by the STS service
}

func (b *BceCredentials) GetAuthType() AuthType {
	return AuthTypeAKSK
}

func (b *BceCredentials) GetSigner() Signer {
	return &BceV1Signer{}
}

func (b *BceCredentials) String() string {
	str := "ak: " + b.AccessKeyId + ", sk: " + b.SecretAccessKey
	if len(b.SessionToken) != 0 {
		return str + ", sessionToken: " + b.SessionToken
	}
	return str
}

func NewBceCredentials(ak, sk string) (*BceCredentials, error) {
	if len(ak) == 0 {
		return nil, errors.New("accessKeyId should not be empty")
	}
	if len(sk) == 0 {
		return nil, errors.New("secretKey should not be empty")
	}

	return &BceCredentials{ak, sk, ""}, nil
}

func NewSessionBceCredentials(ak, sk, token string) (*BceCredentials, error) {
	if len(token) == 0 {
		return nil, errors.New("sessionToken should not be empty")
	}

	result, err := NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	result.SessionToken = token

	return result, nil
}

type tokenCache struct {
	token      string
	expireTime time.Time
	mu         sync.Mutex
}

// AccessTokenCredentials uses AK/SK to fetch and cache an access token from the
// Baidu AI Platform, refreshing it automatically before expiry.
type AccessTokenCredentials struct {
	ak    string
	sk    string
	cache *tokenCache
}

func (a *AccessTokenCredentials) GetAuthType() AuthType {
	return AuthTypeAccessToken
}

func (a *AccessTokenCredentials) GetSigner() Signer {
	return &BceAccessTokenSigner{}
}

func NewAccessTokenCredentials(ak, sk string) (*AccessTokenCredentials, error) {
	if len(ak) == 0 {
		return nil, errors.New("accessKeyId should not be empty")
	}
	if len(sk) == 0 {
		return nil, errors.New("secretKey should not be empty")
	}

	return &AccessTokenCredentials{ak, sk, &tokenCache{}}, nil
}

func (a *AccessTokenCredentials) GetAccessToken() (string, error) {
	a.cache.mu.Lock()
	defer a.cache.mu.Unlock()

	if a.cache.token != "" && time.Now().Before(a.cache.expireTime) {
		return a.cache.token, nil
	}
	token, expireIn, err := fetchAccessToken(a.ak, a.sk)
	if err != nil {
		return "", err
	}
	a.cache.token = token
	a.cache.expireTime = time.Now().Add(time.Duration(expireIn-24*3600) * time.Second)
	return token, nil
}

func fetchAccessToken(ak, sk string) (string, int, error) {
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", ak)
	params.Set("client_secret", sk)

	const maxRetry = 3
	var lastErr error
	for attempt := 1; attempt <= maxRetry; attempt++ {
		resp, err := http.PostForm("https://aip.baidubce.com/oauth/2.0/token", params)
		if err != nil {
			lastErr = fmt.Errorf("failed to fetch access token: %v", err)
			continue
		}

		if resp.StatusCode == 500 {
			resp.Body.Close()
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return "", 0, fmt.Errorf("failed to fetch access token: %v", err)
		}

		var result struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
			Error       string `json:"error"`
			ErrorDesc   string `json:"error_description"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			return "", 0, fmt.Errorf("failed to parse token response: %v", err)
		}
		if result.Error != "" {
			return "", 0, fmt.Errorf("token error: %s - %s", result.Error, result.ErrorDesc)
		}
		return result.AccessToken, result.ExpiresIn, nil
	}
	return "", 0, fmt.Errorf("failed to fetch access token after %d attempts: %v",
		maxRetry, lastErr)
}

// ApiKeyCredentials injects the API Key as a Bearer token in the Authorization header.
type ApiKeyCredentials struct {
	ApiKey string
}

func (a *ApiKeyCredentials) GetAuthType() AuthType {
	return AuthTypeApiKey
}

func (a *ApiKeyCredentials) GetSigner() Signer {
	return &BceApiKeySigner{}
}

func NewApiKeyCredentials(apiKey string) (*ApiKeyCredentials, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("apiKey should not be empty")
	}

	return &ApiKeyCredentials{apiKey}, nil
}
