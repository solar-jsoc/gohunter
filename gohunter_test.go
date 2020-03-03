package gohunter

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_AccountInformation(t *testing.T) {

	client, server := testServer(testAccountInformationHandler(t))

	defer server.Close()

	_, err := client.AccountInformation(context.TODO())
	if err != nil {
		t.Error(err)
	}
}

func TestClient_VerifyEmail(t *testing.T) {
	client, server := testServer(testVerifyEmailHandler(t))

	defer server.Close()

	_, err := client.VerifyEmail(context.TODO(), "steli@close.io")
	if err != nil {
		t.Error(err)
	}
}

func TestClient_DomainSearch(t *testing.T) {
	client, server := testServer(testDomainSearchHandler(t))

	defer server.Close()

	_, err := client.DomainSearch(context.TODO(), "intercom.io", "")
	if err != nil {
		t.Error(err)
	}
}

func TestClient_EmailCount(t *testing.T) {
	client, server := testServer(testEmailCountHandler(t))

	defer server.Close()

	_, err := client.EmailCount(context.TODO(), "intercom.io", "", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestClient_FindEmail(t *testing.T) {
	client, server := testServer(testFindEmailHandler(t))

	defer server.Close()

	_, err := client.FindEmail(context.TODO(), "intercom.io", "intercom", UsingFullName("John Doe"))
	if err != nil {
		t.Error(err)
	}
}

func TestWithDepartment(t *testing.T) {
	opt := WithDepartment(DepartmentSupport, DepartmentHR)

	opts := requestParams(url.Values{})

	opt(opts)

	require.Equal(t, string(DepartmentSupport)+";"+string(DepartmentHR), opts.Get(paramDepartment))
}

func TestWithSeniority(t *testing.T) {
	opt := WithSeniority(SeniorityJunior, SenioritySenior)

	opts := requestParams(url.Values{})

	opt(opts)

	require.Equal(t, string(SeniorityJunior)+";"+string(SenioritySenior), opts.Get(paramSeniority))
}

func TestWithCustomClient(t *testing.T) {

	customHTTPClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	opt := WithCustomClient(customHTTPClient)

	client := Client{}

	opt(&client)

	if client.client.Timeout != 1*time.Second {
		t.Error("custom client not used")
	}
}

func TestCheckResponse(t *testing.T) {
	res := http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Body:          ioutil.NopCloser(bytes.NewBuffer([]byte(domainSearchPayload))),
		ContentLength: 0,
	}

	err := checkResponse(&res)
	if err != nil {
		t.Error("error must be nil")
	}

	res = http.Response{
		Status:        "429 Too Many Requests",
		StatusCode:    429,
		Body:          ioutil.NopCloser(bytes.NewBuffer([]byte(tooManyRequestsPayload))),
		ContentLength: 0,
	}

	err = checkResponse(&res)
	if err == nil {
		t.Error("error must be not nil")
	}
}

func testAccountInformationHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/account" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := rw.Write([]byte(accountInformationPayload))
		if err != nil {
			t.Error(err)
		}
	}
}

func testVerifyEmailHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/email-verifier" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := rw.Write([]byte(verifyEmailPayload))
		if err != nil {
			t.Error(err)
		}
	}
}

func testDomainSearchHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/domain-search" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := rw.Write([]byte(domainSearchPayload))
		if err != nil {
			t.Error(err)
		}
	}
}

func testEmailCountHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/email-count" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := rw.Write([]byte(emailCountPayload))
		if err != nil {
			t.Error(err)
		}
	}
}

func testFindEmailHandler(t *testing.T) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/email-finder" {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := rw.Write([]byte(emailFinderPayload))
		if err != nil {
			t.Error(err)
		}
	}
}

func testServer(handler http.HandlerFunc) (*Client, *httptest.Server) {

	server := httptest.NewServer(handler)

	client := NewClient("token")
	client.baseEndpoint = server.URL + "/"

	return client, server
}
