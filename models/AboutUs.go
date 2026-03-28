package models

type Image struct {
	Src    string `json:"src" bson:"src"`
	Alt    string `json:"alt" bson:"alt"`
	Width  int    `json:"width,omitempty" bson:"width,omitempty"`
	Height int    `json:"height,omitempty" bson:"height,omitempty"`
}

type Avatar struct {
	Src string `json:"src" bson:"src"`
	Alt string `json:"alt" bson:"alt"`
}

type AboutSection struct {
	Image                  Image    `json:"image" bson:"image"`
	ReviewAvatars          []Avatar `json:"reviewAvatars" bson:"reviewAvatars"`
	ReviewText             string   `json:"reviewText" bson:"reviewText"`
	Title                  string   `json:"title" bson:"title"`
	IntroPrefix            string   `json:"introPrefix" bson:"introPrefix"`
	IntroHighlight         string   `json:"introHighlight" bson:"introHighlight"`
	IntroSuffix            string   `json:"introSuffix" bson:"introSuffix"`
	Paragraphs             []string `json:"paragraphs" bson:"paragraphs"`
	PrimaryHighlightText   string   `json:"primaryHighlightText" bson:"primaryHighlightText"`
	ExecutionHighlightText string   `json:"executionHighlightText" bson:"executionHighlightText"`
}

type Counter struct {
	ID          string `json:"id" bson:"id"`
	Count       int    `json:"count" bson:"count"`
	Suffix      string `json:"suffix" bson:"suffix"`
	Label       string `json:"label" bson:"label"`
	Description string `json:"description" bson:"description"`
}

type VideoSection struct {
	Style2               bool      `json:"style2" bson:"style2"`
	PopupBackgroundImage string    `json:"popupBackgroundImage" bson:"popupBackgroundImage"`
	PopupVideoUrl        string    `json:"popupVideoUrl" bson:"popupVideoUrl"`
	PlayIcon             string    `json:"playIcon" bson:"playIcon"`
	Counters             []Counter `json:"counters" bson:"counters"`
}

type WorkProcessStep struct {
	Number      string `json:"number" bson:"number"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

type WorkprocessSection struct {
	Title       string            `json:"title" bson:"title"`
	Description string            `json:"description" bson:"description"`
	Tagline     string            `json:"tagline" bson:"tagline"`
	Steps       []WorkProcessStep `json:"steps" bson:"steps"`
}

type TeamMember struct {
	ID          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Designation string `json:"designation" bson:"designation"`
	Img         string `json:"img" bson:"img"`
	ProfileHref string `json:"profileHref" bson:"profileHref"`
}

type TeamSection struct {
	Title        string       `json:"title" bson:"title"`
	ViewAllLabel string       `json:"viewAllLabel" bson:"viewAllLabel"`
	ViewAllHref  string       `json:"viewAllHref" bson:"viewAllHref"`
	ViewAllIcon  string       `json:"viewAllIcon" bson:"viewAllIcon"`
	Members      []TeamMember `json:"members" bson:"members"`
}

type CTAButton struct {
	Label string `json:"label" bson:"label"`
	Href  string `json:"href" bson:"href"`
	Icon  string `json:"icon,omitempty" bson:"icon,omitempty"`
}

type CTASection struct {
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Button      CTAButton `json:"button" bson:"button"`
}

type AboutPage struct {
	Breadcrumb         Breadcrumb         `json:"breadcrumb" bson:"breadcrumb"`
	AboutSection       AboutSection       `json:"aboutSection" bson:"aboutSection"`
	VideoSection       VideoSection       `json:"videoSection" bson:"videoSection"`
	WorkprocessSection WorkprocessSection `json:"workprocessSection" bson:"workprocessSection"`
	TeamSection        TeamSection        `json:"teamSection" bson:"teamSection"`
	CTA                CTASection         `json:"cta" bson:"cta"`
}
