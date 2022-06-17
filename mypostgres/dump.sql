BEGIN;

SET client_encoding = 'LATIN1';

CREATE TABLE user (
    id integer NOT NULL,
    name text,
    tgusername text,
    tgid integer,
    tinkofftoken text
);