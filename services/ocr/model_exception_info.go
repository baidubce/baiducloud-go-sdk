package ocr

type ExceptionInfo struct {
	Addreason            *string `json:"addreason,omitempty"`
	Adddate              *string `json:"adddate,omitempty"`
	Removereason         *string `json:"removereason,omitempty"`
	Removedate           *string `json:"removedate,omitempty"`
	Decisionoffice       *string `json:"decisionoffice,omitempty"`
	Removedecisionoffice *string `json:"removedecisionoffice,omitempty"`
}
