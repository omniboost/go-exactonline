package purchaseorder

import (
	"encoding/json"

	"github.com/omniboost/go-exactonline/edm"
	"github.com/omniboost/go-exactonline/utils"
)

type PurchaseOrders []PurchaseOrder

type PurchaseOrder struct {
	OrderID                        edm.GUID           `json:"OrderID"`
	AmountDC                       edm.Double         `json:"AmountDC"`
	AmountFC                       edm.Double         `json:"AmountFC"`
	ApprovalStatus                 edm.Int16          `json:"ApprovalStatus"`
	ApprovalStatusDescription      edm.String         `json:"ApprovalStatusDescription"`
	Approved                       edm.DateTime       `json:"Approved"`
	Approver                       edm.GUID           `json:"Approver"`
	ApproverFullName               edm.String         `json:"ApproverFullName"`
	Created                        edm.DateTime       `json:"Created"`
	Creator                        edm.GUID           `json:"Creator"`
	CreatorFullName                edm.String         `json:"CreatorFullName"`
	Currency                       edm.String         `json:"Currency"`
	DeliverTo                      edm.GUID           `json:"DeliverTo"`
	DeliverToContactPerson         edm.GUID           `json:"DeliverToContactPerson"`
	DeliverToContactPersonFullName edm.String         `json:"DeliverToContactPersonFullName"`
	DeliverToName                  edm.String         `json:"DeliverToName"`
	DeliveryAddress                edm.GUID           `json:"DeliveryAddress"`
	DeliveryDate                   edm.DateTime       `json:"DeliveryDate"`
	DeliveryStatus                 edm.Int16          `json:"DeliveryStatus"`
	DeliveryStatusDescription      edm.String         `json:"DeliveryStatusDescription"`
	Description                    edm.String         `json:"Description"`
	Division                       edm.Int32          `json:"Division"`
	Document                       edm.GUID           `json:"Document"`
	DocumentNumber                 edm.Int32          `json:"DocumentNumber"`
	DocumentSubject                edm.String         `json:"DocumentSubject"`
	InvoiceStatus                  edm.Int16          `json:"InvoiceStatus"`
	InvoiceStatusDescription       edm.String         `json:"InvoiceStatusDescription"`
	InvoiceTo                      edm.GUID           `json:"InvoiceTo"`
	InvoiceToContactPerson         edm.GUID           `json:"InvoiceToContactPerson"`
	InvoiceToContactPersonFullName edm.String         `json:"InvoiceToContactPersonFullName"`
	InvoiceToName                  edm.String         `json:"InvoiceToName"`
	Modified                       edm.DateTime       `json:"Modified"`
	Modifier                       edm.GUID           `json:"Modifier"`
	ModifierFullName               edm.String         `json:"ModifierFullName"`
	OrderDate                      edm.DateTime       `json:"OrderDate"`
	OrderedBy                      edm.GUID           `json:"OrderedBy"`
	OrderedByContactPerson         edm.GUID           `json:"OrderedByContactPerson"`
	OrderedByContactPersonFullName edm.String         `json:"OrderedByContactPersonFullName"`
	OrderedByName                  edm.String         `json:"OrderedByName"`
	OrderNumber                    edm.Int32          `json:"OrderNumber"`
	PaymentCondition               edm.String         `json:"PaymentCondition"`
	PaymentConditionDescription    edm.String         `json:"PaymentConditionDescription"`
	PaymentReference               edm.String         `json:"PaymentReference"`
	Remarks                        edm.String         `json:"Remarks"`
	PurchaseOrderLines             PurchaseOrderLines `json:"PurchaseOrderLines"`
	Purchaseperson                 edm.GUID           `json:"Purchaseperson"`
	PurchasepersonFullName         edm.String         `json:"PurchasepersonFullName"`
	ShippingMethod                 edm.GUID           `json:"ShippingMethod"`
	ShippingMethodDescription      edm.String         `json:"ShippingMethodDescription"`
	Status                         edm.Int16          `json:"Status"`
	StatusDescription              edm.String         `json:"StatusDescription"`
	TaxSchedule                    edm.GUID           `json:"TaxSchedule"`
	TaxScheduleCode                edm.String         `json:"TaxScheduleCode"`
	TaxScheduleDescription         edm.String         `json:"TaxScheduleDescription"`
	WarehouseCode                  edm.String         `json:"WarehouseCode"`
	WarehouseDescription           edm.String         `json:"WarehouseDescription"`
	WarehouseID                    edm.GUID           `json:"WarehouseID"`
	YourRef                        edm.String         `json:"YourRef"`
}

