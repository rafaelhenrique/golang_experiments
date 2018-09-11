/*
Reference: https://gist.github.com/Avinash-Bhat/48c4f06b0cc840d9fd6c
*/

package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

type NopLogger struct {
	*log.Logger
}

func (l *NopLogger) Println(v ...interface{}) {
	// noop
}

var nop *NopLogger = &NopLogger{
	log.New(os.Stderr, "", log.LstdFlags),
}

func BenchmarkLog(b *testing.B) {
	log.SetFlags(log.LstdFlags)
	file, err := ioutil.TempFile("", "benchmark-log")
	if err != nil {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(file)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("Benchmark %d", i)
	}
}

func BenchmarkLogIoDiscardWriter(b *testing.B) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(ioutil.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("Benchmark %d\n", i)
	}
}

func BenchmarkLogIoDiscardWriterWithoutFlags(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("Benchmark %d\n", i)
	}
}

func BenchmarkLogCustomNullWriter(b *testing.B) {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(new(NullWriter))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("Benchmark %d\n", i)
	}
}

func BenchmarkLogCustomNullWriterWithoutFlags(b *testing.B) {
	log.SetFlags(0)
	log.SetOutput(new(NullWriter))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Printf("Benchmark %d\n", i)
	}
}

func BenchmarkNopLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nop.Printf("Benchmark %d\n", i)
	}
}

func BenchmarkNopLoggerWithoutFlags(b *testing.B) {
	log.SetFlags(0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nop.Printf("Benchmark %d\n", i)
	}
}
