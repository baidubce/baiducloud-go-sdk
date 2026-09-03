package image

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
)

const ()

// AdvancedGeneral
//
// PARAMS:
//   - request: the arguments to AdvancedGeneral
//
// RETURNS:
//   - AdvancedGeneralResponse: The return type of the AdvancedGeneral interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AdvancedGeneral(request *AdvancedGeneralRequest) (*AdvancedGeneralResponse, error) {
	result := &AdvancedGeneralResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAdvancedGeneralUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AiRetouchingCreateTask
//
// PARAMS:
//   - request: the arguments to AiRetouchingCreateTask
//
// RETURNS:
//   - AiRetouchingCreateTaskResponse: The return type of the AiRetouchingCreateTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AiRetouchingCreateTask(request *AiRetouchingCreateTaskRequest) (*AiRetouchingCreateTaskResponse, error) {
	result := &AiRetouchingCreateTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAiRetouchingCreateTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AiRetouchingQueryTask
//
// PARAMS:
//   - request: the arguments to AiRetouchingQueryTask
//
// RETURNS:
//   - AiRetouchingQueryTaskResponse: The return type of the AiRetouchingQueryTask interface.
//   - error: nil if success otherwise the specific error
func (c *Client) AiRetouchingQueryTask(request *AiRetouchingQueryTaskRequest) (*AiRetouchingQueryTaskResponse, error) {
	result := &AiRetouchingQueryTaskResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAiRetouchingQueryTaskUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Animal
//
// PARAMS:
//   - request: the arguments to Animal
//
// RETURNS:
//   - AnimalResponse: The return type of the Animal interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Animal(request *AnimalRequest) (*AnimalResponse, error) {
	result := &AnimalResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAnimalUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Car
//
// PARAMS:
//   - request: the arguments to Car
//
// RETURNS:
//   - CarResponse: The return type of the Car interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Car(request *CarRequest) (*CarResponse, error) {
	result := &CarResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCarUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ColorEnhance
//
// PARAMS:
//   - request: the arguments to ColorEnhance
//
// RETURNS:
//   - ColorEnhanceResponse: The return type of the ColorEnhance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ColorEnhance(request *ColorEnhanceRequest) (*ColorEnhanceResponse, error) {
	result := &ColorEnhanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getColorEnhanceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Colourize
//
// PARAMS:
//   - request: the arguments to Colourize
//
// RETURNS:
//   - ColourizeResponse: The return type of the Colourize interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Colourize(request *ColourizeRequest) (*ColourizeResponse, error) {
	result := &ColourizeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getColourizeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ContrastEnhance
//
// PARAMS:
//   - request: the arguments to ContrastEnhance
//
// RETURNS:
//   - ContrastEnhanceResponse: The return type of the ContrastEnhance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ContrastEnhance(request *ContrastEnhanceRequest) (*ContrastEnhanceResponse, error) {
	result := &ContrastEnhanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getContrastEnhanceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Dehaze
//
// PARAMS:
//   - request: the arguments to Dehaze
//
// RETURNS:
//   - DehazeResponse: The return type of the Dehaze interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Dehaze(request *DehazeRequest) (*DehazeResponse, error) {
	result := &DehazeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDehazeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Dish
//
// PARAMS:
//   - request: the arguments to Dish
//
// RETURNS:
//   - DishResponse: The return type of the Dish interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Dish(request *DishRequest) (*DishResponse, error) {
	result := &DishResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDishUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DocRepair
//
// PARAMS:
//   - request: the arguments to DocRepair
//
// RETURNS:
//   - DocRepairResponse: The return type of the DocRepair interface.
//   - error: nil if success otherwise the specific error
func (c *Client) DocRepair(request *DocRepairRequest) (*DocRepairResponse, error) {
	result := &DocRepairResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDocRepairUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImageDefinitionEnhance
//
// PARAMS:
//   - request: the arguments to ImageDefinitionEnhance
//
// RETURNS:
//   - ImageDefinitionEnhanceResponse: The return type of the ImageDefinitionEnhance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImageDefinitionEnhance(request *ImageDefinitionEnhanceRequest) (*ImageDefinitionEnhanceResponse, error) {
	result := &ImageDefinitionEnhanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImageDefinitionEnhanceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImageQualityEnhance
//
// PARAMS:
//   - request: the arguments to ImageQualityEnhance
//
// RETURNS:
//   - ImageQualityEnhanceResponse: The return type of the ImageQualityEnhance interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImageQualityEnhance(request *ImageQualityEnhanceRequest) (*ImageQualityEnhanceResponse, error) {
	result := &ImageQualityEnhanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImageQualityEnhanceUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImageUnderstandingGetResult
//
// PARAMS:
//   - request: the arguments to ImageUnderstandingGetResult
//
// RETURNS:
//   - ImageUnderstandingGetResultResponse: The return type of the ImageUnderstandingGetResult interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImageUnderstandingGetResult(request *ImageUnderstandingGetResultRequest) (*ImageUnderstandingGetResultResponse, error) {
	result := &ImageUnderstandingGetResultResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImageUnderstandingGetResultUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ImageUnderstandingRequest
//
// PARAMS:
//   - request: the arguments to ImageUnderstandingRequest
//
// RETURNS:
//   - ImageUnderstandingRequestResponse: The return type of the ImageUnderstandingRequest interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ImageUnderstandingRequest(request *ImageUnderstandingRequestRequest) (*ImageUnderstandingRequestResponse, error) {
	result := &ImageUnderstandingRequestResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getImageUnderstandingRequestUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Ingredient
//
// PARAMS:
//   - request: the arguments to Ingredient
//
// RETURNS:
//   - IngredientResponse: The return type of the Ingredient interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Ingredient(request *IngredientRequest) (*IngredientResponse, error) {
	result := &IngredientResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getIngredientUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Inpainting
//
// PARAMS:
//   - request: the arguments to Inpainting
//
// RETURNS:
//   - InpaintingResponse: The return type of the Inpainting interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Inpainting(request *InpaintingRequest) (*InpaintingResponse, error) {
	result := &InpaintingResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInpaintingUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Landmark
//
// PARAMS:
//   - request: the arguments to Landmark
//
// RETURNS:
//   - LandmarkResponse: The return type of the Landmark interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Landmark(request *LandmarkRequest) (*LandmarkResponse, error) {
	result := &LandmarkResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLandmarkUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Logo
//
// PARAMS:
//   - request: the arguments to Logo
//
// RETURNS:
//   - LogoResponse: The return type of the Logo interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Logo(request *LogoRequest) (*LogoResponse, error) {
	result := &LogoResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLogoUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LogoAdd
//
// PARAMS:
//   - request: the arguments to LogoAdd
//
// RETURNS:
//   - LogoAddResponse: The return type of the LogoAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LogoAdd(request *LogoAddRequest) (*LogoAddResponse, error) {
	result := &LogoAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLogoAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LogoDelete
//
// PARAMS:
//   - request: the arguments to LogoDelete
//
// RETURNS:
//   - LogoDeleteResponse: The return type of the LogoDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) LogoDelete(request *LogoDeleteRequest) (*LogoDeleteResponse, error) {
	result := &LogoDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLogoDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MaterielImageAdd
//
// PARAMS:
//   - request: the arguments to MaterielImageAdd
//
// RETURNS:
//   - MaterielImageAddResponse: The return type of the MaterielImageAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MaterielImageAdd(request *MaterielImageAddRequest) (*MaterielImageAddResponse, error) {
	result := &MaterielImageAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMaterielImageAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MaterielImageDelete
//
// PARAMS:
//   - request: the arguments to MaterielImageDelete
//
// RETURNS:
//   - MaterielImageDeleteResponse: The return type of the MaterielImageDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MaterielImageDelete(request *MaterielImageDeleteRequest) (*MaterielImageDeleteResponse, error) {
	result := &MaterielImageDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMaterielImageDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MaterielImageSearch
//
// PARAMS:
//   - request: the arguments to MaterielImageSearch
//
// RETURNS:
//   - MaterielImageSearchResponse: The return type of the MaterielImageSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MaterielImageSearch(request *MaterielImageSearchRequest) (*MaterielImageSearchResponse, error) {
	result := &MaterielImageSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMaterielImageSearchUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MaterielImageUpdate
//
// PARAMS:
//   - request: the arguments to MaterielImageUpdate
//
// RETURNS:
//   - MaterielImageUpdateResponse: The return type of the MaterielImageUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MaterielImageUpdate(request *MaterielImageUpdateRequest) (*MaterielImageUpdateResponse, error) {
	result := &MaterielImageUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMaterielImageUpdateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MultiObjectDetect
//
// PARAMS:
//   - request: the arguments to MultiObjectDetect
//
// RETURNS:
//   - MultiObjectDetectResponse: The return type of the MultiObjectDetect interface.
//   - error: nil if success otherwise the specific error
func (c *Client) MultiObjectDetect(request *MultiObjectDetectRequest) (*MultiObjectDetectResponse, error) {
	result := &MultiObjectDetectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getMultiObjectDetectUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ObjectDetect
//
// PARAMS:
//   - request: the arguments to ObjectDetect
//
// RETURNS:
//   - ObjectDetectResponse: The return type of the ObjectDetect interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ObjectDetect(request *ObjectDetectRequest) (*ObjectDetectResponse, error) {
	result := &ObjectDetectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getObjectDetectUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PicturebookImageAdd
//
// PARAMS:
//   - request: the arguments to PicturebookImageAdd
//
// RETURNS:
//   - PicturebookImageAddResponse: The return type of the PicturebookImageAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PicturebookImageAdd(request *PicturebookImageAddRequest) (*PicturebookImageAddResponse, error) {
	result := &PicturebookImageAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPicturebookImageAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PicturebookImageDelete
//
// PARAMS:
//   - request: the arguments to PicturebookImageDelete
//
// RETURNS:
//   - PicturebookImageDeleteResponse: The return type of the PicturebookImageDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PicturebookImageDelete(request *PicturebookImageDeleteRequest) (*PicturebookImageDeleteResponse, error) {
	result := &PicturebookImageDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPicturebookImageDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PicturebookImageSearch
//
// PARAMS:
//   - request: the arguments to PicturebookImageSearch
//
// RETURNS:
//   - PicturebookImageSearchResponse: The return type of the PicturebookImageSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PicturebookImageSearch(request *PicturebookImageSearchRequest) (*PicturebookImageSearchResponse, error) {
	result := &PicturebookImageSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPicturebookImageSearchUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PicturebookImageUpdate
//
// PARAMS:
//   - request: the arguments to PicturebookImageUpdate
//
// RETURNS:
//   - PicturebookImageUpdateResponse: The return type of the PicturebookImageUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) PicturebookImageUpdate(request *PicturebookImageUpdateRequest) (*PicturebookImageUpdateResponse, error) {
	result := &PicturebookImageUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPicturebookImageUpdateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Plant
//
// PARAMS:
//   - request: the arguments to Plant
//
// RETURNS:
//   - PlantResponse: The return type of the Plant interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Plant(request *PlantRequest) (*PlantResponse, error) {
	result := &PlantResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPlantUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductImageAdd
//
// PARAMS:
//   - request: the arguments to ProductImageAdd
//
// RETURNS:
//   - ProductImageAddResponse: The return type of the ProductImageAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ProductImageAdd(request *ProductImageAddRequest) (*ProductImageAddResponse, error) {
	result := &ProductImageAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getProductImageAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductImageDelete
//
// PARAMS:
//   - request: the arguments to ProductImageDelete
//
// RETURNS:
//   - ProductImageDeleteResponse: The return type of the ProductImageDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ProductImageDelete(request *ProductImageDeleteRequest) (*ProductImageDeleteResponse, error) {
	result := &ProductImageDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getProductImageDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductImageSearch
//
// PARAMS:
//   - request: the arguments to ProductImageSearch
//
// RETURNS:
//   - ProductImageSearchResponse: The return type of the ProductImageSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ProductImageSearch(request *ProductImageSearchRequest) (*ProductImageSearchResponse, error) {
	result := &ProductImageSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getProductImageSearchUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProductImageUpdate
//
// PARAMS:
//   - request: the arguments to ProductImageUpdate
//
// RETURNS:
//   - ProductImageUpdateResponse: The return type of the ProductImageUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) ProductImageUpdate(request *ProductImageUpdateRequest) (*ProductImageUpdateResponse, error) {
	result := &ProductImageUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getProductImageUpdateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RemoveMoire
//
// PARAMS:
//   - request: the arguments to RemoveMoire
//
// RETURNS:
//   - RemoveMoireResponse: The return type of the RemoveMoire interface.
//   - error: nil if success otherwise the specific error
func (c *Client) RemoveMoire(request *RemoveMoireRequest) (*RemoveMoireResponse, error) {
	result := &RemoveMoireResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getRemoveMoireUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SameImageAdd
//
// PARAMS:
//   - request: the arguments to SameImageAdd
//
// RETURNS:
//   - SameImageAddResponse: The return type of the SameImageAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SameImageAdd(request *SameImageAddRequest) (*SameImageAddResponse, error) {
	result := &SameImageAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSameImageAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SameImageDelete
//
// PARAMS:
//   - request: the arguments to SameImageDelete
//
// RETURNS:
//   - SameImageDeleteResponse: The return type of the SameImageDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SameImageDelete(request *SameImageDeleteRequest) (*SameImageDeleteResponse, error) {
	result := &SameImageDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSameImageDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SameImageSearch
//
// PARAMS:
//   - request: the arguments to SameImageSearch
//
// RETURNS:
//   - SameImageSearchResponse: The return type of the SameImageSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SameImageSearch(request *SameImageSearchRequest) (*SameImageSearchResponse, error) {
	result := &SameImageSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSameImageSearchUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SameImageUpdate
//
// PARAMS:
//   - request: the arguments to SameImageUpdate
//
// RETURNS:
//   - SameImageUpdateResponse: The return type of the SameImageUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SameImageUpdate(request *SameImageUpdateRequest) (*SameImageUpdateResponse, error) {
	result := &SameImageUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSameImageUpdateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Segment
//
// PARAMS:
//   - request: the arguments to Segment
//
// RETURNS:
//   - SegmentResponse: The return type of the Segment interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Segment(request *SegmentRequest) (*SegmentResponse, error) {
	result := &SegmentResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSegmentUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SelfieAnime
//
// PARAMS:
//   - request: the arguments to SelfieAnime
//
// RETURNS:
//   - SelfieAnimeResponse: The return type of the SelfieAnime interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SelfieAnime(request *SelfieAnimeRequest) (*SelfieAnimeResponse, error) {
	result := &SelfieAnimeResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSelfieAnimeUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SimilarImageAdd
//
// PARAMS:
//   - request: the arguments to SimilarImageAdd
//
// RETURNS:
//   - SimilarImageAddResponse: The return type of the SimilarImageAdd interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SimilarImageAdd(request *SimilarImageAddRequest) (*SimilarImageAddResponse, error) {
	result := &SimilarImageAddResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSimilarImageAddUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SimilarImageDelete
//
// PARAMS:
//   - request: the arguments to SimilarImageDelete
//
// RETURNS:
//   - SimilarImageDeleteResponse: The return type of the SimilarImageDelete interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SimilarImageDelete(request *SimilarImageDeleteRequest) (*SimilarImageDeleteResponse, error) {
	result := &SimilarImageDeleteResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSimilarImageDeleteUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SimilarImageSearch
//
// PARAMS:
//   - request: the arguments to SimilarImageSearch
//
// RETURNS:
//   - SimilarImageSearchResponse: The return type of the SimilarImageSearch interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SimilarImageSearch(request *SimilarImageSearchRequest) (*SimilarImageSearchResponse, error) {
	result := &SimilarImageSearchResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSimilarImageSearchUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SimilarImageUpdate
//
// PARAMS:
//   - request: the arguments to SimilarImageUpdate
//
// RETURNS:
//   - SimilarImageUpdateResponse: The return type of the SimilarImageUpdate interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SimilarImageUpdate(request *SimilarImageUpdateRequest) (*SimilarImageUpdateResponse, error) {
	result := &SimilarImageUpdateResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSimilarImageUpdateUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StretchRestore
//
// PARAMS:
//   - request: the arguments to StretchRestore
//
// RETURNS:
//   - StretchRestoreResponse: The return type of the StretchRestore interface.
//   - error: nil if success otherwise the specific error
func (c *Client) StretchRestore(request *StretchRestoreRequest) (*StretchRestoreResponse, error) {
	result := &StretchRestoreResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getStretchRestoreUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StyleTrans
//
// PARAMS:
//   - request: the arguments to StyleTrans
//
// RETURNS:
//   - StyleTransResponse: The return type of the StyleTrans interface.
//   - error: nil if success otherwise the specific error
func (c *Client) StyleTrans(request *StyleTransRequest) (*StyleTransResponse, error) {
	result := &StyleTransResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getStyleTransUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VehicleDetect
//
// PARAMS:
//   - request: the arguments to VehicleDetect
//
// RETURNS:
//   - VehicleDetectResponse: The return type of the VehicleDetect interface.
//   - error: nil if success otherwise the specific error
func (c *Client) VehicleDetect(request *VehicleDetectRequest) (*VehicleDetectResponse, error) {
	result := &VehicleDetectResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getVehicleDetectUri()).
		WithFormBody(request).
		WithHeader("Content-Type", "application/x-www-form-urlencoded").
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
