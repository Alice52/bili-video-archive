create table archived_video
(
    archived_type tinyint    not null default 0 comment '0: fav, 1: coin, 2: like, 3: view',
    sync_status tinyint    not null default 0 comment '0: 未同步, 1: 同步中, 2: 同步完成',
    sync_time    datetime(3) null comment '同步时间',

    bvid        varchar(64)         not null,
    season_id  bigint               not null default 0 comment 'bili season_id',
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    aid        bigint    not null default 0 comment 'bili aid',
    cid         bigint    not null default 0 comment 'bili cid',
    tid         bigint    not null default 0 comment 'bili tid',
    cover       varchar(256)     null comment 'video cover',
    ctime       bigint           not null  default 0  comment 'video create time',
    pubdate      bigint           not null default 0  comment 'video publish time',
    duration    bigint           not null default 0  comment 'video duration',
    title       text      null comment 'video title',
    intro       text      null comment 'video intro',

    upper_mid   bigint           not null default 0  comment 'video upper mid',
    upper_name  varchar(256)      not null default 0  comment 'video upper name',
    face_name  varchar(256)      not null default 0  comment 'video upper face',

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

    honor_reply json                   null comment 'video honor reply',
    resp        json             null,
    primary key (bvid, archived_type)
) comment '视频信息';
