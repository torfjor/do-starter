package apix

import "fmt"

var (
	ErrInvalidAPIResponseCode = fmt.Errorf("invalid api response code")
	ErrInvalidContentType     = fmt.Errorf("invalid api response content type")
)
