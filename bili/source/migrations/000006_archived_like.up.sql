create table archived_like
(
    bvid       varchar(64)       primary key ,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    aid         bigint           not null comment 'bili aid',
    cid         bigint           not null comment 'bili cid',
    cover       varchar(256)     null comment 'video cover',
    duration    bigint           not null comment 'video duration',
    like_time    bigint           not null comment 'video like time',

    season_id bigint  not null default 0 comment 'bili season id',
    intro       text      null comment 'video intro',
    title       text      null comment 'video title',
    type        bigint           not null comment 'video type',
    owner       json             null comment '{"mid": 173986740, "name": "这个月-"}',
    resp        json             null,
    cnt_info    json             null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'


) comment '点赞视频';
