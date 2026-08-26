package bcm

import "github.com/baidubce/baiducloud-go-sdk/bce"

type DescribeMetricCatalogsResponse struct {
	bce.BaseResponse
	Success                            *bool            `json:"success,omitempty"`
	Code                               *string          `json:"code,omitempty"`
	Message                            *string          `json:"message,omitempty"`
	Catalogs                           []*MetricCatalog `json:"catalogs,omitempty"`
	CatalogsName                       *string          `json:"catalogs[].name,omitempty"`
	CatalogsLabel                      *string          `json:"catalogs[].label,omitempty"`
	CatalogsCatalogs                   []*MetricCatalog `json:"catalogs[].catalogs,omitempty"`
	CatalogsMetrics                    []*Metric        `json:"catalogs[].metrics,omitempty"`
	CatalogsMetricsName                *string          `json:"catalogs[].metrics[].name,omitempty"`
	CatalogsMetricsLabel               *string          `json:"catalogs[].metrics[].label,omitempty"`
	CatalogsMetricsResourceIdentifiers []*string        `json:"catalogs[].metrics[].resourceIdentifiers,omitempty"`
	CatalogsMetricsMetricDimensions    []*string        `json:"catalogs[].metrics[].metricDimensions,omitempty"`
	CatalogsMetricsPeriod              *float64         `json:"catalogs[].metrics[].period,omitempty"`
	CatalogsMetricsPeriodUnit          *string          `json:"catalogs[].metrics[].periodUnit,omitempty"`
	CatalogsMetricsUnit                *string          `json:"catalogs[].metrics[].unit,omitempty"`
}
