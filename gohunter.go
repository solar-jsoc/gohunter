package gohunter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseEndpoint           = "https://api.hunter.io/v2/"
	domainSearchPath       = "domain-search?"
	emailFindPath          = "email-finder?"
	emailVerifyPath        = "email-verifier?"
	emailCountPath         = "email-count?"
	accountInformationPath = "account?"
)

// Client is an API client
type Client struct {
	token        string
	baseEndpoint string
	client       *http.Client
}

// NewClient returns a pointer to a client
// To call API methods you must first create a client.
// You could use your custom http.Client using WithCustomClient option.
func NewClient(token string, opts ...clientOption) *Client {

	client := Client{
		token:        token,
		client:       http.DefaultClient,
		baseEndpoint: baseEndpoint,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(&client)
		}
	}

	return &client
}

// DomainSearch implements "domain-search" hunter.io method.
// for "params" use exported functions that returns optionalParam type.
// You need to specify at least domain or companyName for this call.
// method docs: https://hunter.io/api-documentation/v2#domain-search
func (c *Client) DomainSearch(ctx context.Context, domain, companyName string, params ...requestOptionalParam) (*DomainSearchResponse, error) {

	pp := requestParams(url.Values{})

	err := addSearchOrgParam(pp, domain, companyName)
	if err != nil {
		return nil, err
	}

	for _, param := range params {
		param(pp)
	}

	req, err := c.newGetRequest(ctx, domainSearchPath, pp)
	if err != nil {
		return nil, err
	}

	var domainSearchResponse DomainSearchResponse

	err = c.sendRequest(req, &domainSearchResponse)
	return &domainSearchResponse, err
}

// FindEmail implements "email-finder" hunter.io method.
// You need to specify at least domain or companyName for this call.
// method docs: https://hunter.io/api-documentation/v2#email-finder
func (c *Client) FindEmail(ctx context.Context, domain, companyName string, withNameParam emailSearchPersonOption) (*EmailFinderReponse, error) {

	params := requestParams(url.Values{})
	withNameParam(params)

	err := addSearchOrgParam(params, domain, companyName)
	if err != nil {
		return nil, err
	}

	req, err := c.newGetRequest(ctx, emailFindPath, params)
	if err != nil {
		return nil, err
	}

	var emailFinderResponse EmailFinderReponse

	err = c.sendRequest(req, &emailFinderResponse)
	return &emailFinderResponse, err
}

// VerifyEmail implements email-verifier hunter.io method.
// method docs: https://hunter.io/api-documentation/v2#email-verifier
func (c *Client) VerifyEmail(ctx context.Context, email string) (*EmailVerifierResponse, error) {
	if email == "" {
		return nil, fmt.Errorf("error: email to verify must be specified")
	}

	params := requestParams(url.Values{})

	params.Add(paramEmail, email)

	req, err := c.newGetRequest(ctx, emailVerifyPath, params)
	if err != nil {
		return nil, err
	}

	var emailVerifyResponse EmailVerifierResponse

	err = c.sendRequest(req, &emailVerifyResponse)
	return &emailVerifyResponse, err
}

// EmailCount implements "email-count" hunter.io method.
// You need to specify at least domain or companyName for this call.
// emailType is not required for this call, just pass nil if it doesn't needed!
// method docs: https://hunter.io/api-documentation/v2#email-count
func (c *Client) EmailCount(ctx context.Context, domain, companyName string, emailType requestOptionalParam) (*EmailCountResponse, error) {
	params := requestParams(url.Values{})

	if emailType != nil {
		emailType(params)
	}

	err := addSearchOrgParam(params, domain, companyName)
	if err != nil {
		return nil, err
	}

	req, err := c.newGetRequest(ctx, emailCountPath, params)
	if err != nil {
		return nil, err
	}

	var emailCountResponse EmailCountResponse

	err = c.sendRequest(req, &emailCountResponse)
	return &emailCountResponse, err
}

// AccountInformation implements "account" hunter.io method.
// For more information: https://hunter.io/api-documentation/v2#account
func (c *Client) AccountInformation(ctx context.Context) (*AccountInformationResponse, error) {

	req, err := c.newGetRequest(ctx, accountInformationPath, url.Values{})
	if err != nil {
		return nil, err
	}

	var accountInformationResponse AccountInformationResponse

	err = c.sendRequest(req, &accountInformationResponse)
	return &accountInformationResponse, err
}

func addSearchOrgParam(params requestParams, domain, companyName string) error {
	if domain != "" && companyName != "" {
		params.Add(paramDomain, domain)
		params.Add(paramCompany, companyName)
	} else if domain != "" {
		params.Add(paramDomain, domain)
	} else if companyName != "" {
		params.Add(paramCompany, companyName)
	} else {
		return fmt.Errorf("error at least domain or company name must be specified")
	}

	return nil
}

func checkResponse(res *http.Response) error {
	if res.StatusCode >= 400 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		e := HunterErrors{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return fmt.Errorf("unmarshal hunter.io error response error: %w", err)
		}

		return &e.Errors[0]
	}

	return nil
}

func (c *Client) sendRequest(req *http.Request, dest interface{}) error {

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	err = checkResponse(res)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, dest)
}

func (c *Client) newGetRequest(ctx context.Context, method string, params requestParams) (*http.Request, error) {
	link := c.baseEndpoint + method + paramAPIKey + "=" + c.token + "&" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", link, nil)

	return req, err
}
