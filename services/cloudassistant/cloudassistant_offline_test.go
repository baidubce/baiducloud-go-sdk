package cloudassistant

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func strPtr(value string) *string { return &value }
func int32Ptr(value int32) *int32 { return &value }
func int64Ptr(value int64) *int64 { return &value }
func boolPtr(value bool) *bool    { return &value }

func newCloudassistantTestClient(t *testing.T, handler http.HandlerFunc) (*Client, func()) {
	t.Helper()
	server := httptest.NewServer(handler)
	client, err := NewClient("test-ak", "test-sk", server.URL)
	if err != nil {
		server.Close()
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	if client.GetBceClientConfig().Endpoint != server.URL {
		server.Close()
		t.Fatalf("NewClient() endpoint = %q, want %q", client.GetBceClientConfig().Endpoint, server.URL)
	}
	return client, server.Close
}

func decodeRequestBody(t *testing.T, r *http.Request) map[string]interface{} {
	t.Helper()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("ReadAll(request body) error = %v", err)
	}
	if len(body) == 0 {
		return map[string]interface{}{}
	}
	var got map[string]interface{}
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("request body = %q, want valid JSON: %v", string(body), err)
	}
	return got
}

func assertNestedString(t *testing.T, body map[string]interface{}, path []string, want string) {
	t.Helper()
	var current interface{} = body
	for _, key := range path {
		object, ok := current.(map[string]interface{})
		if !ok {
			t.Fatalf("body path %v parent = %T, want object", path, current)
		}
		current, ok = object[key]
		if !ok {
			t.Fatalf("body path %v missing key %q in %#v", path, key, object)
		}
	}
	got, ok := current.(string)
	if !ok || got != want {
		t.Fatalf("body path %v = %#v, want %q", path, current, want)
	}
}

func writeJSONResponse(t *testing.T, w http.ResponseWriter, body string) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write([]byte(body)); err != nil {
		t.Fatalf("Write(response) error = %v", err)
	}
}

