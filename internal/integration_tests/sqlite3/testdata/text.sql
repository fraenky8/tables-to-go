DROP TABLE IF EXISTS text_ref;
CREATE TABLE text_ref
(
    text_ref text UNIQUE
);

DROP TABLE IF EXISTS text_table;
CREATE TABLE text_table
(
    t              text,
    text_nn        text NOT NULL,
    text_unique    text UNIQUE,
    text_check     text CHECK (length(text_check) > 1),
    text_ref       text REFERENCES text_ref (text_ref),
    text_def_const text DEFAULT 'abc',
    text_pk        text PRIMARY KEY
);

DROP TABLE IF EXISTS text_nn_pk;
CREATE TABLE text_nn_pk
(
    text_nn_pk text NOT NULL PRIMARY KEY
);
