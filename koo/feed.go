package koo

import (
	"net/http"

	"github.com/dghubble/sling"
)

type FeedItem struct {
	Id                         string  `json:"id"`
	CreatorId                  string  `json:"creatorId"`
	Title                      string  `json:"title"`
	CreatedAt                  int     `json:"createdAt"`
	KuUrlDB                    string  `json:"kuUrlDB"`
	Handle                     string  `json:"handle"`
	ProfileImage               string  `json:"profileImage"`
	Name                       string  `json:"name"`
	ContentType                int     `json:"contentType"`
	Score                      float32 `json:"score"`
	Type                       int     `json:"type"`
	IsHidden                   bool    `json:"isHidden"`
	UserScore                  float32 `json:"userScore"`
	IsRekoo                    bool    `json:"isRekoo"`
	CommentPreference          int     `json:"commentPreference"`
	UserKooCount               int     `json:"userKooCount"`
	IsExclusive                bool    `json:"isExclusive"`
	Status                     int     `json:"status"`
	ViewStatus                 int     `json:"viewStatus"`
	DeletionReason             string  `json:"deletionReason"`
	IsPinnedKoo                bool    `json:"isPinnedKoo"`
	IsKlip                     bool    `json:"isKlip"`
	ThreadId                   string  `json:"threadId"`
	ThreadIndex                float32 `json:"threadIndex"`
	ThreadSize                 int     `json:"threadSize"`
	IsSpam                     bool    `json:"isSpam"`
	RankerUuid                 string  `json:"rankerUuid"`
	IdentifiedLangCode         string  `json:"identifiedLangCode"`
	ReactionScore              int     `json:"reactionScore"`
	IsFollowDisabled           bool    `json:"isFollowDisabled"`
	KuAllowed                  bool    `json:"kuAllowed"`
	IsFollowing                bool    `json:"isFollowing"`
	IsCommentDisabled          bool    `json:"isCommentDisabled"`
	IsLiked                    bool    `json:"isLiked"`
	IsBlocked                  bool    `json:"isBlocked"`
	IsReported                 bool    `json:"isReported"`
	AllowUnfollowForBulkFollow bool    `json:"allowUnfollowForBulkFollow"`
	KuUrl                      string  `json:"kuUrl"`
	KooEncodedId               string  `json:"kooEncodedId"`
	Badge                      string  `json:"badge"`
	IsFavourite                bool    `json:"isFavourite"`
	CanUnfollow                bool    `json:"canUnfollow"`
	NKus                       int     `json:"nKus"`
	NLikes                     int     `json:"nLikes"`
	NMediaImpressions          int     `json:"nMediaImpressions"`
	NImpressions               int     `json:"nImpressions"`
	NReKoos                    int     `json:"nreKoos"`
}

type FeedResponse struct {
	Id           string     `json:"id"`
	Index        int        `json:"index"`
	SectionType  string     `json:"sectionType"`
	UiItemType   string     `json:"uiItemType"`
	SectionTitle string     `json:"sectionTitle"`
	Items        []FeedItem `json:"items"`
}

type Feed struct {
	Feed []FeedResponse `json:"feed"`
}

type FeedService struct {
	sling *sling.Sling
}

type FeedParams struct {
	Limit              int    `url:"limit"`
	Offset             int    `url:"offset,omitempty"`
	ShowPoll           bool   `url:"showPoll,omitempty"`
	FeedType           string `url:"feedType,omitempty"`
	IsBulkFollow       bool   `url:"isBulkFollow,omitempty"`
	MixedReactions     bool   `url:"mixedReactions"`
	ShowBanner         bool   `url:"showBanner"`
	ShowCeleb          bool   `url:"showCeleb"`
	ShowSampleFeed     bool   `url:"showSampleFeed"`
	ShowTimeline       bool   `url:"showTimeLine"`
	ShowHashtag        bool   `url:"showHashTag"`
	ShowGif            bool   `url:"showGif"`
	ShowTrendingKoo    bool   `url:"showTrendingKoo"`
	ShowInActiveKooers bool   `url:"showInActiveKooers"`
	ShowThreads        bool   `url:"showThreads"`
}

func newFeedService(sling *sling.Sling) *FeedService {
	return &FeedService{
		sling: sling.Path("/consumer_feed"),
	}
}

func (s *FeedService) Fetch(params *FeedParams) (*FeedResponse, *http.Response, error) {
	if params == nil {
		params = &FeedParams{}
	}

	feedRes := &FeedResponse{}
	feedError := &APIError{}
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(feedRes, feedError)

	return feedRes, resp, err
}
