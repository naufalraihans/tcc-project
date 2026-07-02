package midtrans

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	serverKey string
	snapBase  string
	apiBase   string
	http      *http.Client
}

func New(serverKey string, production bool) *Client {
	snap, api := "https://app.sandbox.midtrans.com", "https://api.sandbox.midtrans.com"
	if production {
		snap, api = "https://app.midtrans.com", "https://api.midtrans.com"
	}
	return &Client{serverKey: serverKey, snapBase: snap, apiBase: api, http: &http.Client{}}
}

type Customer struct {
	FirstName string `json:"first_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Item struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
}

type SnapRequest struct {
	OrderID     string
	GrossAmount float64
	Customer    Customer
	Item        Item
}

func (c *Client) CreateSnapToken(req SnapRequest) (string, string, error) {
	body, _ := json.Marshal(map[string]any{
		"transaction_details": map[string]any{"order_id": req.OrderID, "gross_amount": req.GrossAmount},
		"customer_details":    req.Customer,
		"item_details":        []Item{req.Item},
	})
	data, err := c.post(c.snapBase+"/snap/v1/transactions", body)
	if err != nil {
		return "", "", err
	}
	var out struct {
		Token       string `json:"token"`
		RedirectURL string `json:"redirect_url"`
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return "", "", err
	}
	return out.Token, out.RedirectURL, nil
}

func (c *Client) VerifySignature(orderID, statusCode, grossAmount, signature string) bool {
	sum := sha512.Sum512([]byte(orderID + statusCode + grossAmount + c.serverKey))
	return hex.EncodeToString(sum[:]) == signature
}

func (c *Client) Refund(orderID string) error {
	body, _ := json.Marshal(map[string]any{"reason": "kuota penuh"})
	_, err := c.post(c.apiBase+"/v2/"+orderID+"/refund", body)
	return err
}

func (c *Client) post(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.serverKey+":")))
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("midtrans %d: %s", res.StatusCode, string(data))
	}
	return data, nil
}
