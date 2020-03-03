package gohunter

import "fmt"

// HunterErrors implements hunter.io errors response
type HunterErrors struct {
	Errors []HunterError `json:"errors"`
}

// HunterError implements any hunter.io API returned error.
type HunterError struct {
	ID      string `json:"id"`
	Code    int    `json:"code"`
	Details string `json:"details"`
}

func (e HunterError) Error() string {
	return fmt.Sprintf("id: %s, code: %d, details: %s", e.ID, e.Code, e.Details)
}

//	Common "Meta" types

// Meta contains meta information about response and request.
type Meta struct {
	Results int        `json:"results"`
	Limit   int        `json:"limit"`
	Offset  int        `json:"offset"`
	Params  MetaParams `json:"params"`
}

// MetaParams meta params.
type MetaParams struct {
	Domain     string `json:"domain"`
	Company    string `json:"company"`
	Type       string `json:"type"`
	Seniority  string `json:"seniority"`
	Department string `json:"department"`
}

// EmailSource contains data about email source. It's common and embedded in some other structs.
type EmailSource struct {
	Domain      string `json:"domain"`
	URI         string `json:"uri"`
	ExtractedOn string `json:"extracted_on"`
	LastSeenOn  string `json:"last_seen_on"`
	StillOnPage bool   `json:"still_on_page"`
}

//	domain-search method response types

// DomainSearchResponse is a response from domain-search method call.
type DomainSearchResponse struct {
	Data DomainSearchData `json:"data"`
	Meta Meta             `json:"meta"`
}

// DomainSearchData contains data from domain-search method response.
type DomainSearchData struct {
	Domain       string              `json:"domain"`
	Disposable   bool                `json:"disposable"`
	Webmail      bool                `json:"webmail"`
	AcceptAll    bool                `json:"accept_all"`
	Pattern      string              `json:"pattern"`
	Organization string              `json:"organization"`
	Country      string              `json:"country"`
	State        string              `json:"state"`
	Emails       []DomainSearchEmail `json:"emails"`
}

// DomainSearchEmail contains data about specific email.
type DomainSearchEmail struct {
	Value       string        `json:"value"`
	Type        string        `json:"type"`
	Confidence  int           `json:"confidence"`
	Sources     []EmailSource `json:"sources"`
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Position    string        `json:"position"`
	Seniority   string        `json:"seniority"`
	Department  string        `json:"department"`
	Linkedin    string        `json:"linkedin"`
	Twitter     string        `json:"twitter"`
	PhoneNumber string        `json:"phone_number"`
}

//	email-finder method response types

// EmailFinderReponse is a response from email-verifier method call
type EmailFinderReponse struct {
	EmailFinderData EmailFinderData `json:"data"`
	Meta            Meta            `json:"meta"`
}

// EmailFinderData contains data from email-finder method response
type EmailFinderData struct {
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Email       string        `json:"email"`
	Score       int           `json:"score"`
	Domain      string        `json:"domain"`
	AcceptAll   bool          `json:"accept_all"`
	Position    string        `json:"position"`
	Twitter     string        `json:"twitter"`
	LinkedinURL string        `json:"linkedin_url"`
	PhoneNumber string        `json:"phone_number"`
	Company     string        `json:"company"`
	Sources     []EmailSource `json:"sources"`
}

// email-verifier method response types

// EmailVerifierResponse is a response from email-verifier method call.
type EmailVerifierResponse struct {
	Data EmailVerifierData `json:"data"`
	Meta Meta              `json:"meta"`
}

// EmailVerifierData contains data from email-verifier method response.
type EmailVerifierData struct {
	Result     string        `json:"result"`
	Score      int           `json:"score"`
	Email      string        `json:"email"`
	Regexp     bool          `json:"regexp"`
	Gibberish  bool          `json:"gibberish"`
	Disposable bool          `json:"disposable"`
	Webmail    bool          `json:"webmail"`
	MxRecords  bool          `json:"mx_records"`
	SMTPServer bool          `json:"smtp_server"`
	SMTPCheck  bool          `json:"smtp_check"`
	AcceptAll  bool          `json:"accept_all"`
	Block      bool          `json:"block"`
	Sources    []EmailSource `json:"sources"`
}

// email-count method response types

// EmailCountResponse response from email-count method.
type EmailCountResponse struct {
	Data EmailCountData `json:"data"`
	Meta Meta           `json:"meta"`
}

// EmailCountData contains data from email-count method response.
type EmailCountData struct {
	Total          int                  `json:"total"`
	PersonalEmails int                  `json:"personal_emails"`
	GenericEmails  int                  `json:"generic_emails"`
	Department     EmailCountDepartment `json:"department"`
	Seniority      EmailCountSeniority  `json:"seniority"`
}

// EmailCountDepartment contains information about how many mails of employees of each department are in hunter.io.
type EmailCountDepartment struct {
	Executive     int `json:"executive"`
	It            int `json:"it"`
	Finance       int `json:"finance"`
	Management    int `json:"management"`
	Sales         int `json:"sales"`
	Legal         int `json:"legal"`
	Support       int `json:"support"`
	Hr            int `json:"hr"`
	Marketing     int `json:"marketing"`
	Communication int `json:"communication"`
}

// EmailCountSeniority contains information about how many mails of each seniority level are in hunter.io.
type EmailCountSeniority struct {
	Junior    int `json:"junior"`
	Senior    int `json:"senior"`
	Executive int `json:"executive"`
}

// account-information method response types

// AccountInformationResponse response from "account" method.
type AccountInformationResponse struct {
	Data AccountInformationData `json:"data"`
}

// AccountInformationData contains account information data.
type AccountInformationData struct {
	FirstName string                  `json:"first_name"`
	LastName  string                  `json:"last_name"`
	Email     string                  `json:"email"`
	PlanName  string                  `json:"plan_name"`
	PlanLevel int                     `json:"plan_level"`
	ResetDate string                  `json:"reset_date"`
	TeamID    int                     `json:"team_id"`
	Calls     AccountInformationCalls `json:"calls"`
}

// AccountInformationCalls contain information about used and available calls.
type AccountInformationCalls struct {
	Used      int `json:"used"`
	Available int `json:"available"`
}
