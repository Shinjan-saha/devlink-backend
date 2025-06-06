package models

import "gorm.io/gorm"

type Resource struct {
    gorm.Model
    Title       string
    URL         string
    Description string
    Type        string // article, video, tool
    Tags        string
    SubmittedBy uint
    Approved    bool
}
