package testWithDocker

import (
	"goTestProject/dockerDemo/testEnvironment"
	"os"
	"testing"
)

var mongoURI string

// TestMain test  开始执行, 结束终止
func TestMain(m *testing.M) {
	os.Exit(testEnvironment.RunWithMongoInDocker(m, &mongoURI))
}
