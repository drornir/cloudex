package app

type Product interface {
	Name() string
}

type ZapierProduct struct {
}

func (ZapierProduct) Name() string {
	return "Zapier"
}
