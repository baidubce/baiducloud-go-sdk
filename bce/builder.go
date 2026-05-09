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

// builder.go - defines the RequestBuilder structure for BCE servies

package bce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"

	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

// RequestBuilder holds config data for bce request.
// Some of fields are required and the others are optional.
// The builder pattern can simplify the execution of requests.
type RequestBuilder struct {
	client Client

	url                 string            // required
	method              string            // required
	queryParams         map[string]string // optional
	headers             map[string]string // optional
	body                interface{}       // optional
	result              interface{}       // optional
	isFormEncoded       bool              // optional: true for application/x-www-form-urlencoded
	bodyReader          io.Reader         // optional: raw body reader (for stream/utils)
	useRequestBodyUtils bool              // optional: use RequestBodyUtils for encoding
	isStreamResponse    bool              // optional: true for stream response
	endpoint            string            // optional: 请求级别的 endpoint 覆盖（并发安全）
}

// create RequestBuilder with the client.
func NewRequestBuilder(client Client) *RequestBuilder {
	return &RequestBuilder{
		client: client,
	}
}

func (b *RequestBuilder) WithURL(url string) *RequestBuilder {
	b.url = url
	return b
}

func (b *RequestBuilder) WithMethod(method string) *RequestBuilder {
	b.method = method
	return b
}

// WithEndpoint 设置请求级别的 endpoint 覆盖（并发安全）
func (b *RequestBuilder) WithEndpoint(endpoint string) *RequestBuilder {
	b.endpoint = endpoint
	return b
}

// WithHost 设置请求级别的 Host 参数（并发安全）
func (b *RequestBuilder) WithHost(hostParam string) *RequestBuilder {
	if hostParam == "" {
		return b
	}
	originalEndpoint := b.client.GetBceClientConfig().Endpoint
	b.endpoint = util.BuildHostEndpoint(originalEndpoint, hostParam)
	return b
}

// set query param with the key/value directly.
func (b *RequestBuilder) WithQueryParam(key, value string) *RequestBuilder {
	if b.queryParams == nil {
		b.queryParams = make(map[string]string)
	}
	b.queryParams[key] = value
	return b
}

// set query param with the key/value only when the value is not blank.
func (b *RequestBuilder) WithQueryParamFilter(key, value string) *RequestBuilder {
	if len(value) == 0 {
		return b
	}
	return b.WithQueryParam(key, value)
}

// WithHeaderFilter sets a header only when the value is not blank.
func (b *RequestBuilder) WithHeaderFilter(key, value string) *RequestBuilder {
	if len(value) == 0 {
		return b
	}
	return b.WithHeader(key, value)
}

func (b *RequestBuilder) WithQueryParams(params map[string]string) *RequestBuilder {
	if b.queryParams == nil {
		b.queryParams = params
	} else {
		for key, value := range params {
			b.queryParams[key] = value
		}
	}
	return b
}

func (b *RequestBuilder) WithHeader(key, value string) *RequestBuilder {
	if b.headers == nil {
		b.headers = make(map[string]string)
	}
	b.headers[key] = value
	return b
}

func (b *RequestBuilder) WithHeaders(headers map[string]string) *RequestBuilder {
	if b.headers == nil {
		b.headers = headers
	} else {
		for key, value := range headers {
			b.headers[key] = value
		}
	}
	return b
}

func (b *RequestBuilder) WithBody(body interface{}) *RequestBuilder {
	b.body = body
	b.isFormEncoded = false
	return b
}

// WithFormBody sets the request body to be encoded as application/x-www-form-urlencoded.
// The body will be marshaled to JSON first, then converted to URL-encoded form data.
func (b *RequestBuilder) WithFormBody(body interface{}) *RequestBuilder {
	b.body = body
	b.isFormEncoded = true
	return b
}

func (b *RequestBuilder) WithResult(result interface{}) *RequestBuilder {
	b.result = result
	return b
}

// WithStreamResult sets the result to receive a stream response (no JSON parsing).
// The result should be a *StreamResponse pointer.
func (b *RequestBuilder) WithStreamResult(result interface{}) *RequestBuilder {
	b.result = result
	b.isStreamResponse = true
	return b
}

