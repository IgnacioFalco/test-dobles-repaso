package directorio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpyDB struct {
	BuscarPorTelefonoWasCalled bool
}

func (s *SpyDB) BuscarPorNombre(nombre string) string {
	return ""
}

func (s *SpyDB) BuscarPorTelefono(telefono string) string {
	s.BuscarPorTelefonoWasCalled = true
	return ""
}

func (s *SpyDB) AgregarEntrada(nombre, telefono string) error {
	return nil
}

func TestFindByTelephone(t *testing.T) {
	//arrange

	mySpyDB := SpyDB{BuscarPorTelefonoWasCalled: false}
	motor := NewEngine(&mySpyDB)
	telefono := "12345678"

	//act
	motor.FindByTelephone(telefono)

	//assert
	assert.True(t, mySpyDB.BuscarPorTelefonoWasCalled)
}

func TestFindByTelephoneNotCalled(t *testing.T) {
	//arrange

	mySpyDB := SpyDB{}
	motor := NewEngine(&mySpyDB)

	//act
	motor.FindByTelephone("1234")

	//assert
	assert.False(t, mySpyDB.BuscarPorTelefonoWasCalled)
}
