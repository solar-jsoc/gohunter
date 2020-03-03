package gohunter

import (
	"net/http"
	"strconv"
)

const (
	// Seniority levels

	// SeniorityJunior is a junior seniority level
	SeniorityJunior seniorityLevel = "junior"
	// SeniorityExecutive is a executive seniority level
	SeniorityExecutive seniorityLevel = "executive"
	// SenioritySenior is a senior seniority level
	SenioritySenior seniorityLevel = "senior"

	// Possible departments

	// DepartmentExecutive is an executive department
	DepartmentExecutive department = "executive"
	// DepartmentIT is an it department
	DepartmentIT department = "it"
	// DepartmentFinance is a finance department
	DepartmentFinance department = "finance"
	// DepartmentManagement is a management department
	DepartmentManagement department = "management"
	// DepartmentSales is a sales department
	DepartmentSales department = "sales"
	// DepartmentLegal is a legal department
	DepartmentLegal department = "legal"
	// DepartmentSupport is a support department
	DepartmentSupport department = "support"
	// DepartmentHR is a hr department
	DepartmentHR department = "hr"
	// DepartmentMarketing is a marketing department
	DepartmentMarketing department = "marketing"
	// DepartmentCommunication is a communication department
	DepartmentCommunication department = "communication"

	// Emails type

	// EmailTypePersonal is a personal emails of specific employees.
	EmailTypePersonal emailType = "personal"
	// EmailTypeGeneric is a generic email addresses.
	EmailTypeGeneric emailType = "generic"

	// URL params
	paramSeniority  = "seniority"
	paramDepartment = "department"
	paramDomain     = "domain"
	paramCompany    = "company"
	paramLimit      = "limit"
	paramOffset     = "offset"
	paramEmailType  = "type"
	paramFirstName  = "first_name"
	paramLastName   = "last_name"
	paramFullName   = "full_name"
	paramAPIKey     = "api_key"
	paramEmail      = "email"
)

type (
	seniorityLevel string
	department     string
	emailType      string
)

type requestParams interface {
	Add(key, value string)
	Encode() string
	Get(key string) string
}

type requestOptionalParams interface {
	requestParams
}

type requestOptionalParam func(requestOptionalParams)

type emailSearchPersonOptions interface {
	requestParams
}

type emailSearchPersonOption func(emailSearchPersonOptions)

type clientOption func(*Client)

// WithCustomClient used in NewClient function. It allows use custom http.Client for method calls.
func WithCustomClient(client *http.Client) clientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithSeniority allows to get only email addresses for people with the selected seniority level.
// The possible values are junior, senior or executive.
// Several seniority levels can be selected (delimited by a comma).
// Please pass exported const for this method as an arguments.
func WithSeniority(levels ...seniorityLevel) requestOptionalParam {
	return func(params requestOptionalParams) {
		if len(levels) <= 0 {
			return
		}

		var lvls string

		for i := range levels {
			if i == len(levels)-1 {
				lvls += string(levels[i])
			} else {
				lvls += string(levels[i] + ";")
			}
		}

		params.Add(paramSeniority, lvls)
	}
}

// WithDepartment allows to get only email addresses for people working in the selected department(s).
// The possible values are executive, it, finance, management, sales, legal, support, hr, marketing or communication.
// Several departments can be selected (comma-delimited).
// Please pass exported const for this method as an arguments.
func WithDepartment(departments ...department) requestOptionalParam {
	return func(params requestOptionalParams) {
		if len(departments) <= 0 {
			return
		}

		var dd string

		for i := range departments {
			if i == len(departments)-1 {
				dd += string(departments[i])
			} else {
				dd += string(departments[i]) + ";"
			}
		}
		params.Add(paramDepartment, dd)
	}
}

// WithEmailType allows to get only personal or generic email addresses.
// Please pass exported const for this method as an arguments.
func WithEmailType(emailType emailType) requestOptionalParam {
	return func(params requestOptionalParams) {
		params.Add(paramEmailType, string(emailType))
	}
}

// WithLimit specifies the max number of email addresses to return. The default is 10.
func WithLimit(limit int) requestOptionalParam {
	return func(params requestOptionalParams) {
		params.Add(paramLimit, strconv.Itoa(limit))
	}
}

// WithOffset specifies the number of email addresses to skip. The default is 0.
func WithOffset(offset int) requestOptionalParam {
	return func(params requestOptionalParams) {
		params.Add(paramOffset, strconv.Itoa(offset))
	}
}

// UsingFullName search hunter.io person using full name separated with space.
func UsingFullName(fullName string) emailSearchPersonOption {
	return func(params emailSearchPersonOptions) {
		params.Add(paramFullName, fullName)
	}
}

// UsingFirstLastName search hunter.io person using first and last name.
func UsingFirstLastName(firtsName, lastName string) emailSearchPersonOption {
	return func(params emailSearchPersonOptions) {
		params.Add(paramFirstName, firtsName)
		params.Add(paramLastName, lastName)
	}
}
