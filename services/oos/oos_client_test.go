package oos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	stdhttp "net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/bce"
	bcehttp "github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("ak", "sk", "")
	if err != nil {
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	if got := client.GetBceClientConfig().Endpoint; got != DEFAULT_ENDPOINT {
		t.Errorf("NewClient() endpoint = %q, want %q", got, DEFAULT_ENDPOINT)
	}

	if _, err := NewClient("", "sk", "example.com"); err == nil || err.Error() != "accessKeyId should not be empty" {
		t.Errorf("NewClient() empty ak error = %v, want accessKeyId should not be empty", err)
	}
	if _, err := NewClient("ak", "", "example.com"); err == nil || err.Error() != "secretKey should not be empty" {
		t.Errorf("NewClient() empty sk error = %v, want secretKey should not be empty", err)
	}
}

func TestOOSV2URIHelpers(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"check template", getCheckTemplateV2Uri(VERSION_V2), "/v2/template/check"},
		{"create execution", getCreateExecutionV2Uri(VERSION_V2), "/v2/execution"},
		{"create template", getCreateTemplateV2Uri(VERSION_V2), "/v2/template"},
		{"delete template", getDeleteTemplateV2Uri(VERSION_V2), "/v2/template"},
		{"get execution detail", getGetExecutionDetailV2Uri(VERSION_V2), "/v2/execution"},
		{"get execution list", getGetExecutionListV2Uri(VERSION_V2), "/v2/execution/list"},
		{"get operator list", getGetOperatorListV2Uri(VERSION_V2), "/v2/operator/list"},
		{"get task children list", getGetTaskChildrenListV2Uri(VERSION_V2), "/v2/task/children/list"},
		{"get task detail", getGetTaskDetailV2Uri(VERSION_V2), "/v2/task"},
		{"get template detail", getGetTemplateDetailV2Uri(VERSION_V2), "/v2/template"},
		{"get template list", getGetTemplateListV2Uri(VERSION_V2), "/v2/template/list"},
		{"update template", getUpdateTemplateV2Uri(VERSION_V2), "/v2/template"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("uri = %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestOOSClientV2Requests(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		path       string
		query      map[string]string
		bodyFields map[string]interface{}
		call       func(*Client) (interface{}, error)
	}{
		{
			name:       "check template posts template body",
			method:     bcehttp.POST,
			path:       "/v2/template/check",
			bodyFields: map[string]interface{}{"name": "tmpl-check", "description": "validate template", "linear": true, "parallelism": float64(2)},
			call: func(client *Client) (interface{}, error) {
				return client.CheckTemplateV2(&CheckTemplateV2Request{
					Name:        util.PtrString("tmpl-check"),
					Description: util.PtrString("validate template"),
					Linear:      util.PtrBool(true),
					Parallelism: util.PtrInt32(2),
				})
			},
		},
		{
			name:       "create execution sends locale query and body",
			method:     bcehttp.POST,
			path:       "/v2/execution",
			query:      map[string]string{"locale": "zh-cn"},
			bodyFields: map[string]interface{}{"description": "run template", "parallelism": float64(3), "manually": true},
			call: func(client *Client) (interface{}, error) {
				return client.CreateExecutionV2(&CreateExecutionV2Request{
					Locale:      util.PtrString("zh-cn"),
					Description: util.PtrString("run template"),
					Parallelism: util.PtrInt32(3),
					Manually:    util.PtrBool(true),
				})
			},
		},
		{
			name:       "create execution omits blank locale query",
			method:     bcehttp.POST,
			path:       "/v2/execution",
			bodyFields: map[string]interface{}{"description": "run without locale"},
			call: func(client *Client) (interface{}, error) {
				return client.CreateExecutionV2(&CreateExecutionV2Request{
					Locale:      util.PtrString(""),
					Description: util.PtrString("run without locale"),
				})
			},
		},
		{
			name:       "create template posts template body",
			method:     bcehttp.POST,
			path:       "/v2/template",
			bodyFields: map[string]interface{}{"name": "tmpl-create", "description": "new template", "linear": false, "parallelism": float64(1)},
			call: func(client *Client) (interface{}, error) {
				return client.CreateTemplateV2(&CreateTemplateV2Request{
					Name:        util.PtrString("tmpl-create"),
					Description: util.PtrString("new template"),
					Linear:      util.PtrBool(false),
					Parallelism: util.PtrInt32(1),
				})
			},
		},
		{
			name:   "delete template sends namespace and id query",
			method: bcehttp.DELETE,
			path:   "/v2/template",
			query:  map[string]string{"namespace": "default", "id": "tmpl-001"},
			call: func(client *Client) (interface{}, error) {
				return client.DeleteTemplateV2(&DeleteTemplateV2Request{
					Namespace: util.PtrString("default"),
					Id:        util.PtrString("tmpl-001"),
				})
			},
		},
		{
			name:   "get execution detail sends all filters",
			method: bcehttp.GET,
			path:   "/v2/execution",
			query:  map[string]string{"namespace": "default", "id": "exec-001", "withLog": "true", "locale": "en-us"},
			call: func(client *Client) (interface{}, error) {
				return client.GetExecutionDetailV2(&GetExecutionDetailV2Request{
					Namespace: util.PtrString("default"),
					Id:        util.PtrString("exec-001"),
					WithLog:   util.PtrString("true"),
					Locale:    util.PtrString("en-us"),
				})
			},
		},
		{
			name:       "get execution list posts locale query and filters",
			method:     bcehttp.POST,
			path:       "/v2/execution/list",
			query:      map[string]string{"locale": "zh-cn"},
			bodyFields: map[string]interface{}{"namespace": "default", "state": "SUCCESS", "pageNo": float64(2), "pageSize": float64(20)},
			call: func(client *Client) (interface{}, error) {
				return client.GetExecutionListV2(&GetExecutionListV2Request{
					Locale:    util.PtrString("zh-cn"),
					Namespace: util.PtrString("default"),
					State:     util.PtrString("SUCCESS"),
					PageNo:    util.PtrInt32(2),
					PageSize:  util.PtrInt32(20),
				})
			},
		},
		{
			name:       "get operator list posts paging body",
			method:     bcehttp.POST,
			path:       "/v2/operator/list",
			query:      map[string]string{"locale": "zh-cn"},
			bodyFields: map[string]interface{}{"sort": "name", "ascending": true, "pageNo": float64(1), "pageSize": float64(10)},
			call: func(client *Client) (interface{}, error) {
				return client.GetOperatorListV2(&GetOperatorListV2Request{
					Locale:    util.PtrString("zh-cn"),
					Sort:      util.PtrString("name"),
					Ascending: util.PtrBool(true),
					PageNo:    util.PtrInt32(1),
					PageSize:  util.PtrInt32(10),
				})
			},
		},
		{
			name:       "get task children list posts task filters",
			method:     bcehttp.POST,
			path:       "/v2/task/children/list",
			query:      map[string]string{"locale": "zh-cn"},
			bodyFields: map[string]interface{}{"executionId": "exec-001", "taskId": "task-001", "namespace": "default", "pageNo": float64(1), "pageSize": float64(5)},
			call: func(client *Client) (interface{}, error) {
				return client.GetTaskChildrenListV2(&GetTaskChildrenListV2Request{
					Locale:      util.PtrString("zh-cn"),
					ExecutionId: util.PtrString("exec-001"),
					TaskId:      util.PtrString("task-001"),
					Namespace:   util.PtrString("default"),
					PageNo:      util.PtrInt32(1),
					PageSize:    util.PtrInt32(5),
				})
			},
		},
		{
			name:   "get task detail sends all query params",
			method: bcehttp.GET,
			path:   "/v2/task",
			query:  map[string]string{"namespace": "default", "dagId": "dag-001", "taskId": "task-001", "ignoreChildren": "false", "locale": "zh-cn"},
			call: func(client *Client) (interface{}, error) {
				return client.GetTaskDetailV2(&GetTaskDetailV2Request{
					Namespace:      util.PtrString("default"),
					DagId:          util.PtrString("dag-001"),
					TaskId:         util.PtrString("task-001"),
					IgnoreChildren: util.PtrString("false"),
					Locale:         util.PtrString("zh-cn"),
				})
			},
		},
		{
			name:   "get template detail sends identity query params",
			method: bcehttp.GET,
			path:   "/v2/template",
			query:  map[string]string{"namespace": "default", "id": "tmpl-001", "name": "template", "type": "private", "locale": "zh-cn"},
			call: func(client *Client) (interface{}, error) {
				return client.GetTemplateDetailV2(&GetTemplateDetailV2Request{
					Namespace: util.PtrString("default"),
					Id:        util.PtrString("tmpl-001"),
					Name:      util.PtrString("template"),
					Type:      util.PtrString("private"),
					Locale:    util.PtrString("zh-cn"),
				})
			},
		},
		{
			name:       "get template list posts filters",
			method:     bcehttp.POST,
			path:       "/v2/template/list",
			query:      map[string]string{"locale": "zh-cn"},
			bodyFields: map[string]interface{}{"namespace": "default", "name": "template", "id": "tmpl-001", "type": "private", "supportedInstanceType": "bcc"},
			call: func(client *Client) (interface{}, error) {
				return client.GetTemplateListV2(&GetTemplateListV2Request{
					Locale:                util.PtrString("zh-cn"),
					Namespace:             util.PtrString("default"),
					Name:                  util.PtrString("template"),
					Id:                    util.PtrString("tmpl-001"),
					OosType:               util.PtrString("private"),
					SupportedInstanceType: util.PtrString("bcc"),
				})
			},
		},
		{
			name:       "update template puts complete body",
			method:     bcehttp.PUT,
			path:       "/v2/template",
			bodyFields: map[string]interface{}{"namespace": "default", "id": "tmpl-001", "name": "updated", "description": "updated template", "linear": true},
			call: func(client *Client) (interface{}, error) {
				return client.UpdateTemplateV2(&UpdateTemplateV2Request{
					Namespace:   util.PtrString("default"),
					Id:          util.PtrString("tmpl-001"),
					Name:        util.PtrString("updated"),
					Description: util.PtrString("updated template"),
					Linear:      util.PtrBool(true),
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
				if r.Method != tt.method {
					t.Errorf("method = %q, want %q", r.Method, tt.method)
				}
				if r.URL.Path != tt.path {
					t.Errorf("path = %q, want %q", r.URL.Path, tt.path)
				}
				assertQuery(t, r, tt.query)
				assertJSONBody(t, r, tt.bodyFields)

				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{"success":true,"msg":"ok","code":200,"result":{}}`)
			}))
			defer server.Close()

			client, err := NewClient("ak", "sk", server.URL)
			if err != nil {
				t.Fatalf("NewClient() error = %v, want nil", err)
			}
			result, err := tt.call(client)
			if err != nil {
				t.Fatalf("client call error = %v, want nil", err)
			}
			assertResponseFields(t, result)
		})
	}
}

func TestOOSClientV2ReturnsServiceError(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != bcehttp.GET {
			t.Errorf("method = %q, want %q", r.Method, bcehttp.GET)
		}
		if r.URL.Path != "/v2/template" {
			t.Errorf("path = %q, want /v2/template", r.URL.Path)
		}
		w.Header().Set(bcehttp.BCE_REQUEST_ID, "req-001")
		w.WriteHeader(stdhttp.StatusForbidden)
		fmt.Fprint(w, `{"code":"AccessDenied","message":"denied by test"}`)
	}))
	defer server.Close()

	client, err := NewClient("ak", "sk", server.URL)
	if err != nil {
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	result, err := client.GetTemplateDetailV2(&GetTemplateDetailV2Request{Id: util.PtrString("tmpl-001")})
	if result != nil {
		t.Errorf("GetTemplateDetailV2() result = %#v, want nil", result)
	}
	serviceErr, ok := err.(*bce.BceServiceError)
	if !ok {
		t.Fatalf("GetTemplateDetailV2() error type = %T, want *bce.BceServiceError", err)
	}
	if serviceErr.StatusCode != stdhttp.StatusForbidden {
		t.Errorf("service error status = %d, want %d", serviceErr.StatusCode, stdhttp.StatusForbidden)
	}
	if serviceErr.Code != "AccessDenied" {
		t.Errorf("service error code = %q, want AccessDenied", serviceErr.Code)
	}
	if serviceErr.Message != "denied by test" {
		t.Errorf("service error message = %q, want denied by test", serviceErr.Message)
	}
}

func assertQuery(t *testing.T, r *stdhttp.Request, want map[string]string) {
	t.Helper()
	values := r.URL.Query()
	if len(values) != len(want) {
		t.Errorf("query length = %d (%v), want %d (%v)", len(values), values, len(want), want)
	}
	for key, wantValue := range want {
		if got := values.Get(key); got != wantValue {
			t.Errorf("query[%s] = %q, want %q", key, got, wantValue)
		}
	}
}

func assertJSONBody(t *testing.T, r *stdhttp.Request, want map[string]interface{}) {
	t.Helper()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("ReadAll(body) error = %v, want nil", err)
	}
	if len(want) == 0 {
		if len(body) != 0 {
			t.Errorf("body = %s, want empty", string(body))
		}
		return
	}

	var got map[string]interface{}
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal(%s) error = %v, want nil", string(body), err)
	}
	for key, wantValue := range want {
		if !reflect.DeepEqual(got[key], wantValue) {
			t.Errorf("body[%s] = %#v, want %#v", key, got[key], wantValue)
		}
	}
	if _, ok := got["locale"]; ok {
		t.Errorf("body contains locale = %#v, want locale only in query", got["locale"])
	}
}

func assertResponseFields(t *testing.T, result interface{}) {
	t.Helper()
	if result == nil {
		t.Fatal("result = nil, want response object")
	}
	value := reflect.Indirect(reflect.ValueOf(result))
	if !value.IsValid() {
		t.Fatal("result is invalid, want response object")
	}

	success := value.FieldByName("Success")
	if !success.IsValid() || success.IsNil() || !success.Elem().Bool() {
		t.Fatalf("Success = %#v, want true", success)
	}
	msg := value.FieldByName("Msg")
	if !msg.IsValid() || msg.IsNil() || msg.Elem().String() != "ok" {
		t.Fatalf("Msg = %#v, want ok", msg)
	}
	code := value.FieldByName("Code")
	if code.IsValid() && !code.IsNil() && code.Elem().Int() != 200 {
		t.Fatalf("Code = %d, want 200", code.Elem().Int())
	}
}
