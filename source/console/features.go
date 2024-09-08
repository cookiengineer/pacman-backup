package console

const (
	FeatureAll = 0
	FeatureGroup = 1
	FeatureLog = 2
	FeatureInfo = 3
	FeatureWarn = 4
	FeatureError = 5
	FeatureInspect = 6
	FeatureProgress = 7
)

var features map[int]bool

func init() {

	features = make(map[int]bool)

	features[FeatureGroup] = true
	features[FeatureLog] = true
	features[FeatureInfo] = true
	features[FeatureWarn] = true
	features[FeatureError] = true
	features[FeatureInspect] = true
	features[FeatureProgress] = true

}


