package serializer

import (
	proto_gen "github.com/negarciacamilo/go_grpc/proto-gen"
	"github.com/negarciacamilo/go_grpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	writeLaptop := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(writeLaptop, binaryFile)
	require.NoError(t, err)

	readLaptop := &proto_gen.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, readLaptop)
	require.NoError(t, err)
	require.True(t, proto.Equal(writeLaptop, readLaptop))

	err = WriteProtobufToJSONFile(writeLaptop, jsonFile)
	require.NoError(t, err)
}
