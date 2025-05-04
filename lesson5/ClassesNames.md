# Имена классов

Далее переходим к теме классов и будем впитывать рекомендации по именованию их.

1. Используйте имена существительные и их комбинации

Ни в коем случае не надо использовать глаголы в названиях классов, так смысл недостаточно ясен и понятен, также не смотрится достаточно выразительно. Ведь класс должен отображать сущность - существительное, а вот его методы, то что данная сущность может делать - глаголы. Также старайтесь не использовать такие имена классов как: 
`Manager, Processor, Data или Info`

```go
// bad example ❌
type Customized struct {
	// ...fields
}

// good example ✅
type Customer struct {
	// ...fields
}

```

2. Для всей кодовой базы подберите единообразный лексикон

Таким образом вы упростите себе жизнь и жизни других разработчиков, которым не понадобится думать над разницей значений слов `Manager` и `Controller`

```go
// bad example ❌
type User struct {
	// ...fields
}

func (user *User) UserManager {}

type Vehicle struct {
	// ...fields
}

func (vehicle *Vehicle) VehicleController {}

// good example ✅
type User struct {
	// ...fields
}

func (user *User) UserController {}

type Vehicle struct {
	// ...fields
}

func (vehicle *Vehicle) VehicleController {}

```

Задания:

3.1. Улучшите пять имён классов в вашем коде.

[**3.1_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson5/3.1_example1.go)

```go
Object - Tender
// сущность торга на торговой площадке

Equipment_to_object - EquipmentToTender
// сущность для оборудований по определенному торгу

ActivitiesAndTypes - StagesForTenderTypes
// сущность для стадий, которые проходит определенный тип оборудования торга
// при работе с ним

Description - TenderDescription
// сущность описательной части торга, которая содержит в себе описательные 
// данные для торга

Equipment - TenderEquipment
// сущность для оборудования по определенному торгу

Equipment_type - TenderEquipmentType
// сущность для определенного типа оборудования по торгу

```

3.2. Улучшите семь имён методов и объектов по схеме из пункта 2.

[**3.2_example1.go**](https://github.com/aaboyarchukov/clean_code/blob/master/lesson5/3.2_example1.go)

```go

Manager - DriveAgency
// агенство по оказанию услуг перевозки

Receive - Get
Fetch - Get
Append - Add
Connect - Add
Conclude - Add

context:

// old name: ReceiveAllDrivers
func (agency *DriveAgency) GetAllDrivers()         {}

// old name: FetchVehicle
func (agency *DriveAgency) GetVehicle(vehicleID int) {}

// old name: RecieveCharge
func (agency *DriveAgency) GetCharge(driverID int, vehicleID int) {}

// old name: ConcludeContract
func (agency *DriveAgency) AddContract(user User, charge VehicleAndDriver) {}

// old name: AppendVehicle
func (agency *DriveAgency) AddVehicle(vehicle Vehicle) {}

// old name: ConnectDriverToVehicle
func (agency *DriveAgency) AddDriverToVehicle(vehicleID int, driverID int) {}


// Provide - Get
// Give - Get

context:

// old name: ProvideLicense
func (driver *Driver) GetLicense() {}

// old name: GiveVehicleType
func (vehicle *Vehicle) GetVehicleType() {}

```