// WithBodyUsingUtils uses RequestBodyUtils to encode JSON request body
func (b *RequestBuilder) WithBodyUsingUtils(body interface{}) *RequestBuilder {
	b.body = body
	b.isFormEncoded = false
	b.useRequestBodyUtils = true
	return b
}

// WithBodyAndContentTypeUsingUtils uses RequestBodyUtils to encode JSON request body with custom Content-Type
func (b *RequestBuilder) WithBodyAndContentTypeUsingUtils(body interface{}, contentType string) *RequestBuilder {
	b.body = body
	b.isFormEncoded = false
	b.useRequestBodyUtils = true
	if contentType != "" {
		b.WithHeader("Content-Type", contentType)
	}
	return b
}

// WithFormBodyUsingUtils uses RequestBodyUtils to encode form request body
func (b *RequestBuilder) WithFormBodyUsingUtils(body interface{}) *RequestBuilder {
	b.body = body
	b.isFormEncoded = true
	b.useRequestBodyUtils = true
	return b
}

// WithFormBodyAndContentTypeUsingUtils uses RequestBodyUtils to encode form request body with custom Content-Type
func (b *RequestBuilder) WithFormBodyAndContentTypeUsingUtils(body interface{}, contentType string) *RequestBuilder {
	b.body = body
	b.isFormEncoded = true
	b.useRequestBodyUtils = true
	if contentType != "" {
		b.WithHeader("Content-Type", contentType)
	}
	return b
}

// WithStreamBodyUsingUtils uses RequestBodyUtils to encode stream request body
func (b *RequestBuilder) WithStreamBodyUsingUtils(reader io.Reader) *RequestBuilder {
	b.bodyReader = reader
	b.useRequestBodyUtils = true
	return b
}

// WithStreamBodyAndContentTypeUsingUtils uses RequestBodyUtils to encode stream request body with custom Content-Type
func (b *RequestBuilder) WithStreamBodyAndContentTypeUsingUtils(reader io.Reader, contentType string) *RequestBuilder {
	b.bodyReader = reader
	b.useRequestBodyUtils = true
	if contentType != "" {
		b.WithHeader("Content-Type", contentType)
	}
	return b
}

// WithBodyReader sets raw body reader directly (for advanced use cases)
func (b *RequestBuilder) WithBodyReader(reader io.Reader) *RequestBuilder {
	b.bodyReader = reader
	return b
}

// SetBody implements util.RequestFiller interface
// 使得 RequestBuilder 可以直接被 util.FillRequestAsStream 等方法使用
func (b *RequestBuilder) SetBody(body io.Reader, contentType string, contentLength int64) {
	b.bodyReader = body
	b.useRequestBodyUtils = false // 已经处理过了，不需要再次处理
	if contentType != "" {
		b.WithHeader("Content-Type", contentType)
	}
	if contentLength >= 0 {
		b.WithHeader("Content-Length", fmt.Sprintf("%d", contentLength))
	}
}

// SetHeader implements util.RequestFiller interface
func (b *RequestBuilder) SetHeader(key, value string) {
	b.WithHeader(key, value)
}

// Do will send request to bce and get result with the builder's parameters.
func (b *RequestBuilder) Do() error {
	if err := b.validate(); err != nil {
		return err
	}

	// build BceRequest
	req, err := b.buildBceRequest()
	if err != nil {
		return err
	}

	// get result from BceResponse
	if err := b.buildBceResponse(req); err != nil {
		return err
	}

	return nil
}

// Validate if the required fields are providered.
func (b *RequestBuilder) validate() error {
	if len(b.url) == 0 {
		return fmt.Errorf("The url can't be null.")
	}
	if len(b.method) == 0 {
		return fmt.Errorf("The method can't be null.")
	}
	if b.client == nil {
		return fmt.Errorf("The client can't be null.")
	}
	return nil
}

