package models

import (
	"time"
)

type ListingPage struct {
	Breadcrumb      Breadcrumb `json:"breadcrumb" bson:"breadcrumb"`
	BlogsPerPage    int        `json:"blogsPerPage" bson:"blogsPerPage"`
	ReadMoreLabel   string     `json:"readMoreLabel" bson:"readMoreLabel"`
	Widgets         Widgets    `json:"widgets" bson:"widgets"`
	PaginationIcons Icons      `json:"paginationIcons" bson:"paginationIcons"`
}

type DetailsPage struct {
	Breadcrumb    Breadcrumb `json:"breadcrumb" bson:"breadcrumb"`
	Widgets       Widgets    `json:"widgets" bson:"widgets"`
	TagsLabel     string     `json:"tagsLabel" bson:"tagsLabel"`
	PreviousLabel string     `json:"previousPostLabel" bson:"previousPostLabel"`
	NextLabel     string     `json:"nextPostLabel" bson:"nextPostLabel"`
	PagerIcons    Icons      `json:"pagerIcons" bson:"pagerIcons"`
	QuoteIcon     string     `json:"quoteIcon" bson:"quoteIcon"`
}

type Widgets struct {
	SearchTitle       string `json:"searchTitle" bson:"searchTitle"`
	SearchPlaceholder string `json:"searchPlaceholder" bson:"searchPlaceholder"`
	SearchIcon        string `json:"searchIcon" bson:"searchIcon"`
	CategoriesTitle   string `json:"categoriesTitle" bson:"categoriesTitle"`
	RecentPostTitle   string `json:"recentPostTitle" bson:"recentPostTitle"`
	TagCloudTitle     string `json:"tagCloudTitle" bson:"tagCloudTitle"`
}

type Icons struct {
	Previous string `json:"previous" bson:"previous"`
	Next     string `json:"next" bson:"next"`
}

type Category struct {
	Name  string `json:"name" bson:"name"`
	Slug  string `json:"slug" bson:"slug"`
	Count int    `json:"count" bson:"count"`
}

type BlogSection struct {
	Introduction string `json:"introduction" bson:"introduction"`
	Sections     []struct {
		Heading string `json:"heading" bson:"heading"`
		Content string `json:"content" bson:"content"`
	} `json:"sections" bson:"sections"`
	Quote      string `json:"quote" bson:"quote"`
	Conclusion string `json:"conclusion" bson:"conclusion"`
}

type BlogPost struct {
	ID          string      `json:"id" bson:"_id,omitempty"`
	Slug        string      `json:"slug" bson:"slug"`
	Title       string      `json:"title" bson:"title"`
	ShortTitle  string      `json:"shortTitle" bson:"shortTitle"`
	Excerpt     string      `json:"excerpt" bson:"excerpt"`
	Content     string      `json:"content" bson:"content"`
	Category    string      `json:"category" bson:"category"`
	Date        string      `json:"date" bson:"date"`
	Author      string      `json:"author" bson:"author"`
	Img         string      `json:"img" bson:"img"`
	Tags        []string    `json:"tags" bson:"tags"`
	FullContent BlogSection `json:"fullContent" bson:"fullContent"`
	Status      string      `json:"status" bson:"status"`
	CreatedAt   time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type BlogPage struct {
	ID          string      `json:"id,omitempty" bson:"_id,omitempty"`
	ListingPage ListingPage `json:"listingPage" bson:"listingPage"`
	DetailsPage DetailsPage `json:"detailsPage" bson:"detailsPage"`
	Categories  []Category  `json:"categories" bson:"categories"`
	BlogPosts   []BlogPost  `json:"blogPosts" bson:"blogPosts"`
	CreatedAt   time.Time   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (b BlogPage) TableName() string {
	return "blogs"
}
