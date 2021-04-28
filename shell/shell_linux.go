package shell

const (
	// Shell constants
	bash = "/bin/bash"
	sh   = "/bin/sh"
)

func GetSystemShell() (string, string) {
	if exists(bash) {
		return bash, "-c"
	}
	return sh, "-c"
}
