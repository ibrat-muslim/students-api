create table students (
    id serial primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    username varchar(30) not null,
    email varchar(50) not null,
    phone_number varchar(20) not null,
    created_at timestamp with time zone not null default current_timestamp
);