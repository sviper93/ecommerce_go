package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// PurchaseOrder model of table purchase_orders
type PurchaseOrder struct {
	ID        uuid.UUID       `json:"id"`
	UserID    uuid.UUID       `json:"user_id"`
	Products  json.RawMessage `json:"products"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
}

func (p PurchaseOrder) HasID() bool {
	return p.ID != uuid.Nil
}

// Validate valida que el json enviado desde la orden de compra ("Products" json.RawMessage) sea un json
// válido, para eso validamos si
func (p PurchaseOrder) Validate() error {
	if len(p.Products) == 0 { // viene vacío devolvemos un error, indicando que los productos no pueden
		return errors.New("the list of products can't be empty") // estar vacíos
	}

	var ptps []ProductToPurchase             // Y si en caso no viene vacío, almacenamos el JSON dentro de
	err := json.Unmarshal(p.Products, &ptps) // este array, con "Unmarshal" convertimos el JSON al slice de
	// productos definido en la estructura "ProductToPurchase" de abajo.
	if err != nil {
		return fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	for _, v := range ptps { // recorremos cada producto que está dentro del JSON y validamos los campos
		if v.ProductID == uuid.Nil { // ProductID no puede estar vacío
			return errors.New("the product id can´t be empty")
		}
		if v.Amount < 1 { // No puede ser menor a 1
			return errors.New("the amount of products can't be less than 1")
		}
		if v.UnitPrice < 0 { // No puede ser menor a 0 porque podríamos tener un producto gratuito y tenemos
			// que dar en la orden de compra el precio 0
			return errors.New("the unit price can't be negative")
		}
	}

	return nil
}

// TotalAmount nos permite recorrer el array de objetos que están dentro de JSON
func (p PurchaseOrder) TotalAmount() float64 {
	if len(p.Products) == 0 {
		return 0
	}

	var ptps []ProductToPurchase
	err := json.Unmarshal(p.Products, &ptps)
	if err != nil {
		return 0
	}

	var total float64
	for _, v := range ptps {
		total += float64(v.Amount) * v.UnitPrice // Este total es con el cual vamos a validar que el pago
		// que han enviado desde el PayPal sea el correcto
	}

	return total
}

// PurchaseOrders slice of PurchaseOrder
type PurchaseOrders []PurchaseOrder

func (p PurchaseOrders) IsEmpty() bool { return len(p) == 0 }

type ProductToPurchase struct {
	ProductID uuid.UUID `json:"product_id"`
	Amount    uint      `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
}

type ProductToPurchases []ProductToPurchase
