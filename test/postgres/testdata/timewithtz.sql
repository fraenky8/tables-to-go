DROP TABLE IF EXISTS time_with_time_zone_ref CASCADE;
CREATE TABLE time_with_time_zone_ref (
    time_with_time_zone_ref time with time zone UNIQUE
);

DROP TABLE IF EXISTS time_with_time_zone;
CREATE TABLE time_with_time_zone (
    time_with_time_zone time with time zone,
    time_with_time_zone_nn time with time zone NOT NULL,
    time_with_time_zone_nn_unique time with time zone NOT NULL UNIQUE,
    time_with_time_zone_nn_check time with time zone NOT NULL CHECK ( time_with_time_zone > '12:34:56+8' ),
    time_with_time_zone_nn_ref time with time zone NOT NULL REFERENCES time_with_time_zone_ref(time_with_time_zone_ref),
    time_with_time_zone_nn_def_const time with time zone NOT NULL DEFAULT '12:34:56+8',
    time_with_time_zone_nn_def_func time with time zone NOT NULL DEFAULT now(),
    time_with_time_zone_nn_unique_check time with time zone NOT NULL UNIQUE CHECK ( time_with_time_zone > '12:34:56+8' ),

    time_with_time_zone_unique time with time zone UNIQUE,
    time_with_time_zone_unique_check time with time zone UNIQUE CHECK ( time_with_time_zone > '12:34:56+8' ),
    time_with_time_zone_unique_ref time with time zone UNIQUE REFERENCES time_with_time_zone_ref(time_with_time_zone_ref),
    time_with_time_zone_unique_def_const time with time zone UNIQUE DEFAULT '12:34:56+8',
    time_with_time_zone_unique_def_func time with time zone UNIQUE DEFAULT now(),

    time_with_time_zone_check time with time zone CHECK ( time_with_time_zone > '12:34:56+8' ),
    time_with_time_zone_check_ref time with time zone CHECK ( time_with_time_zone > '12:34:56+8' ) REFERENCES time_with_time_zone_ref(time_with_time_zone_ref),
    time_with_time_zone_check_def_const time with time zone CHECK ( time_with_time_zone > '12:34:56+8' ) DEFAULT '12:34:56+8',
    time_with_time_zone_check_def_func time with time zone CHECK ( time_with_time_zone > '12:34:56+8' ) DEFAULT now(),

    time_with_time_zone_ref time with time zone REFERENCES time_with_time_zone_ref(time_with_time_zone_ref),
    time_with_time_zone_ref_def_const time with time zone REFERENCES time_with_time_zone_ref(time_with_time_zone_ref) DEFAULT '12:34:56+8',
    time_with_time_zone_ref_def_func time with time zone REFERENCES time_with_time_zone_ref(time_with_time_zone_ref) DEFAULT now(),
    time_with_time_zone_ref_unique_check time with time zone UNIQUE CHECK ( time_with_time_zone > '12:34:56+8' ) REFERENCES time_with_time_zone_ref(time_with_time_zone_ref),

    time_with_time_zone_def_const time with time zone DEFAULT '12:34:56+8',
    time_with_time_zone_def_const_unique_check time with time zone UNIQUE CHECK ( time_with_time_zone > '12:34:56+8' )DEFAULT '12:34:56+8',

    time_with_time_zone_def_func time with time zone DEFAULT now(),
    time_with_time_zone_def_func_unique_check time with time zone UNIQUE CHECK ( time_with_time_zone > '12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS time_with_time_zone_pk;
CREATE TABLE time_with_time_zone_pk (
    time_with_time_zone_pk time with time zone PRIMARY KEY
);

DROP TABLE IF EXISTS time_with_time_zone_pk_ref;
CREATE TABLE time_with_time_zone_pk_ref (
    time_with_time_zone_pk_ref time with time zone PRIMARY KEY REFERENCES time_with_time_zone_ref(time_with_time_zone_ref)
);

DROP TABLE IF EXISTS time_with_time_zone_pk_def_const;
CREATE TABLE time_with_time_zone_pk_def_const (
    time_with_time_zone_pk_def_const time with time zone PRIMARY KEY DEFAULT '12:34:56+8'
);

DROP TABLE IF EXISTS time_with_time_zone_pk_def_func;
CREATE TABLE time_with_time_zone_pk_def_func (
    time_with_time_zone_pk_def_func time with time zone PRIMARY KEY DEFAULT now()
);

DROP TABLE IF EXISTS time_with_time_zone_nn_pk;
CREATE TABLE time_with_time_zone_nn_pk (
    time_with_time_zone_nn_pk time with time zone NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS time_with_time_zone_nn_unique_check_pk;
CREATE TABLE time_with_time_zone_nn_unique_check_pk (
    time_with_time_zone_nn_unique_check_pk time with time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( time_with_time_zone_nn_unique_check_pk > '12:34:56+8' )
);

DROP TABLE IF EXISTS time_with_time_zone_nn_unique_check_pk_ref;
CREATE TABLE time_with_time_zone_nn_unique_check_pk_ref (
    time_with_time_zone_nn_unique_check_pk_ref time with time zone PRIMARY KEY NOT NULL UNIQUE CHECK ( time_with_time_zone_nn_unique_check_pk_ref > '12:34:56+8' ) REFERENCES time_with_time_zone_ref(time_with_time_zone_ref)
);

DROP TABLE IF EXISTS time_with_time_zone_unique_pk;
CREATE TABLE time_with_time_zone_unique_pk (
    time_with_time_zone_unique_pk time with time zone PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS time_with_time_zone_unique_check_pk;
CREATE TABLE time_with_time_zone_unique_check_pk (
    time_with_time_zone_unique_check_pk time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_unique_check_pk > '12:34:56+8' )
);

DROP TABLE IF EXISTS time_with_time_zone_unique_check_pk_ref;
CREATE TABLE time_with_time_zone_unique_check_pk_ref (
    time_with_time_zone_unique_check_pk_ref time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_unique_check_pk_ref > '12:34:56+8' ) REFERENCES time_with_time_zone_ref(time_with_time_zone_ref)
);

DROP TABLE IF EXISTS time_with_time_zone_check_pk;
CREATE TABLE time_with_time_zone_check_pk (
    time_with_time_zone_check_pk time with time zone PRIMARY KEY CHECK ( time_with_time_zone_check_pk > '12:34:56+8' )
);

DROP TABLE IF EXISTS time_with_time_zone_def_const_unique_check_pk;
CREATE TABLE time_with_time_zone_def_const_unique_check_pk (
    time_with_time_zone_def_const_unique_check_pk time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_def_const_unique_check_pk > '12:34:56+8' ) DEFAULT '12:34:56+8'
);

DROP TABLE IF EXISTS time_with_time_zone_def_const_unique_check_pk_ref;
CREATE TABLE time_with_time_zone_def_const_unique_check_pk_ref (
    time_with_time_zone_def_const_unique_check_pk_ref time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_def_const_unique_check_pk_ref > '12:34:56+8' ) DEFAULT '12:34:56+8' REFERENCES time_with_time_zone_ref(time_with_time_zone_ref)
);

DROP TABLE IF EXISTS time_with_time_zone_def_func_unique_check_pk;
CREATE TABLE time_with_time_zone_def_func_unique_check_pk (
    time_with_time_zone_def_func_unique_check_pk time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_def_func_unique_check_pk > '12:34:56+8' ) DEFAULT now()
);

DROP TABLE IF EXISTS time_with_time_zone_def_func_unique_check_pk_ref;
CREATE TABLE time_with_time_zone_def_func_unique_check_pk_ref (
    time_with_time_zone_def_func_unique_check_pk_ref time with time zone PRIMARY KEY UNIQUE CHECK ( time_with_time_zone_def_func_unique_check_pk_ref > '12:34:56+8' ) DEFAULT now() REFERENCES time_with_time_zone_ref(time_with_time_zone_ref)
);
