package purchaseorder

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sviper93/ecommerce_go/model"
)

// PurchaseOrder implements UseCase
type PurchaseOrder struct {
	storage Storage
}

// New returns a new PurchaseOrder
func New(s Storage) PurchaseOrder {
	return PurchaseOrder{storage: s}
}

// Create creates a model.PurchaseOrder
// Cuando recibamos la creación de una orden de compra, es decir cuando el clic de clic en el botón "comprar"
// vamos a generar una orden de compra, la almacenados en la BD con el usuario que está comprando.
func (p PurchaseOrder) Create(m *model.PurchaseOrder) error {
	if err := m.Validate(); err != nil { // Luego validamos que esa orden de compra tenga su JSON de los
		// objetos con los productos que va a comprar estén correctos, es decir acá estamos validando la
		// estructura.
		// Tarea: Validar que los objetos tienen el ProductID correctos porque si no lo valido podrían enviar
		// un UUID válido como estructura pero puede ser que no exista ese producto, así aparte de validar la
		// estructura debemos validar que los productos que están ingresados son correctos, es decir que existan
		// y en lugar de recibir el valor del cliente.
		return fmt.Errorf("purchaseorder: %w", err)
	}

	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	m.ID = ID
	m.CreatedAt = time.Now().Unix()

	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

func (p PurchaseOrder) GetByID(ID uuid.UUID) (model.PurchaseOrder, error) {
	purchaseOrder, err := p.storage.GetByID(ID)
	if err != nil {
		return model.PurchaseOrder{}, fmt.Errorf("purchaseorder: %w", err)
	}

	return purchaseOrder, nil
}
