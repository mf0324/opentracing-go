package main

import (
	"fmt"
	"reflect"

	"github.com/opentracing/api-golang/opentracing"
)

// TrivialRecorder implements the opentracing.Recorder interface.
type TrivialRecorder struct {
	processName string
	tags        map[string]string
}

// NewTrivialRecorder returns a TrivialRecorder for the given `processName`.
func NewTrivialRecorder(processName string) *TrivialRecorder {
	return &TrivialRecorder{
		processName: processName,
		tags:        make(map[string]string),
	}
}

// ProcessName complies with the opentracing.ProcessIdentifier interface.
func (t *TrivialRecorder) ProcessName() string { return t.processName }

// SetTag complies with the opentracing.ProcessIdentifier interface.
func (t *TrivialRecorder) SetTag(key string, val interface{}) opentracing.ProcessIdentifier {
	t.tags[key] = fmt.Sprint(val)
	return t
}

// RecordSpan complies with the opentracing.Recorder interface.
func (t *TrivialRecorder) RecordSpan(span *opentracing.RawSpan) {
	fmt.Printf(
		"RecordSpan: %v[%v, %v us] --> %v logs. trace context: %v\n",
		span.Operation, span.Start, span.Duration, len(span.Logs),
		span.TraceContext)
	for i, l := range span.Logs {
		fmt.Printf(
			"    log %v: %v --> %v\n", i, l.Message, reflect.TypeOf(l.Payload))
	}
}
