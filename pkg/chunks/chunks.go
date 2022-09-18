package chunks

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)

type Chunk struct {
    Data     []byte
    Checksum string
}

func (c *Chunk) String() string {
    return fmt.Sprintf("{Data: \"%s\", Checksum: %s}", c.Data, c.Checksum)
}

func NewChunk(data []byte) *Chunk {
    tmp := md5.Sum(data)
    return &Chunk{
        Data: data,
        Checksum: hex.EncodeToString(tmp[:]),
    }
}

func SplitData(data []byte, chunkSize int) *[]*Chunk {
    chunks := []*Chunk{}
    // temporary storage for collection chunk data
    cache := []byte{}
    for _, b := range data {
        if len(cache) < chunkSize {
            cache = append(cache, b)
        } else {
            chunk := NewChunk(cache)
            chunks = append(chunks, chunk)
            cache = []byte{b}
        }
    }
    if len(cache) != 0 {
        // append unfinished chunk data
        chunk := NewChunk(cache)
        chunks = append(chunks, chunk)
    }
    return &chunks
}
