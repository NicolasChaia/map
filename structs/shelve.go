package structs

type Shelve struct {
	files         []File
	filesMapFinished int
	filesFinished int
}

func NewShelve() Shelve {
	return Shelve{
		files:         []File{},
		filesMapFinished: 0,
		filesFinished: 0,
	}
}

func (s *Shelve) AddFiles(file_path []string) {
	for _, path := range file_path {
		s.files = append(s.files, NewFile(path))
	}
}

func (s *Shelve) GetNextFileMap() *File {
	for i := range s.files {
		if s.files[i].getStatus() == NotProcessed {
			s.files[i].setStatus(MapInProgress)
			return &s.files[i]
		}
	}
	return nil
}

func (s *Shelve) GetNextFileReduce() *File {
	for i := range s.files {
		if s.files[i].getStatus() == Mapped {
			s.files[i].setStatus(ReduceInProgress)
			return &s.files[i]
		}
	}
	return nil
}

func (s *Shelve) MarkMapFinished (file *File) {
	s.filesMapFinished++
	file.setStatus(Mapped)
}

func (s *Shelve) MarkFileFinished(file *File) {
	file.status = Finished
	s.filesFinished++
}

func (s *Shelve) AllFilesFinished() bool {
	return s.filesFinished == len(s.files)
}

func (s *Shelve) AllFilesMapped() bool {
	return s.filesMapFinished == len(s.files)
}