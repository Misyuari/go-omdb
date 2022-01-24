create table log
(
    id              int auto_increment,
    module          varchar(20) not null,
    request_payload json        not null,
    server_type     varchar(4)  not null,
    timestamp       datetime    not null,
    constraint log_pk
        primary key (id)
);
