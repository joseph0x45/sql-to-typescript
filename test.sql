-- Simplest table ever
create table User IF NOT EXISTS SOMETHING (
  name text not null,
  age integer not null,
  /*
  yeah you know what Im saying
  */
  other_names text[] not null,
  is_alive boolean not null,
  birthdate date not null,
  dna bytea not null,
  metadata json not null
);
