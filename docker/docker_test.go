package docker

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"testing"
)

func TestDocker(t *testing.T) {
	Convey("Docker Support", t, func() {
		Convey("Detecting (our own) container id", func() {
			data := "10:net_prio:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"9:perf_event:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"8:net_cls:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"7:freezer:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"6:devices:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"5:memory:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"4:blkio:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"3:cpuacct:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"2:cpu:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4\n" +
				"1:cpuset:/docker/ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4"

			cgroup := &bytes.Buffer{}
			io.WriteString(cgroup, data)

			client := DockerClient{
				cgroup: cgroup,
			}

			So(client.ContainerId(), ShouldEqual, "ce18c255cbab70caec36a81f948fba7cca856f90ebc0e2664f590c89b0fbeff4")
		})
	})
}
