package entity

import (
	coreSwitchProto "github.com/hdkef/hadoop/proto/coreSwitch"
)

type WriteDto struct {
	file []byte
}

func (w *WriteDto) Tokenize(numParts int) [][]byte {
	// Calculate the size of each part
	partSize := (len(w.file) + numParts - 1) / numParts // Round up to ensure coverage of all bytes
	tokens := make([][]byte, 0)
	for i := 0; i < len(w.file); i += partSize {
		end := i + partSize
		if end > len(w.file) {
			end = len(w.file)
		}
		tokens = append(tokens, w.file[i:end])
	}
	return tokens
}

func (w *WriteDto) NewFromProto(req *coreSwitchProto.WriteReq) {

}
