CREATE DATABASE auth
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

create extension pgcrypto;

create table users (
	id UUID not null primary key default gen_random_uuid(),
	first_name varchar(50),
	last_name varchar(50),
	full_name varchar(100),
    account_id UUID,
    username varchar(100),
    salt bytea,
    hashed_password bytea
);