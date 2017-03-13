package page

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCalcTotalPages(t *testing.T) {
	var totalPages int64
	var err error
	Convey("Calculate total pages", t, func() {
		Convey("Invalid parameters", func() {
			totalPages, err = calcTotalPages(-1, 1)
			So(err, ShouldNotBeNil)
			So(totalPages, ShouldEqual, 0)

			totalPages, err = calcTotalPages(0, 1)
			So(err, ShouldNotBeNil)
			So(totalPages, ShouldEqual, 0)

			totalPages, err = calcTotalPages(10, -1)
			So(err, ShouldNotBeNil)
			So(totalPages, ShouldEqual, 0)
		})

		Convey("Valid parameters", func() {
			totalPages, err = calcTotalPages(3, 10)
			So(err, ShouldBeNil)
			So(totalPages, ShouldEqual, 4)

			totalPages, err = calcTotalPages(2, 10)
			So(err, ShouldBeNil)
			So(totalPages, ShouldEqual, 5)

			totalPages, err = calcTotalPages(2, 0)
			So(err, ShouldBeNil)
			So(totalPages, ShouldEqual, 0)
		})
	})
}
