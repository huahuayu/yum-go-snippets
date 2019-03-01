package judgement

type Foo interface {
	ForceToImplementA()
	ForceToImplementB()
}

type Bar struct {}

func (bar Bar) ForceToImplementA(){

}

//func (bar Bar) ForceToImplementB(){
//
//}

//开启这行可判断结构体是否实现了接口Foo. (如果实现了接口，goland的左侧边栏也有接口小图标显示)
//var _ Foo = (*Bar)(nil)

func main() {

}