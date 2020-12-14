

create type virtual_vending.UserRolesEnum as enum ('manager', 'engineer', 'admin');

create table if not exists virtual_vending.users (
    id serial primary key,
    username varchar(63) not null,
    email varchar(63) not null,
    password varchar not null,
    address varchar,
    role userrolesenum not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint unique_username unique (username)
);


create table if not exists virtual_vending.products (
    id serial primary  key,
    name varchar not null,
    description varchar not null,
    cost int
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint unique_product_name unique (name)
)

alter table add column created_at timestamp  default current_timestamp, add column update_at timestamp  default current_timestamp;