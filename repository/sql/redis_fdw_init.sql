
CREATE EXTENSION IF NOT EXISTS redis_fdw;
CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_fdw OPTIONS (host 'redis', port '6379');
CREATE USER MAPPING FOR PUBLIC SERVER redis_server OPTIONS (password '');

--- all the object value is a json string.

--- user info is single, dont need zset to sort.
create foreign table if not exists rd_users (key text, value text) server redis_server options (database '0'); 
--- this function is for User table, every insert update will trigger this, (del is not support)
--- all interaction (except del) will update this key-value in redis
create or replace function UserInteraction() returns trigger as $trigger_user_interaction$
begin
  insert into rd_users values ('user_'||new.id, row_to_json(new));
  return new;
end;
$trigger_user_interaction$ language plpgsql;
create or replace trigger user_interaction after insert or update on users for each row execute function UserInteraction();



create foreign table if not exists rd_videos (key text, value text) server redis_server options (database '0'); 
 --- use the create time as score
create foreign table if not exists rd_videos_timeseq (score bigint, member text, index int) server redis_server options(database '1', tabletype 'zset', key 'video_time_seq');
--- use create time as score, key 'user_videos_{user_id}'
create foreign table if not exists rd_user_videos (key text, member text, score bigint, index int) server redis_server options (database '1', tabletype 'zset'); 

--- trigger for video insert and update.
create or replace function VideoInsert() returns trigger as $trigger_video_insert$
begin 
  insert into rd_videos values ('video_'||new.id, row_to_json(new));
  insert into rd_videos_timeseq values (new.created_at, ''||new.id, 0);
  insert into rd_user_videos values ('user_videos_'||new.user_id, new.id, new.created_at, 0);
  return new;
end;
$trigger_video_insert$ language plpgsql;
create or replace function VideoUpdate() returns trigger as $trigger_video_update$
begin 
  insert into rd_videos values ('video_'||new.id, row_to_json(new));
  return new;
end;
$trigger_video_update$ language plpgsql;
create or replace trigger video_insert after insert on videos for each row execute function VideoInsert();
create or replace trigger video_update after update on videos for each row execute function VideoUpdate();


---! due to the rw-redis_fdw api, the delete operation cant be execute.
--- No support for update (app dont have this function)
-- create foreign table rd_comment (key text, value text) server redis_server options (database '1');
--- use the 'video_comment_{videos_id}' as the key, use create time as score
create foreign table if not exists rd_video_comment (key text, member text, score bigint, index int) server redis_server options (database '1', tabletype 'zset'); 
create or replace function OnInsertComment() returns trigger as $trigger_comment_insert$
begin
--   insert into rd_comment values ('comment_'||new.id, row_to_json(new));
  insert into rd_video_comment values ('video_comment_'||new.video_id, row_to_json(new), new.created_at, 0);
  return new;
end;
$trigger_comment_insert$ language plpgsql;
create or replace trigger video_comment_add after insert on comments for each row execute function OnInsertComment();

---! due to the rw-redis_fdw api, the delete operation cant be execute.
--- use create time as score, key 'user_favorite_{user_id}'
create foreign table if not exists rd_user_favorite (key text, member text, score bigint, index int, expiry bigint) server redis_server options (database '0', tabletype 'zset'); 
create or replace function DoFavorite() returns trigger as $trigger_do_favorite$
begin 
  insert into rd_user_favorite values ('user_favorite_'||new.user_id, ''||new.video_id, new.created_at, 0);
  return new;
end;
$trigger_do_favorite$ language plpgsql;
create or replace trigger user_do_favorite after insert on favorites for each row execute function DoFavorite();


--- use create time as score, key 'chat_{chat-key}'
create foreign table if not exists rd_user_chat (key text, member text, score bigint, index int, expiry bigint) server redis_server options (database '0', tabletype 'zset');
create or replace function ChatAdd() returns trigger as $chat_add$
begin
  insert into rd_user_chat values ('chat_'||new.chat_key, row_to_json(new), new.created_at);
  return new;
end;
$chat_add$ language plpgsql;
create or replace trigger chat_add after insert on chat_records for each row execute function ChatAdd();

---! due to the rw-redis_fdw api, the delete operation cant be execute.
--- use create time as score, key 'user_follow_{user_id}'
create foreign table if not exists rd_user_follow (key text, member text, score bigint, index int, expiry bigint) server redis_server options (database '0', tabletype 'zset');
--- use create time as score, key 'user_follower_{user_id}'
create foreign table if not exists rd_user_follower (key text, member text, score bigint, index int, expiry bigint) server redis_server options (database '0', tabletype 'zset');
create or replace function DoFollow() returns trigger as $do_follow$
begin 
  insert into rd_user_follow values ('user_follow_'||new.from_id, ''||new.to_id, new.created_at);
  insert into rd_user_follower values ('user_follower_'||new.to_id, ''||new.from_id, new.created_at);
  return new;
end;
$do_follow$ language plpgsql;
create or replace trigger do_follow after insert on relations for each row execute function DoFollow();




