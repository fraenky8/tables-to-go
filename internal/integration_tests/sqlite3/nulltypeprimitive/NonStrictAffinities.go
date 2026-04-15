package dto

type NonStrictAffinities struct {
	AffinityInt      *string `db:"affinity_int"`
	AffinityBigint   *string `db:"affinity_bigint"`
	AffinitySmallint *string `db:"affinity_smallint"`
	AffinityDecimal  *string `db:"affinity_decimal"`
	AffinityDouble   *string `db:"affinity_double"`
	AffinityBool     *bool   `db:"affinity_bool"`
	AffinityDate     *string `db:"affinity_date"`
	AffinityDatetime *string `db:"affinity_datetime"`
	AffinityVarchar  *string `db:"affinity_varchar"`
	AffinityChar     *string `db:"affinity_char"`
	AffinityClob     *string `db:"affinity_clob"`
	AffinityNchar    *string `db:"affinity_nchar"`
	AffinityNvarchar *string `db:"affinity_nvarchar"`
	AffinityJSON     *string `db:"affinity_json"`
	AffinityUUID     *string `db:"affinity_uuid"`
	AffinityBinary   *string `db:"affinity_binary"`
	AffinityVarbin   *string `db:"affinity_varbin"`
}
