## HTTP Service, which provides API for working with user balances (transactions), namely:

### POST:

```/api/``` - CREATE USER BALANCE 

*(Reason: Opening)*

Body:
{"balance": }

### PUT:

```/api/``` - UPDATE USERS BALANCES: CREATE NEW TRANSACTION 

*(Reason: Transfer from one user to another)*

Body:
{"userId":,
"Amount":",
"TransferId:}

```/api/:id``` - UPDATE USER BALANCE BY ID: CREATE NEW TRANSACTION 

*(Reason: Replenishment or Withdrawal depending on the amount)*

Body:
{"Amount":"}

### GET:

```/api/``` - GET ALL USERS BALANCES

```/api/id?id=id&currency=CCY``` - GET USER BALANCE

*you can specify currency in which data will be reflected (on current rate) *

```/transactions/id?id=id&currency=CCY``` - GET USER'S LIST OF TRANSACTIONS

*you can specify currency in which data will be reflected (on current rate)*

### DELETE:

```/api/:id``` - DELETE USER BALANCE BY ID *(and all transactions of this user)*

When the balance changes, a validation is performed, if the balance becomes less than zero, the transaction is canceled.

**Used:** *Postgres, REST API principles, gin-gonic framework, docker-compose, work with external API.*

### To run an app:

```
make build && make run
```

If you are running the application for the first time, you must apply the migrations to the database:

```
make migrate
```