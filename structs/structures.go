package structs

type result struct {
	Spellcheck     Spellcheck
	ResponseHeader ResponseHeader
	Response       Response
}

type Spellcheck struct {
	Suggestions map[string]any
	Collations  map[string]any
}

type ResponseHeader struct {
	zkConnected bool
	Status      int
	QTime       int
	Params      Params
	TotalTime   int
}

type Params struct {
	MatchType       string
	FL              map[string]any
	Qop             string `json:"q.op"`
	FQ              map[string]any
	QOrig           string `json:"q.orig"`
	BQ              map[string]any
	SpellCheckQ     string `json:"spellcheck.q"`
	Trace           map[string]any
	DefType         string
	QF              string
	Req_type        string
	WT              string
	MM              string
	LWPipelineID    string `json:"lw.pipelineId"`
	Start           string
	isFusionQuery   string
	Sort            string
	Rows            string
	Version         string
	FusionQueryID   string
	ReqID           string
	isInternal      string
	Q               string
	EnableElevation string
	SpellCheck      string
	ElevateIDs      string
}

type Response struct {
	NumFound int
	Start    int
	MaxScore float32
	Docs     map[string]any
}
