DROP TABLE IF EXISTS strict_types_ref;
CREATE TABLE strict_types_ref
(
    strict_types_ref integer UNIQUE
) STRICT;

DROP TABLE IF EXISTS strict_types;
CREATE TABLE strict_types
(
    strict_id        integer PRIMARY KEY,
    strict_int       integer,
    strict_real      real,
    strict_text      text,
    strict_blob      blob,
    strict_any       any,
    strict_nn        text NOT NULL,
    strict_unique    real UNIQUE,
    strict_check     integer CHECK (strict_check > 0),
    strict_ref       integer REFERENCES strict_types_ref (strict_types_ref),
    strict_def_const text DEFAULT 'strict'
) STRICT;
