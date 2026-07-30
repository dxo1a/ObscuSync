package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dxo1a/obscusync/internal/manifest"
)

func (c *Client) FetchManifest(profile string) (manifest.Manifest, error) {
	var m manifest.Manifest

	if profile == "" {
		return m, fmt.Errorf("profile is required")
	}

	req, err := http.NewRequest(http.MethodGet, c.url("/api/v1/manifest/"+profile), nil)
	if err != nil {
		return m, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return m, fmt.Errorf("fetch manifest: %w\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return m, errStatus(resp, "fetch manifest")
	}

	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return m, fmt.Errorf("decode manifest: %w", err)
	}

	return m, nil
}
