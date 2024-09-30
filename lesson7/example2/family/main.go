package family

type Family struct {
	Name            string
	NumberOfMembers int
	Assets          *Assets
}

func NewFamily(name string, numberOfMembers int, assets *Assets) *Family {
	return &Family{
		Name:            name,
		NumberOfMembers: numberOfMembers,
		Assets:          assets,
	}
}

type Assets struct {
	NumberOfCars   int
	NumberOfAssets int
}

func NewAssets(numberOfCars, numberOfAssets int) *Assets {
	return &Assets{
		NumberOfCars:   numberOfCars,
		NumberOfAssets: numberOfAssets,
	}
}
