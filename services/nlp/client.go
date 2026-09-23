package nlp

import (
	"github.com/baidubce/baiducloud-go-sdk/bce"
)

const (
	DEFAULT_ENDPOINT = "nlp." + bce.DEFAULT_REGION + ".baidubce.com"

	CONSTANT_RPC = "rpc"

	CONSTANT_2_0 = "2.0"

	CONSTANT_NLP = "nlp"

	CONSTANT_V2 = "v2"

	CONSTANT_TEXT_CORRECTION = "text_correction"

	CONSTANT_V1 = "v1"

	CONSTANT_ECNET = "ecnet"

	CONSTANT_ADDRESS = "address"

	CONSTANT_NEWS_SUMMARY = "news_summary"

	CONSTANT_COMMENT_TAG = "comment_tag"

	CONSTANT_LEXER = "lexer"

	CONSTANT_EMOTION = "emotion"

	CONSTANT_TXT_KEYWORDS_EXTRACTION = "txt_keywords_extraction"

	CONSTANT_TOPIC = "topic"

	CONSTANT_TXT_MONET = "txt_monet"

	CONSTANT_ENTITY_ANALYSIS = "entity_analysis"

	CONSTANT_SENTIMENT_CLASSIFY = "sentiment_classify"

	CONSTANT_KEYWORD = "keyword"

	CONSTANT_SIMNET = "simnet"
)

// Client of nlp service is a kind of BceClient, so derived from BceClient
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

func getAddressUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ADDRESS
}
func getCommentTagUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_COMMENT_TAG
}
func getEcnetUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ECNET
}
func getEmotionUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_EMOTION
}
func getEntityAnalysisUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_ENTITY_ANALYSIS
}
func getKeywordUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_KEYWORD
}
func getLexerUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_LEXER
}
func getNewsSummaryUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_NEWS_SUMMARY
}
func getSentimentClassifyUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_SENTIMENT_CLASSIFY
}
func getSimnetUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_SIMNET
}
func getTextCorrectionUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V2 + bce.URI_PREFIX + CONSTANT_TEXT_CORRECTION
}
func getTopicUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TOPIC
}
func getTxtKeywordsExtractionUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TXT_KEYWORDS_EXTRACTION
}
func getTxtMonetUri() string {
	return bce.URI_PREFIX + CONSTANT_RPC + bce.URI_PREFIX + CONSTANT_2_0 + bce.URI_PREFIX + CONSTANT_NLP + bce.URI_PREFIX + CONSTANT_V1 + bce.URI_PREFIX + CONSTANT_TXT_MONET
}
