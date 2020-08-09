CREATE DATABASE account
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

create extension pgcrypto;

create table accounts (
	id UUID not null primary key default gen_random_uuid(),
	owner_id UUID,
	balance decimal,
    credit_limit decimal,
	card_id UUID
);

create table transactions (
	id UUID not null primary key default gen_random_uuid(),
    purchaser_account_id UUID,
    purchased_from_account_id UUID,
    amount decimal
);