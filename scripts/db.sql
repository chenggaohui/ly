-- auto-generated definition
create table product
(
    id              int auto_increment
        primary key,
    name            varchar(255)   not null,
    count           int            not null,
    production_date timestamp      not null,
    shelf_life      int            not null,
    expiration_date varchar(255)   null,
    is_expired      int default 1  null,
    price           decimal(10, 2) not null,
    warn_date int            not null,
    product_type_id int not null,
    ware_host_id int not null
);


create table product_type
(
    id              int auto_increment
        primary key,
    name            varchar(255)   not null
);

create table ware_host
(
    id              int auto_increment
        primary key,
    name            varchar(255)   not null
);

