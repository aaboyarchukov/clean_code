package lesson5

import "time"

// 3.1
// old name: Object
// new name: Tender
type Tender struct {
	TenderID   string `gorm:"primaryKey"`
	FolderName string
}

// 3.1
// old name: Equipment_to_object
// new name: EquipmentToTender
type EquipmentToTender struct {
	ID               uint `gorm:"primaryKey"`
	ObjectReferID    string
	ObjectID         Tender `gorm:"foreignKey:ObjectReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EquipmentReferID uint
	EquipmentID      TenderEquipment `gorm:"foreignKey:EquipmentReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// 3.1
// old name: ActivitiesAndTypes
// new name: StagesForTenderTypes
type StagesForTenderTypes struct {
	ID            uint `gorm:"primaryKey"`
	ObjectReferID string
	ObjectID      Tender `gorm:"foreignKey:ObjectReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TypeReferID   uint
	TypeID        TenderEquipmentType `gorm:"foreignKey:TypeReferID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Activities    pq.StringArray      `gorm:"type:text[]"`
}

// 3.1
// old name: Description
// new name: TenderDescription
type TenderDescription struct {
	DescriptionID   uint   `gorm:"primaryKey"`
	ObjectRefer     string `gorm:"unique"`
	ObjectID        Tender `gorm:"foreignKey:ObjectRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Number          string
	WorkName        string
	FullName        string
	FullAdress      string
	WorkAdress      string
	Link_to_sbis    string
	Link_to_zakupki string
	ContractNumber  string
	ContractLink    string
	ContractObject  string
	Name_Documents  pq.StringArray `gorm:"type:text[]"`
	Documents       pq.StringArray `gorm:"type:text[]"`
	Delivery_term   string
	Data            time.Time
}

// 3.1
// old name: Equipment
// new name: TenderEquipment
type TenderEquipment struct {
	EquipmentID            uint `gorm:"primaryKey"`
	ArticleNumber          string
	Code                   string
	Resource_name          string
	Resource_code          string
	Deadline               string
	Delivery               string
	PaymentDay             string
	Units                  string
	Price_per_unit_now     float64
	New_Price_per_unit_now float64
	Cost                   float64
	NewCost                float64
	Count                  float64
	NewCount               float64
	NDS                    float64
	INN                    string
	Specifications         string
	LastMovePrice          string
	LastMoveCount          string
	EquipmentKind          string
	EquipmentDocumentsPath *string
}

// 3.1
// old name: Equipment_type
// new name: TenderEquipmentType
type TenderEquipmentType struct {
	Equipment_typeID uint   `gorm:"primaryKey"`
	Name             string `gorm:"unique"`
}
