package nifcloud

// Region is struct of Region
type Region struct {
	Name              string
	ComputingEndpoint string
}

// DefaultRegion is jp-east-1
var DefaultRegion = JPEast1

// Regions is map that <regionName, Region>
var Regions = map[string]Region{
	JPEast1.Name: JPEast1,
	JPEast2.Name: JPEast2,
	JPEast3.Name: JPEast3,
	JPEast4.Name: JPEast4,
	JPWest1.Name: JPWest1,
	USEast1.Name: USEast1,
}

// JPEast1 is jp-east-1 region
var JPEast1 = Region{
	"jp-east-1",
	"https://computing.jp-east-1.api.cloud.nifty.com/api/",
}

// JPEast2 is jp-east-2 region
var JPEast2 = Region{
	"jp-east-2",
	"https://computing.jp-east-2.api.cloud.nifty.com/api/",
}

// JPEast3 is jp-east-3 region
var JPEast3 = Region{
	"jp-east-3",
	"https://computing.jp-east-3.api.cloud.nifty.com/api/",
}

// JPEast4 is jp-east-4 region
var JPEast4 = Region{
	"jp-east-4",
	"https://computing.jp-east-4.api.cloud.nifty.com/api/",
}

// JPWest1 is jp-west-1 region
var JPWest1 = Region{
	"jp-west-1",
	"https://computing.jp-west-1.api.cloud.nifty.com/api/",
}

// USEast1 is us-east-1 region
var USEast1 = Region{
	"us-east-1",
	"https://computing.us-east-1.api.cloud.nifty.com/api/",
}

// NewRegion returs Region
func NewRegion(regionName string) (region Region) {
	for key, value := range Regions {
		if key == regionName {
			region = value
			return
		}
	}

	region = DefaultRegion
	return
}
