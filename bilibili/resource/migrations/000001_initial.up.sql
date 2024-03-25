create table archived_video (
    id bigserial primary key,
    type varchar,
    stream_url varchar,
    title varchar,
    tags varchar[],
    bvid varchar,
    author varchar,
    status bool,
    pan_url varchar,
    created_at timestamptz default current_timestamp,
    created_by varchar(128),
    updated_at timestamptz default current_timestamp,
    updated_by varchar(128),
    deleted_at bool
);