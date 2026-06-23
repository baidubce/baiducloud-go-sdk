package cprom

import "github.com/baidubce/baiducloud-go-sdk/bce"

type CreateCustomScrapeTaskResponse struct {
	bce.BaseResponse
	ScrapeJobId *string `json:"scrapeJobId,omitempty"`
}
