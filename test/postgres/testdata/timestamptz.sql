DROP TABLE IF EXISTS timestamptz_ref CASCADE;
CREATE TABLE timestamptz_ref (
    timestamptz_ref timestamptz UNIQUE
);

DROP TABLE IF EXISTS timestamptz;
CREATE TABLE timestamptz (
    timestamptz timestamptz,
    timestamptz_nn timestamptz NOT NULL,
    timestamptz_nn_unique timestamptz NOT NULL UNIQUE,
    timestamptz_nn_check timestamptz NOT NULL CHECK ( timestamptz > '2020-03-01 12:34:56+8' ),
    timestamptz_nn_ref timestamptz NOT NULL REFERENCES timestamptz_ref(timestamptz_ref),
    timestamptz_nn_def_const timestamptz NOT NULL DEFAULT '2020-03-01 12:34:56+8',
    timestamptz_nn_def_func timestamptz NOT NULL DEFAULT now(),
    timestamptz_nn_unique_check timestamptz NOT NULL UNIQUE CHECK ( timestamptz > '2020-03-01 12:34:56+8' ),

    timestamptz_unique timestamptz UNIQUE,
    timestamptz_unique_check timestamptz UNIQUE CHECK ( timestamptz > '2020-03-01 12:34:56+8' ),
    timestamptz_unique_ref timestamptz UNIQUE REFERENCES timestamptz_ref(timestamptz_ref),
    timestamptz_unique_def_const timestamptz UNIQUE DEFAULT '2020-03-01 12:34:56+8',
    timestamptz_unique_def_func timestamptz UNIQUE DEFAULT now(),

    timestamptz_check timestamptz CHECK ( timestamptz > '2020-03-01 12:34:56+8' ),
    timestamptz_check_ref timestamptz CHECK ( timestamptz > '2020-03-01 12:34:56+8' ) REFERENCES timestamptz_ref(timestamptz_ref),
    timestamptz_check_def_const timestamptz CHECK ( timestamptz > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8',
    timestamptz_check_def_func timestamptz CHECK ( timestamptz > '2020-03-01 12:34:56+8' ) DEFAULT now(),

    timestamptz_ref timestamptz REFERENCES timestamptz_ref(timestamptz_ref),
    timestamptz_ref_def_const timestamptz REFERENCES timestamptz_ref(timestamptz_ref) DEFAULT '2020-03-01 12:34:56+8',
    timestamptz_ref_def_func timestamptz REFERENCES timestamptz_ref(timestamptz_ref) DEFAULT now(),
    timestamptz_ref_unique_check timestamptz UNIQUE CHECK ( timestamptz > '2020-03-01 12:34:56+8' ) REFERENCES timestamptz_ref(timestamptz_ref),

    timestamptz_def_const timestamptz DEFAULT '2020-03-01 12:34:56+8',
    timestamptz_def_const_unique_check timestamptz UNIQUE CHECK ( timestamptz > '2020-03-01 12:34:56+8' )DEFAULT '2020-03-01 12:34:56+8',

    timestamptz_def_func timestamptz DEFAULT now(),
    timestamptz_def_func_unique_check timestamptz UNIQUE CHECK ( timestamptz > '2020-03-01 12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamptz_pk;
CREATE TABLE timestamptz_pk (
    timestamptz_pk timestamptz PRIMARY KEY
);

DROP TABLE IF EXISTS timestamptz_pk_ref;
CREATE TABLE timestamptz_pk_ref (
    timestamptz_pk_ref timestamptz PRIMARY KEY REFERENCES timestamptz_ref(timestamptz_ref)
);

DROP TABLE IF EXISTS timestamptz_pk_def_const;
CREATE TABLE timestamptz_pk_def_const (
    timestamptz_pk_def_const timestamptz PRIMARY KEY DEFAULT '2020-03-01 12:34:56+8'
);

DROP TABLE IF EXISTS timestamptz_pk_def_func;
CREATE TABLE timestamptz_pk_def_func (
    timestamptz_pk_def_func timestamptz PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS timestamptz_nn_pk;
CREATE TABLE timestamptz_nn_pk (
    timestamptz_nn_pk timestamptz NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS timestamptz_nn_unique_check_pk;
CREATE TABLE timestamptz_nn_unique_check_pk (
    timestamptz_nn_unique_check_pk timestamptz PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamptz_nn_unique_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamptz_nn_unique_check_pk_ref;
CREATE TABLE timestamptz_nn_unique_check_pk_ref (
    timestamptz_nn_unique_check_pk_ref timestamptz PRIMARY KEY NOT NULL UNIQUE CHECK ( timestamptz_nn_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) REFERENCES timestamptz_ref(timestamptz_ref)
);

DROP TABLE IF EXISTS timestamptz_unique_pk;
CREATE TABLE timestamptz_unique_pk (
    timestamptz_unique_pk timestamptz PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS timestamptz_unique_check_pk;
CREATE TABLE timestamptz_unique_check_pk (
    timestamptz_unique_check_pk timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_unique_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamptz_unique_check_pk_ref;
CREATE TABLE timestamptz_unique_check_pk_ref (
    timestamptz_unique_check_pk_ref timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) REFERENCES timestamptz_ref(timestamptz_ref)
);

DROP TABLE IF EXISTS timestamptz_check_pk;
CREATE TABLE timestamptz_check_pk (
    timestamptz_check_pk timestamptz PRIMARY KEY CHECK ( timestamptz_check_pk > '2020-03-01 12:34:56+8' )
);

DROP TABLE IF EXISTS timestamptz_def_const_unique_check_pk;
CREATE TABLE timestamptz_def_const_unique_check_pk (
    timestamptz_def_const_unique_check_pk timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_def_const_unique_check_pk > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8'
);

DROP TABLE IF EXISTS timestamptz_def_const_unique_check_pk_ref;
CREATE TABLE timestamptz_def_const_unique_check_pk_ref (
    timestamptz_def_const_unique_check_pk_ref timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_def_const_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) DEFAULT '2020-03-01 12:34:56+8' REFERENCES timestamptz_ref(timestamptz_ref)
);

DROP TABLE IF EXISTS timestamptz_def_func_unique_check_pk;
CREATE TABLE timestamptz_def_func_unique_check_pk (
    timestamptz_def_func_unique_check_pk timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_def_func_unique_check_pk > '2020-03-01 12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS timestamptz_def_func_unique_check_pk_ref;
CREATE TABLE timestamptz_def_func_unique_check_pk_ref (
    timestamptz_def_func_unique_check_pk_ref timestamptz PRIMARY KEY UNIQUE CHECK ( timestamptz_def_func_unique_check_pk_ref > '2020-03-01 12:34:56+8' ) DEFAULT now() REFERENCES timestamptz_ref(timestamptz_ref)
);
