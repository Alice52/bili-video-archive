-- auto-generated definition
create table archived_coin
(
    bvid        varchar(64)      not null primary key,
    create_time datetime(3)      null,
    update_time datetime(3)      null,
    delete_time datetime(3)      null,
    aid         bigint           not null comment 'bili aid',
    cid         bigint           not null comment 'bili cid',
    cover       varchar(256)     null comment 'video cover',
    duration    bigint           not null comment 'video duration',
   coined_time   bigint           not null comment 'video coin time',
    season_id   bigint default 0 not null comment 'bili season id',
    intro       text             null comment 'video intro',
    title       text             null comment 'video title',
    `type`        bigint           not null comment 'video type',
    coins bigint              not null comment 'video coins',
    `owner`       json             null comment '{"mid": 173986740, "name": "这个月-"}',
    resp        json             null,
    cnt_info    json             null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'
)
    comment '投币视频';

