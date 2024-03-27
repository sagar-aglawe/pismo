DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    document_number VARCHAR(50)
);

DROP TABLE IF EXISTS operation_types;

CREATE TABLE operation_types(
    id SERIAL PRIMARY KEY,
    description VARCHAR(50)
);


DROP TABLE IF EXISTS transactions;

CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    account_id SERIAL,
    operation_type_id SERIAL,
    amount NUMERIC,
    event_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);