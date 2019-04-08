package pluginf

type PLUGTYPE string

/**
 * Info(info *PlugInfo){}
 */
type PlugInfo struct {
	Version  string
	PlugType PLUGTYPE
	Name     string
	Describe string
}

const (
	REGISTER PLUGTYPE = "REGISTER" //register plugin
	FILTER   PLUGTYPE = "FILTER"   //filter plugin
)

const (
	REGISTERFUNCNAME = "RegInit"
	FILTERFUNCNAME   = "RegInit"
	PLUGINFO         = "Info"
)
