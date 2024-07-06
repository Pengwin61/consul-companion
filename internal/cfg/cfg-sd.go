package cfg

var (
	ROOT_PROJECT_PATH string
	CONFDIR           string
	TMP_DIR           = "/tmp/consul-companion"
)

func SDConfig(search string, confdir string) {

	ROOT_PROJECT_PATH = search
	CONFDIR = confdir

}
