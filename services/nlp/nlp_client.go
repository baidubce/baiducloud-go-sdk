package nlp

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
	"github.com/baidubce/baiducloud-go-sdk/core/http"
	"github.com/baidubce/baiducloud-go-sdk/core/util"
)

const ()

// Address
//
// PARAMS:
//   - request: the arguments to Address
//
// RETURNS:
//   - AddressResponse: The return type of the Address interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Address(request *AddressRequest) (*AddressResponse, error) {
	result := &AddressResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddressUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CommentTag
//
// PARAMS:
//   - request: the arguments to CommentTag
//
// RETURNS:
//   - CommentTagResponse: The return type of the CommentTag interface.
//   - error: nil if success otherwise the specific error
func (c *Client) CommentTag(request *CommentTagRequest) (*CommentTagResponse, error) {
	result := &CommentTagResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCommentTagUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Ecnet
//
// PARAMS:
//   - request: the arguments to Ecnet
//
// RETURNS:
//   - EcnetResponse: The return type of the Ecnet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Ecnet(request *EcnetRequest) (*EcnetResponse, error) {
	result := &EcnetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEcnetUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Emotion
//
// PARAMS:
//   - request: the arguments to Emotion
//
// RETURNS:
//   - EmotionResponse: The return type of the Emotion interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Emotion(request *EmotionRequest) (*EmotionResponse, error) {
	result := &EmotionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEmotionUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EntityAnalysis
//
// PARAMS:
//   - request: the arguments to EntityAnalysis
//
// RETURNS:
//   - EntityAnalysisResponse: The return type of the EntityAnalysis interface.
//   - error: nil if success otherwise the specific error
func (c *Client) EntityAnalysis(request *EntityAnalysisRequest) (*EntityAnalysisResponse, error) {
	result := &EntityAnalysisResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEntityAnalysisUri()).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Keyword
//
// PARAMS:
//   - request: the arguments to Keyword
//
// RETURNS:
//   - KeywordResponse: The return type of the Keyword interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Keyword(request *KeywordRequest) (*KeywordResponse, error) {
	result := &KeywordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getKeywordUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Lexer
//
// PARAMS:
//   - request: the arguments to Lexer
//
// RETURNS:
//   - LexerResponse: The return type of the Lexer interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Lexer(request *LexerRequest) (*LexerResponse, error) {
	result := &LexerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getLexerUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// NewsSummary
//
// PARAMS:
//   - request: the arguments to NewsSummary
//
// RETURNS:
//   - NewsSummaryResponse: The return type of the NewsSummary interface.
//   - error: nil if success otherwise the specific error
func (c *Client) NewsSummary(request *NewsSummaryRequest) (*NewsSummaryResponse, error) {
	result := &NewsSummaryResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNewsSummaryUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SentimentClassify
//
// PARAMS:
//   - request: the arguments to SentimentClassify
//
// RETURNS:
//   - SentimentClassifyResponse: The return type of the SentimentClassify interface.
//   - error: nil if success otherwise the specific error
func (c *Client) SentimentClassify(request *SentimentClassifyRequest) (*SentimentClassifyResponse, error) {
	result := &SentimentClassifyResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSentimentClassifyUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Simnet
//
// PARAMS:
//   - request: the arguments to Simnet
//
// RETURNS:
//   - SimnetResponse: The return type of the Simnet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Simnet(request *SimnetRequest) (*SimnetResponse, error) {
	result := &SimnetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSimnetUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TextCorrection
//
// PARAMS:
//   - request: the arguments to TextCorrection
//
// RETURNS:
//   - TextCorrectionResponse: The return type of the TextCorrection interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TextCorrection(request *TextCorrectionRequest) (*TextCorrectionResponse, error) {
	result := &TextCorrectionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTextCorrectionUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Topic
//
// PARAMS:
//   - request: the arguments to Topic
//
// RETURNS:
//   - TopicResponse: The return type of the Topic interface.
//   - error: nil if success otherwise the specific error
func (c *Client) Topic(request *TopicRequest) (*TopicResponse, error) {
	result := &TopicResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTopicUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TxtKeywordsExtraction
//
// PARAMS:
//   - request: the arguments to TxtKeywordsExtraction
//
// RETURNS:
//   - TxtKeywordsExtractionResponse: The return type of the TxtKeywordsExtraction interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TxtKeywordsExtraction(request *TxtKeywordsExtractionRequest) (*TxtKeywordsExtractionResponse, error) {
	result := &TxtKeywordsExtractionResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTxtKeywordsExtractionUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TxtMonet
//
// PARAMS:
//   - request: the arguments to TxtMonet
//
// RETURNS:
//   - TxtMonetResponse: The return type of the TxtMonet interface.
//   - error: nil if success otherwise the specific error
func (c *Client) TxtMonet(request *TxtMonetRequest) (*TxtMonetResponse, error) {
	result := &TxtMonetResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTxtMonetUri()).
		WithQueryParamFilter("charset", util.StringValue(request.Charset)).
		WithBody(request).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
