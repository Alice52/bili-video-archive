create table archived_coin
(
    id          bigint primary key auto_increment,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    fid         varchar(64) null comment 'bili folder',
    vid         varchar(64) null comment 'bili avid',
    cover       varchar(256) null comment 'video cover',
    duration    bigint null comment 'video duration',
    coin_time   bigint not null comment 'video favor time',
    intro       varchar(64) null comment 'video intro',
    title       varchar(64) null comment 'video title',
    `type`      varchar(64) null comment 'video type',
    season      varchar(64) null comment 'video season',
    upper_mid   json null comment '{"mid": 173986740, "name": "这个月-"}',
    cnt_info    json null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'
) comment '投币视频';