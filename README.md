# Loan Service

## Create Loans
Flow when borrower propose the loans to system

<img width="396" alt="gambar" src="https://github.com/user-attachments/assets/8c4e8096-cbf0-47e8-b9be-93453cd217aa">

## Approve Loan
Flow when the staff approve the loan request

<img width="374" alt="gambar" src="https://github.com/user-attachments/assets/1b0ee4fc-e488-4356-95a0-0f08f1a46450">

## Invest Loan
Flow when the lender invest the amount of money into some loans.
If the total investment already == principal, publish the loanID into some MQ and notification service will responsible to email the lender

<img width="390" alt="gambar" src="https://github.com/user-attachments/assets/4c4f1146-8fde-4d4f-adc5-9690f28890b7">

## Disburse Loan
Change loan status into disburse when the money already disbursed to borrower

<img width="388" alt="gambar" src="https://github.com/user-attachments/assets/9e7daad2-e74f-4eac-b515-255958b92a28">

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
* [Technical] The code is still draft of API, so expected that the code still not working
* [Technical] Can improve the API docs by providing swagger / openAPI
* [Technical] Can improve the code by using unitest
* [Technical] Can improve overall flow by using CI/CD
* [Technical] Can add API to GET the detail
* [Business] Can add flow to reject the loans
* [Business] Can add detail of the loans tenure
* [Business] Can add flow when borrower failed to pay the loan



