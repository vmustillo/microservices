CREATE DATABASE cards
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

create extension pgcrypto;

create table cards (
	id UUID not null primary key default gen_random_uuid(),
	card_number varchar(50),
	account_id varchar(50)
);