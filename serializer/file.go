package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message")
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write proto message to file")
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read file")
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal proto message")
	}

	return nil
}

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message")
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write proto message to file")
	}

	return nil
}
