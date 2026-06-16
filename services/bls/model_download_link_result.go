package bls

type DownloadLinkResult struct {
	FileDir  *string `json:"fileDir,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	Link     *string `json:"link,omitempty"`
}
