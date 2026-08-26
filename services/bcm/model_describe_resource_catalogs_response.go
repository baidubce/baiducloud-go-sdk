package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeResourceCatalogsResponse struct {
	bce.BaseResponse
	Success                            *bool                  `json:"success,omitempty"`
	Code                               *string                `json:"code,omitempty"`
	Message                            *string                `json:"message,omitempty"`
	Catalogs                           []*ResourceCatalog     `json:"catalogs,omitempty"`
	CatalogsScope                      *string                `json:"catalogs[].scope,omitempty"`
	CatalogsScopeLabel                 *string                `json:"catalogs[].scopeLabel,omitempty"`
	CatalogsResources                  []*ResourceCatalogItem `json:"catalogs[].resources,omitempty"`
	CatalogsResourcesResourceType      *string                `json:"catalogs[].resources[].resourceType,omitempty"`
	CatalogsResourcesResourceTypeLabel *string                `json:"catalogs[].resources[].resourceTypeLabel,omitempty"`
	CatalogsRegions                    []*string              `json:"catalogs[].regions,omitempty"`
}
