# gohunter

Client library for hunter.io

# Methods

Supported methods:

- [x] GET domain-search
- [x] GET email-finder
- [x] GET email-verifier
- [x] GET email-count
- [x] GET account
- [ ] GET leads (get all leads)
- [ ] GET lead/{number} (get one lead)
- [ ] POST leads (create lead)
- [ ] PUT lead/{number} (update lead)
- [ ] DELETE lead/{number} (delete lead)

# Installation

```
    go github.com/solar-jsoc/gohunter
```

# Examples

- client init:
 
    ```
    client := gohunter.NewClient("token", gohunter.WithCustomClient(http.DefaultClient))
    ```
    
- domain-search
    
    ```
    result, err := client.DomainSearch(context.TODO(), "intercom.io", "intercom", 
    	gohunter.WithLimit(20), 
    	gohunter.WithDepartment(DepartmentIT, DepartmentFinance), 
    	gohunter.WithSeniority(SenioritySenior),
    	gohunter.WithEmailType(EmailTypePersonal),
    	gohunter.WithOffset(10))
    if err != nil {
        ...
    }
    ```
    	
- email-finder

    ```
  	result, err := client.FindEmail(context.TODO(), "intercom.io", "intercom", gohunter.UsingFullName("John Doe"))
  	if err != nil {
  		...
  	}
    ```

- email-verifier
    
    ```
  	result, err := client.VerifyEmail(context.TODO(), "ciaran@intercom.io")
  	if err != nil {
  		...
  	}
    ```
  
- email-count
    
    ```
  	result, err := client.EmailCount(context.TODO(), "", "intercom", gohunter.EmailTypePersonal)
  	if err != nil {
  		...
  	}
    ```
  
- account
    
    ```
    result, err := client.AccountInformation(context.TODO())
    if err != nil {
    	...
    }
    ```