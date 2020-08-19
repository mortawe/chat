create table users
(
    id         serial                                 not null
        constraint user_pk
            primary key,
    username   varchar                                not null,
    created_at timestamp with time zone default now() not null
);

alter table users
    owner to "user";

create unique index user_id_uindex
    on users (id);

create unique index user_username_uindex
    on users (username);

create table chats
(
    id         serial                                 not null
        constraint chat_pk
            primary key,
    name       varchar,
    created_at timestamp with time zone default now() not null
);

alter table chats
    owner to "user";

create unique index chat_id_uindex
    on chats (id);

create unique index chat_name_uindex
    on chats (name);

create table messages
(
    id         serial                                 not null
        constraint messages_pk
            primary key,
    chat_id    integer                                not null
        constraint messages_chat_id_fk
            references chats
            on update cascade on delete cascade,
    author_id  integer                                not null
        constraint messages_user_id_fk
            references users
            on update cascade on delete cascade,
    text       varchar,
    created_at timestamp with time zone default now() not null
);

alter table messages
    owner to "user";

create unique index messages_id_uindex
    on messages (id);

create trigger add_users_to_chat
    after update
    on messages
    for each row
execute procedure add_users_to_chat_after_create_msg();

create table chat_users
(
    chat_id integer not null
        constraint chat_users_chats_id_fk
            references chats
            on update cascade on delete cascade,
    user_id integer not null
        constraint chat_users_users_id_fk
            references users
            on update cascade on delete cascade
);

alter table chat_users
    owner to "user";

create unique index chat_users_chat_id_user_id_uindex
    on chat_users (chat_id, user_id);

