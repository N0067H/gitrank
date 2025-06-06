package githubclient

import "errors"

var (
	ErrInternalServerError   = errors.New("Internal Server Error")
	ErrGithubAPIError        = errors.New("Github API Error")
	ErrCreatingRequestFailed = errors.New("Creating Request Failed")
	ErrReadingResponseFailed = errors.New("Reading Response Failed")
)
