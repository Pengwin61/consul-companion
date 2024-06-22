package cfg

import "consul-companion/internal/core"

func SDConfig(search string, confdir string) {

	core.RootProjectPath = search
	core.CONFDIR = confdir

}
