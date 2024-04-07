CREATE TABLE users
(
    id bigserial NOT NULL,
    login character varying(256) NOT NULL,
    password_hash character varying(2048) NOT NULL,
    points_balance numeric NOT NULL default 0,
    PRIMARY KEY (id)
);

ALTER TABLE users
    ADD CONSTRAINT users_login_unq UNIQUE (login);
