package main

import (
    "testing"
    "github.com/Tch1b0/datastream/pkg/chunks"
)

func TestSplitter(t *testing.T) {
    t.Run("Test the Data Splitter", func(t *testing.T) {
        x := chunks.SplitData([]byte("Hello, this is a test message used for a unit test!"), 10)
        if len(*x) != 6 {
            t.Errorf("Expected length of x to be 6, but value is %d", len(*x))
        }
    })
}

func TestChunks(t *testing.T) {
    t.Run("Test chunks", func(t *testing.T) {
        c := chunks.NewChunk([]byte("This is a test"))
        b := chunks.NewChunk([]byte("Another Test"))
        g := chunks.NewChunk([]byte("This is a test"))
        if c.Checksum == b.Checksum {
            t.Error("Checksum of c and b are equal")
        }
        if c.Checksum != g.Checksum {
            t.Errorf("Checksum of c %s is not equal to g %s", c.Checksum, g.Checksum)
        }
    })
}
