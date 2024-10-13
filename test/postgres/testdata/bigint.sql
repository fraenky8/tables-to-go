DROP TABLE IF EXISTS bigint_ref CASCADE;
CREATE TABLE bigint_ref (
    bigint_ref bigint UNIQUE
);

DROP TABLE IF EXISTS bigint;
CREATE TABLE bigint (
    bigint bigint,
    bigint_nn bigint NOT NULL,
    bigint_nn_unique bigint NOT NULL UNIQUE,
    bigint_nn_check bigint NOT NULL CHECK ( bigint > 0 ),
    bigint_nn_ref bigint NOT NULL REFERENCES bigint_ref(bigint_ref),
    bigint_nn_def_const bigint NOT NULL DEFAULT 42,
    bigint_nn_def_func bigint NOT NULL DEFAULT pi(),
    bigint_nn_unique_check bigint NOT NULL UNIQUE CHECK ( bigint > 0 ),

    bigint_unique bigint UNIQUE,
    bigint_unique_check bigint UNIQUE CHECK ( bigint > 0 ),
    bigint_unique_ref bigint UNIQUE REFERENCES bigint_ref(bigint_ref),
    bigint_unique_def_const bigint UNIQUE DEFAULT 42,
    bigint_unique_def_func bigint UNIQUE DEFAULT pi(),

    bigint_check bigint CHECK ( bigint > 0 ),
    bigint_check_ref bigint CHECK ( bigint > 0 ) REFERENCES bigint_ref(bigint_ref),
    bigint_check_def_const bigint CHECK ( bigint > 0 ) DEFAULT 42,
    bigint_check_def_func bigint CHECK ( bigint > 0 ) DEFAULT pi(),

    bigint_ref bigint REFERENCES bigint_ref(bigint_ref),
    bigint_ref_def_const bigint REFERENCES bigint_ref(bigint_ref) DEFAULT 42,
    bigint_ref_def_func bigint REFERENCES bigint_ref(bigint_ref) DEFAULT pi(),
    bigint_ref_unique_check bigint UNIQUE CHECK ( bigint > 0 ) REFERENCES bigint_ref(bigint_ref),

    bigint_def_const bigint DEFAULT 42,
    bigint_def_const_unique_check bigint UNIQUE CHECK ( bigint > 0 )DEFAULT 42,

    bigint_def_func bigint DEFAULT pi(),
    bigint_def_func_unique_check bigint UNIQUE CHECK ( bigint > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS bigint_pk;
CREATE TABLE bigint_pk (
    bigint_pk bigint PRIMARY KEY
);

DROP TABLE IF EXISTS bigint_pk_ref;
CREATE TABLE bigint_pk_ref (
    bigint_pk_ref bigint PRIMARY KEY REFERENCES bigint_ref(bigint_ref)
);

DROP TABLE IF EXISTS bigint_pk_def_const;
CREATE TABLE bigint_pk_def_const (
    bigint_pk_def_const bigint PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS bigint_pk_def_func;
CREATE TABLE bigint_pk_def_func (
    bigint_pk_def_func bigint PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS bigint_nn_pk;
CREATE TABLE bigint_nn_pk (
    bigint_nn_pk bigint NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS bigint_nn_unique_check_pk;
CREATE TABLE bigint_nn_unique_check_pk (
    bigint_nn_unique_check_pk bigint PRIMARY KEY NOT NULL UNIQUE CHECK ( bigint_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS bigint_nn_unique_check_pk_ref;
CREATE TABLE bigint_nn_unique_check_pk_ref (
    bigint_nn_unique_check_pk_ref bigint PRIMARY KEY NOT NULL UNIQUE CHECK ( bigint_nn_unique_check_pk_ref > 0) REFERENCES bigint_ref(bigint_ref)
);

DROP TABLE IF EXISTS bigint_unique_pk;
CREATE TABLE bigint_unique_pk (
    bigint_unique_pk bigint PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS bigint_unique_check_pk;
CREATE TABLE bigint_unique_check_pk (
    bigint_unique_check_pk bigint PRIMARY KEY UNIQUE CHECK ( bigint_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS bigint_unique_check_pk_ref;
CREATE TABLE bigint_unique_check_pk_ref (
    bigint_unique_check_pk_ref bigint PRIMARY KEY UNIQUE CHECK ( bigint_unique_check_pk_ref > 0) REFERENCES bigint_ref(bigint_ref)
);

DROP TABLE IF EXISTS bigint_check_pk;
CREATE TABLE bigint_check_pk (
    bigint_check_pk bigint PRIMARY KEY CHECK ( bigint_check_pk > 0 )
);

DROP TABLE IF EXISTS bigint_def_const_unique_check_pk;
CREATE TABLE bigint_def_const_unique_check_pk (
    bigint_def_const_unique_check_pk bigint PRIMARY KEY UNIQUE CHECK ( bigint_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS bigint_def_const_unique_check_pk_ref;
CREATE TABLE bigint_def_const_unique_check_pk_ref (
    bigint_def_const_unique_check_pk_ref bigint PRIMARY KEY UNIQUE CHECK ( bigint_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES bigint_ref(bigint_ref)
);

DROP TABLE IF EXISTS bigint_def_func_unique_check_pk;
CREATE TABLE bigint_def_func_unique_check_pk (
    bigint_def_func_unique_check_pk bigint PRIMARY KEY UNIQUE CHECK ( bigint_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS bigint_def_func_unique_check_pk_ref;
CREATE TABLE bigint_def_func_unique_check_pk_ref (
    bigint_def_func_unique_check_pk_ref bigint PRIMARY KEY UNIQUE CHECK ( bigint_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES bigint_ref(bigint_ref)
);
