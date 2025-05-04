package lesson5

// service drivers

type User struct {
	userID   int
	name     string
	surname  string
	passport string
}

type UserGetRentalCharges struct {
	chargeID           int
	userID             int
	vehicleAndDriverID int
}

// 3.2
// old name: Manager
// new name: DriveAgency
type DriveAgency struct {
	agencyID           int
	name               string
	contactInformation string
}

// 3.2
// Receive - Get
// Fetch - Get
// Append - Add
// Connect - Add
// Conclude - Add

// old name: ReceiveAllDrivers
func (agency *DriveAgency) GetAllDrivers()         {}
func (agency *DriveAgency) GetDriver(driverID int) {}

// old name: FetchVehicle
func (agency *DriveAgency) GetVehicle(vehicleID int) {}

// old name: RecieveCharge
func (agency *DriveAgency) GetCharge(driverID int, vehicleID int) {}

// old name: ConcludeContract
func (agency *DriveAgency) AddContract(user User, charge VehicleAndDriver) {}
func (agency *DriveAgency) AddDriver(driverInfo Driver)                    {}

// old name: AppendVehicle
func (agency *DriveAgency) AddVehicle(vehicle Vehicle) {}

// old name: ConnectDriverToVehicle
func (agency *DriveAgency) AddDriverToVehicle(vehicleID int, driverID int) {}

type Driver struct {
	driverID int
	name     string
	surname  string
	license  bool
}

// Provide - Get
// Give - Get

// old name: ProvideLicense
func (driver *Driver) GetLicense() {}

type VehicleAndDriver struct {
	id        int
	vehicleID int
	driverID  int
}

type Vehicle struct {
	vehicleID   int
	vehicleType VehicleType
}

// old name: GiveVehicleType
func (vehicle *Vehicle) GetVehicleType() {}

type VehicleType struct {
	vehicleTypeID int
	vehicleType   string
}
