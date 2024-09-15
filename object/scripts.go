package object

type Scripts struct {
	DefaultJScript []byte
	JScriptVersion []byte
}

func NewScripts() *Scripts {
	return &Scripts{
		DefaultJScript: nil,
		JScriptVersion: nil,
	}
}

func (s *Scripts) GetDefaultJScript() []byte {
	return s.DefaultJScript
}

func (s *Scripts) SetDefaultJScript(defaultJScript []byte) {
	s.DefaultJScript = defaultJScript
}

func (s *Scripts) GetJScriptVersion() []byte {
	return s.JScriptVersion
}

func (s *Scripts) SetJScriptVersion(jScriptVersion []byte) {
	s.JScriptVersion = jScriptVersion
}

func (s *Scripts) Copy(from *Scripts) {
	s.DefaultJScript = from.DefaultJScript
	s.JScriptVersion = from.JScriptVersion
}
