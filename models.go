package poe_api

type CreateBot struct {
	Handle            string
	Prompt            string
	DisplayName       string
	BaseModel         string
	Description       string
	IntroMessage      string
	ApiKey            *string
	ApiBot            bool
	ApiUrl            *string
	PromptPublic      *bool
	PfpUrl            *string
	Linkification     bool
	MarkdownRendering *bool
	SuggestedReplies  bool
	Private           bool
	Temperature       string
}

type Point interface {
	string | bool
}

func GetPoint[T Point](s T) *T {
	return &s
}
