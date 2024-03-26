create table archived_view_history
(
    id          bigint primary key auto_increment,
    create_time datetime(3) null,
    update_time datetime(3) null,
    delete_time datetime(3) null
) comment '浏览历史记录';
