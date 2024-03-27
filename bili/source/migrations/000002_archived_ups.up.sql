create table archived_ups
(
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,
    tag_id      bigint not null default 0 comment 'my group',
    sign        varchar(2048) null comment 'up desc',
    uname       varchar(128) null comment 'up name',
    mid         bigint not null default 0 comment 'up uid',
    `level`     varchar(3) null comment 'up level',
    `rank`      varchar(30) null comment 'up rank',
    `following` varchar(30) null comment 'up following',
    follower    varchar(30) null comment 'up follower',
    `view`      varchar(30) null comment 'up view',
    likes       varchar(30) null comment 'up likes',
    resp        json             null,
    video       varchar(30) null comment 'up video count',
    PRIMARY KEY (tag_id, mid)
) comment '关注的UP';
