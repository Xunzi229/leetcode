package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

type Person struct {
	Name  string `gorm:"column:name"`
	Age   int    `gorm:"column:age"`
	Heigh int
}

func (p Person) GormTableField() []string {
	l := reflect.TypeOf(p)

	var field []string
	for i := 0; i < l.NumField(); i++ {
		name := l.Field(i).Tag.Get("gorm")
		if len(name) != 0 {
			field = append(field, parseGormFieldName(name))
		}
	}
	return field
}

func main() {
	p := Person{}

	field := p.GormTableField()
	fmt.Println(field)

	Wait()
}

func parseGormFieldName(str string) string {
	if len(str) == 0 {
		return ""
	}

	cs := strings.Split(str, ";")
	for i := 0; i < len(cs); i++ {
		c := strings.Split(cs[i], ":")
		if c[0] == "column" && len(c) == 2 {
			return c[1]
		}
	}
	return ""
}

func Wait() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("开始", time.Now())
	go func() {
		time.Sleep(4 * time.Second)
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("结束", time.Now())
}
