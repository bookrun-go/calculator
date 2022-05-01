package scanner

type ScannerOption func(*Scanner)

func WithAddScanners(scanner ...IScanner) ScannerOption {
	return func(s *Scanner) {
		s.scanners = append(s.scanners, scanner...)
	}
}
