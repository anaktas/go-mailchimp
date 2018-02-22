package mailchimp

import (
	"net/url"
)

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	CheckSubscription(listID string, email string) (*MemberResponse, error)
	Subscribe(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	UpdateSubscription(listID string, email string, mergeFields map[string]interface{}) (*MemberResponse, error)
	RemoveSubscription(listID string, email string, mergeFields map[string]interface{}) error
	SetBaseURL(baseURL *url.URL)
	GetBaseURL() *url.URL
}
