DROP TABLE IF EXISTS real_ref CASCADE;
CREATE TABLE real_ref (
    real_ref real UNIQUE
);

DROP TABLE IF EXISTS real;
CREATE TABLE real (
    real real,
    real_nn real NOT NULL,
    real_nn_unique real NOT NULL UNIQUE,
    real_nn_check real NOT NULL CHECK ( real > 0 ),
    real_nn_ref real NOT NULL REFERENCES real_ref(real_ref),
    real_nn_def_const real NOT NULL DEFAULT 42,
    real_nn_def_func real NOT NULL DEFAULT pi(),
    real_nn_unique_check real NOT NULL UNIQUE CHECK ( real > 0 ),

    real_unique real UNIQUE,
    real_unique_check real UNIQUE CHECK ( real > 0 ),
    real_unique_ref real UNIQUE REFERENCES real_ref(real_ref),
    real_unique_def_const real UNIQUE DEFAULT 42,
    real_unique_def_func real UNIQUE DEFAULT pi(),

    real_check real CHECK ( real > 0 ),
    real_check_ref real CHECK ( real > 0 ) REFERENCES real_ref(real_ref),
    real_check_def_const real CHECK ( real > 0 ) DEFAULT 42,
    real_check_def_func real CHECK ( real > 0 ) DEFAULT pi(),

    real_ref real REFERENCES real_ref(real_ref),
    real_ref_def_const real REFERENCES real_ref(real_ref) DEFAULT 42,
    real_ref_def_func real REFERENCES real_ref(real_ref) DEFAULT pi(),
    real_ref_unique_check real UNIQUE CHECK ( real > 0 ) REFERENCES real_ref(real_ref),

    real_def_const real DEFAULT 42,
    real_def_const_unique_check real UNIQUE CHECK ( real > 0 )DEFAULT 42,

    real_def_func real DEFAULT pi(),
    real_def_func_unique_check real UNIQUE CHECK ( real > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS real_pk;
CREATE TABLE real_pk (
    real_pk real PRIMARY KEY
);

DROP TABLE IF EXISTS real_pk_ref;
CREATE TABLE real_pk_ref (
    real_pk_ref real PRIMARY KEY REFERENCES real_ref(real_ref)
);

DROP TABLE IF EXISTS real_pk_def_const;
CREATE TABLE real_pk_def_const (
    real_pk_def_const real PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS real_pk_def_func;
CREATE TABLE real_pk_def_func (
    real_pk_def_func real PRIMARY KEY DEFAULT pi()
);

DROP TABLE IF EXISTS real_nn_pk;
CREATE TABLE real_nn_pk (
    real_nn_pk real NOT NULL PRIMARY KEY
);

DROP TABLE IF EXISTS real_nn_unique_check_pk;
CREATE TABLE real_nn_unique_check_pk (
    real_nn_unique_check_pk real PRIMARY KEY NOT NULL UNIQUE CHECK ( real_nn_unique_check_pk > 0)
);

DROP TABLE IF EXISTS real_nn_unique_check_pk_ref;
CREATE TABLE real_nn_unique_check_pk_ref (
    real_nn_unique_check_pk_ref real PRIMARY KEY NOT NULL UNIQUE CHECK ( real_nn_unique_check_pk_ref > 0) REFERENCES real_ref(real_ref)
);

DROP TABLE IF EXISTS real_unique_pk;
CREATE TABLE real_unique_pk (
    real_unique_pk real PRIMARY KEY UNIQUE
);

DROP TABLE IF EXISTS real_unique_check_pk;
CREATE TABLE real_unique_check_pk (
    real_unique_check_pk real PRIMARY KEY UNIQUE CHECK ( real_unique_check_pk > 0 )
);

DROP TABLE IF EXISTS real_unique_check_pk_ref;
CREATE TABLE real_unique_check_pk_ref (
    real_unique_check_pk_ref real PRIMARY KEY UNIQUE CHECK ( real_unique_check_pk_ref > 0) REFERENCES real_ref(real_ref)
);

DROP TABLE IF EXISTS real_check_pk;
CREATE TABLE real_check_pk (
    real_check_pk real PRIMARY KEY CHECK ( real_check_pk > 0 )
);

DROP TABLE IF EXISTS real_def_const_unique_check_pk;
CREATE TABLE real_def_const_unique_check_pk (
    real_def_const_unique_check_pk real PRIMARY KEY UNIQUE CHECK ( real_def_const_unique_check_pk > 0 ) DEFAULT 42
);

DROP TABLE IF EXISTS real_def_const_unique_check_pk_ref;
CREATE TABLE real_def_const_unique_check_pk_ref (
    real_def_const_unique_check_pk_ref real PRIMARY KEY UNIQUE CHECK ( real_def_const_unique_check_pk_ref > 0 ) DEFAULT 42 REFERENCES real_ref(real_ref)
);

DROP TABLE IF EXISTS real_def_func_unique_check_pk;
CREATE TABLE real_def_func_unique_check_pk (
    real_def_func_unique_check_pk real PRIMARY KEY UNIQUE CHECK ( real_def_func_unique_check_pk > 0 ) DEFAULT pi()
);

DROP TABLE IF EXISTS real_def_func_unique_check_pk_ref;
CREATE TABLE real_def_func_unique_check_pk_ref (
    real_def_func_unique_check_pk_ref real PRIMARY KEY UNIQUE CHECK ( real_def_func_unique_check_pk_ref > 0 ) DEFAULT pi() REFERENCES real_ref(real_ref)
);
