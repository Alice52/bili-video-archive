create table archived_ups_tag
(
    id         varchar(64)   primary key comment 'tagid',
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    name         varchar(64)  null comment 'name',
    `count`      varchar(128) null comment 'count',
    tip      varchar(128) null comment 'tip'
) comment '关注的UP主分组';
create table archived_ups
(
    id           bigint auto_increment  primary key,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    tag_id         varchar(64)  null comment 'my group',
    sign         varchar(2048)  null comment 'up desc',
    uname      varchar(128) null comment 'up name',
    mid      varchar(128) null comment 'up uid',
    `level`      varchar(3) null comment 'up level',
    `rank` varchar(30) null comment 'up rank',
    following varchar(30) null comment 'up following',
    follower      varchar(30) null comment 'up follower',
    `view`      varchar(30) null comment 'up view',
    likes      varchar(30) null comment 'up likes',
    video      varchar(30) null comment 'up video count'
)  comment '关注的UP';

create table archived_fav_folders(
    fid         varchar(64)    primary key comment 'bili folder id' ,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    mid         varchar(64)  null comment 'bili uid',
    media_count varchar(64) null comment 'media count',
    title         varchar(64)  null comment 'title'
) comment '收藏文件夹';
create table archived_fav(
    id           bigint auto_increment  primary key,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    fid         varchar(64)  null comment 'bili folder',
    vid         varchar(64)  null comment 'bili avid',
    cover varchar(256)  null comment 'video cover',
    duration bigint  null comment 'video duration',
    fav_time bigint not null  comment 'video favor time',
    intro   varchar(64)  null comment 'video intro',
    title   varchar(64)  null comment 'video title',
    type   varchar(64)  null comment 'video type',
    season   varchar(64)  null comment 'video season',
    upper_mid  json   null comment '{"mid": 173986740, "name": "这个月-"}',
    cnt_info json null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'
) comment '收藏视频';

create table archived_coin(
    id           bigint auto_increment  primary key,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    fid         varchar(64)  null comment 'bili folder',
    vid         varchar(64)  null comment 'bili avid',
    cover varchar(256)  null comment 'video cover',
    duration bigint  null comment 'video duration',
    coin_time bigint not null  comment 'video favor time',
    intro   varchar(64)  null comment 'video intro',
    title   varchar(64)  null comment 'video title',
    type   varchar(64)  null comment 'video type',
    season   varchar(64)  null comment 'video season',
    upper_mid  json   null comment '{"mid": 173986740, "name": "这个月-"}',
    cnt_info json null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'
) comment '投币视频';
create table archived_like(
    id           bigint auto_increment  primary key,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null,

    fid         varchar(64)  null comment 'bili folder',
    vid         varchar(64)  null comment 'bili avid',
    cover varchar(256)  null comment 'video cover',
    duration bigint  null comment 'video duration',
    coin_time bigint not null  comment 'video favor time',
    intro   varchar(64)  null comment 'video intro',
    title   varchar(64)  null comment 'video title',
    type   varchar(64)  null comment 'video type',
    season   varchar(64)  null comment 'video season',
    upper_mid  json   null comment '{"mid": 173986740, "name": "这个月-"}',
    cnt_info json null comment '{"collect": 73600, "play": 1068474, "danmaku": 2632, "vt": 0, "play_switch": 0, "reply": 0, "view_text_1": "106.8万" }'
) comment '点赞视频';

create table archived_view_history(
    id         bigint  primary key,
    create_time   bigint  null,
    update_time   bigint null,
    delete_time   bigint  null


) comment '浏览历史记录';
