CREATE TABLE users_balances
(
    id serial not null unique,
    balance int
);

CREATE TABLE

(
    id serial not null unique,
    user_id int references users_balances(id) not null,
    amount int not null,
    reason varchar(255),
    transaction_date date not null
);