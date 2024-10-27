package block

import (
	"io"

	"github.com/63square/cm2go2/builder"
)

type Creation struct {
	blocks      []byte
	connections []byte

	blockIndex uint64
}

func (c *Creation) Add(blockType byte, x float64, y float64, z float64, properties []float64) uint64 {
	c.blocks = builder.MakeBlock(c.blocks, blockType, x, y, z, properties)
	c.blockIndex += 1

	return c.blockIndex
}

func (c *Creation) Connect(src uint64, dst uint64) {
	c.connections = builder.ConnectBlock(c.connections, src, dst)
}

func (c Creation) Compile(writer io.Writer) (err error) {
	_, err = writer.Write(c.blocks[:len(c.blocks)-1])
	if err != nil {
		return
	}

	_, err = writer.Write([]byte("?"))
	if err != nil {
		return
	}

	if len(c.connections) > 0 {
		_, err = writer.Write(c.connections[:len(c.connections)-1])
		if err != nil {
			return
		}
	}

	_, err = writer.Write([]byte("??"))
	if err != nil {
		return
	}

	return
}

func NewCreation() Creation {
	return Creation{
		blocks:      make([]byte, 0),
		connections: make([]byte, 0),
		blockIndex:  0,
	}
}
