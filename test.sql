-- Simplest table ever
create table User (
  name text not null,
  age integer not null,
  other_names text[] not null,
  is_alive boolean not null,
  birthdate date not null,
  dna bytea not null,
  metadata json not null
);
