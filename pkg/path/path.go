package path

import "path"

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (*Path) PathIsValid(uri string) bool {
	cleanPath := path.Clean(uri) // the .Clean() clean the function by 
	// parsing single and double ellipses ("." & "..")
	// and replace "//" ==> to "/"
	return uri == cleanPath
}