package structs



type File struct {
	Path	string
	status  StatusFile
	ReducePaths []string
	OutputPaths map[int]string


}
func NewFile(path string) File {
	return File{
		Path: path,
		status: NotProcessed,
		ReducePaths: []string{},
		OutputPaths: make(map[int]string),
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
