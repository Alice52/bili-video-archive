create table archived_season
(
    season_id        bigint   primary key default 0 ,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    cover       varchar(256)     null comment 'season cover',
    title        varchar(256)     null comment 'season title',
    intro       text      null comment 'season evaluate',
    episodes_count int           null comment 'season episodes count',
    link          varchar(256)     null comment 'season link',
    media_id      bigint        not   null default 0 comment 'season media id',
    ctime       bigint           not null  default 0  comment 'video create time',
    pubdate      bigint           not null default 0  comment 'video publish time',
    rating_count   bigint           not null default 0  comment 'video rating count',
    rating_score decimal               not null default 0  comment 'video rating score',

    stat_view    bigint           not null default 0  comment 'video view count',
    stat_danmaku  bigint           not null default 0  comment 'video danmaku count',
    stat_reply   bigint           not null default 0  comment 'video reply count',
    stat_favorite bigint           not null default 0  comment 'video favorite count',
    stat_coin     bigint           not null default 0  comment 'video coin count',
    stat_share    bigint           not null default 0  comment 'video share count',
    stat_like     bigint           not null default 0  comment 'video like count',
    stat_dislike  bigint           not null default 0  comment 'video dislike count',
    stat_now_rank  bigint           not null default 0  comment 'video now rank',
    stat_his_rank  bigint           not null default 0  comment 'video history rank',
    stat_evaluation varchar(32)      not null comment 'video evaluation',
    stat_vt  bigint           not null default 0  comment 'video vt',

    styles json      not null comment 'video styles',

    upper_mid   bigint           not null default 0  comment 'video upper mid',
    upper_name  varchar(256)      not null default 0  comment 'video upper name',
    face_name  varchar(256)      not null default 0  comment 'video upper face',

    resp        json             null
) comment '视频信息';