func (b *RequestBuilder) buildBceRequest() (*BceRequest, error) {
	// Build the request
	req := &BceRequest{}
	req.SetUri(b.url)
	req.SetMethod(b.method)

	// 设置 endpoint（并发安全：优先使用请求级别的 endpoint）
	if b.endpoint != "" {
		req.SetEndpoint(b.endpoint)
	} else {
		req.SetEndpoint(b.client.GetBceClientConfig().Endpoint)
	}

	if b.headers != nil {
		req.SetHeaders(b.headers)
	}
	if b.queryParams != nil {
		req.SetParams(b.queryParams)
	}

	// Handle body encoding
	if b.bodyReader != nil {
		// Use raw body reader (stream or already encoded)
		if b.useRequestBodyUtils {
			// Use RequestBodyUtils for stream
			customContentType := ""
			if b.headers != nil {
				if ct, ok := b.headers["Content-Type"]; ok {
					customContentType = ct
				}
			}

			reader, contentType, contentLength, err := util.FillPayloadAsStreamWithContentType(
				b.bodyReader, 0, customContentType)
			if err != nil {
				return nil, fmt.Errorf("failed to encode stream body with RequestBodyUtils: %v", err)
			}

			// Read all from reader into bytes for Body
			bodyBytes, err := io.ReadAll(reader)
			if err != nil {
				return nil, fmt.Errorf("failed to read from stream: %v", err)
			}

			body, err := NewBodyFromBytes(bodyBytes)
			if err != nil {
				return nil, err
			}
			req.SetBody(body)

			// Set headers
			b.WithHeader("Content-Type", contentType)
			if contentLength >= 0 {
				b.WithHeader("Content-Length", fmt.Sprintf("%d", contentLength))
			}
			req.SetHeaders(b.headers)
		} else {
			// Use raw reader as-is
			bodyBytes, err := io.ReadAll(b.bodyReader)
			if err != nil {
				return nil, fmt.Errorf("failed to read from body reader: %v", err)
			}
			body, err := NewBodyFromBytes(bodyBytes)
			if err != nil {
				return nil, err
			}
			req.SetBody(body)
		}
	} else if b.body != nil {
		var bodyBytes []byte
		var err error

		if b.useRequestBodyUtils {
			// Use RequestBodyUtils for encoding
			var reader io.Reader
			var contentType string
			var contentLength int64

			customContentType := ""
			if b.headers != nil {
				if ct, ok := b.headers["Content-Type"]; ok {
					customContentType = ct
				}
			}

			if b.isFormEncoded {
				// Form encoding with RequestBodyUtils
				if customContentType != "" {
					reader, contentType, contentLength, err = util.FillPayloadAsFormWithContentType(b.body, customContentType)
				} else {
					reader, contentType, contentLength, err = util.FillPayloadAsForm(b.body)
				}
			} else {
				// JSON encoding with RequestBodyUtils
				if customContentType != "" {
					reader, contentType, contentLength, err = util.FillPayloadAsJSONWithContentType(b.body, customContentType)
				} else {
					reader, contentType, contentLength, err = util.FillPayloadAsJSON(b.body)
				}
			}

			if err != nil {
				return nil, fmt.Errorf("failed to encode body with RequestBodyUtils: %v", err)
			}

			// Read from reader
			bodyBytes, err = io.ReadAll(reader)
			if err != nil {
				return nil, fmt.Errorf("failed to read from RequestBodyUtils reader: %v", err)
			}

			// Set headers
			b.WithHeader("Content-Type", contentType)
			b.WithHeader("Content-Length", fmt.Sprintf("%d", contentLength))
			req.SetHeaders(b.headers)
		} else {
			// Legacy encoding (for backward compatibility)
			if b.isFormEncoded {
				// Form encoding
				bodyBytes, err = encodeFormData(b.body)
			} else {
				// JSON encoding (default)
				bodyBytes, err = json.Marshal(b.body)
			}

			if err != nil {
				return nil, err
			}
		}

		body, err := NewBodyFromBytes(bodyBytes)
		if err != nil {
			return nil, err
		}
		req.SetBody(body)
	}

	return req, nil
}

