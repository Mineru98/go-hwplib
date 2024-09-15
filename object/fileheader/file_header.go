package fileheader

// FileHeader represents the file recognition information. It is stored in the "FileHeader" stream within the HWP file.
type FileHeader struct {
	version                          *FileVersion
	compressed                       bool
	hasPassword                      bool
	isDistribution                   bool
	saveScript                       bool
	isDRMDocument                    bool
	hasXMLTemplate                   bool
	hasDocumentHistory               bool
	hasSignature                     bool
	encryptPublicCertification       bool
	savePrepareSignature             bool
	isPublicCertificationDRMDocument bool
	isCCLDocument                    bool
}

// NewFileHeader creates a new instance of FileHeader
func NewFileHeader() *FileHeader {
	return &FileHeader{
		version: NewFileVersion(),
	}
}

// GetFileSignature returns the signature of the HWP file. The signature is used to check whether the file is an HWP file.
func GetFileSignature() []byte {
	return []byte{0x48, 0x57, 0x50, 0x20, 0x44, 0x6f, 0x63, 0x75, 0x6d,
		0x65, 0x6e, 0x74, 0x20, 0x46, 0x69, 0x6c, 0x65, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00}
}

// GetVersion returns the version information object.
func (fh *FileHeader) GetVersion() *FileVersion {
	return fh.version
}

// IsCompressed returns whether the file is compressed.
func (fh *FileHeader) IsCompressed() bool {
	return fh.compressed
}

// SetCompressed sets whether the file is compressed.
func (fh *FileHeader) SetCompressed(compressed bool) {
	fh.compressed = compressed
}

// HasPassword returns whether a password is set.
func (fh *FileHeader) HasPassword() bool {
	return fh.hasPassword
}

// SetHasPassword sets whether a password is set.
func (fh *FileHeader) SetHasPassword(hasPassword bool) {
	fh.hasPassword = hasPassword
}

// IsDistribution returns whether the file is for distribution.
func (fh *FileHeader) IsDistribution() bool {
	return fh.isDistribution
}

// SetDistribution sets whether the file is for distribution.
func (fh *FileHeader) SetDistribution(isDistribution bool) {
	fh.isDistribution = isDistribution
}

// IsSaveScript returns whether the script is saved.
func (fh *FileHeader) IsSaveScript() bool {
	return fh.saveScript
}

// SetSaveScript sets whether the script is saved.
func (fh *FileHeader) SetSaveScript(saveScript bool) {
	fh.saveScript = saveScript
}

// IsDRMDocument returns whether the file is a DRM-secured document.
func (fh *FileHeader) IsDRMDocument() bool {
	return fh.isDRMDocument
}

// SetDRMDocument sets whether the file is a DRM-secured document.
func (fh *FileHeader) SetDRMDocument(isDRMDocument bool) {
	fh.isDRMDocument = isDRMDocument
}

// HasXMLTemplate returns whether the XMLTemplate storage exists.
func (fh *FileHeader) HasXMLTemplate() bool {
	return fh.hasXMLTemplate
}

// SetHasXMLTemplate sets whether the XMLTemplate storage exists.
func (fh *FileHeader) SetHasXMLTemplate(hasXMLTemplate bool) {
	fh.hasXMLTemplate = hasXMLTemplate
}

// HasDocumentHistory returns whether the document history exists.
func (fh *FileHeader) HasDocumentHistory() bool {
	return fh.hasDocumentHistory
}

// SetHasDocumentHistory sets whether the document history exists.
func (fh *FileHeader) SetHasDocumentHistory(hasDocumentHistory bool) {
	fh.hasDocumentHistory = hasDocumentHistory
}

// HasSignature returns whether the electronic signature information exists.
func (fh *FileHeader) HasSignature() bool {
	return fh.hasSignature
}

// SetHasSignature sets whether the electronic signature information exists.
func (fh *FileHeader) SetHasSignature(hasSignature bool) {
	fh.hasSignature = hasSignature
}

// IsEncryptPublicCertification returns whether the public certification is encrypted.
func (fh *FileHeader) IsEncryptPublicCertification() bool {
	return fh.encryptPublicCertification
}

// SetEncryptPublicCertification sets whether the public certification is encrypted.
func (fh *FileHeader) SetEncryptPublicCertification(encryptPublicCertification bool) {
	fh.encryptPublicCertification = encryptPublicCertification
}

// IsSavePrepareSignature returns whether the electronic signature is preliminarily saved.
func (fh *FileHeader) IsSavePrepareSignature() bool {
	return fh.savePrepareSignature
}

// SetSavePrepareSignature sets whether the electronic signature is preliminarily saved.
func (fh *FileHeader) SetSavePrepareSignature(savePrepareSignature bool) {
	fh.savePrepareSignature = savePrepareSignature
}

// IsPublicCertificationDRMDocument returns whether the document is a public certification DRM-secured document.
func (fh *FileHeader) IsPublicCertificationDRMDocument() bool {
	return fh.isPublicCertificationDRMDocument
}

// SetPublicCertificationDRMDocument sets whether the document is a public certification DRM-secured document.
func (fh *FileHeader) SetPublicCertificationDRMDocument(isPublicCertificationDRMDocument bool) {
	fh.isPublicCertificationDRMDocument = isPublicCertificationDRMDocument
}

// IsCCLDocument returns whether the document is a CCL document.
func (fh *FileHeader) IsCCLDocument() bool {
	return fh.isCCLDocument
}

// SetCCLDocument sets whether the document is a CCL document.
func (fh *FileHeader) SetCCLDocument(isCCLDocument bool) {
	fh.isCCLDocument = isCCLDocument
}

// Copy copies the FileHeader from another instance.
func (fh *FileHeader) Copy(from *FileHeader) {
	fh.version.mm = from.version.mm
	fh.version.nn = from.version.nn
	fh.version.pp = from.version.pp
	fh.version.rr = from.version.rr

	fh.compressed = from.compressed
	fh.hasPassword = from.hasPassword
	fh.isDistribution = from.isDistribution
	fh.saveScript = from.saveScript
	fh.isDRMDocument = from.isDRMDocument
	fh.hasXMLTemplate = from.hasXMLTemplate
	fh.hasDocumentHistory = from.hasDocumentHistory
	fh.hasSignature = from.hasSignature
	fh.encryptPublicCertification = from.encryptPublicCertification
	fh.savePrepareSignature = from.savePrepareSignature
	fh.isPublicCertificationDRMDocument = from.isPublicCertificationDRMDocument
	fh.isCCLDocument = from.isCCLDocument
}
