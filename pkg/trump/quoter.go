package trump

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/torfjor/do-starter/pkg/apix"
)

const endpoint = "https://api.whatdoestrumpthink.com/api/v1/quotes/random"

type Quoter struct {
}

func (r *Quoter) GetQuote(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: %d", apix.ErrInvalidAPIResponseCode, res.StatusCode)
	}

	ct := res.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		return "", fmt.Errorf("%w: %q", apix.ErrInvalidContentType, ct)
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
