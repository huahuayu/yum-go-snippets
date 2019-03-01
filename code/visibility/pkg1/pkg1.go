package pkg1

var (
	Name string
	Age  *int
)

func InitVar() {
	Name = "shiming"
	age := 20
	Age = &age
}
