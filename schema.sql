create table users (
    user_id serial primary key,
    first_name varchar(100) not null,
    last_name varchar(100) not null,
    email varchar(150) not null unique,
    password varchar(150) not null,
    registered_at timestamp not null,
    last_logged_in timestamp
);

create table definitions (
    definition_id serial primary key,
    from_lang int references user_languages(u_lang_id) not null,
    new_word text not null,
    to_lang int references user_languages(u_lang_id) not null,
    meaning text not null,
    notebook int references notebooks(notebook_id),
    added_at timestamp not null
);

create table user_languages (
    u_lang_id serial primary key,
    lang varchar(100) not null,
    lang_abbr varchar(3) not null,
    user_ int references users (user_id) not null,
);

create table notebooks (
    notebook_id serial primary key,
    notebook_name varchar(150) not null,
    owner int references users (user_id),
    is_public boolean default 'false',
    created_at timestamp not null
);

create table user_favourite_definitions (
    user_ int references users(user_id) not null,
    definition int references definitions(definition_id) not null,
    primary key (user_, definition)
);