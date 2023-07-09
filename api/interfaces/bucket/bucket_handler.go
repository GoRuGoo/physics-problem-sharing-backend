package bucket

import "os"

type BucketHandler interface {
	WriteExecute(string, os.File) error
}