func TestNewClientUsesDefaultEndpoint(t *testing.T) {
	client, err := NewClient("test-ak", "test-sk", "")
	if err != nil {
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	if client.GetBceClientConfig().Endpoint != DEFAULT_ENDPOINT {
		t.Fatalf("NewClient() endpoint = %q, want %q", client.GetBceClientConfig().Endpoint, DEFAULT_ENDPOINT)
	}
}

func TestCloudassistantURIHelpers(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "action list", got: getActionListUri(VERSION_V1), want: "/v1/ca/action/list"},
		{name: "action log", got: getActionLogUri(VERSION_V1), want: "/v1/ca/log"},
		{name: "action run", got: getActionRunUri(VERSION_V1), want: "/v1/ca/actionRun"},
		{name: "action run list", got: getActionRunListUri(VERSION_V1), want: "/v1/ca/actionRun/list"},
		{name: "batch get agent", got: getBatchGetAgentUri(VERSION_V1), want: "/v1/ca/agent/batch"},
		{name: "create action", got: getCreateActionUri(VERSION_V1), want: "/v1/ca/action"},
		{name: "delete action", got: getDeleteActionUri(VERSION_V1, "act-delete"), want: "/v1/ca/action/act-delete"},
		{name: "get action", got: getGetActionUri(VERSION_V1), want: "/v1/ca/action"},
		{name: "get action run", got: getGetActionRunUri(VERSION_V1), want: "/v1/ca/actionRun"},
		{name: "update action", got: getUpdateActionUri(VERSION_V1), want: "/v1/ca/action"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Fatalf("uri = %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestClientMethodsSendExpectedRequestsAndDecodeResponses(t *testing.T) {
	tests := []struct {
		name        string
		wantMethod  string
		wantPath    string
		wantLocale  string
		wantQuery   map[string]string
		assertBody  func(*testing.T, map[string]interface{})
		respond     string
		call        func(*Client) (interface{}, error)
		assertReply func(*testing.T, interface{})
	}{
		{
			name:       "ActionList",
			wantPath:   "/v1/ca/action/list",
			wantLocale: "zh-cn",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"action", "name"}, "disk-check")
			},
			respond: `{"requestId":"req-list","code":"OK","message":"success","success":true,"result":{"pageNo":1,"pageSize":10,"totalCount":1,"data":[{"id":"act-1","name":"disk-check"}]}}`,
			call: func(client *Client) (interface{}, error) {
				return client.ActionList(&ActionListRequest{
					Locale: strPtr("zh-cn"),
					PageNo: int32Ptr(1),
					Action: &ActionFilter{Name: strPtr("disk-check")},
				})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*ActionListResponse)
				if *resp.RequestId != "req-list" || !*resp.Success || *resp.Result.Data[0].Name != "disk-check" {
					t.Fatalf("ActionList response = %#v, want decoded page with disk-check", resp)
				}
			},
		},
		{
			name:     "ActionLog",
			wantPath: "/v1/ca/log",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"runId"}, "run-1")
			},
			respond: `{"requestId":"req-log","code":"OK","message":"success","success":true,"result":{"nextCursor":8,"state":"Done","childId":"child-1"}}`,
			call: func(client *Client) (interface{}, error) {
				return client.ActionLog(&ActionLogRequest{RunId: strPtr("run-1")})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*ActionLogResponse)
				if *resp.RequestId != "req-log" || *resp.Result.NextCursor != 8 || *resp.Result.State != "Done" {
					t.Fatalf("ActionLog response = %#v, want decoded log result", resp)
				}
			},
		},
		{
			name:       "ActionRun",
			wantPath:   "/v1/ca/actionRun",
			wantLocale: "zh-cn",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"action", "name"}, "act-1")
			},
			respond: `{"requestId":"req-run","code":"OK","message":"success","success":true,"result":{"runId":"run-1"}}`,
			call: func(client *Client) (interface{}, error) {
				action := interface{}(map[string]string{"name": "act-1"})
				target := interface{}(map[string]string{"instanceId": "i-1"})
				return client.ActionRun(&ActionRunRequest{Locale: strPtr("zh-cn"), Action: &action, Targets: []*interface{}{&target}})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*ActionRunResponse)
				result := (*resp.Result).(map[string]interface{})
				if *resp.RequestId != "req-run" || result["runId"] != "run-1" {
					t.Fatalf("ActionRun response = %#v, want decoded run id", resp)
				}
			},
		},
		{
			name:       "ActionRunList",
			wantPath:   "/v1/ca/actionRun/list",
			wantLocale: "zh-cn",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"state"}, "Running")
			},
			respond: `{"requestId":"req-run-list","code":"OK","message":"success","success":true,"result":{"pageNo":2,"pageSize":5,"totalCount":1,"data":[{"id":"run-1","state":"Running"}]}}`,
			call: func(client *Client) (interface{}, error) {
				return client.ActionRunList(&ActionRunListRequest{Locale: strPtr("zh-cn"), PageNo: int32Ptr(2), State: strPtr("Running")})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*ActionRunListResponse)
				if *resp.RequestId != "req-run-list" || *resp.Result.PageNo != 2 || *resp.Result.Data[0].State != "Running" {
					t.Fatalf("ActionRunList response = %#v, want decoded run page", resp)
				}
			},
		},
		{
			name:     "BatchGetAgent",
			wantPath: "/v1/ca/agent/batch",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				hosts, ok := body["hosts"].([]interface{})
				if !ok || len(hosts) != 2 {
					t.Fatalf("hosts = %#v, want two hosts", body["hosts"])
				}
				firstHost := hosts[0].(map[string]interface{})
				if firstHost["instanceId"] != "i-1" {
					t.Fatalf("first host instanceId = %#v, want i-1", firstHost["instanceId"])
				}
			},
			respond: `{"requestId":"req-agent","code":"OK","message":"success","success":true,"result":[{"host":{"instanceId":"i-1"},"state":"Online","version":"1.0.0"}]}`,
			call: func(client *Client) (interface{}, error) {
				return client.BatchGetAgent(&BatchGetAgentRequest{Hosts: []*Host{{InstanceId: strPtr("i-1")}, {InstanceId: strPtr("i-2")}}})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*BatchGetAgentResponse)
				if *resp.RequestId != "req-agent" || *resp.Result[0].Host.InstanceId != "i-1" || *resp.Result[0].State != "Online" {
					t.Fatalf("BatchGetAgent response = %#v, want decoded agent", resp)
				}
			},
		},
		{
			name:     "CreateAction",
			wantPath: "/v1/ca/action",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"action", "name"}, "install-agent")
			},
			respond: `{"requestId":"req-create","code":"OK","message":"success","success":true,"result":{"actionId":"act-2","actionName":"install-agent","runId":"run-2"}}`,
			call: func(client *Client) (interface{}, error) {
				return client.CreateAction(&CreateActionRequest{Execution: strPtr("bash"), Action: &Action{Name: strPtr("install-agent"), Command: &Command{Content: strPtr("date")}}})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*CreateActionResponse)
				if *resp.RequestId != "req-create" || *resp.Result.ActionId != "act-2" || *resp.Result.ActionName != "install-agent" {
					t.Fatalf("CreateAction response = %#v, want decoded action index", resp)
				}
			},
		},
		{
			name:       "DeleteAction",
			wantMethod: http.MethodDelete,
			wantPath:   "/v1/ca/action/act-2",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				if len(body) != 0 {
					t.Fatalf("body = %#v, want empty delete body", body)
				}
			},
			respond: `{"requestId":"req-delete","code":"OK","message":"deleted","success":true}`,
			call: func(client *Client) (interface{}, error) {
				return client.DeleteAction(&DeleteActionRequest{Id: strPtr("act-2")})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*DeleteActionResponse)
				if *resp.RequestId != "req-delete" || *resp.Message != "deleted" || !*resp.Success {
					t.Fatalf("DeleteAction response = %#v, want decoded delete status", resp)
				}
			},
		},
		{
			name:       "GetAction",
			wantMethod: http.MethodGet,
			wantPath:   "/v1/ca/action",
			wantQuery:  map[string]string{"id": "act-3", "locale": "zh-cn"},
			assertBody: func(t *testing.T, body map[string]interface{}) {
				if len(body) != 0 {
					t.Fatalf("body = %#v, want empty get body", body)
				}
			},
			respond: `{"requestId":"req-get-action","code":"OK","message":"success","success":true,"result":{"id":"act-3","name":"cleanup","timeoutSecond":30}}`,
			call: func(client *Client) (interface{}, error) {
				return client.GetAction(&GetActionRequest{Id: strPtr("act-3"), Locale: strPtr("zh-cn")})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*GetActionResponse)
				if *resp.RequestId != "req-get-action" || *resp.Result.Id != "act-3" || *resp.Result.TimeoutSecond != 30 {
					t.Fatalf("GetAction response = %#v, want decoded action", resp)
				}
			},
		},
		{
			name:       "GetActionRun",
			wantMethod: http.MethodGet,
			wantPath:   "/v1/ca/actionRun",
			wantQuery:  map[string]string{"id": "run-3", "withLog": "true", "pageNo": "1", "pageSize": "50", "childRunState": "Done", "locale": "zh-cn"},
			assertBody: func(t *testing.T, body map[string]interface{}) {
				if len(body) != 0 {
					t.Fatalf("body = %#v, want empty get body", body)
				}
			},
			respond: `{"requestId":"req-get-run","code":"OK","message":"success","success":true,"result":{"id":"run-3","state":"Done","totalCount":4}}`,
			call: func(client *Client) (interface{}, error) {
				return client.GetActionRun(&GetActionRunRequest{Id: strPtr("run-3"), WithLog: strPtr("true"), PageNo: int32Ptr(1), PageSize: int32Ptr(50), ChildRunState: strPtr("Done"), Locale: strPtr("zh-cn")})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*GetActionRunResponse)
				if *resp.RequestId != "req-get-run" || *resp.Result.Id != "run-3" || *resp.Result.TotalCount != 4 {
					t.Fatalf("GetActionRun response = %#v, want decoded action run", resp)
				}
			},
		},
		{
			name:       "UpdateAction",
			wantMethod: http.MethodPut,
			wantPath:   "/v1/ca/action",
			assertBody: func(t *testing.T, body map[string]interface{}) {
				assertNestedString(t, body, []string{"action", "name"}, "cleanup-new")
			},
			respond: `{"requestId":"req-update","code":"OK","message":"updated","success":true}`,
			call: func(client *Client) (interface{}, error) {
				return client.UpdateAction(&UpdateActionRequest{Action: &Action{Name: strPtr("cleanup-new"), Command: &Command{User: strPtr("root")}}})
			},
			assertReply: func(t *testing.T, got interface{}) {
				resp := got.(*UpdateActionResponse)
				if *resp.RequestId != "req-update" || *resp.Message != "updated" || !*resp.Success {
					t.Fatalf("UpdateAction response = %#v, want decoded update status", resp)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newCloudassistantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				wantMethod := tt.wantMethod
				if wantMethod == "" {
					wantMethod = http.MethodPost
				}
				if r.Method != wantMethod {
					t.Fatalf("method = %q, want %q", r.Method, wantMethod)
				}
				if r.URL.Path != tt.wantPath {
					t.Fatalf("path = %q, want %q", r.URL.Path, tt.wantPath)
				}
				if tt.wantLocale != "" && r.URL.Query().Get("locale") != tt.wantLocale {
					t.Fatalf("locale = %q, want %q", r.URL.Query().Get("locale"), tt.wantLocale)
				}
				for key, want := range tt.wantQuery {
					if got := r.URL.Query().Get(key); got != want {
						t.Fatalf("query %s = %q, want %q", key, got, want)
					}
				}
				if wantMethod == http.MethodPost || wantMethod == http.MethodPut {
					if contentType := r.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
						t.Fatalf("Content-Type = %q, want application/json", contentType)
					}
				}
				body := decodeRequestBody(t, r)
				tt.assertBody(t, body)
				writeJSONResponse(t, w, tt.respond)
			})
			defer cleanup()

			got, err := tt.call(client)
			if err != nil {
				t.Fatalf("%s() error = %v, want nil", tt.name, err)
			}
			tt.assertReply(t, got)
		})
	}
}

