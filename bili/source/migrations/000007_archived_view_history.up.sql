create table archived_view_history
(
    bvid       varchar(64)       primary key ,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null,


    resp        json             null
) comment '浏览历史记录';
