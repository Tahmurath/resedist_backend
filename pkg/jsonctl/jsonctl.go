package jsonctl

import (
	"github.com/gin-gonic/gin"
	"resedist/pkg/config"
)

type JsonCtl struct {
	Status        string
	Error_message string
	Error_code    string
	Pagination    string
	Data          string
	c             *gin.Context
}

func New(c *gin.Context) *JsonCtl {
	cfg := config.Get().Jsonkey

	ctl := NewJson(JsonCtl{
		Status:        cfg.Status,
		Error_message: cfg.Error_message,
		Error_code:    cfg.Error_code,
		c:             c,
	})
	return ctl
}

func NewJson(cfg JsonCtl) *JsonCtl {
	return &JsonCtl{
		Status:        cfg.Status,
		Error_message: cfg.Error_message,
		Error_code:    cfg.Error_code,
		c:             cfg.c,
	}
}

func (jc *JsonCtl) Json(http int) {
	jc.c.JSON(http, gin.H{
		jc.Status:        "failed",
		jc.Error_message: "Opps, there is an error with Query bind",
		jc.Error_code:    "",
	})
}

//
//type Person struct {
//	name string
//	age  int
//}
//
//type Option func(*Person)
//
//func WithName(name string) Option {
//	return func(p *Person) {
//		p.name = name
//	}
//}
//
//func WithAge(age int) Option {
//	return func(p *Person) {
//		p.age = age
//	}
//}
//
//func NewPerson(options ...Option) *Person {
//	p := &Person{} // پیش‌فرض‌ها رو می‌تونید اینجا تنظیم کنید
//	for _, opt := range options {
//		opt(p)
//	}
//	return p
//}
//
//func main() {
//	person := NewPerson(
//		WithName("Ali"),
//		WithAge(30),
//	)
//	fmt.Printf("Name: %s, Age: %d\n", person.name, person.age)
//}