func TestNewClientReturnsCredentialError(t *testing.T) {
	client, err := NewClient("", "test-sk", "http://127.0.0.1")
	if err == nil {
		t.Fatal("NewClient() error = nil, want access key error")
	}
	if client != nil {
		t.Fatalf("NewClient() client = %#v, want nil", client)
	}
	if err.Error() != "accessKeyId should not be empty" {
		t.Fatalf("NewClient() error = %q, want accessKeyId should not be empty", err.Error())
	}
}

func TestClientMethodsReturnErrorsFromServer(t *testing.T) {
	tests := []struct {
		name string
		call func(*Client) (interface{}, error)
	}{
		{name: "ActionList", call: func(client *Client) (interface{}, error) {
			return client.ActionList(&ActionListRequest{Locale: strPtr("zh-cn")})
		}},
		{name: "ActionLog", call: func(client *Client) (interface{}, error) {
			return client.ActionLog(&ActionLogRequest{RunId: strPtr("run-error")})
		}},
		{name: "ActionRun", call: func(client *Client) (interface{}, error) {
			action := interface{}(map[string]string{"name": "act-error"})
			return client.ActionRun(&ActionRunRequest{Action: &action})
		}},
		{name: "ActionRunList", call: func(client *Client) (interface{}, error) {
			return client.ActionRunList(&ActionRunListRequest{State: strPtr("Failed")})
		}},
		{name: "BatchGetAgent", call: func(client *Client) (interface{}, error) {
			return client.BatchGetAgent(&BatchGetAgentRequest{Hosts: []*Host{{InstanceId: strPtr("i-error")}}})
		}},
		{name: "CreateAction", call: func(client *Client) (interface{}, error) {
			return client.CreateAction(&CreateActionRequest{Action: &Action{Name: strPtr("create-error")}})
		}},
		{name: "DeleteAction", call: func(client *Client) (interface{}, error) {
			return client.DeleteAction(&DeleteActionRequest{Id: strPtr("delete-error")})
		}},
		{name: "GetAction", call: func(client *Client) (interface{}, error) {
			return client.GetAction(&GetActionRequest{Id: strPtr("get-error")})
		}},
		{name: "GetActionRun", call: func(client *Client) (interface{}, error) {
			return client.GetActionRun(&GetActionRunRequest{Id: strPtr("run-error")})
		}},
		{name: "UpdateAction", call: func(client *Client) (interface{}, error) {
			return client.UpdateAction(&UpdateActionRequest{Action: &Action{Name: strPtr("update-error")}})
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, cleanup := newCloudassistantTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, `{"code":"InternalError","message":"forced failure"}`, http.StatusInternalServerError)
			})
			defer cleanup()

			got, err := tt.call(client)
			if err == nil {
				t.Fatalf("%s() error = nil, want server error", tt.name)
			}
			if got != nil && !reflect.ValueOf(got).IsNil() {
				t.Fatalf("%s() result = %#v, want nil on error", tt.name, got)
			}
			if !strings.Contains(err.Error(), "500") && !strings.Contains(err.Error(), "InternalError") {
				t.Fatalf("%s() error = %q, want HTTP 500/InternalError detail", tt.name, err.Error())
			}
		})
	}
}

