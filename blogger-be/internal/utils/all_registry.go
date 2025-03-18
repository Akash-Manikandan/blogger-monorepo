package utils

import (
	"log"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func LogAllMethods() {
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		for i := range fd.Services().Len() {
			service := fd.Services().Get(i)
			for j := range service.Methods().Len() {
				method := service.Methods().Get(j)
				log.Printf("Registered gRPC Method: %s/%s", service.FullName(), method.Name())
			}
		}
		return true
	})
}
