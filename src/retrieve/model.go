package retrieve

import "time"

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
