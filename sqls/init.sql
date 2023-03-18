CREATE TABLE instance_tab (
    id bigserial not null PRIMARY KEY,
    name varchar(64) not null,
    enable_http bool not null,
    enable_https bool not null,
    cert_name varchar(64) not null,
    updated_by varchar(64) not null,
    created_at int not null,
    updated_at int not null,
    deleted_at int not null,
    UNIQUE (name, deleted_at)
);
