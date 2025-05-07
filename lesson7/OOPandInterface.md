# ООП и интерфейсы

Продолжаем изучать и впитывать полезные рекомендации.

1. При перегрузке конструкторов классов, лучше давать ясные и понятные имена

При создании конструктора классов, перегружая их, лучше всего давать имена, которые описывают аргументы, а не те же имена, что и сам класс. В Go, как таковых классов нет, как и конструкторов с перегрузкой, но описывая на псевдо-Go языке, это должно выглядеть примерно так. Допустим мы можем задать окружность, с помощью радиуса, диаметра или длины:

```go
// bad example ❌
func (circle *Circle) NewCircle(radius float64) Circle {}
func (circle *Circle) NewCircle(diamenr float64) Circle {}
func (circle *Circle) NewCircle(length float64) Circle {}

// good example ✅
func (circle *Circle) ByRadius(radius float64) Circle {}
func  (circle *Circle) ByDiametr(diametr float64) Circle {}
func  (circle *Circle) ByLength(length float64) Circle {}
```

2. Именование интерфейсом

Давайте имена интерфейсам такие же ясные, как их реализациям, нет никакого смысла в виде таких имен, как:
- IShapeFactory
- ShapeFactoryImp
- CShapeFactory

```go
// bad example ❌
type IVehicle interface {

}

// good example ✅
type Vehile interface {

}
```

Задания:

3.1. Сделайте в своём коде три примера наглядных методов- фабрик.

[3.1_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson7/3.1_example1.go)

```go
// old name: NewRectangle
// new name: BySides

NewRectangle - BySides

func (rectangle *Rectangle) BySides(firstSide, secondSide, thirdSide int) *Rectangle {
    return &Rectangle{}
}
// rectangle := Rectungle.BySides(...)

NewCircle1 - ByRadius
NewCircle2 - ByDiametr
NewCircle3 - ByLength

type Circle struct {
}

func (circle *Circle) ByRadius(radius float64) Circle {
    return Circle{}
}

func (circle *Circle) ByDiametr(diametr float64) Circle {
    return Circle{}
}

func (circle *Circle) ByLength(length float64) Circle {
    return Circle{}
}

// circle := Circle.ByRadius(4)
```

3.2. Если вы когда-нибудь использовали интерфейсы или абстрактные классы, напишите несколько примеров их правильного именования.

[3.2_example1.go](https://github.com/aaboyarchukov/clean_code/blob/master/lesson7/3.2_example1.go)

```go
// old name: IPostgres
// new name: Storage
type Storage interface{}
// интерфейс хранилища
  

// old name: IAuth
// new name: AuthService
type AuthService interface{}
// интерфейс сервиса аутентификации
  

// old name: IWriter
// new name: Writer
type FileWriter interface{}
// интерфейс приложения, записывающее данные в файл
```