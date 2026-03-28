package models

type LogoHome struct {
	Src    string `json:"src"`
	Alt    string `json:"alt"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type SocialLinkHome struct {
	Label string `json:"label"`
	Href  string `json:"href"`
	Icon  string `json:"icon"`
}

type CTAHome struct {
	Label string `json:"label"`
	Href  string `json:"href"`
	Icon  string `json:"icon,omitempty"`
}

type SiteHome struct {
	BrandName string `json:"brandName"`
}

type SubMenuHome struct {
	Link  string `json:"link"`
	Title string `json:"title"`
}

type NavItemHome struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Link        string        `json:"link"`
	HasDropdown bool          `json:"has_dropdown"`
	SubMenus    []SubMenuHome `json:"sub_menus,omitempty"`
}

type HeaderHome struct {
	Logo                  LogoHome `json:"logo"`
	MobileToggleAriaLabel string   `json:"mobileToggleAriaLabel"`
	MobileToggleIcon      string   `json:"mobileToggleIcon"`
	SideMenuButtonIcon    string   `json:"sideMenuButtonIcon"`
}

type ContactItemHome struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Href  string `json:"href,omitempty"`
	Icon  string `json:"icon"`
}

type OffcanvasHome struct {
	Logo           LogoHome          `json:"logo"`
	Title          string            `json:"title"`
	Description    string            `json:"description"`
	ContactItems   []ContactItemHome `json:"contactItems"`
	SocialHeading  string            `json:"socialHeading"`
	CloseAriaLabel string            `json:"closeAriaLabel"`
	CTA            CTAHome           `json:"cta"`
}

type HeroHeadlineHome struct {
	Top             string `json:"top"`
	MiddlePrefix    string `json:"middlePrefix"`
	MiddleHighlight string `json:"middleHighlight"`
	Bottom          string `json:"bottom"`
}

type HeroImageHome struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type HeroHome struct {
	Headline    HeroHeadlineHome `json:"headline"`
	Description string           `json:"description"`
	Image       HeroImageHome    `json:"image"`
	CTA         CTAHome          `json:"cta"`
}

type FeatureHome struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}

type SocialProofHome struct {
	SectionTitle string        `json:"sectionTitle"`
	Eyebrow      string        `json:"eyebrow"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Quote        string        `json:"quote"`
	Features     []FeatureHome `json:"features"`
}

type ServiceItemHome struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImgSrc      string   `json:"imgSrc"`
	Features    []string `json:"features"`
	HookText    string   `json:"hookText"`
}

type ServicesSectionHome struct {
	SectionTitle string            `json:"sectionTitle"`
	Items        []ServiceItemHome `json:"items"`
}

type WhyItemHome struct {
	Number      string `json:"number"`
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BannerHome struct {
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Highlight string `json:"highlight"`
	Suffix    string `json:"suffix"`
}

type WhyStratixa struct {
	SectionTitle string        `json:"sectionTitle"`
	Items        []WhyItemHome `json:"items"`
	Banner       BannerHome    `json:"banner"`
}

type WhoWeWorkWithItemHome struct {
	BgImage string `json:"bgImage"`
	Icon    string `json:"icon"`
	Label   string `json:"label"`
}

type WhoWeWorkWithHome struct {
	SectionTitle string                  `json:"sectionTitle"`
	Items        []WhoWeWorkWithItemHome `json:"items"`
}

type ProcessStepHome struct {
	ID    int    `json:"id"`
	Icon  string `json:"icon"`
	Title string `json:"title"`
}
type ProcessHome struct {
	SectionTitle    string            `json:"sectionTitle"`
	StepLabelPrefix string            `json:"stepLabelPrefix"`
	Steps           []ProcessStepHome `json:"steps"`
}
type ClientItemHome struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Industry    string   `json:"industry"`
	Services    []string `json:"services"`
	Icon        string   `json:"icon"`
	Description string   `json:"description"`
}

type PartnershipSegmentHome struct {
	Text     string `json:"text"`
	Emphasis string `json:"emphasis,omitempty"`
}

type PartnershipHome struct {
	Icon     string                   `json:"icon"`
	Title    string                   `json:"title"`
	Segments []PartnershipSegmentHome `json:"segments"`
}
type ClienteleHome struct {
	SectionTitle      string           `json:"sectionTitle"`
	Intro             string           `json:"intro"`
	DeliverablesLabel string           `json:"deliverablesLabel"`
	Items             []ClientItemHome `json:"items"`
	Partnership       PartnershipHome  `json:"partnership"`
}

type ResultItemHome struct {
	Icon        string `json:"icon"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type ResultsDrivenHome struct {
	SectionTitle string           `json:"sectionTitle"`
	Items        []ResultItemHome `json:"items"`
}

type CategoryMarqueeHome struct {
	Items []string `json:"items"`
}

type MissionImageHome struct {
	Src    string `json:"src"`
	Alt    string `json:"alt"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type ExperienceHome struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type AccordionItemHome struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DefaultOpen bool   `json:"defaultOpen"`
}

type MissionHome struct {
	Image          MissionImageHome    `json:"image"`
	Experience     ExperienceHome      `json:"experience"`
	Title          string              `json:"title"`
	Description    string              `json:"description"`
	AccordionItems []AccordionItemHome `json:"accordionItems"`
}

type CTASectionHome struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Button      CTAHome `json:"button"`
}

type LegalLinkHome struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}

type FooterHome struct {
	Logo             LogoHome        `json:"logo"`
	Description      string          `json:"description"`
	ServicesHeading  string          `json:"servicesHeading"`
	ResourcesHeading string          `json:"resourcesHeading"`
	SocialHeading    string          `json:"socialHeading"`
	CopyrightPrefix  string          `json:"copyrightPrefix"`
	LegalLinks       []LegalLinkHome `json:"legalLinks"`
}

type Home struct {
	Site            SiteHome            `json:"site"`
	Navigation      []NavItemHome       `json:"navigation"`
	SocialLinks     []SocialLinkHome    `json:"socialLinks"`
	Header          HeaderHome          `json:"header"`
	Offcanvas       OffcanvasHome       `json:"offcanvas"`
	Hero            HeroHome            `json:"hero"`
	SocialProof     SocialProofHome     `json:"socialProof"`
	ServicesSection ServicesSectionHome `json:"servicesSection"`
	WhyStratixa     WhyStratixa         `json:"whyStratixa"`
	WhoWeWorkWith   WhoWeWorkWithHome   `json:"whoWeWorkWith"`
	Process         ProcessHome         `json:"process"`
	Clientele       ClienteleHome       `json:"clientele"`
	ResultsDriven   ResultsDrivenHome   `json:"resultsDriven"`
	CategoryMarquee CategoryMarqueeHome `json:"categoryMarquee"`
	Mission         MissionHome         `json:"mission"`
	CTA             CTASectionHome      `json:"cta"`
	Footer          FooterHome          `json:"footer"`
}
