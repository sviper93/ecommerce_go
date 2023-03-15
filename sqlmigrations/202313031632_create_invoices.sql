CREATE TABLE invoices (
	id UUID NOT NULL,
	user_id UUID NOT NULL,
	purchase_order_id UUID NOT NULL, -- Cada factura tiene un orden de compra, incluso una orden de compra --
	-- pueden tener varias facturas porque talvez una persona que haya comprado quiera comprar exactamente --
	-- lo mismo sobre la misma orden de compra, aunque esto no debería pasar porque se supone que una orden --
	-- de compra es inmutable por lo tanto es mejor generar una nueva orden de compra --
	created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
	updated_at INTEGER,
	CONSTRAINT invoices_id_pk PRIMARY KEY (id),
	CONSTRAINT invoices_user_id_fk FOREIGN KEY (user_id)
            REFERENCES users (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    CONSTRAINT invoices_purchase_order_id_fk FOREIGN KEY (purchase_order_id)
            REFERENCES purchase_orders (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);

COMMENT ON TABLE invoices IS 'Storage the head of the invoices';
