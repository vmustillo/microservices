CREATE DATABASE payments
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

create extension pgcrypto;

create table payments (
	id UUID not null primary key default gen_random_uuid(),
    purchaser_account_id UUID,
    purchased_from_account_id UUID,
    amount decimal
);