package constant

type BaseMap map[string]interface{}

type StringMap map[string]string

type MsgMap map[string]map[int]string

type TableMap map[string]struct {
	Name    string
	Columns interface{}
}

type FieldMap map[string][]string

type SqlWhereInMap map[string][]interface{}

type SqlBetweenInMap map[string][2]interface{}

type SqlOrderByMap map[string]string

type SqlBuild struct {
	Where   BaseMap
	WhereIn SqlWhereInMap
	Between SqlBetweenInMap
	Like    BaseMap
	OrderBy SqlOrderByMap
	Limit   int
	Offset  int
}