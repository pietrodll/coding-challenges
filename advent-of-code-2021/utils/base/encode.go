package base

type Coder interface {
	Encode(codable Codable) int
	Decode(encoded int) Codable
}

type Codable interface {
	Encode(coder Coder) int
}
