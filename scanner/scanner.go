package scanner

import (
	"errors"
	"sort"

	"github.com/bookrun-go/calculator/token"
)

func NewScanner(fomula string, opts ...ScannerOption) *Scanner {
	scnanner := &Scanner{
		fomula:     []rune(fomula),
		currentPos: -1,
	}
	scnanner.fomulaLen = len(fomula)

	for _, opt := range opts {
		opt(scnanner)
	}

	return scnanner
}

type Scanner struct {
	scanners []IScanner
	fomula   []rune

	// 临时变量，随遍历推移
	currentScanner IScanner
	currentPos     int
	fomulaLen      int
}

func (s *Scanner) Scan() ([]*token.TokenValue, error) {
	var tokens []*token.TokenValue

	for {
		char, ok := s.currentPosition()
		if !ok {
			return tokens, nil
		}
		if s.skipWhitespace(char) {
			continue
		}

		err := s.setCurrentScanner(char)
		if err != nil {
			return nil, err
		}

		tv, nextPos := s.currentScanner.Scan(s.fomula, s.currentPos)
		if tv.Tok.IsIllegal() {
			return nil, nil
		}
		s.currentPos = nextPos

		tokens = append(tokens, tv)
	}
}

func (s *Scanner) setCurrentScanner(char rune) error {
	for _, scanner := range s.scanners {
		ok := scanner.Of(char)
		if ok {
			s.currentScanner = scanner
			return nil
		}
	}

	return errors.New("not found")
}

func (s *Scanner) currentPosition() (rune, bool) {
	if s.currentPos > s.fomulaLen-1 {
		return 0, false
	}

	return s.fomula[s.currentPos], true
}

func (s *Scanner) skipWhitespace(char rune) bool {
	for char == ' ' || char == '\t' || char == '\n' || char == '\r' {
		s.currentPos++
		return true
	}

	return false
}

// 扫描器按优先级从小到大排序
// 每次有新扫描器加入后需要执行一下排序
func (s *Scanner) Sort() {
	if len(s.scanners) == 0 {
		return
	}

	sort.Slice(s.scanners, func(i, j int) bool {
		return s.scanners[i].Precedence() > s.scanners[j].Precedence()
	})
}
