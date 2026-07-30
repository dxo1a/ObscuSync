package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ProgressFunc вызывается во время скачивания (written — уже записано, total — Content-Length или 0).
type ProgressFunc func(written, total int64)

// DownloadFile качает файл с сервера и сохраняет в destPath.
// relPath — путь как в манифесте (mods/foo.zip), со слешами.
// onProgress может быть nil.
func (c *Client) DownloadFile(profile, relPath, destPath string, onProgress ProgressFunc) error {
	if profile == "" || relPath == "" {
		return fmt.Errorf("profile and path are required")
	}

	// URL-path: каждый сегмент кодируем, слеши оставляем
	escaped := escapePath(relPath)
	reqURL := c.url("/api/v1/files/" + profile + "/" + escaped)

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("download %s: %w", relPath, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errStatus(resp, "download "+relPath)
	}

	total := resp.ContentLength

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// we write to a temporary file next to it, then rename it because it's more atomic
	tmpPath := destPath + ".tmp"
	out, err := os.Create(tmpPath)
	if err != nil {
		return err
	}

	var reader io.Reader = resp.Body
	if onProgress != nil {
		reader = &progressReader{
			r:          resp.Body,
			total:      total,
			onProgress: onProgress,
		}
	}

	_, copyErr := io.Copy(out, reader)
	closeErr := out.Close()

	if copyErr != nil {
		_ = os.Remove(tmpPath)
		return fmt.Errorf("write %s: %w", relPath, copyErr)
	}
	if closeErr != nil {
		_ = os.Remove(tmpPath)
		return closeErr
	}

	if err := os.Rename(tmpPath, destPath); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	return nil
}

type progressReader struct {
	r          io.Reader
	total      int64
	written    int64
	onProgress ProgressFunc
}

func (p *progressReader) Read(b []byte) (int, error) {
	n, err := p.r.Read(b)
	if n > 0 {
		p.written += int64(n)
		p.onProgress(p.written, p.total)
	}
	return n, err
}

func escapePath(relPath string) string {
	relPath = path.Clean("/" + strings.ReplaceAll(relPath, "\\", "/"))
	relPath = strings.TrimPrefix(relPath, "/")

	parts := strings.Split(relPath, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return strings.Join(parts, "/")
}
