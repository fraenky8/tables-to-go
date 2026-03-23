DROP TABLE IF EXISTS blob_ref CASCADE;
CREATE TABLE blob_ref
(
    blob_ref blob NOT NULL,
    KEY      blob_ref_len_key (blob_ref(10))
);

DROP TABLE IF EXISTS blob_table;
CREATE TABLE blob_table
(
    col                  blob,

    blob_def_const       blob,
    blob_def_func        blob,

    blob_ref             blob REFERENCES blob_ref (blob_ref),

    blob_nn              blob NOT NULL,
    blob_nn_check_cmp    blob NOT NULL,
    blob_nn_check_fn     blob NOT NULL,
    blob_nn_ref          blob NOT NULL REFERENCES blob_ref (blob_ref),
    blob_nn_def_const    blob NOT NULL,
    blob_nn_def_func     blob NOT NULL,

    blob_check           blob,
    blob_check_ref       blob REFERENCES blob_ref (blob_ref),
    blob_check_def_const blob,
    blob_check_def_func  blob
);
