CREATE TABLE link
(
    id         BIGSERIAL PRIMARY KEY,
    link       TEXT      NOT NULL UNIQUE,
    short_link TEXT      NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL

);