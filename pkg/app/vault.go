package app

// the place to export API_KEYs to
type HashicorpVault struct{}

func (v *HashicorpVault) SaveSecret(key, value string) error {
	panic("unimplemented")
}
