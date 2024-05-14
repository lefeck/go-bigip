package bigip

//const (
//	BasePath      = "mgmt"
//	SHAREResoucre = "shared"
//	TMResource    = "tm"
//)

type Resource struct {
	BaseResourceName  string
	ShareResourceName string
	TMResourceName    string
}

func NewResourceName() Resource {
	return Resource{}
}

func (rn *Resource) BaseResourceNameString() string {
	return "mgmt"
}

func (rn *Resource) ShareResourceNameString() string {
	return "shared"
}

func (rn *Resource) TMResourceNameString() string {
	return "tm"
}
