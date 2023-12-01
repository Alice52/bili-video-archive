create table download (
    id bigserial primary key,
    type varchar,
    stream_url varchar,
    title varchar,
    tags varchar[],
    bvid varchar,
    author varchar,
    created_at timestamptz default current_timestamp,
    created_by varchar(128),
    updated_at timestamptz default current_timestamp,
    updated_by varchar(128),
    deleted_at bool
);