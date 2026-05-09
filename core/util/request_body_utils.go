/*
 * Copyright 2024 Baidu, Inc.
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

package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	ContentTypeJSON = "application/json; charset=utf-8"

	ContentTypeForm = "application/x-www-form-urlencoded; charset=utf-8"

	ContentTypeOctetStream = "application/octet-stream"

	ContentTypeMultipart = "multipart/form-data"
)

// BodyType represents the type of request body encoding
type BodyType string

const (
	BodyTypeJSON BodyType = "application/json; charset=utf-8"

	BodyTypeForm BodyType = "application/x-www-form-urlencoded; charset=utf-8"

	BodyTypeBinary BodyType = "application/octet-stream"

	BodyTypeXML BodyType = "application/xml; charset=utf-8"

	BodyTypeNone BodyType = ""
)

// GetContentType returns the Content-Type header value for this body type
func (bt BodyType) GetContentType() string {
	return string(bt)
}

// HasContentType checks if this body type has a content type
func (bt BodyType) HasContentType() bool {
	return bt != BodyTypeNone && bt != ""
}

func FillPayload(request interface{}, bodyType BodyType) (io.Reader, string, int64, error) {
	return FillPayloadWithCustomContentType(request, bodyType, "")
}

func FillPayloadWithCustomContentType(request interface{}, bodyType BodyType, customContentType string) (io.Reader, string, int64, error) {
	if !bodyType.HasContentType() {
		return nil, "", 0, nil
	}

	var bodyBytes []byte
	var err error

	switch bodyType {
	case BodyTypeJSON:
		bodyBytes, err = encodeAsJSON(request)
		if err != nil {
			return nil, "", 0, fmt.Errorf("failed to encode request as JSON: %w", err)
		}
	case BodyTypeForm:
		bodyBytes, err = encodeAsForm(request)
		if err != nil {
			return nil, "", 0, fmt.Errorf("failed to encode request as form: %w", err)
		}
	case BodyTypeXML:
		return nil, "", 0, fmt.Errorf("XML body type is not yet implemented")
	default:
		return nil, "", 0, fmt.Errorf("unsupported BodyType: %s", bodyType)
	}

	// 如果用户指定了自定义 contentType，使用自定义值；否则使用 bodyType 的默认值
	finalContentType := customContentType
	if finalContentType == "" {
		finalContentType = bodyType.GetContentType()
	}

	return bytes.NewReader(bodyBytes), finalContentType, int64(len(bodyBytes)), nil
}

func FillPayloadAsStreamWithBodyType(reader io.Reader, bodyType BodyType) (io.Reader, string, int64, error) {
	return FillPayloadAsStreamWithBodyTypeAndLengthAndContentType(reader, 0, bodyType, "")
}

func FillPayloadAsStreamWithBodyTypeAndLength(reader io.Reader, contentLength int64, bodyType BodyType) (io.Reader, string, int64, error) {
	return FillPayloadAsStreamWithBodyTypeAndLengthAndContentType(reader, contentLength, bodyType, "")
}

func FillPayloadAsStreamWithBodyTypeAndLengthAndContentType(reader io.Reader, contentLength int64, bodyType BodyType, customContentType string) (io.Reader, string, int64, error) {
	if reader == nil {
		return nil, "", 0, errors.New("reader cannot be nil")
	}

	// 如果用户没有指定 contentLength,自动计算
	if contentLength == 0 {
		var err error
		contentLength, err = calculateContentLength(reader)
		if err != nil {
			return nil, "", 0, err
		}
	}

	// 如果用户指定了自定义 contentType，使用自定义值；否则使用 bodyType 的默认值
	finalContentType := customContentType
	if finalContentType == "" {
		finalContentType = bodyType.GetContentType()
	}

	return reader, finalContentType, contentLength, nil
}

func FillPayloadAsByteArrayWithBodyType(data []byte, bodyType BodyType) (io.Reader, string, int64, error) {
	if data == nil {
		return nil, "", 0, errors.New("data cannot be nil")
	}

	return bytes.NewReader(data), bodyType.GetContentType(), int64(len(data)), nil
}

func FillPayloadAsJSON(request interface{}) (io.Reader, string, int64, error) {
	return FillPayload(request, BodyTypeJSON)
}

func FillPayloadAsJSONWithContentType(request interface{}, contentType string) (io.Reader, string, int64, error) {
	return FillPayloadWithCustomContentType(request, BodyTypeJSON, contentType)
}

func FillPayloadAsForm(request interface{}) (io.Reader, string, int64, error) {
	return FillPayload(request, BodyTypeForm)
}

func FillPayloadAsFormWithContentType(request interface{}, contentType string) (io.Reader, string, int64, error) {
	return FillPayloadWithCustomContentType(request, BodyTypeForm, contentType)
}

func FillPayloadAsStream(reader io.Reader) (io.Reader, string, int64, error) {
	return FillPayloadAsStreamWithContentType(reader, 0, "")
}

func FillPayloadAsStreamWithContentType(reader io.Reader, contentLength int64, contentType string) (io.Reader, string, int64, error) {
	if reader == nil {
		return nil, "", 0, errors.New("reader cannot be nil")
	}

	if contentType == "" {
		contentType = ContentTypeOctetStream
	}

	// 如果用户没有指定 contentLength,自动计算
	if contentLength == 0 {
		var err error
		contentLength, err = calculateContentLength(reader)
		if err != nil {
			return nil, "", 0, err
		}
	}

	return reader, contentType, contentLength, nil
}

func FillPayloadAsByteArray(data []byte, contentType string) (io.Reader, string, int64, error) {
	if data == nil {
		return nil, "", 0, errors.New("data cannot be nil")
	}

	if contentType == "" {
		contentType = ContentTypeOctetStream
	}

	return bytes.NewReader(data), contentType, int64(len(data)), nil
}

// encodeAsJSON 将请求对象编码为 JSON 格式
func encodeAsJSON(request interface{}) ([]byte, error) {
	if request == nil {
		return []byte("{}"), nil
	}

	return json.Marshal(request)
}

func encodeAsForm(request interface{}) ([]byte, error) {
	if request == nil {
		return []byte(""), nil
	}

	// 1. 将对象转换为 JSON,然后解析为 map
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	var fieldMap map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &fieldMap); err != nil {
		return nil, err
	}

	// 2. 构建表单格式: key1=value1&key2=value2
	formValues := url.Values{}
	for key, value := range fieldMap {
		// 跳过 nil 值
		if value == nil {
			continue
		}

		// 转换值为字符串
		strValue := formatFormValue(value)
		if strValue != "" {
			formValues.Add(key, strValue)
		}
	}

	return []byte(formValues.Encode()), nil
}

// formatFormValue 格式化表单值
func formatFormValue(value interface{}) string {
	if value == nil {
		return ""
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	default:
		// 复杂类型转换为 JSON 字符串
		if jsonBytes, err := json.Marshal(value); err == nil {
			return string(jsonBytes)
		}
		return fmt.Sprintf("%v", value)
	}
}

func calculateContentLength(reader io.Reader) (int64, error) {
	if buf, ok := reader.(*bytes.Buffer); ok {
		return int64(buf.Len()), nil
	}
	if r, ok := reader.(*bytes.Reader); ok {
		return int64(r.Len()), nil
	}
	if file, ok := reader.(*os.File); ok {
		fileInfo, err := file.Stat()
		if err != nil {
			return 0, fmt.Errorf("failed to get file info: %w", err)
		}
		return fileInfo.Size(), nil
	}

	if seeker, ok := reader.(io.ReadSeeker); ok {
		currentPos, err := seeker.Seek(0, io.SeekCurrent)
		if err != nil {
			return 0, fmt.Errorf("failed to get current position: %w", err)
		}

		endPos, err := seeker.Seek(0, io.SeekEnd)
		if err != nil {
			return 0, fmt.Errorf("failed to seek to end: %w", err)
		}
		_, err = seeker.Seek(currentPos, io.SeekStart)
		if err != nil {
			return 0, fmt.Errorf("failed to restore position: %w", err)
		}

		return endPos - currentPos, nil
	}
	return -1, nil
}

// ========== Host 处理工具方法 ==========
func BuildHostEndpoint(endpoint string, hostParam string) string {
	if hostParam == "" {
		return endpoint
	}

	// 检查是否包含协议
	if idx := strings.Index(endpoint, "://"); idx != -1 {
		protocol := endpoint[:idx]
		host := endpoint[idx+3:]
		return protocol + "://" + hostParam + "." + host
	}

	return hostParam + "." + endpoint
}

// ========== 请求体填充方法（直接操作请求，类似 Java 的 RequestBodyUtils）==========
type RequestFiller interface {
	SetBody(body io.Reader, contentType string, contentLength int64)
	SetHeader(key, value string)
}

func FillRequestAsJSON(filler RequestFiller, request interface{}) error {
	return FillRequestAsJSONWithContentType(filler, request, "")
}

// FillRequestAsJSONWithContentType 使用 JSON 格式填充请求体（自定义 Content-Type）
func FillRequestAsJSONWithContentType(filler RequestFiller, request interface{}, contentType string) error {
	reader, ct, length, err := FillPayloadAsJSONWithContentType(request, contentType)
	if err != nil {
		return err
	}
	filler.SetBody(reader, ct, length)
	return nil
}

// FillRequestAsForm 使用表单格式填充请求体（类似 Java 的 fillPayloadAsForm）
func FillRequestAsForm(filler RequestFiller, request interface{}) error {
	return FillRequestAsFormWithContentType(filler, request, "")
}

// FillRequestAsFormWithContentType 使用表单格式填充请求体（自定义 Content-Type）
func FillRequestAsFormWithContentType(filler RequestFiller, request interface{}, contentType string) error {
	reader, ct, length, err := FillPayloadAsFormWithContentType(request, contentType)
	if err != nil {
		return err
	}
	filler.SetBody(reader, ct, length)
	return nil
}

// FillRequestAsStream 使用二进制流格式填充请求体（类似 Java 的 fillPayloadAsStream）
func FillRequestAsStream(filler RequestFiller, reader io.Reader) error {
	return FillRequestAsStreamWithContentType(filler, reader, "")
}

// FillRequestAsStreamWithContentType 使用二进制流格式填充请求体（自定义 Content-Type）
func FillRequestAsStreamWithContentType(filler RequestFiller, reader io.Reader, contentType string) error {
	r, ct, length, err := FillPayloadAsStreamWithContentType(reader, 0, contentType)
	if err != nil {
		return err
	}
	filler.SetBody(r, ct, length)
	return nil
}

// FillRequestAsByteArray 使用字节数组填充请求体
func FillRequestAsByteArray(filler RequestFiller, data []byte, contentType string) error {
	reader, ct, length, err := FillPayloadAsByteArray(data, contentType)
	if err != nil {
		return err
	}
	filler.SetBody(reader, ct, length)
	return nil
}
