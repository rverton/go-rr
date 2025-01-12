-- name: AuthorsList :many
select * from authors;

-- name: AuthorsInsert :one
insert into authors (name) values (:name) returning *;
