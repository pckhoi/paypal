package paypal

import (
	"fmt"
	"time"
)

type (
	// SubscriptionRequest struct
	SubscriptionRequest struct {
		PlanID             string              `json:"plan_id"`
		StartTime          *time.Time          `json:"start_time,omitempty"`
		Quantity           string              `json:"quantity"`
		ShippingAmount     *Money              `json:"shipping_amount,omitempty"`
		Subscriber         *Subscriber         `json:"subscriber,omitempty"`
		AutoRenewal        bool                `json:"auto_renewal,omitempty"`
		ApplicationContext *ApplicationContext `json:"application_context,omitempty"`
	}

	// SubscriptionReviseRequest struct
	SubscriptionReviseRequest struct {
		PlanID             string              `json:"plan_id"`
		EffectiveTime      *time.Time          `json:"effective_time,omitempty"`
		Quantity           string              `json:"quantity"`
		ShippingAmount     *Money              `json:"shipping_amount,omitempty"`
		ShippingAddress    *ShippingDetail     `json:"shipping_address,omitempty"`
		ApplicationContext *ApplicationContext `json:"application_context,omitempty"`
	}

	// SubscriptionReviseResp struct
	SubscriptionReviseResp struct {
		PlanID          string            `json:"plan_id"`
		EffectiveTime   *time.Time        `json:"effective_time,omitempty"`
		Quantity        string            `json:"quantity"`
		ShippingAmount  *Money            `json:"shipping_amount,omitempty"`
		ShippingAddress *ShippingDetail   `json:"shipping_address,omitempty"`
		Links           []LinkDescription `json:"links"`
	}

	// SubscriptionTransactionsList struct
	SubscriptionTransactionsList struct {
		Transactions []SubscriptionTransaction `json:"transaction"`
		TotalItems   int                       `json:"total_items,omitempty"`
		TotalPages   int                       `json:"total_pages,omitempty"`
		Links        []LinkDescription         `json:"links"`
	}
)

// CreateSubscription creates a subscription in Paypal
// Endpoint: POST /v1/billing/subscriptions
func (c *Client) CreateSubscription(request SubscriptionRequest) (*Subscription, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s/v1/billing/subscriptions", c.APIBase), request)
	if err != nil {
		return nil, err
	}
	response := &Subscription{}
	err = c.SendWithAuth(req, response)
	return response, err
}

// GetSubscription shows details for a subscription, by ID.
// Endpoint: GET /v1/billing/subscriptions/ID
func (c *Client) GetSubscription(subscriptionID string) (*Subscription, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s/v1/billing/subscriptions/%s", c.APIBase, subscriptionID), nil)
	if err != nil {
		return nil, err
	}
	response := &Subscription{}
	err = c.SendWithAuth(req, response)
	return response, err
}

// ActivateSubscription activates the subscription.
// Endpoint: POST /v1/billing/subscriptions/ID/activate
func (c *Client) ActivateSubscription(subscriptionID, reason string) error {
	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("%s/v1/billing/subscriptions/%s/activate", c.APIBase, subscriptionID),
		ReasonedRequest{Reason: reason},
	)
	if err != nil {
		return err
	}
	return c.SendWithAuth(req, nil)
}

// CancelSubscription cancels the subscription.
// Endpoint: POST /v1/billing/subscriptions/ID/cancel
func (c *Client) CancelSubscription(subscriptionID, reason string) error {
	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("%s/v1/billing/subscriptions/%s/cancel", c.APIBase, subscriptionID),
		ReasonedRequest{Reason: reason},
	)
	if err != nil {
		return err
	}
	return c.SendWithAuth(req, nil)
}

// ReviseSubscription revises the subscription.
// Endpoint: POST /v1/billing/subscriptions/ID/revise
func (c *Client) ReviseSubscription(subscriptionID string, request SubscriptionReviseRequest) (*SubscriptionReviseResp, error) {
	req, err := c.NewRequest(
		"POST", fmt.Sprintf("%s/v1/billing/subscriptions/%s/revise", c.APIBase, subscriptionID), request,
	)
	if err != nil {
		return nil, err
	}
	response := &SubscriptionReviseResp{}
	err = c.SendWithAuth(req, response)
	return response, err
}

// SuspendSubscription suspends the subscription.
// Endpoint: POST /v1/billing/subscriptions/ID/suspend
func (c *Client) SuspendSubscription(subscriptionID, reason string) error {
	req, err := c.NewRequest(
		"POST",
		fmt.Sprintf("%s/v1/billing/subscriptions/%s/suspend", c.APIBase, subscriptionID),
		ReasonedRequest{Reason: reason},
	)
	if err != nil {
		return err
	}
	return c.SendWithAuth(req, nil)
}

// ListSubscriptionTransactions lists transactions the subscription.
// Endpoint: POST /v1/billing/subscriptions/ID/transactions
func (c *Client) ListSubscriptionTransactions(subscriptionID string) (*SubscriptionTransactionsList, error) {
	req, err := c.NewRequest(
		"GET", fmt.Sprintf("%s/v1/billing/subscriptions/%s/transactions", c.APIBase, subscriptionID), nil,
	)
	if err != nil {
		return nil, err
	}
	response := &SubscriptionTransactionsList{}
	err = c.SendWithAuth(req, response)
	return response, err
}
