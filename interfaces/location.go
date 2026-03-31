package interfaces

type LocationProvider interface {
	GetLocation(cep string) (string, error)
}
