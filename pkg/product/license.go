package product

type License interface {
	Product() Product
	Credentials() string
}

type LicenseMeta struct {
	ID     int64
	UserID int64
}

type LicenseAndMeta struct {
	License License
	Meta    LicenseMeta
}
