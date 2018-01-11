package nifcloud

type Region struct {
	Name              string
	ComputingEndpoint string
}

var DefaultRegion = JPEast1

var Regions = map[string]Region{
	JPEast1.Name: JPEast1,
	JPEast2.Name: JPEast2,
	JPEast3.Name: JPEast3,
	JPEast4.Name: JPEast4,
	JPWest1.Name: JPWest1,
	USEast1.Name: USEast1,
}

var JPEast1 = Region{
	"jp-east-1",
	"https://computing.jp-east-1.api.cloud.nifty.com/api/",
}

var JPEast2 = Region{
	"jp-east-2",
	"https://computing.jp-east-2.api.cloud.nifty.com/api/",
}

var JPEast3 = Region{
	"jp-east-3",
	"https://computing.jp-east-3.api.cloud.nifty.com/api/",
}

var JPEast4 = Region{
	"jp-east-4",
	"https://computing.jp-east-4.api.cloud.nifty.com/api/",
}

var JPWest1 = Region{
	"jp-west-1",
	"https://computing.jp-west-1.api.cloud.nifty.com/api/",
}

var USEast1 = Region{
	"us-east-1",
	"https://computing.us-east-1.api.cloud.nifty.com/api/",
}

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
