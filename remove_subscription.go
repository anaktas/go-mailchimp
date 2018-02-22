package mailchimp

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/RichardKnop/go-mailchimp/status"
)

// UpdateSubscription ...
func (c *Client) RemoveSubscription(listID string, email string, mergeFields map[string]interface{}) error {
	// Hash email
	emailMD5 := fmt.Sprintf("%x", md5.Sum([]byte(email)))
	// Make request
	params := map[string]interface{}{
		"email_address": email,
		"status":        status.Unsubscribed,
		"merge_fields":  mergeFields,
	}
	log.Print("Params: ")
	log.Println(params)
	resp, err := c.do(
		"DELETE",
		fmt.Sprintf("/lists/%s/members/%s", listID, emailMD5),
		&params,
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Print("Response body: ")
	log.Println(resp.Body)
	// Allow any success status (2xx)
	if resp.StatusCode/100 == 2 {
		return nil
	}

	// Request failed
	errorResponse, err := extractError(data)
	if err != nil {
		return err
	}
	return errorResponse
}
