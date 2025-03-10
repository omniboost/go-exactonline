package logistics

import (
	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/omitempty"
)

type Items []Item

type Item struct {
	ID                      edm.GUID     `json:"ID"`
	Barcode                 edm.String   `json:"Barcode"`
	Class01                 edm.String   `json:"Class_01"`
	Class02                 edm.String   `json:"Class_02"`
	Class03                 edm.String   `json:"Class_03"`
	Class04                 edm.String   `json:"Class_04"`
	Class05                 edm.String   `json:"Class_05"`
	Class06                 edm.String   `json:"Class_06"`
	Class07                 edm.String   `json:"Class_07"`
	Class08                 edm.String   `json:"Class_08"`
	Class09                 edm.String   `json:"Class_09"`
	Class10                 edm.String   `json:"Class_10"`
	Code                    edm.String   `json:"Code"`
	CopyRemarks             edm.Byte     `json:"CopyRemarks"`
	CostPriceCurrency       edm.String   `json:"CostPriceCurrency"`
	CostPriceNew            edm.Double   `json:"CostPriceNew"`
	CostPriceStandard       edm.Double   `json:"CostPriceStandard"`
	Created                 edm.DateTime `json:"Created"`
	Creator                 edm.GUID     `json:"Creator"`
	CreatorFullName         edm.String   `json:"CreatorFullName"`
	Description             edm.String   `json:"Description"`
	Division                edm.Int32    `json:"Division"`
	EndDate                 edm.DateTime `json:"EndDate"`
	ExtraDescription        edm.String   `json:"ExtraDescription"`
	FreeBoolField01         edm.Boolean  `json:"FreeBoolField_01"`
	FreeBoolField02         edm.Boolean  `json:"FreeBoolField_02"`
	FreeBoolField03         edm.Boolean  `json:"FreeBoolField_03"`
	FreeBoolField04         edm.Boolean  `json:"FreeBoolField_04"`
	FreeBoolField05         edm.Boolean  `json:"FreeBoolField_05"`
	FreeDateField01         edm.DateTime `json:"FreeDateField_01"`
	FreeDateField02         edm.DateTime `json:"FreeDateField_02"`
	FreeDateField03         edm.DateTime `json:"FreeDateField_03"`
	FreeDateField04         edm.DateTime `json:"FreeDateField_04"`
	FreeDateField05         edm.DateTime `json:"FreeDateField_05"`
	FreeNumberField01       edm.Double   `json:"FreeNumberField_01"`
	FreeNumberField02       edm.Double   `json:"FreeNumberField_02"`
	FreeNumberField03       edm.Double   `json:"FreeNumberField_03"`
	FreeNumberField04       edm.Double   `json:"FreeNumberField_04"`
	FreeNumberField05       edm.Double   `json:"FreeNumberField_05"`
	FreeNumberField06       edm.Double   `json:"FreeNumberField_06"`
	FreeNumberField07       edm.Double   `json:"FreeNumberField_07"`
	FreeNumberField08       edm.Double   `json:"FreeNumberField_08"`
	FreeTextField01         edm.String   `json:"FreeTextField_01"`
	FreeTextField02         edm.String   `json:"FreeTextField_02"`
	FreeTextField03         edm.String   `json:"FreeTextField_03"`
	FreeTextField04         edm.String   `json:"FreeTextField_04"`
	FreeTextField05         edm.String   `json:"FreeTextField_05"`
	FreeTextField06         edm.String   `json:"FreeTextField_06"`
	FreeTextField07         edm.String   `json:"FreeTextField_07"`
	FreeTextField08         edm.String   `json:"FreeTextField_08"`
	FreeTextField09         edm.String   `json:"FreeTextField_09"`
	FreeTextField10         edm.String   `json:"FreeTextField_10"`
	GLCosts                 edm.GUID     `json:"GLCosts"`
	GLCostsCode             edm.String   `json:"GLCostsCode"`
	GLCostsDescription      edm.String   `json:"GLCostsDescription"`
	GLRevenue               edm.GUID     `json:"GLRevenue"`
	GLRevenueCode           edm.String   `json:"GLRevenueCode"`
	GLRevenueDescription    edm.String   `json:"GLRevenueDescription"`
	GLStock                 edm.GUID     `json:"GLStock"`
	GLStockCode             edm.String   `json:"GLStockCode"`
	GLStockDescription      edm.String   `json:"GLStockDescription"`
	GrossWeight             edm.Double   `json:"GrossWeight"`
	IsBatchItem             edm.Byte     `json:"IsBatchItem"`
	IsBatchNumberItem       edm.Byte     `json:"IsBatchNumberItem"`
	IsFractionAllowedItem   edm.Boolean  `json:"IsFractionAllowedItem"`
	IsMakeItem              edm.Byte     `json:"IsMakeItem"`
	IsNewContract           edm.Byte     `json:"IsNewContract"`
	IsOnDemandItem          edm.Byte     `json:"IsOnDemandItem"`
	IsPackageItem           edm.Boolean  `json:"IsPackageItem"`
	IsPurchaseItem          edm.Boolean  `json:"IsPurchaseItem"`
	IsRegistrationCodeItem  edm.Byte     `json:"IsRegistrationCodeItem"`
	IsSalesItem             edm.Boolean  `json:"IsSalesItem"`
	IsSerialItem            edm.Boolean  `json:"IsSerialItem"`
	IsSerialNumberItem      edm.Boolean  `json:"IsSerialNumberItem"`
	IsStockItem             edm.Boolean  `json:"IsStockItem"`
	IsSubcontractedItem     edm.Boolean  `json:"IsSubcontractedItem"`
	IsTaxableItem           edm.Byte     `json:"IsTaxableItem"`
	IsTime                  edm.Byte     `json:"IsTime"`
	IsWebshopItem           edm.Byte     `json:"IsWebshopItem"`
	ItemGroup               edm.GUID     `json:"ItemGroup"`
	ItemGroupCode           edm.String   `json:"ItemGroupCode"`
	ItemGroupDescription    edm.String   `json:"ItemGroupDescription"`
	Modified                edm.DateTime `json:"Modified"`
	Modifier                edm.GUID     `json:"Modifier"`
	ModifierFullName        edm.String   `json:"ModifierFullName"`
	NetWeight               edm.Double   `json:"NetWeight"`
	NetWeightUnit           edm.String   `json:"NetWeightUnit"`
	Notes                   edm.String   `json:"Notes"`
	PictureName             edm.String   `json:"PictureName"`
	PictureThumbnailUrl     edm.String   `json:"PictureThumbnailUrl"`
	PictureUrl              edm.String   `json:"PictureUrl"`
	SalesVatCode            edm.String   `json:"SalesVatCode"`
	SalesVatCodeDescription edm.String   `json:"SalesVatCodeDescription"`
	SearchCode              edm.String   `json:"SearchCode"`
	SecurityLevel           edm.Int32    `json:"SecurityLevel"`
	StartDate               edm.DateTime `json:"StartDate"`
	Stock                   edm.Double   `json:"Stock"`
	Unit                    edm.String   `json:"Unit"`
	UnitDescription         edm.String   `json:"UnitDescription"`
	UnitType                UnitType     `json:"UnitType"`
}

