package face

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// FaceDelete
//
// PARAMS:
//   - request: the arguments to FaceDelete
//
// RETURNS:
//   - FaceDeleteResponse: The return type of the FaceDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceDelete(request *FaceDeleteRequest) (*FaceDeleteResponse, error) {
	result := &FaceDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceDeleteUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceDetect
//
// PARAMS:
//   - request: the arguments to FaceDetect
//
// RETURNS:
//   - FaceDetectResponse: The return type of the FaceDetect interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceDetect(request *FaceDetectRequest) (*FaceDetectResponse, error) {
	result := &FaceDetectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceDetectUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceEditAttr
//
// PARAMS:
//   - request: the arguments to FaceEditAttr
//
// RETURNS:
//   - FaceEditAttrResponse: The return type of the FaceEditAttr interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceEditAttr(request *FaceEditAttrRequest) (*FaceEditAttrResponse, error) {
	result := &FaceEditAttrResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceEditAttrUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceGetList
//
// PARAMS:
//   - request: the arguments to FaceGetList
//
// RETURNS:
//   - FaceGetListResponse: The return type of the FaceGetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceGetList(request *FaceGetListRequest) (*FaceGetListResponse, error) {
	result := &FaceGetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceGetListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceLandmark
//
// PARAMS:
//   - request: the arguments to FaceLandmark
//
// RETURNS:
//   - FaceLandmarkResponse: The return type of the FaceLandmark interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceLandmark(request *FaceLandmarkRequest) (*FaceLandmarkResponse, error) {
	result := &FaceLandmarkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceLandmarkUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceMerge
//
// PARAMS:
//   - request: the arguments to FaceMerge
//
// RETURNS:
//   - FaceMergeResponse: The return type of the FaceMerge interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceMerge(request *FaceMergeRequest) (*FaceMergeResponse, error) {
	result := &FaceMergeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceMergeUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceMultiSearch
//
// PARAMS:
//   - request: the arguments to FaceMultiSearch
//
// RETURNS:
//   - FaceMultiSearchResponse: The return type of the FaceMultiSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceMultiSearch(request *FaceMultiSearchRequest) (*FaceMultiSearchResponse, error) {
	result := &FaceMultiSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceMultiSearchUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FacePersonVerify
//
// PARAMS:
//   - request: the arguments to FacePersonVerify
//
// RETURNS:
//   - FacePersonVerifyResponse: The return type of the FacePersonVerify interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FacePersonVerify(request *FacePersonVerifyRequest) (*FacePersonVerifyResponse, error) {
	result := &FacePersonVerifyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFacePersonVerifyUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceSearch
//
// PARAMS:
//   - request: the arguments to FaceSearch
//
// RETURNS:
//   - FaceSearchResponse: The return type of the FaceSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceSearch(request *FaceSearchRequest) (*FaceSearchResponse, error) {
	result := &FaceSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceSearchUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceVerify
//
// PARAMS:
//   - request: the arguments to FaceVerify
//
// RETURNS:
//   - FaceVerifyResponse: The return type of the FaceVerify interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceVerify() (*FaceVerifyResponse, error) {
	result := &FaceVerifyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceVerifyUri()).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FaceVerifyDate
//
// PARAMS:
//   - request: the arguments to FaceVerifyDate
//
// RETURNS:
//   - FaceVerifyDateResponse: The return type of the FaceVerifyDate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) FaceVerifyDate(request *FaceVerifyDateRequest) (*FaceVerifyDateResponse, error) {
	result := &FaceVerifyDateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFaceVerifyDateUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GroupAdd
//
// PARAMS:
//   - request: the arguments to GroupAdd
//
// RETURNS:
//   - GroupAddResponse: The return type of the GroupAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GroupAdd(request *GroupAddRequest) (*GroupAddResponse, error) {
	result := &GroupAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupAddUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GroupDelete
