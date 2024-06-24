package task

import "fmt"

// определяем тип Gender как int
type Gender int

const (
	// используем iota для определения констант Male и Female
	Male Gender = iota
	Female
)

// метод String для типа Gender, возвращающий строковое представление
func (g Gender) String() string {
	return [2]string{"Male", "Female"}[g]
}

// определяем структуру Human с полями Name, Age и Gender
type Human struct {
	Name   string
	Age    int
	Gender Gender
}

// метод String для структуры Human, возвращающий строковое представление
func (h Human) String() string {
	return fmt.Sprintf("%s (%s) is %d years old", h.Name, h.Gender.String(), h.Age)
}

// определяем структуру Action с полем Actor типа Human
type Action struct {
	Actor Human
}

// метод Act для структуры Action, выводящий строковое представление Actor
func (a Action) Act() {
	fmt.Println(a.Actor.String())
}

func TaskN1() {
	// создаем экземпляр Human для Male
	humanMale := Human{
		"John",
		25,
		Male,
	}

	// выводим строковое представление Human для Male
	fmt.Println("Human struct (Male):", humanMale.String())

	// создаем экземпляр Action с встраиванием humanMale
	actionMale := Action{
		humanMale,
	}

	// вызываем метод Act для Action для Male
	fmt.Print("Action struct (Male): ")
	actionMale.Act()

	// создаем экземпляр Human для Female
	humanFemale := Human{
		"Jane",
		30,
		Female,
	}

	// выводим строковое представление Human для Female
	fmt.Println("Human struct (Female):", humanFemale.String())

	// создаем экземпляр Action с встраиванием humanFemale
	actionFemale := Action{
		humanFemale,
	}

	// вызываем метод Act для Action для Female
	fmt.Print("Action struct (Female): ")
	actionFemale.Act()
}
