DROP TABLE IF EXISTS integer_ref;
CREATE TABLE integer_ref
(
    integer_ref integer UNIQUE
);

DROP TABLE IF EXISTS integer_table;
CREATE TABLE integer_table
(
    i                 integer,
    integer_nn        integer NOT NULL,
    integer_unique    integer UNIQUE,
    integer_check     integer CHECK (integer_check > 0),
    integer_ref       integer REFERENCES integer_ref (integer_ref),
    integer_def_const integer DEFAULT 42,
    integer_pk        integer PRIMARY KEY
);

DROP TABLE IF EXISTS integer_nn_pk;
CREATE TABLE integer_nn_pk
(
    integer_nn_pk integer NOT NULL PRIMARY KEY
);
