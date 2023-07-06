CREATE TABLE users
(
    id         bigserial,
    email      varchar NOT NULL,
    full_name  varchar NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    CONSTRAINT users_pk PRIMARY KEY (id),
    CONSTRAINT users_unq_email UNIQUE (email)
);
