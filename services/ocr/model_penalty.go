package ocr

type Penalty struct {
	Docno       *string `json:"docno,omitempty"`
	Penaltytype *string `json:"penaltytype,omitempty"`
	Officename  *string `json:"officename,omitempty"`
	Content     *string `json:"content,omitempty"`
	Penaltydate *string `json:"penaltydate,omitempty"`
	Publicdate  *string `json:"publicdate,omitempty"`
	Remark      *string `json:"remark,omitempty"`
}
