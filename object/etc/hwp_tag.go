package etc

const (
	// tag start value
	BEGIN = 0x10

	// for "DocInfo" stream
	DOCUMENT_PROPERTIES = BEGIN + iota // 문서 속성 tag
	ID_MAPPINGS                        // 아이디 매핑 헤더 tag
	BIN_DATA                           // BinData tag
	FACE_NAME                          // Typeface Name tag
	BORDER_FILL                        // 테두리/배경 tag
	CHAR_SHAPE                         // 글자 모양 tag
	TAB_DEF                            // 탭 정의 tag
	NUMBERING                          // 번호 정의 tag
	BULLET                             // 글머리표 tag
	PARA_SHAPE                         // 문단 모양 tag
	STYLE                              // 스타일 tag
	DOC_DATA                           // 문서의 임의의 데이터 tag
	DISTRIBUTE_DOC_DATA                // 배포용 문서 데이터 tag
	// Skipping one value (BEGIN + 13 is not used)
	COMPATIBLE_DOCUMENT  = BEGIN + 14 // 호환 문서 tag
	LAYOUT_COMPATIBILITY = BEGIN + 15 // 레이아웃 호환성 tag
	TRACKCHANGE          = BEGIN + 16 // 변경 추적 정보 tag
	MEMO_SHAPE           = BEGIN + 76 // 메모 모양 tag
	FORBIDDEN_CHAR       = BEGIN + 78 // 금칙처리 문자 tag
	TRACK_CHANGE         = BEGIN + 80 // 변경 추적 내용 및 모양 tag
	TRACK_CHANGE_AUTHOR  = BEGIN + 81 // 변경 추적 작성자 tag

	// for "BodyText" storages
	PARA_HEADER               = BEGIN + 50 // 문단 헤더 tag
	PARA_TEXT                 = BEGIN + 51 // 문단의 텍스트 tag
	PARA_CHAR_SHAPE           = BEGIN + 52 // 문단의 글자 모양 tag
	PARA_LINE_SEG             = BEGIN + 53 // 문단의 레이아웃 tag
	PARA_RANGE_TAG            = BEGIN + 54 // 문단의 영역 태그 tag
	CTRL_HEADER               = BEGIN + 55 // 컨트롤 헤더 tag
	LIST_HEADER               = BEGIN + 56 // 문단 리스트 헤더 tag
	PAGE_DEF                  = BEGIN + 57 // 용지 설정 tag
	FOOTNOTE_SHAPE            = BEGIN + 58 // 각주/미주 모양 tag
	PAGE_BORDER_FILL          = BEGIN + 59 // 쪽 테두리/배경 tag
	SHAPE_COMPONENT           = BEGIN + 60 // 개체 tag
	TABLE                     = BEGIN + 61 // 표 개체 tag
	SHAPE_COMPONENT_LINE      = BEGIN + 62 // 직선 개체 tag
	SHAPE_COMPONENT_RECTANGLE = BEGIN + 63 // 사각형 개체 tag
	SHAPE_COMPONENT_ELLIPSE   = BEGIN + 64 // 타원 개체 tag
	SHAPE_COMPONENT_ARC       = BEGIN + 65 // 호 개체 tag
	SHAPE_COMPONENT_POLYGON   = BEGIN + 66 // 다각형 개체 tag
	SHAPE_COMPONENT_CURVE     = BEGIN + 67 // 곡선 개체 tag
	SHAPE_COMPONENT_OLE       = BEGIN + 68 // OLE 개체 tag
	SHAPE_COMPONENT_PICTURE   = BEGIN + 69 // 그림 개체 tag
	SHAPE_COMPONENT_CONTAINER = BEGIN + 70 // 컨테이너 개체 tag
	CTRL_DATA                 = BEGIN + 71 // 컨트롤 임의의 데이터 tag
	EQEDIT                    = BEGIN + 72 // 수식 개체 tag
	SHAPE_COMPONENT_TEXTART   = BEGIN + 74 // 글맵시 tag
	FORM_OBJECT               = BEGIN + 75 // 양식 개체 tag
	MEMO_LIST                 = BEGIN + 77 // 메모 리스트 헤더 tag
	CHART_DATA                = BEGIN + 79 // 차트 데이터 tag
	VIDEO_DATA                = BEGIN + 82 // 비디오 데이터 tag
	SHAPE_COMPONENT_UNKNOWN   = BEGIN + 99 // Unknown tag
)
