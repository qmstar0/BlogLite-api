package CQRS

type Topic interface {
	Topic() string
}

type TypeEvent interface {
	Topic
}

type TypeCommand interface {
	Topic
}
