package structs

type Result struct {
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
	Docs     []KCS
}

type KCS struct {
	DocumentKind          string   `json:"documentKind"`
	URI                   string   `json:"uri"`
	ResourceURI           string   `json:"resource_uri"`
	Path                  string   `json:"path"`
	ViewURI               string   `json:"view_uri"`
	Title                 string   `json:"title"`
	AllTitle              string   `json:"allTitle"`
	SortTitle             string   `json:"sortTitle"`
	ID                    string   `json:"id"`
	SolutionID            string   `json:"solution.id"`
	AuthorSSOName         string   `json:"authorSSOName"`
	LastModifiedBySSOName string   `json:"lastModifiedBySSOName"`
	HasPublishedRevision  string   `json:"hasPublishedRevision"`
	Body                  []string `json:"body"`
	Issue                 []string `json:"issue"`
	IssueTag              []string `json:"issueTag"`
	LastModifiedDate      string   `json:"lastModifiedDate"`
	DisplayDate           string   `json:"displayDate"`
	CreatedDate           string   `json:"createdDate"`
	KCSState              string   `json:"kcsState"`
	ModerationState       string   `json:"ModerationState"`
	AccessState           string   `json:"accessState"`
	SetLanguage           string   `json:"setLanguage"`
	Abstract              string   `json:"abstract"`
	DetectedLanguage      []string `json:"detectedLanguage"`
	Language              string   `json:"language"`
	SBRs                  []string `json:"sbr"`
	Tags                  []string `json:"tag"`
	Product               []string `json:"product"`
	PublishedTitle        string   `json:"publishedTitle"`
	PublishedAbstract     string   `json:"publishedAbstract"`
	InternalTags          []string `json:"internalTags"`
	CaseCount             int      `json:"caseCount"`
	CaseCount365          int      `json:"caseCount_365"`
	KCSRateUp             int      `json:"kcsRateUp"`
	KCSRateDown           int      `json:"kcsRateDown"`
	Category              []string `json:"category"`
	PublicationState      string   `json:"publication_state"`
	RequiresSubscription  bool     `json:"requires_subscription"`
	ArticleType           string   `json:"article_type"`
	BoostProduct          string   `json:"boostProduct"`
	DetectedProducts      []string `json:"detectedProducts"`
	DuplicateResourceID   string   `json:"duplicateResourceId"`
	InferredTag           []string `json:"inferred_tag"`
	Timestamp             string   `json:"timestamp"`
	Component             []string `json:"component"`
	Version               int      `json:"_version_"`
	Elevated              bool     `json:"[elevated]"`
}
