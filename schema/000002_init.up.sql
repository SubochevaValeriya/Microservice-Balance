CREATE TABLE users_balances
(
    id serial not null unique,
    balance int
);

CREATE TABLE transactions

(
    id serial not null unique,
    user_id int references users_balances(id) not null,
    amount int not null,
    reason varchar(255),
    transfer_id int references users_balances(id),
    transaction_date date not null
);