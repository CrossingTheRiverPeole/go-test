package package_test

import (
	"go_learning/src/ch15/client_package"
	"testing"
)

func TestPackage(t *testing.T)  {
	t.Log(client_package.Add(1,2))
}
