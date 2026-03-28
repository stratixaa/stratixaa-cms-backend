package models

import "time"

type PageContent struct {
	Breadcrumb         Breadcrumb `json:"breadcrumb" bson:"breadcrumb"`
	HeroEyebrow        string     `json:"heroEyebrow" bson:"heroEyebrow"`
	SystemLabel        string     `json:"systemLabel" bson:"systemLabel"`
	PricingLabel       string     `json:"pricingLabel" bson:"pricingLabel"`
	PricingSubtext     string     `json:"pricingSubtext" bson:"pricingSubtext"`
	PricingTiers       []string   `json:"pricingTiers" bson:"pricingTiers"`
	FeaturedRibbon     string     `json:"featuredRibbon" bson:"featuredRibbon"`
	PricingNoteIcon    string     `json:"pricingNoteIcon" bson:"pricingNoteIcon"`
	FAQHeading         string     `json:"faqHeading" bson:"faqHeading"`
	ServiceListHeading string     `json:"serviceListHeading" bson:"serviceListHeading"`
	SidebarLinkIcon    string     `json:"sidebarLinkIcon" bson:"sidebarLinkIcon"`
}

type Breadcrumb struct {
	Title    string `json:"title" bson:"title"`
	Subtitle string `json:"subtitle" bson:"subtitle"`
	BgImage  string `json:"bgImage,omitempty" bson:"bgImage,omitempty"`
}

type Offer struct {
	Title string   `json:"title" bson:"title"`
	Items []string `json:"items" bson:"items"`
}

type Process struct {
	Title string   `json:"title" bson:"title"`
	Items []string `json:"items" bson:"items"`
}

type SystemFollow struct {
	Title string   `json:"title" bson:"title"`
	Items []string `json:"items" bson:"items"`
}

type PackageItem struct {
	Name  string `json:"name" bson:"name"`
	Price string `json:"price" bson:"price"`
}

type Packages struct {
	Title string        `json:"title" bson:"title"`
	Items []PackageItem `json:"items" bson:"items"`
	Note  string        `json:"note,omitempty" bson:"note,omitempty"`
}

type FAQItem struct {
	Question string `json:"question" bson:"question"`
	Answer   string `json:"answer" bson:"answer"`
}

type Service struct {
	Slug            string       `json:"slug" bson:"slug"`
	Title           string       `json:"title" bson:"title"`
	H3Title         string       `json:"h3Title" bson:"h3Title"`
	Description     string       `json:"description" bson:"description"`
	Offer           Offer        `json:"offer" bson:"offer"`
	Process         Process      `json:"process" bson:"process"`
	System          SystemFollow `json:"system" bson:"system"`
	Packages        Packages     `json:"packages" bson:"packages"`
	FAQs            []FAQItem    `json:"faqs" bson:"faqs"`
	MetaTitle       string       `json:"metaTitle,omitempty" bson:"metaTitle,omitempty"`
	MetaDescription string       `json:"metaDescription,omitempty" bson:"metaDescription,omitempty"`
}

type ServicePage struct {
	ID          string      `json:"id,omitempty" bson:"_id,omitempty"`
	PageContent PageContent `json:"pageContent" bson:"pageContent"`
	Services    []Service   `json:"services" bson:"services"`
	CreatedAt   time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