const (
	UnitTypeArea   UnitType = "A" // Area
	UnitTypeLength UnitType = "L" // Length
	UnitTypeOther  UnitType = "O" // Other
	UnitTypeTime   UnitType = "T" // Time
	UnitTypeVolume UnitType = "V" // Volume
	UnitTypeWeight UnitType = "W" // Weight
)

type UnitType edm.String

func (u *UnitType) String() string {
	switch string(*u) {
	case "A":
		return "Area"
	case "L":
		return "Length"
	case "O":
		return "Other"
	case "T":
		return "Time"
	case "V":
		return "Volume"
	case "W":
		return "Weight"
	}
	return ""
}

type NewItem struct {
	ID                    edm.GUID     `json:"ID,omitempty"`                    // A guid that is the unique identifier of the item
	AverageCost           edm.Double   `json:"AverageCost,omitempty"`           // The current average cost price
	Barcode               edm.String   `json:"Barcode,omitempty"`               // Barcode of the item (numeric string)
	Class_01              edm.String   `json:"Class_01,omitempty"`              // Item class code referring to ItemClasses with ClassID 1
	Class_02              edm.String   `json:"Class_02,omitempty"`              // Item class code referring to ItemClasses with ClassID 2
	Class_03              edm.String   `json:"Class_03,omitempty"`              // Item class code referring to ItemClasses with ClassID 3
	Class_04              edm.String   `json:"Class_04,omitempty"`              // Item class code referring to ItemClasses with ClassID 4
	Class_05              edm.String   `json:"Class_05,omitempty"`              // Item class code referring to ItemClasses with ClassID 5
	Class_06              edm.String   `json:"Class_06,omitempty"`              // Item class code referring to ItemClasses with ClassID 6
	Class_07              edm.String   `json:"Class_07,omitempty"`              // Item class code referring to ItemClasses with ClassID 7
	Class_08              edm.String   `json:"Class_08,omitempty"`              // Item class code referring to ItemClasses with ClassID 8
	Class_09              edm.String   `json:"Class_09,omitempty"`              // Item class code referring to ItemClasses with ClassID 9
	Class_10              edm.String   `json:"Class_10,omitempty"`              // Item class code referring to ItemClasses with ClassID 10
	Code                  edm.String   `json:"Code"`                            // Item code
	CopyRemarks           edm.Byte     `json:"CopyRemarks,omitempty"`           // Copy sales remarks to sales lines
	CostPriceCurrency     edm.String   `json:"CostPriceCurrency,omitempty"`     // The currency of the current and proposed cost price
	CostPriceNew          edm.Double   `json:"CostPriceNew,omitempty"`          // Proposed cost price
	CostPriceStandard     edm.Double   `json:"CostPriceStandard,omitempty"`     // The current standard cost price
	Description           edm.String   `json:"Description"`                     // Description of the item
	EndDate               edm.DateTime `json:"EndDate,omitempty"`               // Together with StartDate this determines if the item is active
	ExtraDescription      edm.String   `json:"ExtraDescription,omitempty"`      // Extra description text, slightly longer than the regular description (255 instead of 60)
	FreeBoolField_01      edm.Boolean  `json:"FreeBoolField_01,omitempty"`      // Free boolean field 1
	FreeBoolField_02      edm.Boolean  `json:"FreeBoolField_02,omitempty"`      // Free boolean field 2
	FreeBoolField_03      edm.Boolean  `json:"FreeBoolField_03,omitempty"`      // Free boolean field 3
	FreeBoolField_04      edm.Boolean  `json:"FreeBoolField_04,omitempty"`      // Free boolean field 4
	FreeBoolField_05      edm.Boolean  `json:"FreeBoolField_05,omitempty"`      // Free boolean field 5
	FreeDateField_01      edm.DateTime `json:"FreeDateField_01,omitempty"`      // Free date field 1
	FreeDateField_02      edm.DateTime `json:"FreeDateField_02,omitempty"`      // Free date field 2
	FreeDateField_03      edm.DateTime `json:"FreeDateField_03,omitempty"`      // Free date field 3
	FreeDateField_04      edm.DateTime `json:"FreeDateField_04,omitempty"`      // Free date field 4
	FreeDateField_05      edm.DateTime `json:"FreeDateField_05,omitempty"`      // Free date field 5
	FreeNumberField_01    edm.Double   `json:"FreeNumberField_01,omitempty"`    // Free numeric field 1
	FreeNumberField_02    edm.Double   `json:"FreeNumberField_02,omitempty"`    // Free numeric field 2
	FreeNumberField_03    edm.Double   `json:"FreeNumberField_03,omitempty"`    // Free numeric field 3
	FreeNumberField_04    edm.Double   `json:"FreeNumberField_04,omitempty"`    // Free numeric field 4
	FreeNumberField_05    edm.Double   `json:"FreeNumberField_05,omitempty"`    // Free numeric field 5
	FreeNumberField_06    edm.Double   `json:"FreeNumberField_06,omitempty"`    // Free numeric field 6
	FreeNumberField_07    edm.Double   `json:"FreeNumberField_07,omitempty"`    // Free numeric field 7
	FreeNumberField_08    edm.Double   `json:"FreeNumberField_08,omitempty"`    // Free numeric field 8
	FreeTextField_01      edm.String   `json:"FreeTextField_01,omitempty"`      // Free text field 1
	FreeTextField_02      edm.String   `json:"FreeTextField_02,omitempty"`      // Free text field 2
	FreeTextField_03      edm.String   `json:"FreeTextField_03,omitempty"`      // Free text field 3
	FreeTextField_04      edm.String   `json:"FreeTextField_04,omitempty"`      // Free text field 4
	FreeTextField_05      edm.String   `json:"FreeTextField_05,omitempty"`      // Free text field 5
	FreeTextField_06      edm.String   `json:"FreeTextField_06,omitempty"`      // Free text field 6
	FreeTextField_07      edm.String   `json:"FreeTextField_07,omitempty"`      // Free text field 7
	FreeTextField_08      edm.String   `json:"FreeTextField_08,omitempty"`      // Free text field 8
	FreeTextField_09      edm.String   `json:"FreeTextField_09,omitempty"`      // Free text field 9
	FreeTextField_10      edm.String   `json:"FreeTextField_10,omitempty"`      // Free text field 10
	GLCosts               edm.GUID     `json:"GLCosts,omitempty"`               // GL account the cost entries will be booked on. This overrules the GL account from the item group. If the license contains 'Intuit integration' this property overrides the value in Settings, not the item group.
	GLRevenue             edm.GUID     `json:"GLRevenue,omitempty"`             // GL account the revenue will be booked on. This overrules the GL account from the item group. If the license contains 'Intuit integration' this property overrides the value in Settings, not the item group.
	GLStock               edm.GUID     `json:"GLStock,omitempty"`               // GL account the stock entries will be booked on. This overrules the GL account from the item group. If the license contains 'Intuit integration' this property overrides the value in Settings, not the item group.
	GrossWeight           edm.Double   `json:"GrossWeight,omitempty"`           // Gross weight for international goods shipments
	IsBatchItem           edm.Byte     `json:"IsBatchItem,omitempty"`           // Indicates if batches are used for this item
	IsFractionAllowedItem edm.Boolean  `json:"IsFractionAllowedItem,omitempty"` // Indicates if fractions (for example 0.35) are allowed for quantities of this item
	IsMakeItem            edm.Byte     `json:"IsMakeItem,omitempty"`            // Indicates that an Item is produced to Inventory, not purchased
	IsNewContract         edm.Byte     `json:"IsNewContract,omitempty"`         // Only used for packages (IsPackageItem=1). To indicate if this package is a new contract type package
	IsOnDemandItem        edm.Byte     `json:"IsOnDemandItem,omitempty"`        // Is On demand Item
	IsPackageItem         edm.Boolean  `json:"IsPackageItem,omitempty"`         // Indicates if the item is a package item. Can only be created in the hosting administration
	IsPurchaseItem        edm.Boolean  `json:"IsPurchaseItem,omitempty"`        // Indicates if the item can be purchased
	IsSalesItem           edm.Boolean  `json:"IsSalesItem,omitempty"`           // Indicates if the item can be sold
	IsSerialItem          edm.Boolean  `json:"IsSerialItem,omitempty"`          // Indicates that serial numbers are used for this item
	IsStockItem           edm.Boolean  `json:"IsStockItem,omitempty"`           // If you have the Trade or Manufacturing license and you check this property the item will be shown in the stock positions overview, stock counts and transaction lists. If you have the Invoice module and you check this property you will get a general journal entry based on the Stock and Costs G/L accounts of the item group. If you don’t want the general journal entry to be created you should change the Stock/Costs G/L account on the Item group page to the type Costs instead of Inventory.
	IsSubcontractedItem   edm.Boolean  `json:"IsSubcontractedItem,omitempty"`   // Indicates if the item is provided by an outside supplier
	IsTaxableItem         edm.Byte     `json:"IsTaxableItem,omitempty"`         // Indicates if tax needs to be calculated for this item
	IsTime                edm.Byte     `json:"IsTime,omitempty"`                // Indicates if the item is a time unit item (for example a labor hour item)
	IsWebshopItem         edm.Byte     `json:"IsWebshopItem,omitempty"`         // Indicates if the item can be exported to a web shop
	ItemGroup             edm.GUID     `json:"ItemGroup,omitempty"`             // GUID of Item group of the item
	NetWeight             edm.Double   `json:"NetWeight,omitempty"`             // Net weight for international goods shipments
	NetWeightUnit         edm.String   `json:"NetWeightUnit,omitempty"`         // Net Weight unit for international goods shipment, only available in manufacturing packages
	Notes                 edm.String   `json:"Notes,omitempty"`                 // Notes
	SalesVatCode          edm.String   `json:"SalesVatCode,omitempty"`          // Code of SalesVat
	SearchCode            edm.String   `json:"SearchCode,omitempty"`            // Search code of the item
	SecurityLevel         edm.Int32    `json:"SecurityLevel,omitempty"`         // Security level (0 - 100)
	StartDate             edm.DateTime `json:"StartDate,omitempty"`             // Together with EndDate this determines if the item is active
	Unit                  edm.String   `json:"Unit,omitempty"`                  // The standard unit of this item
}

func (i NewItem) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type Units []Unit

type Unit struct {
	ID          edm.GUID    `json:"ID,omitempty"`          // Primary key
	Active      edm.Boolean `json:"Active,omitempty"`      // Indicates whether a unit is in use
	Code        edm.String  `json:"Code,omitempty"`        // Unique code for the unit
	Description edm.String  `json:"Description,omitempty"` // Description
	Division    edm.Int32   `json:"Division,omitempty"`    // Division code
	Main        edm.Byte    `json:"Main,omitempty"`        // Indicates the main unit per division. Will be used when creating new item
	TimeUnit    edm.String  `json:"TimeUnit,omitempty"`    // If Type = 'T' (time) then this fields indicates the type of time frame. yy = Year, mm = Month, wk = Week, dd = Day, hh = Hour, mi = Minute, ss = Second
	Type        edm.String  `json:"Type,omitempty"`        // Type of unit. Type 'Time' is especially important for contracts.
}

type SalesItemPrices []SalesItemPrice

type SalesItemPrice struct {
}
