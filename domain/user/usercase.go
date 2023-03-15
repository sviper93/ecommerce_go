package user

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sviper93/ecommerce_go/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

func (u User) Create(m *model.User) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}
	m.ID = ID
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("%s %w", "GenerateFromPassword()", err)
	}
	m.Password = string(password)

	if m.Details == nil {
		m.Details = []byte("{}")
	}

	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	m.Password = ""
	return nil
}

func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "storage.GetByEmail()", err)
	}

	return user, nil
}

func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}

	return users, nil
}

/*
En los métodos "GetByEmail" y "GetAll" estamos devolviendo el password y si bien este está encriptado
no está bien que los devolvamos, así que vamos a hacer que en "GetAll" no se devuelva el password y en
"GetByEmail" si se devuelva porque este método servirá para hacer el método login.

Esto lo podemos hacer de varias formas, una de ellas es darla la responsabilidad al dominio de encargarse
de recorrer cada uno de los elementos del "GetAll" y limpiar el password, sin embargo, al hacer el "GetAll"
de la BD, ya se está haciendo una iteración por todos los registros, así que, no tiene sentido volver a darle
la responsabilidad al dominio de volver a recorrer los registros. Por lo que, la tarea de limpiar el
password se le asignará al paquete de BD por eficiencia, y así evitamos volver a recorrer el array de users.
*/
