package fileheader

// FileVersion represents a file version
type FileVersion struct {
	mm, nn, pp, rr uint8
}

// NewFileVersion creates a new FileVersion instance
func NewFileVersion() *FileVersion {
	return &FileVersion{}
}
