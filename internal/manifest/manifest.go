package manifest

import "time"

//FileInfo не подходит т.к. FileInfo клиентский, когда как Manifest работает с сетью
type File struct {
	Path   string `json:"path"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

type Manifest struct {
	Version   int       `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	Files     []File    `json:"files"`
}
