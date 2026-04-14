DROP TABLE IF EXISTS char_ref CASCADE;
CREATE TABLE char_ref
(
    char_ref char UNIQUE
);

DROP TABLE IF EXISTS char_table;
CREATE TABLE char_table
(
    col                         char,

    char_cap                    char(255),
    char_nn                     char NOT NULL,
    char_nn_unique              char NOT NULL UNIQUE,
    char_nn_check_cmp           char NOT NULL,
    char_nn_check_fn            char NOT NULL,
    char_nn_ref                 char NOT NULL REFERENCES char_ref (char_ref),
    char_nn_def_const           char NOT NULL DEFAULT '4',
    char_nn_def_func            char NOT NULL DEFAULT '5',
    char_nn_unique_check        char NOT NULL UNIQUE,

    char_unique                 char UNIQUE,
    char_unique_check           char UNIQUE,
    char_unique_ref             char UNIQUE REFERENCES char_ref (char_ref),
    char_unique_def_const       char UNIQUE   DEFAULT '4',
    char_unique_def_func        char UNIQUE   DEFAULT '5',

    char_check                  char,
    char_check_ref              char REFERENCES char_ref (char_ref),
    char_check_def_const        char          DEFAULT '4',
    char_check_def_func         char          DEFAULT '5',

    char_ref                    char REFERENCES char_ref (char_ref),
    char_ref_unique_check       char UNIQUE REFERENCES char_ref (char_ref),

    char_def_const              char          DEFAULT '4',
    char_def_const_unique_check char UNIQUE   DEFAULT '4',

    char_def_func               char          DEFAULT '5',
    char_def_func_unique_check  char UNIQUE   DEFAULT '5'
);

DROP TABLE IF EXISTS char_pk;
CREATE TABLE char_pk
(
    char_pk char PRIMARY KEY
);

DROP TABLE IF EXISTS char_pk_ref;
CREATE TABLE char_pk_ref
(
    char_pk_ref char PRIMARY KEY REFERENCES char_ref (char_ref)
);

DROP TABLE IF EXISTS char_pk_def_const;
CREATE TABLE char_pk_def_const
(
    char_pk_def_const char PRIMARY KEY DEFAULT '4'
);

DROP TABLE IF EXISTS char_pk_def_func;
CREATE TABLE char_pk_def_func
(
    char_pk_def_func char PRIMARY KEY DEFAULT '5'
);

DROP TABLE IF EXISTS char_nn_pk;
CREATE TABLE char_nn_pk
(
    char_nn_pk char NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS char_nn_unique_check_pk;
CREATE TABLE char_nn_unique_check_pk
(
    char_nn_unique_check_pk char PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS char_nn_unique_check_pk_ref;
CREATE TABLE char_nn_unique_check_pk_ref
(
    char_nn_unique_check_pk_ref char PRIMARY KEY NOT NULL UNIQUE REFERENCES char_ref (char_ref)
);

DROP TABLE IF EXISTS char_unique_pk;
CREATE TABLE char_unique_pk
(
    char_unique_pk char PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS char_unique_check_pk;
CREATE TABLE char_unique_check_pk
(
    char_unique_check_pk char PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS char_unique_check_pk_ref;
CREATE TABLE char_unique_check_pk_ref
(
    char_unique_check_pk_ref char PRIMARY KEY UNIQUE REFERENCES char_ref (char_ref)
);

DROP TABLE IF EXISTS char_check_pk;
CREATE TABLE char_check_pk
(
    char_check_pk char PRIMARY KEY
);

DROP TABLE IF EXISTS char_def_const_unique_check_pk;
CREATE TABLE char_def_const_unique_check_pk
(
    char_def_const_unique_check_pk char PRIMARY KEY UNIQUE DEFAULT '4'
);

DROP TABLE IF EXISTS char_def_const_unique_check_pk_ref;
CREATE TABLE char_def_const_unique_check_pk_ref
(
    char_def_const_unique_check_pk_ref char PRIMARY KEY UNIQUE DEFAULT '4' REFERENCES char_ref (char_ref)
);

DROP TABLE IF EXISTS char_def_func_unique_check_pk;
CREATE TABLE char_def_func_unique_check_pk
(
    char_def_func_unique_check_pk char PRIMARY KEY UNIQUE DEFAULT '5'
);

DROP TABLE IF EXISTS char_def_func_unique_check_pk_ref;
CREATE TABLE char_def_func_unique_check_pk_ref
(
    char_def_func_unique_check_pk_ref char PRIMARY KEY UNIQUE DEFAULT '5' REFERENCES char_ref (char_ref)
);