func (b *RequestBuilder) buildBceResponse(req *BceRequest) error {
	// Send request and get response
	resp := &BceResponse{}
	if err := b.client.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	if b.result == nil {
		resp.Body().Close()
		return nil
	}

	// Handle stream response
	if b.isStreamResponse {
		streamResp := resp.ParseStreamBody()
		if sr, ok := b.result.(*StreamResponse); ok {
			sr.Body = streamResp.Body
			sr.ContentType = streamResp.ContentType
			sr.ContentLength = streamResp.ContentLength
			sr.Metadata = streamResp.Metadata
			return nil
		}

		// If result has a SetStream method, call it
		if receiver, ok := b.result.(interface {
			SetStream(*StreamResponse)
		}); ok {
			receiver.SetStream(streamResp)
			return nil
		}

		// Fallback: close the body and return error
		resp.Body().Close()
		return fmt.Errorf("result type does not support stream response, expected *StreamResponse or implement SetStream(*StreamResponse)")
	}

	defer resp.Body().Close()

	// Check if response body is empty - some APIs (like PutObject) return empty body
	// In this case, skip JSON parsing and just inject metadata
	contentLength := resp.Header(http.CONTENT_LENGTH)
	if contentLength == "0" || contentLength == "" {
		// Try to read a byte to check if body is actually empty
		body := resp.Body()
		buf := make([]byte, 1)
		n, err := body.Read(buf)
		if err == io.EOF || n == 0 {
			// Body is empty, skip JSON parsing
			// Inject MetaData if supported
			if receiver, ok := b.result.(interface {
				SetMetaData(map[string]string)
			}); ok {
				// Layer 2: Set known metadata
				receiver.SetMetaData(resp.Headers())
			}
			return nil
		}
		// Body is not empty, we need to put the byte back and continue parsing
		// Create a new reader that includes the byte we read
		combinedReader := io.MultiReader(bytes.NewReader(buf[:n]), body)
		// Parse using combined reader
		jsonDecoder := json.NewDecoder(combinedReader)
		if err := jsonDecoder.Decode(b.result); err != nil {
			// If JSON parsing fails, it might be an empty or non-JSON response
			// Just return nil for compatibility
			if err == io.EOF {
				// Inject MetaData if supported
				if receiver, ok := b.result.(interface {
					SetMetaData(map[string]string)
				}); ok {
					receiver.SetMetaData(resp.Headers())
				}
				return nil
			}
			return err
		}
		// Inject MetaData
		if receiver, ok := b.result.(interface {
			SetMetaData(map[string]string)
		}); ok {
			receiver.SetMetaData(resp.Headers())
		}
		return nil
	}

	// Parse JSON body
	if err := resp.ParseJsonBody(b.result); err != nil {
		// If JSON parsing fails with EOF, it means body is empty
		if err == io.EOF {
			// Inject MetaData if supported
			if receiver, ok := b.result.(interface {
				SetMetaData(map[string]string)
			}); ok {
				receiver.SetMetaData(resp.Headers())
			}
			return nil
		}
		return err
	}

	// Inject MetaData
	if receiver, ok := b.result.(interface {
		SetMetaData(map[string]string)
	}); ok {
		receiver.SetMetaData(resp.Headers())
	}

	return nil
}

// encodeFormData converts a struct or map to application/x-www-form-urlencoded format.
// It first marshals the data to JSON, then converts to a map, and finally URL-encodes it.
func encodeFormData(data interface{}) ([]byte, error) {
	// Marshal to JSON first to normalize the data structure
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data to JSON: %v", err)
	}

	// Unmarshal JSON to map
	var dataMap map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &dataMap); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to map: %v", err)
	}

	// Build URL-encoded form data
	values := url.Values{}
	for key, value := range dataMap {
		// Skip nil and empty values
		if value == nil {
			continue
		}

		// Convert value to string
		strValue := formatFormValue(value)
		if strValue == "" {
			continue
		}

		values.Add(key, strValue)
	}

	return []byte(values.Encode()), nil
}

// formatFormValue converts various types to string for form encoding
func formatFormValue(value interface{}) string {
	if value == nil {
		return ""
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		str := v.String()
		if str == "" {
			return ""
		}
		return str
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", v.Float())
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
		// For complex types, convert to JSON string
		// Note: This is not ideal for form data but maintains compatibility
		bytes, err := json.Marshal(value)
		if err != nil {
			return fmt.Sprintf("%v", value)
		}
		return string(bytes)
	default:
		return fmt.Sprintf("%v", value)
	}
}
