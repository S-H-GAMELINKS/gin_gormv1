create table users(id integer primary_key, name text not null);
create table tweets(id integer primary_key, content text not null, user_id integer);
