create table archived_fav_folders
(
    fid         varchar(64) primary key comment 'bili folder id',
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,

    mid         varchar(64) null comment 'bili uid',
    media_count varchar(64) null comment 'media count',
    title       varchar(64) null comment 'title'
) comment '收藏文件夹';