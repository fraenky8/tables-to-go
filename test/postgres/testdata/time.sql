DROP TABLE IF EXISTS time_ref CASCADE;
CREATE TABLE time_ref (
    time_ref time UNIQUE
);

DROP TABLE IF EXISTS time;
CREATE TABLE time (
    time time,
    time_nn time NOT NULL,
    time_nn_unique time NOT NULL UNIQUE,
    time_nn_check time NOT NULL CHECK ( time > '12:34:56' ),
    time_nn_ref time NOT NULL REFERENCES time_ref(time_ref),
    time_nn_def_const time NOT NULL DEFAULT '12:34:56',
    time_nn_def_func time NOT NULL DEFAULT now(),
    time_nn_unique_check time NOT NULL UNIQUE CHECK ( time > '12:34:56' ),

    time_unique time UNIQUE,
    time_unique_check time UNIQUE CHECK ( time > '12:34:56' ),
    time_unique_ref time UNIQUE REFERENCES time_ref(time_ref),
    time_unique_def_const time UNIQUE DEFAULT '12:34:56',
    time_unique_def_func time UNIQUE DEFAULT now(),

    time_check time CHECK ( time > '12:34:56' ),
    time_check_ref time CHECK ( time > '12:34:56' ) REFERENCES time_ref(time_ref),
    time_check_def_const time CHECK ( time > '12:34:56' ) DEFAULT '12:34:56',
    time_check_def_func time CHECK ( time > '12:34:56' ) DEFAULT now(),

    time_ref time REFERENCES time_ref(time_ref),
    time_ref_def_const time REFERENCES time_ref(time_ref) DEFAULT '12:34:56',
    time_ref_def_func time REFERENCES time_ref(time_ref) DEFAULT now(),
    time_ref_unique_check time UNIQUE CHECK ( time > '12:34:56' ) REFERENCES time_ref(time_ref),

    time_def_const time DEFAULT '12:34:56',
    time_def_const_unique_check time UNIQUE CHECK ( time > '12:34:56' )DEFAULT '12:34:56',

    time_def_func time DEFAULT now(),
    time_def_func_unique_check time UNIQUE CHECK ( time > '12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS time_pk;
CREATE TABLE time_pk (
    time_pk time PRIMARY KEY
);

DROP TABLE IF EXISTS time_pk_ref;
CREATE TABLE time_pk_ref (
    time_pk_ref time PRIMARY KEY REFERENCES time_ref(time_ref)
);

DROP TABLE IF EXISTS time_pk_def_const;
CREATE TABLE time_pk_def_const (
    time_pk_def_const time PRIMARY KEY DEFAULT '12:34:56'
);

DROP TABLE IF EXISTS time_pk_def_func;
CREATE TABLE time_pk_def_func (
    time_pk_def_func time PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS time_nn_pk;
CREATE TABLE time_nn_pk (
    time_nn_pk time NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS time_nn_unique_check_pk;
CREATE TABLE time_nn_unique_check_pk (
    time_nn_unique_check_pk time PRIMARY KEY NOT NULL UNIQUE CHECK ( time_nn_unique_check_pk > '12:34:56' )
);

DROP TABLE IF EXISTS time_nn_unique_check_pk_ref;
CREATE TABLE time_nn_unique_check_pk_ref (
    time_nn_unique_check_pk_ref time PRIMARY KEY NOT NULL UNIQUE CHECK ( time_nn_unique_check_pk_ref > '12:34:56' ) REFERENCES time_ref(time_ref)
);

DROP TABLE IF EXISTS time_unique_pk;
CREATE TABLE time_unique_pk (
    time_unique_pk time PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS time_unique_check_pk;
CREATE TABLE time_unique_check_pk (
    time_unique_check_pk time PRIMARY KEY UNIQUE CHECK ( time_unique_check_pk > '12:34:56' )
);

DROP TABLE IF EXISTS time_unique_check_pk_ref;
CREATE TABLE time_unique_check_pk_ref (
    time_unique_check_pk_ref time PRIMARY KEY UNIQUE CHECK ( time_unique_check_pk_ref > '12:34:56' ) REFERENCES time_ref(time_ref)
);

DROP TABLE IF EXISTS time_check_pk;
CREATE TABLE time_check_pk (
    time_check_pk time PRIMARY KEY CHECK ( time_check_pk > '12:34:56' )
);

DROP TABLE IF EXISTS time_def_const_unique_check_pk;
CREATE TABLE time_def_const_unique_check_pk (
    time_def_const_unique_check_pk time PRIMARY KEY UNIQUE CHECK ( time_def_const_unique_check_pk > '12:34:56' ) DEFAULT '12:34:56'
);

DROP TABLE IF EXISTS time_def_const_unique_check_pk_ref;
CREATE TABLE time_def_const_unique_check_pk_ref (
    time_def_const_unique_check_pk_ref time PRIMARY KEY UNIQUE CHECK ( time_def_const_unique_check_pk_ref > '12:34:56' ) DEFAULT '12:34:56' REFERENCES time_ref(time_ref)
);

DROP TABLE IF EXISTS time_def_func_unique_check_pk;
CREATE TABLE time_def_func_unique_check_pk (
    time_def_func_unique_check_pk time PRIMARY KEY UNIQUE CHECK ( time_def_func_unique_check_pk > '12:34:56' ) DEFAULT now()
);

DROP TABLE IF EXISTS time_def_func_unique_check_pk_ref;
CREATE TABLE time_def_func_unique_check_pk_ref (
    time_def_func_unique_check_pk_ref time PRIMARY KEY UNIQUE CHECK ( time_def_func_unique_check_pk_ref > '12:34:56' ) DEFAULT now() REFERENCES time_ref(time_ref)
);
