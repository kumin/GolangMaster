package mysql

type SectionFieldsDAO struct {
	Id      uint64 `gorm:"primaryKey" json:"id,omitempty"`
	Section string `json:"section,omitempty"`
	Fields  string `json:"fields,omitempty"`
}

func (SectionFieldsDAO) TableName() string {
	return "section_fields"
}

type FieldValuesDAO struct {
	Id    uint64 `gorm:"primaryKey" json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

func (FieldValuesDAO) TableName() string {
	return "field_values"
}
