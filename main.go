package main

// import "fmt"

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	cmd := NewCmdflag()
	cmd.Execute(&todos)
	Storage.Save(&todos)
	// fmt.Printf("%+v",todos);
}