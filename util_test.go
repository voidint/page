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

func TestCalcOffset(t *testing.T) {
	var offset int64
	var err error
	Convey("Calculate offset for database", t, func() {
		Convey("Invalid parameters", func() {
			offset, err = calcOffset(0, 10, 10)
			So(err, ShouldNotBeNil)

			offset, err = calcOffset(-1, 10, 10)
			So(err, ShouldNotBeNil)

			offset, err = calcOffset(1, 0, 10)
			So(err, ShouldNotBeNil)

			offset, err = calcOffset(1, -1, 10)
			So(err, ShouldNotBeNil)

			offset, err = calcOffset(1, 10, -1)
			So(err, ShouldNotBeNil)
		})

		Convey("Valid parameters", func() {
			offset, err = calcOffset(1, 10, 0)
			So(err, ShouldBeNil)
			So(offset, ShouldEqual, 0) // begin with 0

			offset, err = calcOffset(1, 10, 10)
			So(err, ShouldBeNil)
			So(offset, ShouldEqual, 0)

			offset, err = calcOffset(2, 3, 10)
			So(err, ShouldBeNil)
			So(offset, ShouldEqual, 3)

			offset, err = calcOffset(2, 10, 5)
			So(err, ShouldBeNil)
			So(offset, ShouldEqual, 0)

			offset, err = calcOffset(2, 3, 10)
			So(err, ShouldBeNil)
			So(offset, ShouldEqual, 3)
		})
	})
}
