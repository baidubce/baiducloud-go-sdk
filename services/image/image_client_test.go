package image

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
	IMAGE_CLIENT *Client
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
	// IMAGE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	// ==== AccessToken 鉴权（API Key / Secret Key 换取 AccessToken）====
	// IMAGE_CLIENT, _ = NewClientWithAccessToken(confObj.ApiKey, confObj.SecretKey, confObj.Endpoint)

	// ==== API Key 鉴权 ====
	IMAGE_CLIENT, _ = NewClientWithApiKey(confObj.ApiKey, confObj.Endpoint)

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

func TestClient_AdvancedGeneral(t *testing.T) {
	advancedGeneralRequest := &AdvancedGeneralRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		BaikeNum: util.PtrInt32(int32(0)),
	}
	result := &AdvancedGeneralResponse{}
	result, err := IMAGE_CLIENT.AdvancedGeneral(advancedGeneralRequest)
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
func TestClient_AiRetouchingCreateTask(t *testing.T) {
	IColorParams := &IColorParams{
		Shadow:            util.PtrFloat64(float64(0)),
		SmartRemoveFog:    util.PtrFloat64(float64(0)),
		Tint:              util.PtrFloat64(float64(0)),
		SkinColorRefresh:  util.PtrInt32(int32(0)),
		AiColor:           util.PtrInt32(int32(0)),
		SmartExposure:     util.PtrFloat64(float64(0)),
		Saturation:        util.PtrFloat64(float64(0)),
		Highlight:         util.PtrFloat64(float64(0)),
		BgEnhance:         util.PtrFloat64(float64(0)),
		White:             util.PtrFloat64(float64(0)),
		SharpenAmount:     util.PtrFloat64(float64(0)),
		Temperature:       util.PtrFloat64(float64(0)),
		LutValue:          util.PtrFloat64(float64(0)),
		AutoWhitebalance:  util.PtrFloat64(float64(0)),
		SharpenRadius:     util.PtrFloat64(float64(0)),
		Black:             util.PtrFloat64(float64(0)),
		HslParams:         []*HslParams{},
		AutoExposure:      util.PtrFloat64(float64(0)),
		Brightness:        util.PtrFloat64(float64(0)),
		Exposure:          util.PtrFloat64(float64(0)),
		Contrast:          util.PtrFloat64(float64(0)),
		Vibrance:          util.PtrFloat64(float64(0)),
		SmartWhitebalance: util.PtrFloat64(float64(0)),
		RemoveFog:         util.PtrFloat64(float64(0)),
		LutId:             util.PtrString(""),
	}
	AllHumanOptions := &AllHumanOptions{
		BodyHeighten:     util.PtrFloat64(float64(0)),
		RemoveBgFlaw:     util.PtrFloat64(float64(0)),
		LegLong:          util.PtrFloat64(float64(0)),
		AllSkinColorSame: util.PtrFloat64(float64(0)),
		RemovePureBgFlaw: util.PtrFloat64(float64(0)),
	}
	PartialHumanOptions := &PartialHumanOptions{
		BodySmoothHighpass:      util.PtrFloat64(float64(0)),
		EyeWidthRight:           util.PtrFloat64(float64(0)),
		SkinRed:                 util.PtrFloat64(float64(0)),
		NoseScale:               util.PtrFloat64(float64(0)),
		MakeupEyelashId:         util.PtrInt32(int32(0)),
		RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
		AiBodyFlowThin:          util.PtrFloat64(float64(0)),
		JawWidthLeft:            util.PtrFloat64(float64(0)),
		CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
		FaceColorSame:           util.PtrFloat64(float64(0)),
		FaceSmall:               util.PtrFloat64(float64(0)),
		RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
		EyeScaleSame:            util.PtrBool(false),
		MakeupShadow:            util.PtrFloat64(float64(0)),
		MakeupEyebrow:           util.PtrFloat64(float64(0)),
		MakeupBlush:             util.PtrFloat64(float64(0)),
		MakeupLipstick:          util.PtrFloat64(float64(0)),
		LegThin:                 util.PtrFloat64(float64(0)),
		RemoveDoubleChin:        util.PtrFloat64(float64(0)),
		FaceSymmetry:            util.PtrFloat64(float64(0)),
		EyeHeight:               util.PtrFloat64(float64(0)),
		EyeWidth:                util.PtrFloat64(float64(0)),
		RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
		FaceSmoothNew:           util.PtrFloat64(float64(0)),
		RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
		TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
		NeckThin:                util.PtrFloat64(float64(0)),
		HairlineHeight:          util.PtrFloat64(float64(0)),
		RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
		RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
		ClothesFlawRemove:       util.PtrFloat64(float64(0)),
		FillHairPart:            util.PtrFloat64(float64(0)),
		MakeupEyeballId:         util.PtrInt32(int32(0)),
		RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
		ShinyEye:                util.PtrFloat64(float64(0)),
		MakeupEyebrowId:         util.PtrInt32(int32(0)),
		NoseHeight:              util.PtrFloat64(float64(0)),
		LipPlumpUp:              util.PtrFloat64(float64(0)),
		FaceThin:                util.PtrFloat64(float64(0)),
		HeadSmall:               util.PtrFloat64(float64(0)),
		SkinPrefer:              util.PtrFloat64(float64(0)),
		RemoveWhiteHair:         util.PtrFloat64(float64(0)),
		EyePosition:             util.PtrFloat64(float64(0)),
		FaceSmoothGray:          util.PtrFloat64(float64(0)),
		SkinWhite:               util.PtrFloat64(float64(0)),
		ForeheadHeight:          util.PtrFloat64(float64(0)),
		EyeScale:                util.PtrFloat64(float64(0)),
		EyebrowDistance:         util.PtrFloat64(float64(0)),
		EyeHeightLeft:           util.PtrFloat64(float64(0)),
		EyeWidthLeft:            util.PtrFloat64(float64(0)),
		RemoveDarkCircles:       util.PtrFloat64(float64(0)),
		MakeupFreckleId:         util.PtrInt32(int32(0)),
		JawWidth:                util.PtrFloat64(float64(0)),
		CheekboneWidthRight:     util.PtrFloat64(float64(0)),
		LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
		BodySmoothFine:          util.PtrFloat64(float64(0)),
		BodyColorSame:           util.PtrFloat64(float64(0)),
		RemoveGlassesReflection: util.PtrFloat64(float64(0)),
		MakeupHighlight:         util.PtrFloat64(float64(0)),
		RemovePolymastia:        util.PtrInt32(int32(0)),
		EyeDistanceRight:        util.PtrFloat64(float64(0)),
		MouthScale:              util.PtrFloat64(float64(0)),
		MouthPosition:           util.PtrFloat64(float64(0)),
		ChinHeight:              util.PtrFloat64(float64(0)),
		EyeHeightRight:          util.PtrFloat64(float64(0)),
		MakeupFaceId:            util.PtrInt32(int32(0)),
		BodySmooth:              util.PtrFloat64(float64(0)),
		SkinColor:               util.PtrFloat64(float64(0)),
		FaceV:                   util.PtrFloat64(float64(0)),
		EyeAngleRight:           util.PtrFloat64(float64(0)),
		LipPlumpDown:            util.PtrFloat64(float64(0)),
		FaceWidth:               util.PtrFloat64(float64(0)),
		EyePositionLeft:         util.PtrFloat64(float64(0)),
		EyePositionRight:        util.PtrFloat64(float64(0)),
		NeckLength:              util.PtrFloat64(float64(0)),
		FaceShadow:              util.PtrFloat64(float64(0)),
		SkinBright:              util.PtrFloat64(float64(0)),
		SkinColorId:             util.PtrInt32(int32(0)),
		NoseBridge:              util.PtrFloat64(float64(0)),
		NoseTip:                 util.PtrFloat64(float64(0)),
		LeftSwanNeck:            util.PtrFloat64(float64(0)),
		FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
		EyeScaleRight:           util.PtrFloat64(float64(0)),
		RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
		MouthWidth:              util.PtrFloat64(float64(0)),
		FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
		RemoveBurstHair:         util.PtrFloat64(float64(0)),
		MakeupEyeShadowId:       util.PtrInt32(int32(0)),
		BodySmoothLowpass:       util.PtrFloat64(float64(0)),
		EyeAngle:                util.PtrFloat64(float64(0)),
		AiBodyThin:              util.PtrFloat64(float64(0)),
		MakeupEyelash:           util.PtrFloat64(float64(0)),
		MakeupFace:              util.PtrFloat64(float64(0)),
		NoseWing:                util.PtrFloat64(float64(0)),
		TeethWhite:              util.PtrFloat64(float64(0)),
		BodyThin:                util.PtrFloat64(float64(0)),
		SkinSharpen:             util.PtrFloat64(float64(0)),
		EyeScaleLeft:            util.PtrFloat64(float64(0)),
		RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
		CalvariaHeight:          util.PtrFloat64(float64(0)),
		JawWidthRight:           util.PtrFloat64(float64(0)),
		WaistThin:               util.PtrFloat64(float64(0)),
		FaceHighlight:           util.PtrFloat64(float64(0)),
		RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
		EyebrowThickness:        util.PtrFloat64(float64(0)),
		TeethRepair:             util.PtrInt32(int32(0)),
		RightSwanNeck:           util.PtrFloat64(float64(0)),
		ArmThin:                 util.PtrFloat64(float64(0)),
		MakeupEyeball:           util.PtrFloat64(float64(0)),
		CheekboneWidth:          util.PtrFloat64(float64(0)),
		EyeAngleLeft:            util.PtrFloat64(float64(0)),
		EyeDistanceLeft:         util.PtrFloat64(float64(0)),
		MakeupEyeShadow:         util.PtrFloat64(float64(0)),
		TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
		FaceSmoothFine:          util.PtrFloat64(float64(0)),
		MakeupFreckle:           util.PtrFloat64(float64(0)),
		MakeupBlushId:           util.PtrInt32(int32(0)),
		RemoveLaughLine:         util.PtrFloat64(float64(0)),
		MakeupLipstickId:        util.PtrInt32(int32(0)),
		FaceSmooth:              util.PtrFloat64(float64(0)),
		EyebrowHeight:           util.PtrFloat64(float64(0)),
		EyeDistance:             util.PtrFloat64(float64(0)),
		RemoveFaceMoles:         util.PtrInt32(int32(0)),
		RemoveStretchMark:       util.PtrFloat64(float64(0)),
		RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
	}
	PartialTemplates := &PartialTemplates{
		MaleOld: &PartialHumanOptions{
			BodySmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeWidthRight:           util.PtrFloat64(float64(0)),
			SkinRed:                 util.PtrFloat64(float64(0)),
			NoseScale:               util.PtrFloat64(float64(0)),
			MakeupEyelashId:         util.PtrInt32(int32(0)),
			RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
			AiBodyFlowThin:          util.PtrFloat64(float64(0)),
			JawWidthLeft:            util.PtrFloat64(float64(0)),
			CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
			FaceColorSame:           util.PtrFloat64(float64(0)),
			FaceSmall:               util.PtrFloat64(float64(0)),
			RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
			EyeScaleSame:            util.PtrBool(false),
			MakeupShadow:            util.PtrFloat64(float64(0)),
			MakeupEyebrow:           util.PtrFloat64(float64(0)),
			MakeupBlush:             util.PtrFloat64(float64(0)),
			MakeupLipstick:          util.PtrFloat64(float64(0)),
			LegThin:                 util.PtrFloat64(float64(0)),
			RemoveDoubleChin:        util.PtrFloat64(float64(0)),
			FaceSymmetry:            util.PtrFloat64(float64(0)),
			EyeHeight:               util.PtrFloat64(float64(0)),
			EyeWidth:                util.PtrFloat64(float64(0)),
			RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
			FaceSmoothNew:           util.PtrFloat64(float64(0)),
			RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
			TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
			NeckThin:                util.PtrFloat64(float64(0)),
			HairlineHeight:          util.PtrFloat64(float64(0)),
			RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
			RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
			ClothesFlawRemove:       util.PtrFloat64(float64(0)),
			FillHairPart:            util.PtrFloat64(float64(0)),
			MakeupEyeballId:         util.PtrInt32(int32(0)),
			RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
			ShinyEye:                util.PtrFloat64(float64(0)),
			MakeupEyebrowId:         util.PtrInt32(int32(0)),
			NoseHeight:              util.PtrFloat64(float64(0)),
			LipPlumpUp:              util.PtrFloat64(float64(0)),
			FaceThin:                util.PtrFloat64(float64(0)),
			HeadSmall:               util.PtrFloat64(float64(0)),
			SkinPrefer:              util.PtrFloat64(float64(0)),
			RemoveWhiteHair:         util.PtrFloat64(float64(0)),
			EyePosition:             util.PtrFloat64(float64(0)),
			FaceSmoothGray:          util.PtrFloat64(float64(0)),
			SkinWhite:               util.PtrFloat64(float64(0)),
			ForeheadHeight:          util.PtrFloat64(float64(0)),
			EyeScale:                util.PtrFloat64(float64(0)),
			EyebrowDistance:         util.PtrFloat64(float64(0)),
			EyeHeightLeft:           util.PtrFloat64(float64(0)),
			EyeWidthLeft:            util.PtrFloat64(float64(0)),
			RemoveDarkCircles:       util.PtrFloat64(float64(0)),
			MakeupFreckleId:         util.PtrInt32(int32(0)),
			JawWidth:                util.PtrFloat64(float64(0)),
			CheekboneWidthRight:     util.PtrFloat64(float64(0)),
			LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
			BodySmoothFine:          util.PtrFloat64(float64(0)),
			BodyColorSame:           util.PtrFloat64(float64(0)),
			RemoveGlassesReflection: util.PtrFloat64(float64(0)),
			MakeupHighlight:         util.PtrFloat64(float64(0)),
			RemovePolymastia:        util.PtrInt32(int32(0)),
			EyeDistanceRight:        util.PtrFloat64(float64(0)),
			MouthScale:              util.PtrFloat64(float64(0)),
			MouthPosition:           util.PtrFloat64(float64(0)),
			ChinHeight:              util.PtrFloat64(float64(0)),
			EyeHeightRight:          util.PtrFloat64(float64(0)),
			MakeupFaceId:            util.PtrInt32(int32(0)),
			BodySmooth:              util.PtrFloat64(float64(0)),
			SkinColor:               util.PtrFloat64(float64(0)),
			FaceV:                   util.PtrFloat64(float64(0)),
			EyeAngleRight:           util.PtrFloat64(float64(0)),
			LipPlumpDown:            util.PtrFloat64(float64(0)),
			FaceWidth:               util.PtrFloat64(float64(0)),
			EyePositionLeft:         util.PtrFloat64(float64(0)),
			EyePositionRight:        util.PtrFloat64(float64(0)),
			NeckLength:              util.PtrFloat64(float64(0)),
			FaceShadow:              util.PtrFloat64(float64(0)),
			SkinBright:              util.PtrFloat64(float64(0)),
			SkinColorId:             util.PtrInt32(int32(0)),
			NoseBridge:              util.PtrFloat64(float64(0)),
			NoseTip:                 util.PtrFloat64(float64(0)),
			LeftSwanNeck:            util.PtrFloat64(float64(0)),
			FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeScaleRight:           util.PtrFloat64(float64(0)),
			RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
			MouthWidth:              util.PtrFloat64(float64(0)),
			FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
			RemoveBurstHair:         util.PtrFloat64(float64(0)),
			MakeupEyeShadowId:       util.PtrInt32(int32(0)),
			BodySmoothLowpass:       util.PtrFloat64(float64(0)),
			EyeAngle:                util.PtrFloat64(float64(0)),
			AiBodyThin:              util.PtrFloat64(float64(0)),
			MakeupEyelash:           util.PtrFloat64(float64(0)),
			MakeupFace:              util.PtrFloat64(float64(0)),
			NoseWing:                util.PtrFloat64(float64(0)),
			TeethWhite:              util.PtrFloat64(float64(0)),
			BodyThin:                util.PtrFloat64(float64(0)),
			SkinSharpen:             util.PtrFloat64(float64(0)),
			EyeScaleLeft:            util.PtrFloat64(float64(0)),
			RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
			CalvariaHeight:          util.PtrFloat64(float64(0)),
			JawWidthRight:           util.PtrFloat64(float64(0)),
			WaistThin:               util.PtrFloat64(float64(0)),
			FaceHighlight:           util.PtrFloat64(float64(0)),
			RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
			EyebrowThickness:        util.PtrFloat64(float64(0)),
			TeethRepair:             util.PtrInt32(int32(0)),
			RightSwanNeck:           util.PtrFloat64(float64(0)),
			ArmThin:                 util.PtrFloat64(float64(0)),
			MakeupEyeball:           util.PtrFloat64(float64(0)),
			CheekboneWidth:          util.PtrFloat64(float64(0)),
			EyeAngleLeft:            util.PtrFloat64(float64(0)),
			EyeDistanceLeft:         util.PtrFloat64(float64(0)),
			MakeupEyeShadow:         util.PtrFloat64(float64(0)),
			TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
			FaceSmoothFine:          util.PtrFloat64(float64(0)),
			MakeupFreckle:           util.PtrFloat64(float64(0)),
			MakeupBlushId:           util.PtrInt32(int32(0)),
			RemoveLaughLine:         util.PtrFloat64(float64(0)),
			MakeupLipstickId:        util.PtrInt32(int32(0)),
			FaceSmooth:              util.PtrFloat64(float64(0)),
			EyebrowHeight:           util.PtrFloat64(float64(0)),
			EyeDistance:             util.PtrFloat64(float64(0)),
			RemoveFaceMoles:         util.PtrInt32(int32(0)),
			RemoveStretchMark:       util.PtrFloat64(float64(0)),
			RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
		},
		FemaleOld: &PartialHumanOptions{
			BodySmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeWidthRight:           util.PtrFloat64(float64(0)),
			SkinRed:                 util.PtrFloat64(float64(0)),
			NoseScale:               util.PtrFloat64(float64(0)),
			MakeupEyelashId:         util.PtrInt32(int32(0)),
			RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
			AiBodyFlowThin:          util.PtrFloat64(float64(0)),
			JawWidthLeft:            util.PtrFloat64(float64(0)),
			CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
			FaceColorSame:           util.PtrFloat64(float64(0)),
			FaceSmall:               util.PtrFloat64(float64(0)),
			RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
			EyeScaleSame:            util.PtrBool(false),
			MakeupShadow:            util.PtrFloat64(float64(0)),
			MakeupEyebrow:           util.PtrFloat64(float64(0)),
			MakeupBlush:             util.PtrFloat64(float64(0)),
			MakeupLipstick:          util.PtrFloat64(float64(0)),
			LegThin:                 util.PtrFloat64(float64(0)),
			RemoveDoubleChin:        util.PtrFloat64(float64(0)),
			FaceSymmetry:            util.PtrFloat64(float64(0)),
			EyeHeight:               util.PtrFloat64(float64(0)),
			EyeWidth:                util.PtrFloat64(float64(0)),
			RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
			FaceSmoothNew:           util.PtrFloat64(float64(0)),
			RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
			TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
			NeckThin:                util.PtrFloat64(float64(0)),
			HairlineHeight:          util.PtrFloat64(float64(0)),
			RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
			RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
			ClothesFlawRemove:       util.PtrFloat64(float64(0)),
			FillHairPart:            util.PtrFloat64(float64(0)),
			MakeupEyeballId:         util.PtrInt32(int32(0)),
			RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
			ShinyEye:                util.PtrFloat64(float64(0)),
			MakeupEyebrowId:         util.PtrInt32(int32(0)),
			NoseHeight:              util.PtrFloat64(float64(0)),
			LipPlumpUp:              util.PtrFloat64(float64(0)),
			FaceThin:                util.PtrFloat64(float64(0)),
			HeadSmall:               util.PtrFloat64(float64(0)),
			SkinPrefer:              util.PtrFloat64(float64(0)),
			RemoveWhiteHair:         util.PtrFloat64(float64(0)),
			EyePosition:             util.PtrFloat64(float64(0)),
			FaceSmoothGray:          util.PtrFloat64(float64(0)),
			SkinWhite:               util.PtrFloat64(float64(0)),
			ForeheadHeight:          util.PtrFloat64(float64(0)),
			EyeScale:                util.PtrFloat64(float64(0)),
			EyebrowDistance:         util.PtrFloat64(float64(0)),
			EyeHeightLeft:           util.PtrFloat64(float64(0)),
			EyeWidthLeft:            util.PtrFloat64(float64(0)),
			RemoveDarkCircles:       util.PtrFloat64(float64(0)),
			MakeupFreckleId:         util.PtrInt32(int32(0)),
			JawWidth:                util.PtrFloat64(float64(0)),
			CheekboneWidthRight:     util.PtrFloat64(float64(0)),
			LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
			BodySmoothFine:          util.PtrFloat64(float64(0)),
			BodyColorSame:           util.PtrFloat64(float64(0)),
			RemoveGlassesReflection: util.PtrFloat64(float64(0)),
			MakeupHighlight:         util.PtrFloat64(float64(0)),
			RemovePolymastia:        util.PtrInt32(int32(0)),
			EyeDistanceRight:        util.PtrFloat64(float64(0)),
			MouthScale:              util.PtrFloat64(float64(0)),
			MouthPosition:           util.PtrFloat64(float64(0)),
			ChinHeight:              util.PtrFloat64(float64(0)),
			EyeHeightRight:          util.PtrFloat64(float64(0)),
			MakeupFaceId:            util.PtrInt32(int32(0)),
			BodySmooth:              util.PtrFloat64(float64(0)),
			SkinColor:               util.PtrFloat64(float64(0)),
			FaceV:                   util.PtrFloat64(float64(0)),
			EyeAngleRight:           util.PtrFloat64(float64(0)),
			LipPlumpDown:            util.PtrFloat64(float64(0)),
			FaceWidth:               util.PtrFloat64(float64(0)),
			EyePositionLeft:         util.PtrFloat64(float64(0)),
			EyePositionRight:        util.PtrFloat64(float64(0)),
			NeckLength:              util.PtrFloat64(float64(0)),
			FaceShadow:              util.PtrFloat64(float64(0)),
			SkinBright:              util.PtrFloat64(float64(0)),
			SkinColorId:             util.PtrInt32(int32(0)),
			NoseBridge:              util.PtrFloat64(float64(0)),
			NoseTip:                 util.PtrFloat64(float64(0)),
			LeftSwanNeck:            util.PtrFloat64(float64(0)),
			FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeScaleRight:           util.PtrFloat64(float64(0)),
			RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
			MouthWidth:              util.PtrFloat64(float64(0)),
			FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
			RemoveBurstHair:         util.PtrFloat64(float64(0)),
			MakeupEyeShadowId:       util.PtrInt32(int32(0)),
			BodySmoothLowpass:       util.PtrFloat64(float64(0)),
			EyeAngle:                util.PtrFloat64(float64(0)),
			AiBodyThin:              util.PtrFloat64(float64(0)),
			MakeupEyelash:           util.PtrFloat64(float64(0)),
			MakeupFace:              util.PtrFloat64(float64(0)),
			NoseWing:                util.PtrFloat64(float64(0)),
			TeethWhite:              util.PtrFloat64(float64(0)),
			BodyThin:                util.PtrFloat64(float64(0)),
			SkinSharpen:             util.PtrFloat64(float64(0)),
			EyeScaleLeft:            util.PtrFloat64(float64(0)),
			RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
			CalvariaHeight:          util.PtrFloat64(float64(0)),
			JawWidthRight:           util.PtrFloat64(float64(0)),
			WaistThin:               util.PtrFloat64(float64(0)),
			FaceHighlight:           util.PtrFloat64(float64(0)),
			RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
			EyebrowThickness:        util.PtrFloat64(float64(0)),
			TeethRepair:             util.PtrInt32(int32(0)),
			RightSwanNeck:           util.PtrFloat64(float64(0)),
			ArmThin:                 util.PtrFloat64(float64(0)),
			MakeupEyeball:           util.PtrFloat64(float64(0)),
			CheekboneWidth:          util.PtrFloat64(float64(0)),
			EyeAngleLeft:            util.PtrFloat64(float64(0)),
			EyeDistanceLeft:         util.PtrFloat64(float64(0)),
			MakeupEyeShadow:         util.PtrFloat64(float64(0)),
			TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
			FaceSmoothFine:          util.PtrFloat64(float64(0)),
			MakeupFreckle:           util.PtrFloat64(float64(0)),
			MakeupBlushId:           util.PtrInt32(int32(0)),
			RemoveLaughLine:         util.PtrFloat64(float64(0)),
			MakeupLipstickId:        util.PtrInt32(int32(0)),
			FaceSmooth:              util.PtrFloat64(float64(0)),
			EyebrowHeight:           util.PtrFloat64(float64(0)),
			EyeDistance:             util.PtrFloat64(float64(0)),
			RemoveFaceMoles:         util.PtrInt32(int32(0)),
			RemoveStretchMark:       util.PtrFloat64(float64(0)),
			RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
		},
		FemaleYoung: &PartialHumanOptions{
			BodySmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeWidthRight:           util.PtrFloat64(float64(0)),
			SkinRed:                 util.PtrFloat64(float64(0)),
			NoseScale:               util.PtrFloat64(float64(0)),
			MakeupEyelashId:         util.PtrInt32(int32(0)),
			RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
			AiBodyFlowThin:          util.PtrFloat64(float64(0)),
			JawWidthLeft:            util.PtrFloat64(float64(0)),
			CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
			FaceColorSame:           util.PtrFloat64(float64(0)),
			FaceSmall:               util.PtrFloat64(float64(0)),
			RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
			EyeScaleSame:            util.PtrBool(false),
			MakeupShadow:            util.PtrFloat64(float64(0)),
			MakeupEyebrow:           util.PtrFloat64(float64(0)),
			MakeupBlush:             util.PtrFloat64(float64(0)),
			MakeupLipstick:          util.PtrFloat64(float64(0)),
			LegThin:                 util.PtrFloat64(float64(0)),
			RemoveDoubleChin:        util.PtrFloat64(float64(0)),
			FaceSymmetry:            util.PtrFloat64(float64(0)),
			EyeHeight:               util.PtrFloat64(float64(0)),
			EyeWidth:                util.PtrFloat64(float64(0)),
			RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
			FaceSmoothNew:           util.PtrFloat64(float64(0)),
			RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
			TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
			NeckThin:                util.PtrFloat64(float64(0)),
			HairlineHeight:          util.PtrFloat64(float64(0)),
			RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
			RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
			ClothesFlawRemove:       util.PtrFloat64(float64(0)),
			FillHairPart:            util.PtrFloat64(float64(0)),
			MakeupEyeballId:         util.PtrInt32(int32(0)),
			RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
			ShinyEye:                util.PtrFloat64(float64(0)),
			MakeupEyebrowId:         util.PtrInt32(int32(0)),
			NoseHeight:              util.PtrFloat64(float64(0)),
			LipPlumpUp:              util.PtrFloat64(float64(0)),
			FaceThin:                util.PtrFloat64(float64(0)),
			HeadSmall:               util.PtrFloat64(float64(0)),
			SkinPrefer:              util.PtrFloat64(float64(0)),
			RemoveWhiteHair:         util.PtrFloat64(float64(0)),
			EyePosition:             util.PtrFloat64(float64(0)),
			FaceSmoothGray:          util.PtrFloat64(float64(0)),
			SkinWhite:               util.PtrFloat64(float64(0)),
			ForeheadHeight:          util.PtrFloat64(float64(0)),
			EyeScale:                util.PtrFloat64(float64(0)),
			EyebrowDistance:         util.PtrFloat64(float64(0)),
			EyeHeightLeft:           util.PtrFloat64(float64(0)),
			EyeWidthLeft:            util.PtrFloat64(float64(0)),
			RemoveDarkCircles:       util.PtrFloat64(float64(0)),
			MakeupFreckleId:         util.PtrInt32(int32(0)),
			JawWidth:                util.PtrFloat64(float64(0)),
			CheekboneWidthRight:     util.PtrFloat64(float64(0)),
			LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
			BodySmoothFine:          util.PtrFloat64(float64(0)),
			BodyColorSame:           util.PtrFloat64(float64(0)),
			RemoveGlassesReflection: util.PtrFloat64(float64(0)),
			MakeupHighlight:         util.PtrFloat64(float64(0)),
			RemovePolymastia:        util.PtrInt32(int32(0)),
			EyeDistanceRight:        util.PtrFloat64(float64(0)),
			MouthScale:              util.PtrFloat64(float64(0)),
			MouthPosition:           util.PtrFloat64(float64(0)),
			ChinHeight:              util.PtrFloat64(float64(0)),
			EyeHeightRight:          util.PtrFloat64(float64(0)),
			MakeupFaceId:            util.PtrInt32(int32(0)),
			BodySmooth:              util.PtrFloat64(float64(0)),
			SkinColor:               util.PtrFloat64(float64(0)),
			FaceV:                   util.PtrFloat64(float64(0)),
			EyeAngleRight:           util.PtrFloat64(float64(0)),
			LipPlumpDown:            util.PtrFloat64(float64(0)),
			FaceWidth:               util.PtrFloat64(float64(0)),
			EyePositionLeft:         util.PtrFloat64(float64(0)),
			EyePositionRight:        util.PtrFloat64(float64(0)),
			NeckLength:              util.PtrFloat64(float64(0)),
			FaceShadow:              util.PtrFloat64(float64(0)),
			SkinBright:              util.PtrFloat64(float64(0)),
			SkinColorId:             util.PtrInt32(int32(0)),
			NoseBridge:              util.PtrFloat64(float64(0)),
			NoseTip:                 util.PtrFloat64(float64(0)),
			LeftSwanNeck:            util.PtrFloat64(float64(0)),
			FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeScaleRight:           util.PtrFloat64(float64(0)),
			RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
			MouthWidth:              util.PtrFloat64(float64(0)),
			FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
			RemoveBurstHair:         util.PtrFloat64(float64(0)),
			MakeupEyeShadowId:       util.PtrInt32(int32(0)),
			BodySmoothLowpass:       util.PtrFloat64(float64(0)),
			EyeAngle:                util.PtrFloat64(float64(0)),
			AiBodyThin:              util.PtrFloat64(float64(0)),
			MakeupEyelash:           util.PtrFloat64(float64(0)),
			MakeupFace:              util.PtrFloat64(float64(0)),
			NoseWing:                util.PtrFloat64(float64(0)),
			TeethWhite:              util.PtrFloat64(float64(0)),
			BodyThin:                util.PtrFloat64(float64(0)),
			SkinSharpen:             util.PtrFloat64(float64(0)),
			EyeScaleLeft:            util.PtrFloat64(float64(0)),
			RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
			CalvariaHeight:          util.PtrFloat64(float64(0)),
			JawWidthRight:           util.PtrFloat64(float64(0)),
			WaistThin:               util.PtrFloat64(float64(0)),
			FaceHighlight:           util.PtrFloat64(float64(0)),
			RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
			EyebrowThickness:        util.PtrFloat64(float64(0)),
			TeethRepair:             util.PtrInt32(int32(0)),
			RightSwanNeck:           util.PtrFloat64(float64(0)),
			ArmThin:                 util.PtrFloat64(float64(0)),
			MakeupEyeball:           util.PtrFloat64(float64(0)),
			CheekboneWidth:          util.PtrFloat64(float64(0)),
			EyeAngleLeft:            util.PtrFloat64(float64(0)),
			EyeDistanceLeft:         util.PtrFloat64(float64(0)),
			MakeupEyeShadow:         util.PtrFloat64(float64(0)),
			TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
			FaceSmoothFine:          util.PtrFloat64(float64(0)),
			MakeupFreckle:           util.PtrFloat64(float64(0)),
			MakeupBlushId:           util.PtrInt32(int32(0)),
			RemoveLaughLine:         util.PtrFloat64(float64(0)),
			MakeupLipstickId:        util.PtrInt32(int32(0)),
			FaceSmooth:              util.PtrFloat64(float64(0)),
			EyebrowHeight:           util.PtrFloat64(float64(0)),
			EyeDistance:             util.PtrFloat64(float64(0)),
			RemoveFaceMoles:         util.PtrInt32(int32(0)),
			RemoveStretchMark:       util.PtrFloat64(float64(0)),
			RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
		},
		MaleYoung: &PartialHumanOptions{
			BodySmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeWidthRight:           util.PtrFloat64(float64(0)),
			SkinRed:                 util.PtrFloat64(float64(0)),
			NoseScale:               util.PtrFloat64(float64(0)),
			MakeupEyelashId:         util.PtrInt32(int32(0)),
			RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
			AiBodyFlowThin:          util.PtrFloat64(float64(0)),
			JawWidthLeft:            util.PtrFloat64(float64(0)),
			CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
			FaceColorSame:           util.PtrFloat64(float64(0)),
			FaceSmall:               util.PtrFloat64(float64(0)),
			RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
			EyeScaleSame:            util.PtrBool(false),
			MakeupShadow:            util.PtrFloat64(float64(0)),
			MakeupEyebrow:           util.PtrFloat64(float64(0)),
			MakeupBlush:             util.PtrFloat64(float64(0)),
			MakeupLipstick:          util.PtrFloat64(float64(0)),
			LegThin:                 util.PtrFloat64(float64(0)),
			RemoveDoubleChin:        util.PtrFloat64(float64(0)),
			FaceSymmetry:            util.PtrFloat64(float64(0)),
			EyeHeight:               util.PtrFloat64(float64(0)),
			EyeWidth:                util.PtrFloat64(float64(0)),
			RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
			FaceSmoothNew:           util.PtrFloat64(float64(0)),
			RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
			TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
			NeckThin:                util.PtrFloat64(float64(0)),
			HairlineHeight:          util.PtrFloat64(float64(0)),
			RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
			RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
			ClothesFlawRemove:       util.PtrFloat64(float64(0)),
			FillHairPart:            util.PtrFloat64(float64(0)),
			MakeupEyeballId:         util.PtrInt32(int32(0)),
			RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
			ShinyEye:                util.PtrFloat64(float64(0)),
			MakeupEyebrowId:         util.PtrInt32(int32(0)),
			NoseHeight:              util.PtrFloat64(float64(0)),
			LipPlumpUp:              util.PtrFloat64(float64(0)),
			FaceThin:                util.PtrFloat64(float64(0)),
			HeadSmall:               util.PtrFloat64(float64(0)),
			SkinPrefer:              util.PtrFloat64(float64(0)),
			RemoveWhiteHair:         util.PtrFloat64(float64(0)),
			EyePosition:             util.PtrFloat64(float64(0)),
			FaceSmoothGray:          util.PtrFloat64(float64(0)),
			SkinWhite:               util.PtrFloat64(float64(0)),
			ForeheadHeight:          util.PtrFloat64(float64(0)),
			EyeScale:                util.PtrFloat64(float64(0)),
			EyebrowDistance:         util.PtrFloat64(float64(0)),
			EyeHeightLeft:           util.PtrFloat64(float64(0)),
			EyeWidthLeft:            util.PtrFloat64(float64(0)),
			RemoveDarkCircles:       util.PtrFloat64(float64(0)),
			MakeupFreckleId:         util.PtrInt32(int32(0)),
			JawWidth:                util.PtrFloat64(float64(0)),
			CheekboneWidthRight:     util.PtrFloat64(float64(0)),
			LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
			BodySmoothFine:          util.PtrFloat64(float64(0)),
			BodyColorSame:           util.PtrFloat64(float64(0)),
			RemoveGlassesReflection: util.PtrFloat64(float64(0)),
			MakeupHighlight:         util.PtrFloat64(float64(0)),
			RemovePolymastia:        util.PtrInt32(int32(0)),
			EyeDistanceRight:        util.PtrFloat64(float64(0)),
			MouthScale:              util.PtrFloat64(float64(0)),
			MouthPosition:           util.PtrFloat64(float64(0)),
			ChinHeight:              util.PtrFloat64(float64(0)),
			EyeHeightRight:          util.PtrFloat64(float64(0)),
			MakeupFaceId:            util.PtrInt32(int32(0)),
			BodySmooth:              util.PtrFloat64(float64(0)),
			SkinColor:               util.PtrFloat64(float64(0)),
			FaceV:                   util.PtrFloat64(float64(0)),
			EyeAngleRight:           util.PtrFloat64(float64(0)),
			LipPlumpDown:            util.PtrFloat64(float64(0)),
			FaceWidth:               util.PtrFloat64(float64(0)),
			EyePositionLeft:         util.PtrFloat64(float64(0)),
			EyePositionRight:        util.PtrFloat64(float64(0)),
			NeckLength:              util.PtrFloat64(float64(0)),
			FaceShadow:              util.PtrFloat64(float64(0)),
			SkinBright:              util.PtrFloat64(float64(0)),
			SkinColorId:             util.PtrInt32(int32(0)),
			NoseBridge:              util.PtrFloat64(float64(0)),
			NoseTip:                 util.PtrFloat64(float64(0)),
			LeftSwanNeck:            util.PtrFloat64(float64(0)),
			FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeScaleRight:           util.PtrFloat64(float64(0)),
			RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
			MouthWidth:              util.PtrFloat64(float64(0)),
			FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
			RemoveBurstHair:         util.PtrFloat64(float64(0)),
			MakeupEyeShadowId:       util.PtrInt32(int32(0)),
			BodySmoothLowpass:       util.PtrFloat64(float64(0)),
			EyeAngle:                util.PtrFloat64(float64(0)),
			AiBodyThin:              util.PtrFloat64(float64(0)),
			MakeupEyelash:           util.PtrFloat64(float64(0)),
			MakeupFace:              util.PtrFloat64(float64(0)),
			NoseWing:                util.PtrFloat64(float64(0)),
			TeethWhite:              util.PtrFloat64(float64(0)),
			BodyThin:                util.PtrFloat64(float64(0)),
			SkinSharpen:             util.PtrFloat64(float64(0)),
			EyeScaleLeft:            util.PtrFloat64(float64(0)),
			RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
			CalvariaHeight:          util.PtrFloat64(float64(0)),
			JawWidthRight:           util.PtrFloat64(float64(0)),
			WaistThin:               util.PtrFloat64(float64(0)),
			FaceHighlight:           util.PtrFloat64(float64(0)),
			RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
			EyebrowThickness:        util.PtrFloat64(float64(0)),
			TeethRepair:             util.PtrInt32(int32(0)),
			RightSwanNeck:           util.PtrFloat64(float64(0)),
			ArmThin:                 util.PtrFloat64(float64(0)),
			MakeupEyeball:           util.PtrFloat64(float64(0)),
			CheekboneWidth:          util.PtrFloat64(float64(0)),
			EyeAngleLeft:            util.PtrFloat64(float64(0)),
			EyeDistanceLeft:         util.PtrFloat64(float64(0)),
			MakeupEyeShadow:         util.PtrFloat64(float64(0)),
			TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
			FaceSmoothFine:          util.PtrFloat64(float64(0)),
			MakeupFreckle:           util.PtrFloat64(float64(0)),
			MakeupBlushId:           util.PtrInt32(int32(0)),
			RemoveLaughLine:         util.PtrFloat64(float64(0)),
			MakeupLipstickId:        util.PtrInt32(int32(0)),
			FaceSmooth:              util.PtrFloat64(float64(0)),
			EyebrowHeight:           util.PtrFloat64(float64(0)),
			EyeDistance:             util.PtrFloat64(float64(0)),
			RemoveFaceMoles:         util.PtrInt32(int32(0)),
			RemoveStretchMark:       util.PtrFloat64(float64(0)),
			RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
		},
		Child: &PartialHumanOptions{
			BodySmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeWidthRight:           util.PtrFloat64(float64(0)),
			SkinRed:                 util.PtrFloat64(float64(0)),
			NoseScale:               util.PtrFloat64(float64(0)),
			MakeupEyelashId:         util.PtrInt32(int32(0)),
			RemoveEyeStreaks:        util.PtrFloat64(float64(0)),
			AiBodyFlowThin:          util.PtrFloat64(float64(0)),
			JawWidthLeft:            util.PtrFloat64(float64(0)),
			CheekboneWidthLeft:      util.PtrFloat64(float64(0)),
			FaceColorSame:           util.PtrFloat64(float64(0)),
			FaceSmall:               util.PtrFloat64(float64(0)),
			RemoveForeheadWrinkles:  util.PtrFloat64(float64(0)),
			EyeScaleSame:            util.PtrBool(false),
			MakeupShadow:            util.PtrFloat64(float64(0)),
			MakeupEyebrow:           util.PtrFloat64(float64(0)),
			MakeupBlush:             util.PtrFloat64(float64(0)),
			MakeupLipstick:          util.PtrFloat64(float64(0)),
			LegThin:                 util.PtrFloat64(float64(0)),
			RemoveDoubleChin:        util.PtrFloat64(float64(0)),
			FaceSymmetry:            util.PtrFloat64(float64(0)),
			EyeHeight:               util.PtrFloat64(float64(0)),
			EyeWidth:                util.PtrFloat64(float64(0)),
			RemoveBurstHairBody:     util.PtrFloat64(float64(0)),
			FaceSmoothNew:           util.PtrFloat64(float64(0)),
			RemoveBodyFlaw:          util.PtrFloat64(float64(0)),
			TeethWhiteAddBright:     util.PtrFloat64(float64(0)),
			NeckThin:                util.PtrFloat64(float64(0)),
			HairlineHeight:          util.PtrFloat64(float64(0)),
			RightEyebrowEnhance:     util.PtrFloat64(float64(0)),
			RemoveFaceGlossy:        util.PtrFloat64(float64(0)),
			ClothesFlawRemove:       util.PtrFloat64(float64(0)),
			FillHairPart:            util.PtrFloat64(float64(0)),
			MakeupEyeballId:         util.PtrInt32(int32(0)),
			RemoveLipWrinkles:       util.PtrFloat64(float64(0)),
			ShinyEye:                util.PtrFloat64(float64(0)),
			MakeupEyebrowId:         util.PtrInt32(int32(0)),
			NoseHeight:              util.PtrFloat64(float64(0)),
			LipPlumpUp:              util.PtrFloat64(float64(0)),
			FaceThin:                util.PtrFloat64(float64(0)),
			HeadSmall:               util.PtrFloat64(float64(0)),
			SkinPrefer:              util.PtrFloat64(float64(0)),
			RemoveWhiteHair:         util.PtrFloat64(float64(0)),
			EyePosition:             util.PtrFloat64(float64(0)),
			FaceSmoothGray:          util.PtrFloat64(float64(0)),
			SkinWhite:               util.PtrFloat64(float64(0)),
			ForeheadHeight:          util.PtrFloat64(float64(0)),
			EyeScale:                util.PtrFloat64(float64(0)),
			EyebrowDistance:         util.PtrFloat64(float64(0)),
			EyeHeightLeft:           util.PtrFloat64(float64(0)),
			EyeWidthLeft:            util.PtrFloat64(float64(0)),
			RemoveDarkCircles:       util.PtrFloat64(float64(0)),
			MakeupFreckleId:         util.PtrInt32(int32(0)),
			JawWidth:                util.PtrFloat64(float64(0)),
			CheekboneWidthRight:     util.PtrFloat64(float64(0)),
			LeftEyebrowEnhance:      util.PtrFloat64(float64(0)),
			BodySmoothFine:          util.PtrFloat64(float64(0)),
			BodyColorSame:           util.PtrFloat64(float64(0)),
			RemoveGlassesReflection: util.PtrFloat64(float64(0)),
			MakeupHighlight:         util.PtrFloat64(float64(0)),
			RemovePolymastia:        util.PtrInt32(int32(0)),
			EyeDistanceRight:        util.PtrFloat64(float64(0)),
			MouthScale:              util.PtrFloat64(float64(0)),
			MouthPosition:           util.PtrFloat64(float64(0)),
			ChinHeight:              util.PtrFloat64(float64(0)),
			EyeHeightRight:          util.PtrFloat64(float64(0)),
			MakeupFaceId:            util.PtrInt32(int32(0)),
			BodySmooth:              util.PtrFloat64(float64(0)),
			SkinColor:               util.PtrFloat64(float64(0)),
			FaceV:                   util.PtrFloat64(float64(0)),
			EyeAngleRight:           util.PtrFloat64(float64(0)),
			LipPlumpDown:            util.PtrFloat64(float64(0)),
			FaceWidth:               util.PtrFloat64(float64(0)),
			EyePositionLeft:         util.PtrFloat64(float64(0)),
			EyePositionRight:        util.PtrFloat64(float64(0)),
			NeckLength:              util.PtrFloat64(float64(0)),
			FaceShadow:              util.PtrFloat64(float64(0)),
			SkinBright:              util.PtrFloat64(float64(0)),
			SkinColorId:             util.PtrInt32(int32(0)),
			NoseBridge:              util.PtrFloat64(float64(0)),
			NoseTip:                 util.PtrFloat64(float64(0)),
			LeftSwanNeck:            util.PtrFloat64(float64(0)),
			FaceSmoothHighpass:      util.PtrFloat64(float64(0)),
			EyeScaleRight:           util.PtrFloat64(float64(0)),
			RemoveFaceFlaw:          util.PtrFloat64(float64(0)),
			MouthWidth:              util.PtrFloat64(float64(0)),
			FaceSmoothLowpass:       util.PtrFloat64(float64(0)),
			RemoveBurstHair:         util.PtrFloat64(float64(0)),
			MakeupEyeShadowId:       util.PtrInt32(int32(0)),
			BodySmoothLowpass:       util.PtrFloat64(float64(0)),
			EyeAngle:                util.PtrFloat64(float64(0)),
			AiBodyThin:              util.PtrFloat64(float64(0)),
			MakeupEyelash:           util.PtrFloat64(float64(0)),
			MakeupFace:              util.PtrFloat64(float64(0)),
			NoseWing:                util.PtrFloat64(float64(0)),
			TeethWhite:              util.PtrFloat64(float64(0)),
			BodyThin:                util.PtrFloat64(float64(0)),
			SkinSharpen:             util.PtrFloat64(float64(0)),
			EyeScaleLeft:            util.PtrFloat64(float64(0)),
			RemoveNeckWrinkles:      util.PtrFloat64(float64(0)),
			CalvariaHeight:          util.PtrFloat64(float64(0)),
			JawWidthRight:           util.PtrFloat64(float64(0)),
			WaistThin:               util.PtrFloat64(float64(0)),
			FaceHighlight:           util.PtrFloat64(float64(0)),
			RemoveEyeAroundWrinkles: util.PtrFloat64(float64(0)),
			EyebrowThickness:        util.PtrFloat64(float64(0)),
			TeethRepair:             util.PtrInt32(int32(0)),
			RightSwanNeck:           util.PtrFloat64(float64(0)),
			ArmThin:                 util.PtrFloat64(float64(0)),
			MakeupEyeball:           util.PtrFloat64(float64(0)),
			CheekboneWidth:          util.PtrFloat64(float64(0)),
			EyeAngleLeft:            util.PtrFloat64(float64(0)),
			EyeDistanceLeft:         util.PtrFloat64(float64(0)),
			MakeupEyeShadow:         util.PtrFloat64(float64(0)),
			TeethWhiteDesYellow:     util.PtrFloat64(float64(0)),
			FaceSmoothFine:          util.PtrFloat64(float64(0)),
			MakeupFreckle:           util.PtrFloat64(float64(0)),
			MakeupBlushId:           util.PtrInt32(int32(0)),
			RemoveLaughLine:         util.PtrFloat64(float64(0)),
			MakeupLipstickId:        util.PtrInt32(int32(0)),
			FaceSmooth:              util.PtrFloat64(float64(0)),
			EyebrowHeight:           util.PtrFloat64(float64(0)),
			EyeDistance:             util.PtrFloat64(float64(0)),
			RemoveFaceMoles:         util.PtrInt32(int32(0)),
			RemoveStretchMark:       util.PtrFloat64(float64(0)),
			RemoveBurstHairBack:     util.PtrFloat64(float64(0)),
		},
	}
	TransformOptions := &TransformOptions{
		AutoCorrectAngle: util.PtrBool(false),
		SizeCompress:     util.PtrBool(false),
	}
	aiRetouchingCreateTaskRequest := &AiRetouchingCreateTaskRequest{
		Image:               util.PtrString(""),
		Url:                 util.PtrString(""),
		CallbackData:        util.PtrString(""),
		IColorParams:        IColorParams,
		AllHumanOptions:     AllHumanOptions,
		PartialHumanOptions: PartialHumanOptions,
		PartialTemplates:    PartialTemplates,
		TransformOptions:    TransformOptions,
	}
	result := &AiRetouchingCreateTaskResponse{}
	result, err := IMAGE_CLIENT.AiRetouchingCreateTask(aiRetouchingCreateTaskRequest)
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
func TestClient_AiRetouchingQueryTask(t *testing.T) {
	aiRetouchingQueryTaskRequest := &AiRetouchingQueryTaskRequest{
		TaskId: util.PtrString(""),
	}
	result := &AiRetouchingQueryTaskResponse{}
	result, err := IMAGE_CLIENT.AiRetouchingQueryTask(aiRetouchingQueryTaskRequest)
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
func TestClient_Animal(t *testing.T) {
	animalRequest := &AnimalRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		TopNum:   util.PtrInt32(int32(0)),
		BaikeNum: util.PtrInt32(int32(0)),
	}
	result := &AnimalResponse{}
	result, err := IMAGE_CLIENT.Animal(animalRequest)
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
func TestClient_Car(t *testing.T) {
	carRequest := &CarRequest{
		Image:       util.PtrString(""),
		Url:         util.PtrString(""),
		TopNum:      util.PtrInt32(int32(0)),
		BaikeNum:    util.PtrInt32(int32(0)),
		OutputBrand: util.PtrBool(false),
	}
	result := &CarResponse{}
	result, err := IMAGE_CLIENT.Car(carRequest)
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
func TestClient_ColorEnhance(t *testing.T) {
	colorEnhanceRequest := &ColorEnhanceRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &ColorEnhanceResponse{}
	result, err := IMAGE_CLIENT.ColorEnhance(colorEnhanceRequest)
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
func TestClient_Colourize(t *testing.T) {
	colourizeRequest := &ColourizeRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &ColourizeResponse{}
	result, err := IMAGE_CLIENT.Colourize(colourizeRequest)
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
func TestClient_ContrastEnhance(t *testing.T) {
	contrastEnhanceRequest := &ContrastEnhanceRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &ContrastEnhanceResponse{}
	result, err := IMAGE_CLIENT.ContrastEnhance(contrastEnhanceRequest)
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
func TestClient_Dehaze(t *testing.T) {
	dehazeRequest := &DehazeRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &DehazeResponse{}
	result, err := IMAGE_CLIENT.Dehaze(dehazeRequest)
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
func TestClient_Dish(t *testing.T) {
	dishRequest := &DishRequest{
		Image:           util.PtrString(""),
		Url:             util.PtrString(""),
		TopNum:          util.PtrInt32(int32(0)),
		FilterThreshold: util.PtrFloat32(float32(0)),
		BaikeNum:        util.PtrInt32(int32(0)),
	}
	result := &DishResponse{}
	result, err := IMAGE_CLIENT.Dish(dishRequest)
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
func TestClient_DocRepair(t *testing.T) {
	docRepairRequest := &DocRepairRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &DocRepairResponse{}
	result, err := IMAGE_CLIENT.DocRepair(docRepairRequest)
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
func TestClient_ImageDefinitionEnhance(t *testing.T) {
	imageDefinitionEnhanceRequest := &ImageDefinitionEnhanceRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &ImageDefinitionEnhanceResponse{}
	result, err := IMAGE_CLIENT.ImageDefinitionEnhance(imageDefinitionEnhanceRequest)
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
func TestClient_ImageQualityEnhance(t *testing.T) {
	imageQualityEnhanceRequest := &ImageQualityEnhanceRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &ImageQualityEnhanceResponse{}
	result, err := IMAGE_CLIENT.ImageQualityEnhance(imageQualityEnhanceRequest)
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
func TestClient_ImageUnderstandingGetResult(t *testing.T) {
	imageUnderstandingGetResultRequest := &ImageUnderstandingGetResultRequest{
		TaskId: util.PtrString(""),
	}
	result := &ImageUnderstandingGetResultResponse{}
	result, err := IMAGE_CLIENT.ImageUnderstandingGetResult(imageUnderstandingGetResultRequest)
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
func TestClient_ImageUnderstandingRequest(t *testing.T) {
	imageUnderstandingRequestRequest := &ImageUnderstandingRequestRequest{
		Question: util.PtrString(""),
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
	}
	result := &ImageUnderstandingRequestResponse{}
	result, err := IMAGE_CLIENT.ImageUnderstandingRequest(imageUnderstandingRequestRequest)
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
func TestClient_Ingredient(t *testing.T) {
	ingredientRequest := &IngredientRequest{
		Image:  util.PtrString(""),
		Url:    util.PtrString(""),
		TopNum: util.PtrInt32(int32(0)),
	}
	result := &IngredientResponse{}
	result, err := IMAGE_CLIENT.Ingredient(ingredientRequest)
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
func TestClient_Inpainting(t *testing.T) {
	inpaintingRequest := &InpaintingRequest{
		Rectangle: util.PtrString(""),
		Image:     util.PtrString(""),
		Url:       util.PtrString(""),
	}
	result := &InpaintingResponse{}
	result, err := IMAGE_CLIENT.Inpainting(inpaintingRequest)
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
func TestClient_Landmark(t *testing.T) {
	landmarkRequest := &LandmarkRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &LandmarkResponse{}
	result, err := IMAGE_CLIENT.Landmark(landmarkRequest)
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
func TestClient_Logo(t *testing.T) {
	logoRequest := &LogoRequest{
		Image:     util.PtrString(""),
		Url:       util.PtrString(""),
		CustomLib: util.PtrBool(false),
	}
	result := &LogoResponse{}
	result, err := IMAGE_CLIENT.Logo(logoRequest)
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
func TestClient_LogoAdd(t *testing.T) {
	logoAddRequest := &LogoAddRequest{
		Brief: util.PtrString(""),
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &LogoAddResponse{}
	result, err := IMAGE_CLIENT.LogoAdd(logoAddRequest)
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
func TestClient_LogoDelete(t *testing.T) {
	logoDeleteRequest := &LogoDeleteRequest{
		Image:    util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &LogoDeleteResponse{}
	result, err := IMAGE_CLIENT.LogoDelete(logoDeleteRequest)
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
func TestClient_MaterielImageAdd(t *testing.T) {
	materielImageAddRequest := &MaterielImageAddRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
		Brief: util.PtrString(""),
		Tags:  util.PtrString(""),
	}
	result := &MaterielImageAddResponse{}
	result, err := IMAGE_CLIENT.MaterielImageAdd(materielImageAddRequest)
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
func TestClient_MaterielImageDelete(t *testing.T) {
	materielImageDeleteRequest := &MaterielImageDeleteRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &MaterielImageDeleteResponse{}
	result, err := IMAGE_CLIENT.MaterielImageDelete(materielImageDeleteRequest)
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
func TestClient_MaterielImageSearch(t *testing.T) {
	materielImageSearchRequest := &MaterielImageSearchRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		Tags:     util.PtrString(""),
		TagLogic: util.PtrInt32(int32(0)),
		Pn:       util.PtrInt32(int32(0)),
		Rn:       util.PtrInt32(int32(0)),
	}
	result := &MaterielImageSearchResponse{}
	result, err := IMAGE_CLIENT.MaterielImageSearch(materielImageSearchRequest)
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
func TestClient_MaterielImageUpdate(t *testing.T) {
	materielImageUpdateRequest := &MaterielImageUpdateRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
		Brief:    util.PtrString(""),
		Tags:     util.PtrString(""),
	}
	result := &MaterielImageUpdateResponse{}
	result, err := IMAGE_CLIENT.MaterielImageUpdate(materielImageUpdateRequest)
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
func TestClient_MultiObjectDetect(t *testing.T) {
	multiObjectDetectRequest := &MultiObjectDetectRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &MultiObjectDetectResponse{}
	result, err := IMAGE_CLIENT.MultiObjectDetect(multiObjectDetectRequest)
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
func TestClient_ObjectDetect(t *testing.T) {
	objectDetectRequest := &ObjectDetectRequest{
		Image:    util.PtrString(""),
		WithFace: util.PtrInt32(int32(0)),
	}
	result := &ObjectDetectResponse{}
	result, err := IMAGE_CLIENT.ObjectDetect(objectDetectRequest)
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
func TestClient_PicturebookImageAdd(t *testing.T) {
	picturebookImageAddRequest := &PicturebookImageAddRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
		Brief: util.PtrString(""),
		Tags:  util.PtrString(""),
	}
	result := &PicturebookImageAddResponse{}
	result, err := IMAGE_CLIENT.PicturebookImageAdd(picturebookImageAddRequest)
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
func TestClient_PicturebookImageDelete(t *testing.T) {
	picturebookImageDeleteRequest := &PicturebookImageDeleteRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &PicturebookImageDeleteResponse{}
	result, err := IMAGE_CLIENT.PicturebookImageDelete(picturebookImageDeleteRequest)
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
func TestClient_PicturebookImageSearch(t *testing.T) {
	picturebookImageSearchRequest := &PicturebookImageSearchRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		Tags:     util.PtrString(""),
		TagLogic: util.PtrInt32(int32(0)),
		Pn:       util.PtrInt32(int32(0)),
		Rn:       util.PtrInt32(int32(0)),
	}
	result := &PicturebookImageSearchResponse{}
	result, err := IMAGE_CLIENT.PicturebookImageSearch(picturebookImageSearchRequest)
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
func TestClient_PicturebookImageUpdate(t *testing.T) {
	picturebookImageUpdateRequest := &PicturebookImageUpdateRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
		Brief:    util.PtrString(""),
		Tags:     util.PtrString(""),
	}
	result := &PicturebookImageUpdateResponse{}
	result, err := IMAGE_CLIENT.PicturebookImageUpdate(picturebookImageUpdateRequest)
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
func TestClient_Plant(t *testing.T) {
	plantRequest := &PlantRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		BaikeNum: util.PtrInt32(int32(0)),
	}
	result := &PlantResponse{}
	result, err := IMAGE_CLIENT.Plant(plantRequest)
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
func TestClient_ProductImageAdd(t *testing.T) {
	productImageAddRequest := &ProductImageAddRequest{
		Brief:    util.PtrString(""),
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ClassId1: util.PtrInt32(int32(0)),
		ClassId2: util.PtrInt32(int32(0)),
	}
	result := &ProductImageAddResponse{}
	result, err := IMAGE_CLIENT.ProductImageAdd(productImageAddRequest)
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
func TestClient_ProductImageDelete(t *testing.T) {
	productImageDeleteRequest := &ProductImageDeleteRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &ProductImageDeleteResponse{}
	result, err := IMAGE_CLIENT.ProductImageDelete(productImageDeleteRequest)
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
func TestClient_ProductImageSearch(t *testing.T) {
	productImageSearchRequest := &ProductImageSearchRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ClassId1: util.PtrInt32(int32(0)),
		ClassId2: util.PtrInt32(int32(0)),
		TagLogic: util.PtrInt32(int32(0)),
		Pn:       util.PtrInt32(int32(0)),
		Rn:       util.PtrInt32(int32(0)),
	}
	result := &ProductImageSearchResponse{}
	result, err := IMAGE_CLIENT.ProductImageSearch(productImageSearchRequest)
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
func TestClient_ProductImageUpdate(t *testing.T) {
	productImageUpdateRequest := &ProductImageUpdateRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
		Brief:    util.PtrString(""),
		ClassId1: util.PtrInt32(int32(0)),
		ClassId2: util.PtrInt32(int32(0)),
	}
	result := &ProductImageUpdateResponse{}
	result, err := IMAGE_CLIENT.ProductImageUpdate(productImageUpdateRequest)
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
func TestClient_RemoveMoire(t *testing.T) {
	removeMoireRequest := &RemoveMoireRequest{
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		PdfFile:    util.PtrString(""),
		PdfFileNum: util.PtrInt32(int32(0)),
	}
	result := &RemoveMoireResponse{}
	result, err := IMAGE_CLIENT.RemoveMoire(removeMoireRequest)
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
func TestClient_SameImageAdd(t *testing.T) {
	sameImageAddRequest := &SameImageAddRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
		Brief: util.PtrString(""),
		Tags:  util.PtrString(""),
	}
	result := &SameImageAddResponse{}
	result, err := IMAGE_CLIENT.SameImageAdd(sameImageAddRequest)
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
func TestClient_SameImageDelete(t *testing.T) {
	sameImageDeleteRequest := &SameImageDeleteRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &SameImageDeleteResponse{}
	result, err := IMAGE_CLIENT.SameImageDelete(sameImageDeleteRequest)
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
func TestClient_SameImageSearch(t *testing.T) {
	sameImageSearchRequest := &SameImageSearchRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		Tags:     util.PtrString(""),
		TagLogic: util.PtrInt32(int32(0)),
		Pn:       util.PtrInt32(int32(0)),
		Rn:       util.PtrInt32(int32(0)),
	}
	result := &SameImageSearchResponse{}
	result, err := IMAGE_CLIENT.SameImageSearch(sameImageSearchRequest)
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
func TestClient_SameImageUpdate(t *testing.T) {
	sameImageUpdateRequest := &SameImageUpdateRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
		Brief:    util.PtrString(""),
		Tags:     util.PtrString(""),
	}
	result := &SameImageUpdateResponse{}
	result, err := IMAGE_CLIENT.SameImageUpdate(sameImageUpdateRequest)
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
func TestClient_Segment(t *testing.T) {
	segmentRequest := &SegmentRequest{
		Method:     util.PtrString(""),
		Image:      util.PtrString(""),
		Url:        util.PtrString(""),
		ReturnForm: util.PtrString(""),
		RefineMask: util.PtrBool(false),
		Position:   util.PtrString(""),
	}
	result := &SegmentResponse{}
	result, err := IMAGE_CLIENT.Segment(segmentRequest)
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
func TestClient_SelfieAnime(t *testing.T) {
	selfieAnimeRequest := &SelfieAnimeRequest{
		Image:     util.PtrString(""),
		Url:       util.PtrString(""),
		ImageType: util.PtrString(""),
		MaskId:    util.PtrString(""),
	}
	result := &SelfieAnimeResponse{}
	result, err := IMAGE_CLIENT.SelfieAnime(selfieAnimeRequest)
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
func TestClient_SimilarImageAdd(t *testing.T) {
	similarImageAddRequest := &SimilarImageAddRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
		Brief: util.PtrString(""),
		Tags:  util.PtrString(""),
	}
	result := &SimilarImageAddResponse{}
	result, err := IMAGE_CLIENT.SimilarImageAdd(similarImageAddRequest)
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
func TestClient_SimilarImageDelete(t *testing.T) {
	similarImageDeleteRequest := &SimilarImageDeleteRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
	}
	result := &SimilarImageDeleteResponse{}
	result, err := IMAGE_CLIENT.SimilarImageDelete(similarImageDeleteRequest)
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
func TestClient_SimilarImageSearch(t *testing.T) {
	similarImageSearchRequest := &SimilarImageSearchRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		Tags:     util.PtrString(""),
		TagLogic: util.PtrInt32(int32(0)),
		Pn:       util.PtrInt32(int32(0)),
		Rn:       util.PtrInt32(int32(0)),
	}
	result := &SimilarImageSearchResponse{}
	result, err := IMAGE_CLIENT.SimilarImageSearch(similarImageSearchRequest)
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
func TestClient_SimilarImageUpdate(t *testing.T) {
	similarImageUpdateRequest := &SimilarImageUpdateRequest{
		Image:    util.PtrString(""),
		Url:      util.PtrString(""),
		ContSign: util.PtrString(""),
		Brief:    util.PtrString(""),
		Tags:     util.PtrString(""),
	}
	result := &SimilarImageUpdateResponse{}
	result, err := IMAGE_CLIENT.SimilarImageUpdate(similarImageUpdateRequest)
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
func TestClient_StretchRestore(t *testing.T) {
	stretchRestoreRequest := &StretchRestoreRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
	}
	result := &StretchRestoreResponse{}
	result, err := IMAGE_CLIENT.StretchRestore(stretchRestoreRequest)
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
func TestClient_StyleTrans(t *testing.T) {
	styleTransRequest := &StyleTransRequest{
		Option: util.PtrString(""),
		Image:  util.PtrString(""),
		Url:    util.PtrString(""),
	}
	result := &StyleTransResponse{}
	result, err := IMAGE_CLIENT.StyleTrans(styleTransRequest)
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
func TestClient_VehicleDetect(t *testing.T) {
	vehicleDetectRequest := &VehicleDetectRequest{
		Image: util.PtrString(""),
		Url:   util.PtrString(""),
		Area:  util.PtrString(""),
	}
	result := &VehicleDetectResponse{}
	result, err := IMAGE_CLIENT.VehicleDetect(vehicleDetectRequest)
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
