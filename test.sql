-- Simplest table ever
/*
some multi line comment
should be ignored
*/
create table User IF NOT EXISTS (
  name text not null,
  age integer not null,
  other_names text[] not null,
  is_alive boolean not null,
  birthdate date not null,
  dna bytea not null,
  metadata json not null
);
