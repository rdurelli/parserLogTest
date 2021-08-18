create table log
(
    ID int,
    date varchar(50) null,
    time varchar(80) null,
    http_code varchar(20) null,
    http_method varchar(40) null,
    path varchar(255) null
);

create unique index log_ID_uindex
    on log (ID);

alter table log
    add constraint log_pk
        primary key (ID);

alter table log modify ID int auto_increment;

