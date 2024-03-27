create table archived_ups_tag
(
    tag_id      bigint primary key default 0 comment 'tagid',
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,
    `name`      varchar(64) null comment 'name',
    `count`     bigint null comment 'count',
    resp        json             null,
    tip         varchar(128) null comment 'tip'
) comment '关注的UP主分组';
