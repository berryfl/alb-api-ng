CREATE TABLE instance_tab (
    id bigserial not null PRIMARY KEY,
    name varchar(64) not null,
    service varchar(128) not null,
    updated_by varchar(64) not null,
    created_at int not null,
    updated_at int not null,
    deleted_at int not null,
    UNIQUE (name, deleted_at)
);

CREATE TABLE router_tab (
    id bigserial not null PRIMARY KEY,
    instance_name varchar(64) not null,
    domain varchar(128) not null,
    enable_http bool not null,
    enable_https bool not null,
    cert_name varchar(128) not null,
    content jsonb not null,
    updated_by varchar(64) not null,
    created_at int not null,
    updated_at int not null,
    deleted_at int not null,
    UNIQUE (instance_name, domain, deleted_at)
);

CREATE INDEX idx_router_content ON router_tab USING GIN (content);

CREATE TABLE target_tab (
    id bigserial not null PRIMARY KEY,
    instance_name varchar(64) not null,
    name varchar(64) not null,
    target_type varchar(64) not null,
    content jsonb not null,
    updated_by varchar(64) not null,
    created_at int not null,
    updated_at int not null,
    deleted_at int not null,
    UNIQUE (instance_name, name, deleted_at)
);
