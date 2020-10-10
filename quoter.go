package do_starter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type RandomQuoter struct {
	client *http.Client
}

const quoteAPIEndpoint = "https://api.whatdoestrumpthink.com/api/v1/quotes/random"

func (r *RandomQuoter) GetQuote(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteAPIEndpoint, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid api response code: %d", res.StatusCode)
	}

	ct := res.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return "", fmt.Errorf("invalid api response content type: %q", ct)
	}

	type apiResponse struct {
		Message string `json:"message"`
	}
	var response apiResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Message, nil
}
