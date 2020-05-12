package logging

import (
	"testing"
)

var _ Logger = &TestLogger{}

// TestLogger implements a test wrapper for testing.T implementing Logger interface
type TestLogger struct {
	t         *testing.T
	neverFail bool
	fatals    int
	errors    int
	panics    int
}

// NewTestLogger returns a new TestLogger
func NewTestLogger(t *testing.T) *TestLogger {
	return &TestLogger{
		t: t,
	}
}

// NewNonFailingTestLogger returns a new TestLogger
func NewNonFailingTestLogger(t *testing.T) *TestLogger {
	return &TestLogger{
		t:         t,
		neverFail: true,
	}
}

// HasFailures returns true if any errors, fatals, or panics have been logged
func (l *TestLogger) HasFailures() bool {
	return l.fatals+l.errors+l.panics > 0
}

// Fatals returns the number of fatals that have been logged
func (l *TestLogger) Fatals() int {
	return l.fatals
}

// Errors returns the number of errors that have been logged
func (l *TestLogger) Errors() int {
	return l.errors
}

// Panics returns the number of panics that have been logged
func (l *TestLogger) Panics() int {
	return l.panics
}

// Debugf implements Logger interface
func (l *TestLogger) Debugf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Infof implements Logger interface
func (l *TestLogger) Infof(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Printf implements Logger interface
func (l *TestLogger) Printf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Warnf implements Logger interface
func (l *TestLogger) Warnf(format string, args ...interface{}) {
	l.t.Logf(format, args...)
}

// Errorf implements Logger interface
func (l *TestLogger) Errorf(format string, args ...interface{}) {
	l.errors++
	if l.neverFail {
		l.t.Logf(format, args...)
		return
	}

	l.t.Errorf(format, args...)
}

// Fatalf implements Logger interface
func (l *TestLogger) Fatalf(format string, args ...interface{}) {
	l.fatals++
	if l.neverFail {
		l.t.Logf(format, args...)
		return
	}

	l.t.Fatalf(format, args...)
}

// Panicf implements Logger interface
func (l *TestLogger) Panicf(format string, args ...interface{}) {
	l.panics++
	if l.neverFail {
		l.t.Logf(format, args...)
		return
	}

	l.t.Logf(format, args...)
	l.t.FailNow()
}

// Debug implements Logger interface
func (l *TestLogger) Debug(args ...interface{}) {
	l.t.Log(args...)
}

// Info implements Logger interface
func (l *TestLogger) Info(args ...interface{}) {
	l.t.Log(args...)
}

// Print implements Logger interface
func (l *TestLogger) Print(args ...interface{}) {
	l.t.Log(args...)
}

// Println implements Logger interface
func (l *TestLogger) Println(args ...interface{}) {
	l.t.Log(args...)
}

// Warn implements Logger interface
func (l *TestLogger) Warn(args ...interface{}) {
	l.t.Log(args...)
}

// Error implements Logger interface
func (l *TestLogger) Error(args ...interface{}) {
	l.errors++
	if l.neverFail {
		l.t.Log(args...)
		return
	}

	l.t.Error(args...)
}

// Fatal implements Logger interface
func (l *TestLogger) Fatal(args ...interface{}) {
	l.fatals++
	if l.neverFail {
		l.t.Log(args...)
		return
	}

	l.t.Fatal(args...)
}

// Panic implements Logger interface
func (l *TestLogger) Panic(args ...interface{}) {
	l.panics++
	if l.neverFail {
		l.t.Log(args...)
		return
	}

	l.t.Log(args...)
	l.t.FailNow()
}
