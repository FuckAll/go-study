package mode

type IPayMode interface {
	Security(uID string) bool
}
