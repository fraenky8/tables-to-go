DROP TABLE IF EXISTS constraint_combo_ref CASCADE;
CREATE TABLE constraint_combo_ref
(
    constraint_combo_ref double precision UNIQUE
);

DROP TABLE IF EXISTS constraint_combo_non_pk CASCADE;
CREATE TABLE constraint_combo_non_pk
(
    not_null_unique_check_def_const double precision NOT NULL UNIQUE CHECK (not_null_unique_check_def_const > 0) DEFAULT 42,
    not_null_unique_check_def_func  double precision NOT NULL UNIQUE CHECK (not_null_unique_check_def_func > 0)  DEFAULT pi(),
    not_null_unique_ref_def_const   double precision NOT NULL UNIQUE                                             DEFAULT 42 REFERENCES constraint_combo_ref (constraint_combo_ref),
    not_null_unique_ref_def_func    double precision NOT NULL UNIQUE                                             DEFAULT pi() REFERENCES constraint_combo_ref (constraint_combo_ref),
    not_null_check_ref_def_const    double precision NOT NULL CHECK (not_null_check_ref_def_const > 0)           DEFAULT 42 REFERENCES constraint_combo_ref (constraint_combo_ref),
    not_null_check_ref_def_func     double precision NOT NULL CHECK (not_null_check_ref_def_func > 0)            DEFAULT pi() REFERENCES constraint_combo_ref (constraint_combo_ref),
    not_null_unique_def_const       double precision NOT NULL UNIQUE                                             DEFAULT 42,
    not_null_unique_def_func        double precision NOT NULL UNIQUE                                             DEFAULT pi(),
    not_null_check_def_const        double precision NOT NULL CHECK (not_null_check_def_const > 0)               DEFAULT 42,
    not_null_check_def_func         double precision NOT NULL CHECK (not_null_check_def_func > 0)                DEFAULT pi(),
    not_null_ref_def_const          double precision NOT NULL                                                    DEFAULT 42 REFERENCES constraint_combo_ref (constraint_combo_ref),
    not_null_ref_def_func           double precision NOT NULL                                                    DEFAULT pi() REFERENCES constraint_combo_ref (constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_unique_pk_def_const CASCADE;
CREATE TABLE constraint_combo_not_null_unique_pk_def_const
(
    constraint_combo_not_null_unique_pk_def_const double precision PRIMARY KEY NOT NULL UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_unique_pk_def_func CASCADE;
CREATE TABLE constraint_combo_not_null_unique_pk_def_func
(
    constraint_combo_not_null_unique_pk_def_func double precision PRIMARY KEY NOT NULL UNIQUE DEFAULT pi()
);

DROP TABLE IF EXISTS constraint_combo_not_null_check_pk_def_const CASCADE;
CREATE TABLE constraint_combo_not_null_check_pk_def_const
(
    constraint_combo_not_null_check_pk_def_const double precision PRIMARY KEY NOT NULL CHECK (constraint_combo_not_null_check_pk_def_const > 0) DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_check_pk_def_func CASCADE;
CREATE TABLE constraint_combo_not_null_check_pk_def_func
(
    constraint_combo_not_null_check_pk_def_func double precision PRIMARY KEY NOT NULL CHECK (constraint_combo_not_null_check_pk_def_func > 0) DEFAULT pi()
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_ref_def_const CASCADE;
CREATE TABLE constraint_combo_not_null_pk_ref_def_const
(
    constraint_combo_not_null_pk_ref_def_const double precision PRIMARY KEY NOT NULL DEFAULT 42 REFERENCES constraint_combo_ref (constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_ref_def_func CASCADE;
CREATE TABLE constraint_combo_not_null_pk_ref_def_func
(
    constraint_combo_not_null_pk_ref_def_func double precision PRIMARY KEY NOT NULL DEFAULT pi() REFERENCES constraint_combo_ref (constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_def_const CASCADE;
CREATE TABLE constraint_combo_not_null_pk_def_const
(
    constraint_combo_not_null_pk_def_const double precision NOT NULL PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_def_func CASCADE;
CREATE TABLE constraint_combo_not_null_pk_def_func
(
    constraint_combo_not_null_pk_def_func double precision NOT NULL PRIMARY KEY DEFAULT pi()
);