//
// PARAMS:
//   - request: the arguments to GroupDelete
//
// RETURNS:
//   - GroupDeleteResponse: The return type of the GroupDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GroupDelete(request *GroupDeleteRequest) (*GroupDeleteResponse, error) {
	result := &GroupDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupDeleteUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GroupGetList
//
// PARAMS:
//   - request: the arguments to GroupGetList
//
// RETURNS:
//   - GroupGetListResponse: The return type of the GroupGetList interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GroupGetList(request *GroupGetListRequest) (*GroupGetListResponse, error) {
	result := &GroupGetListResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupGetListUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GroupGetUsers
//
// PARAMS:
//   - request: the arguments to GroupGetUsers
//
// RETURNS:
//   - GroupGetUsersResponse: The return type of the GroupGetUsers interface.
//   - error: nil if success otherwise the specific error
func (c *Client) GroupGetUsers(request *GroupGetUsersRequest) (*GroupGetUsersResponse, error) {
	result := &GroupGetUsersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupGetUsersUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// IdMatchDate
//
// PARAMS:
//   - request: the arguments to IdMatchDate
//
// RETURNS:
//   - IdMatchDateResponse: The return type of the IdMatchDate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) IdMatchDate(request *IdMatchDateRequest) (*IdMatchDateResponse, error) {
	result := &IdMatchDateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getIdMatchDateUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// IdMatchStatus
//
// PARAMS:
//   - request: the arguments to IdMatchStatus
//
// RETURNS:
//   - IdMatchStatusResponse: The return type of the IdMatchStatus interface.
//   - error: nil if success otherwise the specific error
func (c *Client) IdMatchStatus(request *IdMatchStatusRequest) (*IdMatchStatusResponse, error) {
	result := &IdMatchStatusResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getIdMatchStatusUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PersonIdMatch
//
// PARAMS:
//   - request: the arguments to PersonIdMatch
//
// RETURNS:
//   - PersonIdMatchResponse: The return type of the PersonIdMatch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PersonIdMatch(request *PersonIdMatchRequest) (*PersonIdMatchResponse, error) {
	result := &PersonIdMatchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPersonIdMatchUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SessionCode
//
// PARAMS:
//   - request: the arguments to SessionCode
//
// RETURNS:
//   - SessionCodeResponse: The return type of the SessionCode interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SessionCode(request *SessionCodeRequest) (*SessionCodeResponse, error) {
	result := &SessionCodeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSessionCodeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserAdd
//
// PARAMS:
//   - request: the arguments to UserAdd
//
// RETURNS:
//   - UserAddResponse: The return type of the UserAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserAdd(request *UserAddRequest) (*UserAddResponse, error) {
	result := &UserAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserAddUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserCopy
//
// PARAMS:
//   - request: the arguments to UserCopy
//
// RETURNS:
//   - UserCopyResponse: The return type of the UserCopy interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserCopy(request *UserCopyRequest) (*UserCopyResponse, error) {
	result := &UserCopyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserCopyUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserDelete
//
// PARAMS:
//   - request: the arguments to UserDelete
//
// RETURNS:
//   - UserDeleteResponse: The return type of the UserDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserDelete(request *UserDeleteRequest) (*UserDeleteResponse, error) {
	result := &UserDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserDeleteUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserGet
//
// PARAMS:
//   - request: the arguments to UserGet
//
// RETURNS:
//   - UserGetResponse: The return type of the UserGet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserGet(request *UserGetRequest) (*UserGetResponse, error) {
	result := &UserGetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserGetUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserUpdate
//
// PARAMS:
//   - request: the arguments to UserUpdate
//
// RETURNS:
//   - UserUpdateResponse: The return type of the UserUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) UserUpdate(request *UserUpdateRequest) (*UserUpdateResponse, error) {
	result := &UserUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUserUpdateUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Verify
//
// PARAMS:
//   - request: the arguments to Verify
//
// RETURNS:
//   - VerifyResponse: The return type of the Verify interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Verify(request *VerifyRequest) (*VerifyResponse, error) {
	result := &VerifyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVerifyUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
