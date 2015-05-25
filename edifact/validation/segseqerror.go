package validation

import (
	"fmt"
)

type SegSeqErrorKind string

const (
	missingMandatorySegment SegSeqErrorKind = "missing_mandatory_segment"
	noMoreSegments          SegSeqErrorKind = "no_more_segments"
	maxRepeatCountExceeded  SegSeqErrorKind = "max_repeat_count_exceeded"
	missingGroup            SegSeqErrorKind = "missing_group"
	noSegmentSpecs          SegSeqErrorKind = "no_segment_specs"
	noSegments              SegSeqErrorKind = "no_segments"
	unexpectedSegment       SegSeqErrorKind = "unexpected_segment"
	unexpectedErr           SegSeqErrorKind = "unexpected_err"
)

// An exception that provides an error kind to check for specific error conditions
type SegSeqError struct {
	kind    SegSeqErrorKind
	message string
}

func (e SegSeqError) Error() string {
	return fmt.Sprintf("%s: %s", e.kind, e.message)
}

func NewSegSeqError(kind SegSeqErrorKind, message string) SegSeqError {
	if message == "" {
		message = string(kind)
	}
	return SegSeqError{kind, message}
}
