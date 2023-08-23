-- Simplest table ever
/*
some multi line comment
should be ignored
*/
create table User IF NOT EXISTS (
  id uuid not null primary key,
  name text not null,
  age integer not null,
  other_names text[] not null,
  is_alive boolean not null,
  birthdate date not null,
  metadata json 
);

create table test (
  name text not null,
  age integer
);
