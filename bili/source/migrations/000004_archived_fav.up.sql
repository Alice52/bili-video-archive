-- auto-generated definition
create table archived_fav
(
    id          bigint default 0 not null
    ,
    create_time datetime(3)      null,
    update_time datetime(3)      null,
    delete_time datetime(3)      null,
    fid         bigint           not null comment 'bili folder',
    bvid        varchar(64)      null comment 'bili avid',
    cover       varchar(256)     null comment 'video cover',
    ctime       bigint           not null comment 'video create time',
    duration    bigint           not null comment 'video duration',
    fav_time    bigint           not null comment 'video favor time',
    intro       text      null comment 'video intro',
    title       text      null comment 'video title',
    type        bigint           not null comment 'video type',
    season      json             null comment 'video season',
    upper       json             null comment '{"mid": 173986740, "name": "这个月-"}',
    cnt_info    json             null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }',
    resp        json             null,
    primary key (id, fid)
)
    comment '收藏视频';

