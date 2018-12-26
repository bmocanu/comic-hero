package model

import "time"

type ComicDef struct {
    Id          int
    Name        string
    Description string
    Author      string
    Url         string
}

// An Issue represents one image of a certain comic, from a certain day.
//
// Retrievers are used to fetch the issue for the current date. Each issue contains the time when the issue was
// retrieved, the URL where the image can be found and the title (optional)
type Issue struct {
    Comic string
    Time  time.Time
    Url   string
    Title string
}

// An IssueLink is one link from a linked list, the structure used by the Store part of comic-hero for
// storing the issues gathered for a particular comic
type IssueLink struct {
    Issue      *Issue
    Hash       uint32
    NextLink   *IssueLink
    IssueCount int
}
