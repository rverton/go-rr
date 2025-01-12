create table if not exists authors (
    id integer primary key, -- alias for rowid
    name text not null,

    created_at integer not null default (strftime('%s', 'now'))
) strict;
