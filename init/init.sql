create table user
(
    id         int unsigned auto_increment primary key,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    username   varchar(255) not null,
    password   varchar(255) not null,
    email      varchar(255) null
);