type PurchaseOrderLines []PurchaseOrderLine

// standalone: "PurchaseOrderLines": []
// deferred: "PurchaseOrderLines": {"__deferred": {}}
// embedded: "PurchaseOrderLines": {"results": []}
func (o *PurchaseOrderLines) UnmarshalJSON(data []byte) error {
	type Results PurchaseOrderLines

	type Envelope struct {
		Results  Results         `json:"results"`
		Deferred json.RawMessage `json:"__deferred"`
	}

	// create the json tester
	tester := &utils.JsonTester{}
	err := json.Unmarshal(data, tester)
	if err != nil {
		return err
	}

	// test if json is array (standalone)
	if tester.IsArray() {
		results := &Results{}
		err := json.Unmarshal(data, results)
		if err != nil {
			return err
		}

		*o = PurchaseOrderLines(*results)
		return nil
	}

	envelope := &Envelope{Results: Results(*o)}
	err = json.Unmarshal(data, envelope)
	if err != nil {
		return err
	}

	*o = PurchaseOrderLines(envelope.Results)
	return nil
}

type PurchaseOrderLine struct {
	ID                      edm.GUID     `json:"ID"`
	AmountDC                edm.Double   `json:"AmountDC"`
	AmountFC                edm.Double   `json:"AmountFC"`
	CostCenter              edm.String   `json:"CostCenter"`
	CostCenterDescription   edm.String   `json:"CostCenterDescription"`
	CostPriceFC             edm.Double   `json:"CostPriceFC"`
	CostUnit                edm.String   `json:"CostUnit"`
	CostUnitDescription     edm.String   `json:"CostUnitDescription"`
	DeliveryDate            edm.DateTime `json:"DeliveryDate"`
	Description             edm.String   `json:"Description"`
	Discount                edm.Double   `json:"Discount"`
	Division                edm.Int32    `json:"Division"`
	Item                    edm.GUID     `json:"Item"`
	ItemCode                edm.String   `json:"ItemCode"`
	ItemDescription         edm.String   `json:"ItemDescription"`
	ItemVersion             edm.GUID     `json:"ItemVersion"`
	ItemVersionDescription  edm.String   `json:"ItemVersionDescription"`
	LineNumber              edm.Int32    `json:"LineNumber"`
	Margin                  edm.Double   `json:"Margin"`
	NetPrice                edm.Double   `json:"NetPrice"`
	Notes                   edm.String   `json:"Notes"`
	OrderID                 edm.GUID     `json:"OrderID"`
	OrderNumber             edm.Int32    `json:"OrderNumber"`
	Pricelist               edm.GUID     `json:"Pricelist"`
	PricelistDescription    edm.String   `json:"PricelistDescription"`
	Project                 edm.GUID     `json:"Project"`
	ProjectDescription      edm.String   `json:"ProjectDescription"`
	PurchaseOrder           edm.GUID     `json:"PurchaseOrder"`
	PurchaseOrderLine       edm.GUID     `json:"PurchaseOrderLine"`
	PurchaseOrderLineNumber edm.Int32    `json:"PurchaseOrderLineNumber"`
	PurchaseOrderNumber     edm.Int32    `json:"PurchaseOrderNumber"`
	Quantity                edm.Double   `json:"Quantity"`
	QuantityDelivered       edm.Double   `json:"QuantityDelivered"`
	QuantityInvoiced        edm.Double   `json:"QuantityInvoiced"`
	ShopOrder               edm.GUID     `json:"ShopOrder"`
	TaxSchedule             edm.GUID     `json:"TaxSchedule"`
	TaxScheduleCode         edm.String   `json:"TaxScheduleCode"`
	TaxScheduleDescription  edm.String   `json:"TaxScheduleDescription"`
	UnitCode                edm.String   `json:"UnitCode"`
	UnitDescription         edm.String   `json:"UnitDescription"`
	UnitPrice               edm.Double   `json:"UnitPrice"`
	UseDropShipment         edm.Byte     `json:"UseDropShipment"`
	VATAmount               edm.Double   `json:"VATAmount"`
	VATCode                 edm.String   `json:"VATCode"`
	VATCodeDescription      edm.String   `json:"VATCodeDescription"`
	VATPercentage           edm.Double   `json:"VATPercentage"`
}

const (
	OrderStatusOpen      OrderStatus = 12
	OrderStatusPartial   OrderStatus = 20
	OrderStatusComplete  OrderStatus = 21
	OrderStatusCancelled OrderStatus = 45 // following documentation spelling
)

type OrderStatus edm.Int16
