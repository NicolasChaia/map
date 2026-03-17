package structs
type File struct {
	Path	string
	status  StatusFile
	ReducePaths []string
	OutputPath string


}
func NewFile(path string) File {
	return File{
		Path: path,
		status: NotProcessed,
		ReducePaths: []string{},
	}
}

func (f *File) getStatus() StatusFile {
	return f.status
}

func (f *File) setStatus(status StatusFile) {
	f.status = status
}

func (f *File) addReducePath(path string) {
	f.ReducePaths = append(f.ReducePaths, path)
}
