package fhir

type BaseResource interface {
	MarshalJSON() ([]byte, error)
}
