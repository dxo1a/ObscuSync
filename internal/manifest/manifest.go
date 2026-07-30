package manifest

import "time"

// FileInfo is not suitable because it is client-side when Manifest works with the network
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
