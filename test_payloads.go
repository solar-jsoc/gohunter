package gohunter

const (
	accountInformationPayload = `
		{   
			"data" :{      
				"first_name":"Antoine",
      			"last_name":"Finkelstein",
      			"email":"antoine@hunter.io",
      			"plan_name":"Pro",
      			"plan_level":4,
      			"reset_date":"2020-03-07",
      			"team_id":1,
      			"calls":{
      				"used":28526,
					"available":50000
				}
			}
		}`

	verifyEmailPayload = `
		{
		  "data": {
			"result": "deliverable",
			"score": 91,
			"email": "steli@close.io",
			"regexp": true,
			"gibberish": false,
			"disposable": false,
			"webmail": false,
			"mx_records": true,
			"smtp_server": true,
			"smtp_check": true,
			"accept_all": false,
			"block": false,
			"sources": [
			  {
				"domain": "blog.close.io",
				"uri": "http://blog.close.io/how-to-become-great-at-sales",
				"extracted_on": "2015-01-26",
				"last_seen_on": "2017-02-25",
				"still_on_page": true
			  },
			  {
				"domain": "blog.close.io",
				"uri": "http://blog.close.io/how-to-do-referral-sales",
				"extracted_on": "2015-01-26",
				"last_seen_on": "2016-02-25",
				"still_on_page": false
			  }
			]
		  },
		  "meta": {
			"params": {
			  "email": "steli@close.io"
			}
		  }
		}`

	domainSearchPayload = `
		{
		  "data": {
			"domain": "intercom.io",
			"disposable": false,
			"webmail": false,
			"accept_all": false,
			"pattern": "{first}",
			"organization": "Intercom",
			"country": null,
			"state": null,
			"emails": [
			  {
				"value": "ciaran@intercom.io",
				"type": "personal",
				"confidence": 92,
				"sources": [
				  {
					"domain": "github.com",
					"uri": "http://github.com/ciaranlee",
					"extracted_on": "2015-07-29",
					"last_seen_on": "2017-07-01",
					"still_on_page": true
				  },
				  {
					"domain": "blog.intercom.io",
					"uri": "http://blog.intercom.io/were-hiring-a-support-engineer/",
					"extracted_on": "2015-08-29",
					"last_seen_on": "2017-07-01",
					"still_on_page": true
				  }
				],
				"first_name": "Ciaran",
				"last_name": "Lee",
				"position": "Support Engineer",
				"seniority": "senior",
				"department": "it",
				"linkedin": null,
				"twitter": "ciaran_lee",
				"phone_number": null
			  }
			]
		  },
		  "meta": {
			"results": 35,
			"limit": 10,
			"offset": 0,
			"params": {
			  "domain": "intercom.io",
			  "company": null,
			  "type": null,
			  "seniority": null,
			  "department": null
			}
		  }
		}`

	emailCountPayload = `
		{
		  "data": {
			"total": 81,
			"personal_emails": 65,
			"generic_emails": 16,
			"department": {
			  "executive": 10,
			  "it": 0,
			  "finance": 8,
			  "management": 0,
			  "sales": 0,
			  "legal": 0,
			  "support": 6,
			  "hr": 0,
			  "marketing": 0,
			  "communication": 2
			},
			"seniority": {
			  "junior": 13,
			  "senior": 5,
			  "executive": 2
			}
		  },
		  "meta": {
			"params": {
			  "domain": "stripe.com",
			  "type": null
			}
		  }
		}`

	emailFinderPayload = `
		{
		  "data": {
			"first_name": "Dustin",
			"last_name": "Moskovitz",
			"email": "dustin@asana.com",
			"score": 72,
			"domain": "asana.com",
			"accept_all": false,
			"position": "CEO",
			"twitter": "moskov",
			"linkedin_url": "https://www.linkedin.com/in/dmoskov",
			"phone_number": null,
			"company": "Asana",
			"sources": [
			  {
				"domain": "blog.asana.com",
				"uri": "http://blog.asana.com",
				"extracted_on": "2015-09-27",
				"last_seen_on": "2017-09-01",
				"still_on_page": true
			  }
			]
		  },
		  "meta": {
			"params": {
			  "first_name": "Dustin",
			  "last_name": "Moskovitz",
			  "full_name": null,
			  "domain": "asana.com",
			  "company": null
			}
		  }
		}`

	tooManyRequestsPayload = `
		{
		  "errors": [
			{
			  "id": "too_many_requests",
			  "code": 429,
			  "details": "Too many requests, try again later."
			}
		  ]
		}`
)
