package structs

type Shelve struct {
	files         []File
	filesFinished int
}

func (s *Shelve) AddFiles(file_path []string) {
	for _, path := range file_path {
		s.files = append(s.files, NewFile(path))
	}
}

func (s *Shelve) GetNextFile() *File {
	for i := range s.files {
		if s.files[i].getStatus() == NotProcessed {
			s.files[i].setStatus(InProgress)
			return &s.files[i]
		}
	}
	return nil
}

func (s *Shelve) MarkFileFinished(file *File) {
	file.status = Processed
	s.filesFinished++
}

func (s *Shelve) AllFilesFinished() bool {
	return s.filesFinished == len(s.files)
}
