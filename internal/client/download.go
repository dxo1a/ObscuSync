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

// DownloadFile качает файл с сервера и сохраняет в destPath.
// relPath — путь как в манифесте (mods/foo.zip), со слешами.
func (c *Client) DownloadFile(profile, relPath, destPath string) error {
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

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	// пишем во временный файл рядом, потом rename — атомарнее
	tmpPath := destPath + ".tmp"
	out, err := os.Create(tmpPath)
	if err != nil {
		return err
	}

	_, copyErr := io.Copy(out, resp.Body)
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

func escapePath(relPath string) string {
	relPath = path.Clean("/" + strings.ReplaceAll(relPath, "\\", "/"))
	relPath = strings.TrimPrefix(relPath, "/")

	parts := strings.Split(relPath, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return strings.Join(parts, "/")
}
