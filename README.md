# Loan Service

## Create Loans
Flow when borrower propose the loans to system
```mermaid
sequenceDiagram
    actor Borrower
    Borrower->>LoanSystem: Propose Loan
    LoanSystem->>DB: Save Proposed Loan
```

## Approve Loan
Flow when the staff approve the loan request

```mermaid
sequenceDiagram
    actor Staff
    Staff->>LoanSystem: Approve Loan Request
    LoanSystem->>DB: Update Loan Status
```


## Invest Loan
Flow when the lender invest the amount of money into some loans.
If the total investment already == principal, publish the loanID into some MQ and notification service will responsible to email the lender
```mermaid
sequenceDiagram
    actor Lender
    Lender->>LoanSystem: Invest
    LoanSystem->>DB: Insert Investor
    LoanSystem->>MQ: Publish if Invested
    MQ->>NotificationSystem: 
    NotificationSystem->>Lender: Send Email

```

## Disburse Loan
Change loan status into disburse when the money already disbursed to borrower

```mermaid
sequenceDiagram
    actor Staff
    Staff->>LoanSystem: Update disbursment status
    LoanSystem->>DB: Update State
```


# API List
```
loanGroup := e.Group("loan")
loanGroup.POST("/loans", a.CreateLoan)
loanGroup.PUT("/loans/:id/approve", a.ApproveLoan)
loanGroup.POST("/loans/:id/invest", a.RecordLoanInvestment)
loanGroup.PUT("/loans/:id/disburse", a.RecordLoanDisbursement)
// TODO: Create API To get loan detail
loanGroup.GET("/loans/:id", nil)
loanGroup.GET("/loans", nil)
```

# Improvement
* [Technical] The code is still draft of API
* [Technical] Can improve the API docs by providing swagger / openAPI
* [Technical] Can improve the code by using unitest
* [Technical] Can improve overall flow by using CI/CD
* [Technical] Can add API to GET the detail
* [Technical] Can implement Auth for security
* [Business] Can add flow to reject the loans
* [Business] Can add detail of the loans tenure
* [Business] Can add flow when borrower failed to pay the loan



