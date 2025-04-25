package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cmdflag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdflag() *Cmdflag {
	cf := Cmdflag{}
	flag.StringVar(&cf.Add,"add","","add new todo")
	flag.StringVar(&cf.Edit,"edit","","edit todo")
	flag.IntVar(&cf.Del,"del",-1,"delete a todo")
	flag.IntVar(&cf.Toggle,"toggle",-1,"delete a todo")
	flag.BoolVar(&cf.List,"list",false,"list  todo")
	flag.Parse()
	return &cf
}

func (cf * Cmdflag) Execute (todos *Todos)  {
	switch {
	case cf.List :
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Add != "":
		parts := strings.SplitN(cf.Edit,":",2)
		if len(parts) != 2 {
			fmt.Println("inavlid format")
			os.Exit(1)
		}
		i,err:= strconv.Atoi(parts[0])
		if err!= nil {
			fmt.Println("error occured")
			os.Exit(1)
		}
		todos.update(i,parts[1])
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	default:
		fmt.Println("invalid")
	}


}