func TestCloudassistantModelsJSONRoundTrip(t *testing.T) {
	execParams := interface{}(map[string]string{"mode": "strict"})
	parameters := interface{}(map[string]string{"pkg": "nginx"})
	supportedType := "bcc"
	fileUpload := interface{}(map[string]string{"bucket": "scripts"})
	action := &Action{
		Id:                     strPtr("act-json"),
		Ref:                    strPtr("ref-1"),
		CloudassistantType:     strPtr("command"),
		Name:                   strPtr("json-action"),
		Alias:                  strPtr("alias-json"),
		Description:            strPtr("json description"),
		TimeoutSecond:          int32Ptr(60),
		Command:                &Command{CloudassistantType: strPtr("shell"), Content: strPtr("echo ok"), Scope: strPtr("global"), EnableParameter: boolPtr(true), User: strPtr("root"), WorkDir: strPtr("/tmp"), ExecParams: &execParams, WaitOnAgentMilli: int32Ptr(1000)},
		FileUpload:             &FileUpload{Os: strPtr("linux"), Content: strPtr("echo ok"), Filename: strPtr("script.sh"), Filepath: strPtr("/tmp/script.sh"), BosBucketName: strPtr("bucket"), BosFilePath: strPtr("scripts/script.sh"), BosEtag: strPtr("etag-1"), User: strPtr("root"), Group: strPtr("root"), Mode: strPtr("0755"), Overwrite: boolPtr(true)},
		SupportedInstanceTypes: []*string{&supportedType},
		CreatedTimestamp:       int64Ptr(100),
		UpdatedTimestamp:       int64Ptr(200),
	}
	request := &CreateActionRequest{
		Execution:          strPtr("bash"),
		Action:             action,
		Parameters:         &parameters,
		TargetSelectorType: strPtr("manual"),
		Targets:            []*Target{{InstanceType: strPtr("bcc"), InstanceId: strPtr("i-json"), InstanceName: strPtr("vm-json"), InternalIp: strPtr("10.0.0.1"), ExternalIp: strPtr("1.1.1.1"), Bandwidth: strPtr("100M")}},
		TargetSelector:     &TargetSelector{InstanceType: strPtr("bcc"), Tags: []*Tag{{TagKey: strPtr("role"), TagValue: strPtr("web")}}, ImportInstances: &TargetImport{KeywordType: strPtr("instanceId"), Instances: []*string{strPtr("i-json")}}},
	}

	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Marshal(CreateActionRequest) error = %v, want nil", err)
	}
	jsonText := string(encoded)
	for _, want := range []string{"\"execution\":\"bash\"", "\"type\":\"command\"", "\"supportedInstanceTypes\":[\"bcc\"]", "\"targetSelectorType\":\"manual\"", "\"instanceId\":\"i-json\""} {
		if !strings.Contains(jsonText, want) {
			t.Fatalf("CreateActionRequest JSON = %s, want to contain %s", jsonText, want)
		}
	}
	if strings.Contains(jsonText, "Locale") || strings.Contains(jsonText, "locale") {
		t.Fatalf("CreateActionRequest JSON = %s, want no locale field", jsonText)
	}

	var decoded CreateActionRequest
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal(CreateActionRequest) error = %v, want nil", err)
	}
	if *decoded.Action.Id != "act-json" || *decoded.Action.Command.Content != "echo ok" || *decoded.Targets[0].ExternalIp != "1.1.1.1" {
		t.Fatalf("decoded request = %#v, want nested action and target fields preserved", decoded)
	}

	list := &ActionListRequest{Locale: strPtr("zh-cn"), PageSize: int32Ptr(20), Sort: strPtr("name"), Ascending: boolPtr(false), Action: &ActionFilter{Name: strPtr("json-action"), FileUpload: &fileUpload}}
	listJSON, err := json.Marshal(list)
	if err != nil {
		t.Fatalf("Marshal(ActionListRequest) error = %v, want nil", err)
	}
	if strings.Contains(string(listJSON), "locale") {
		t.Fatalf("ActionListRequest JSON = %s, want locale omitted", string(listJSON))
	}
	if !strings.Contains(string(listJSON), "\"fileUpload\":{\"bucket\":\"scripts\"}") {
		t.Fatalf("ActionListRequest JSON = %s, want fileUpload payload", string(listJSON))
	}
}
