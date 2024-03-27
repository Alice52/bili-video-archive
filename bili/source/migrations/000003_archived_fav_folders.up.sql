create table archived_fav_folders
(
    id bigint primary key,
    fid         bigint  not null comment 'bili folder id',
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    mid         bigint not null comment 'bili uid',
    media_count bigint not null comment 'media count',
    title       varchar(64) null comment 'title',
    resp    json null
) comment '收藏文件夹';