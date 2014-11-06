package consul

import (
	"fmt"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewConsulCollector(t *testing.T) {
	Convey("Create new NewConsulCollector", t, func() {
		consulCollector, err := NewConsulCollector("192.168.93.100:8500")
		fmt.Println(err)
		fmt.Println(consulCollector.DataCenter())
		fmt.Println(consulCollector.ServiceByName("arm"))
	})
}
