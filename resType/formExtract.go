package resType

type FormData struct {
	Name  string
	Type  string
	Value string
}

type FormDatas struct {
	Enctype  string
	Action   string
	Method   string
	FormData []FormData
}
