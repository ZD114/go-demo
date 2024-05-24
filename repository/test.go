package repository

type Test struct {
	Id    int64  `xorm:"bigint 'id' autoincr pk notnull default() comment('ID')" json:"id"`
	Value string `xorm:"varchar 'value' notnull default('') comment('值')" json:"value"`
}
