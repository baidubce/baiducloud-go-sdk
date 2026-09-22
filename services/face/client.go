package face

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "face." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_REST = "rest"

	CONSTANT_2_0 = "2.0"

	CONSTANT_FACE = "face"

	CONSTANT_V1 = "v1"

	CONSTANT_FACELIVENESS = "faceliveness"

	CONSTANT_SESSIONCODE = "sessioncode"

	CONSTANT_V3 = "v3"

	CONSTANT_DETECT = "detect"

	CONSTANT_FACESET = "faceset"

	CONSTANT_GROUP = "group"

	CONSTANT_ADD = "add"

	CONSTANT_GETLIST = "getlist"

	CONSTANT_DELETE = "delete"

	CONSTANT_PERSON = "person"

	CONSTANT_VERIFY = "verify"

	CONSTANT_FACEVERIFY = "faceverify"

	CONSTANT_SEARCH = "search"

	CONSTANT_LANDMARK = "landmark"

	CONSTANT_USER = "user"

	CONSTANT_MULTI_SEARCH = "multi-search"

	CONSTANT_UPDATE = "update"

	CONSTANT_GETUSERS = "getusers"

	CONSTANT_V4 = "v4"

	CONSTANT_VERIFY_DATE = "verify_date"

	CONSTANT_IDMATCH_STATUS = "idmatch_status"

	CONSTANT_GET = "get"

	CONSTANT_COPY = "copy"

	CONSTANT_IDMATCH = "idmatch"

	CONSTANT_IDMATCH_DATE = "idmatch_date"

	CONSTANT_EDITATTR = "editattr"

	CONSTANT_MERGE = "merge"
)

// Client of face service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewClientWithApiKey(apiKey, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithApiKey(apiKey, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func NewClientWithAccessToken(apiKey, secretKey, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAccessToken(apiKey, secretKey, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getFaceDeleteUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_DELETE
}
func getFaceDetectUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_DETECT
}
func getFaceEditAttrUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_EDITATTR
}
func getFaceGetListUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_GETLIST
}
func getFaceLandmarkUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_LANDMARK
}
func getFaceMergeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_MERGE
}
func getFaceMultiSearchUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_MULTI_SEARCH
}
func getFacePersonVerifyUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_PERSON + bce.URI_PREFIX + CONSTANT_VERIFY
}
func getFaceSearchUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_SEARCH
}
func getFaceVerifyUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACEVERIFY
}
func getFaceVerifyDateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V4 + bce.URI_PREFIX + CONSTANT_VERIFY_DATE
}
func getGroupAddUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_ADD
}
func getGroupDeleteUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_DELETE
}
func getGroupGetListUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_GETLIST
}
func getGroupGetUsersUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_GROUP + bce.URI_PREFIX + CONSTANT_GETUSERS
}
func getIdMatchDateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V4 + bce.URI_PREFIX + CONSTANT_IDMATCH_DATE
}
func getIdMatchStatusUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V4 + bce.URI_PREFIX + CONSTANT_IDMATCH_STATUS
}
func getPersonIdMatchUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_PERSON + bce.URI_PREFIX + CONSTANT_IDMATCH
}
func getSessionCodeUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FACELIVENESS + bce.URI_PREFIX + CONSTANT_SESSIONCODE
}
func getUserAddUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_ADD
}
func getUserCopyUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_COPY
}
func getUserDeleteUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_DELETE
}
func getUserGetUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_GET
}
func getUserUpdateUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V3 + bce.URI_PREFIX + CONSTANT_FACESET + bce.URI_PREFIX + CONSTANT_USER + bce.URI_PREFIX + CONSTANT_UPDATE
}
func getVerifyUri() string {
	return bce.URI_PREFIX + CONSTANT_REST + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_FACE + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_FACELIVENESS + bce.URI_PREFIX + CONSTANT_VERIFY
}
