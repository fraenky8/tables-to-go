package dto

type CharacterVaryingPk struct {
	CharacterVaryingPk string `db:"character_varying_pk" gorm:"column:character_varying_pk;primaryKey;not null"`
}
