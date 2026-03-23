DROP TABLE IF EXISTS binary_ref CASCADE;
CREATE TABLE binary_ref
(
    binary_ref binary UNIQUE
);

DROP TABLE IF EXISTS binary_table;
CREATE TABLE binary_table
(
    col                           binary,

    binary_cap                    binary(255),
    binary_nn                     binary NOT NULL,
    binary_nn_unique              binary NOT NULL UNIQUE,
    binary_nn_check_cmp           binary NOT NULL,
    binary_nn_check_fn            binary NOT NULL,
    binary_nn_ref                 binary NOT NULL REFERENCES binary_ref (binary_ref),
    binary_nn_def_const           binary NOT NULL DEFAULT '4',
    binary_nn_def_func            binary NOT NULL DEFAULT '5',
    binary_nn_unique_check        binary NOT NULL UNIQUE,

    binary_unique                 binary UNIQUE,
    binary_unique_check           binary UNIQUE,
    binary_unique_ref             binary UNIQUE REFERENCES binary_ref (binary_ref),
    binary_unique_def_const       binary UNIQUE   DEFAULT '4',
    binary_unique_def_func        binary UNIQUE   DEFAULT '5',

    binary_check                  binary,
    binary_check_ref              binary REFERENCES binary_ref (binary_ref),
    binary_check_def_const        binary          DEFAULT '4',
    binary_check_def_func         binary          DEFAULT '5',

    binary_ref                    binary REFERENCES binary_ref (binary_ref),
    binary_ref_unique_check       binary UNIQUE REFERENCES binary_ref (binary_ref),

    binary_def_const              binary          DEFAULT '4',
    binary_def_const_unique_check binary UNIQUE   DEFAULT '4',

    binary_def_func               binary          DEFAULT '5',
    binary_def_func_unique_check  binary UNIQUE   DEFAULT '5'
);

DROP TABLE IF EXISTS binary_pk;
CREATE TABLE binary_pk
(
    binary_pk binary PRIMARY KEY
);

DROP TABLE IF EXISTS binary_pk_ref;
CREATE TABLE binary_pk_ref
(
    binary_pk_ref binary PRIMARY KEY REFERENCES binary_ref (binary_ref)
);

DROP TABLE IF EXISTS binary_pk_def_const;
CREATE TABLE binary_pk_def_const
(
    binary_pk_def_const binary PRIMARY KEY DEFAULT '4'
);

DROP TABLE IF EXISTS binary_pk_def_func;
CREATE TABLE binary_pk_def_func
(
    binary_pk_def_func binary PRIMARY KEY DEFAULT '5'
);

DROP TABLE IF EXISTS binary_nn_pk;
CREATE TABLE binary_nn_pk
(
    binary_nn_pk binary NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS binary_nn_unique_check_pk;
CREATE TABLE binary_nn_unique_check_pk
(
    binary_nn_unique_check_pk binary PRIMARY KEY NOT NULL UNIQUE
);

DROP TABLE IF EXISTS binary_nn_unique_check_pk_ref;
CREATE TABLE binary_nn_unique_check_pk_ref
(
    binary_nn_unique_check_pk_ref binary PRIMARY KEY NOT NULL UNIQUE REFERENCES binary_ref (binary_ref)
);

DROP TABLE IF EXISTS binary_unique_pk;
CREATE TABLE binary_unique_pk
(
    binary_unique_pk binary PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS binary_unique_check_pk;
CREATE TABLE binary_unique_check_pk
(
    binary_unique_check_pk binary PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS binary_unique_check_pk_ref;
CREATE TABLE binary_unique_check_pk_ref
(
    binary_unique_check_pk_ref binary PRIMARY KEY UNIQUE REFERENCES binary_ref (binary_ref)
);

DROP TABLE IF EXISTS binary_check_pk;
CREATE TABLE binary_check_pk
(
    binary_check_pk binary PRIMARY KEY
);

DROP TABLE IF EXISTS binary_def_const_unique_check_pk;
CREATE TABLE binary_def_const_unique_check_pk
(
    binary_def_const_unique_check_pk binary PRIMARY KEY UNIQUE DEFAULT '4'
);

DROP TABLE IF EXISTS binary_def_const_unique_check_pk_ref;
CREATE TABLE binary_def_const_unique_check_pk_ref
(
    binary_def_const_unique_check_pk_ref binary PRIMARY KEY UNIQUE DEFAULT '4' REFERENCES binary_ref (binary_ref)
);

DROP TABLE IF EXISTS binary_def_func_unique_check_pk;
CREATE TABLE binary_def_func_unique_check_pk
(
    binary_def_func_unique_check_pk binary PRIMARY KEY UNIQUE DEFAULT '5'
);

DROP TABLE IF EXISTS binary_def_func_unique_check_pk_ref;
CREATE TABLE binary_def_func_unique_check_pk_ref
(
    binary_def_func_unique_check_pk_ref binary PRIMARY KEY UNIQUE DEFAULT '5' REFERENCES binary_ref (binary_ref)
);
