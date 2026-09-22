package face

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/baiducloud-go-sdk/core/util"
	"github.com/baidubce/baiducloud-go-sdk/core/util/log"
)

var (
	FACE_CLIENT *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK        string
	SK        string
	Endpoint  string
	ApiKey    string
	SecretKey string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	// ==== AK/SK 鉴权 ====
	// FACE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// FACE_CLIENT, _ = NewClientWithAccessToken(confObj.ApiKey, confObj.SecretKey, confObj.Endpoint)

	// ==== API Key 鉴权 ====
	FACE_CLIENT, _ = NewClientWithApiKey(confObj.ApiKey, confObj.Endpoint)

	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestClient_FaceDelete(t *testing.T) {
	faceDeleteRequest := &FaceDeleteRequest{
		UserId:    util.PtrString(""),
		GroupId:   util.PtrString(""),
		FaceToken: util.PtrString(""),
		LogId:     util.PtrInt64(int64(0)),
	}
	result := &FaceDeleteResponse{}
	result, err := FACE_CLIENT.FaceDelete(faceDeleteRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceDetect(t *testing.T) {
	faceDetectRequest := &FaceDetectRequest{
		Image:            util.PtrString(""),
		ImageType:        util.PtrString(""),
		FaceField:        util.PtrString(""),
		MaxFaceNum:       util.PtrInt32(int32(0)),
		FaceType:         util.PtrString(""),
		LivenessControl:  util.PtrString(""),
		FaceSortType:     util.PtrInt32(int32(0)),
		DisplayCorpImage: util.PtrInt32(int32(0)),
	}
	result := &FaceDetectResponse{}
	result, err := FACE_CLIENT.FaceDetect(faceDetectRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceEditAttr(t *testing.T) {
	faceEditAttrRequest := &FaceEditAttrRequest{
		Image:          util.PtrString(""),
		ImageType:      util.PtrString(""),
		ActionType:     util.PtrString(""),
		Target:         util.PtrInt32(int32(0)),
		QualityControl: util.PtrString(""),
		FaceLocation:   util.PtrString(""),
	}
	result := &FaceEditAttrResponse{}
	result, err := FACE_CLIENT.FaceEditAttr(faceEditAttrRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceGetList(t *testing.T) {
	faceGetListRequest := &FaceGetListRequest{
		UserId:  util.PtrString(""),
		GroupId: util.PtrString(""),
	}
	result := &FaceGetListResponse{}
	result, err := FACE_CLIENT.FaceGetList(faceGetListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceLandmark(t *testing.T) {
	faceLandmarkRequest := &FaceLandmarkRequest{
		Image:      util.PtrString(""),
		ImageType:  util.PtrString(""),
		MaxFaceNum: util.PtrInt32(int32(0)),
		FaceField:  util.PtrString(""),
	}
	result := &FaceLandmarkResponse{}
	result, err := FACE_CLIENT.FaceLandmark(faceLandmarkRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceMerge(t *testing.T) {
	faceMergeRequest := &FaceMergeRequest{
		Version:       util.PtrString(""),
		Alpha:         util.PtrFloat32(float32(0)),
		ImageTemplate: util.PtrString(""),
		ImageTarget:   util.PtrString(""),
		MergeDegree:   util.PtrString(""),
		Position:      util.PtrInt32(int32(0)),
		Language:      util.PtrInt32(int32(0)),
	}
	result := &FaceMergeResponse{}
	result, err := FACE_CLIENT.FaceMerge(faceMergeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceMultiSearch(t *testing.T) {
	faceMultiSearchRequest := &FaceMultiSearchRequest{
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		GroupIdList:     util.PtrString(""),
		MaxFaceNum:      util.PtrInt32(int32(0)),
		MatchThreshold:  util.PtrInt32(int32(0)),
		QualityControl:  util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
		MaxUserNum:      util.PtrInt32(int32(0)),
	}
	result := &FaceMultiSearchResponse{}
	result, err := FACE_CLIENT.FaceMultiSearch(faceMultiSearchRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FacePersonVerify(t *testing.T) {
	facePersonVerifyRequest := &FacePersonVerifyRequest{
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		IdCardNumber:    util.PtrString(""),
		Name:            util.PtrString(""),
		QualityControl:  util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
	}
	result := &FacePersonVerifyResponse{}
	result, err := FACE_CLIENT.FacePersonVerify(facePersonVerifyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceSearch(t *testing.T) {
	faceSearchRequest := &FaceSearchRequest{
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		GroupIdList:     util.PtrString(""),
		QualityControl:  util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
		UserId:          util.PtrString(""),
		MaxUserNum:      util.PtrInt32(int32(0)),
		FaceSortType:    util.PtrInt32(int32(0)),
		MatchThreshold:  util.PtrInt32(int32(0)),
	}
	result := &FaceSearchResponse{}
	result, err := FACE_CLIENT.FaceSearch(faceSearchRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceVerify(t *testing.T) {
	result := &FaceVerifyResponse{}
	result, err := FACE_CLIENT.FaceVerify()
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FaceVerifyDate(t *testing.T) {
	faceVerifyDateRequest := &FaceVerifyDateRequest{
		Name:            util.PtrString(""),
		IdCardNumber:    util.PtrString(""),
		StartDate:       util.PtrString(""),
		EndDate:         util.PtrString(""),
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
		QualityControl:  util.PtrString(""),
	}
	result := &FaceVerifyDateResponse{}
	result, err := FACE_CLIENT.FaceVerifyDate(faceVerifyDateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GroupAdd(t *testing.T) {
	groupAddRequest := &GroupAddRequest{
		GroupId: util.PtrString(""),
	}
	result := &GroupAddResponse{}
	result, err := FACE_CLIENT.GroupAdd(groupAddRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GroupDelete(t *testing.T) {
	groupDeleteRequest := &GroupDeleteRequest{
		GroupId: util.PtrString(""),
	}
	result := &GroupDeleteResponse{}
	result, err := FACE_CLIENT.GroupDelete(groupDeleteRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GroupGetList(t *testing.T) {
	groupGetListRequest := &GroupGetListRequest{
		Start:  util.PtrInt32(int32(0)),
		Length: util.PtrInt32(int32(0)),
	}
	result := &GroupGetListResponse{}
	result, err := FACE_CLIENT.GroupGetList(groupGetListRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GroupGetUsers(t *testing.T) {
	groupGetUsersRequest := &GroupGetUsersRequest{
		GroupId: util.PtrString(""),
		Start:   util.PtrInt32(int32(0)),
		Length:  util.PtrInt32(int32(0)),
	}
	result := &GroupGetUsersResponse{}
	result, err := FACE_CLIENT.GroupGetUsers(groupGetUsersRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_IdMatchDate(t *testing.T) {
	idMatchDateRequest := &IdMatchDateRequest{
		Name:         util.PtrString(""),
		IdCardNumber: util.PtrString(""),
		StartDate:    util.PtrString(""),
		EndDate:      util.PtrString(""),
	}
	result := &IdMatchDateResponse{}
	result, err := FACE_CLIENT.IdMatchDate(idMatchDateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_IdMatchStatus(t *testing.T) {
	idMatchStatusRequest := &IdMatchStatusRequest{
		Name:         util.PtrString(""),
		IdCardNumber: util.PtrString(""),
		StartDate:    util.PtrString(""),
		EndDate:      util.PtrString(""),
	}
	result := &IdMatchStatusResponse{}
	result, err := FACE_CLIENT.IdMatchStatus(idMatchStatusRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_PersonIdMatch(t *testing.T) {
	personIdMatchRequest := &PersonIdMatchRequest{
		IdCardNumber: util.PtrString(""),
		Name:         util.PtrString(""),
	}
	result := &PersonIdMatchResponse{}
	result, err := FACE_CLIENT.PersonIdMatch(personIdMatchRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_SessionCode(t *testing.T) {
	sessionCodeRequest := &SessionCodeRequest{
		FaceType:      util.PtrString(""),
		MinCodeLength: util.PtrString(""),
		MaxCodeLength: util.PtrString(""),
	}
	result := &SessionCodeResponse{}
	result, err := FACE_CLIENT.SessionCode(sessionCodeRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserAdd(t *testing.T) {
	userAddRequest := &UserAddRequest{
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		GroupId:         util.PtrString(""),
		UserId:          util.PtrString(""),
		UserInfo:        util.PtrString(""),
		QualityControl:  util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
		ActionType:      util.PtrString(""),
		FaceSortType:    util.PtrInt32(int32(0)),
	}
	result := &UserAddResponse{}
	result, err := FACE_CLIENT.UserAdd(userAddRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserCopy(t *testing.T) {
	userCopyRequest := &UserCopyRequest{
		UserId:     util.PtrString(""),
		SrcGroupId: util.PtrString(""),
		DstGroupId: util.PtrString(""),
	}
	result := &UserCopyResponse{}
	result, err := FACE_CLIENT.UserCopy(userCopyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserDelete(t *testing.T) {
	userDeleteRequest := &UserDeleteRequest{
		GroupId: util.PtrString(""),
		UserId:  util.PtrString(""),
	}
	result := &UserDeleteResponse{}
	result, err := FACE_CLIENT.UserDelete(userDeleteRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserGet(t *testing.T) {
	userGetRequest := &UserGetRequest{
		UserId:  util.PtrString(""),
		GroupId: util.PtrString(""),
	}
	result := &UserGetResponse{}
	result, err := FACE_CLIENT.UserGet(userGetRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UserUpdate(t *testing.T) {
	userUpdateRequest := &UserUpdateRequest{
		Image:           util.PtrString(""),
		ImageType:       util.PtrString(""),
		GroupId:         util.PtrString(""),
		UserId:          util.PtrString(""),
		UserInfo:        util.PtrString(""),
		QualityControl:  util.PtrString(""),
		LivenessControl: util.PtrString(""),
		SpoofingControl: util.PtrString(""),
		ActionType:      util.PtrString(""),
	}
	result := &UserUpdateResponse{}
	result, err := FACE_CLIENT.UserUpdate(userUpdateRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_Verify(t *testing.T) {
	verifyRequest := &VerifyRequest{
		VideoBase64:  util.PtrString(""),
		TypeIdentify: util.PtrString(""),
		SessionId:    util.PtrString(""),
		LipIdentify:  util.PtrString(""),
		FaceField:    util.PtrString(""),
	}
	result := &VerifyResponse{}
	result, err := FACE_CLIENT.Verify(verifyRequest)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
