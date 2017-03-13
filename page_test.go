package page

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPager(t *testing.T) {
	Convey("Business layer pagination", t, func() {
		Convey("Invalid parameters and use default value", func() {
			totalRecords, _ := countUsers()
			pager := NewPager(reflect.TypeOf(&User{}), -1, -1, totalRecords)
			limiter := pager.BuildLimiter()
			users, _ := getUsersFromDB(limiter.Offset, limiter.Limit)
			for i := range users {
				pager.AddRecords(&users[i])
			}
			page := pager.BuildPage()
			So(page, ShouldNotBeNil)
			So(page.Page, ShouldEqual, 1)
			So(page.PageSize, ShouldEqual, 10)
			So(page.TotalPages, ShouldEqual, 1)
			So(page.TotalRecords, ShouldEqual, totalRecords)
			So(len(page.Records), ShouldEqual, 10)

		})
		Convey("Valid parameters", func() {
			totalRecords, _ := countUsers()
			pager := NewPager(reflect.TypeOf(User{}), 2, 3, totalRecords)
			limiter := pager.BuildLimiter()
			users, _ := getUsersFromDB(limiter.Offset, limiter.Limit)
			for i := range users {
				pager.AddRecords(users[i])
			}
			So(pager.AddRecords(&User{}), ShouldNotBeNil)

			page := pager.BuildPage()
			So(page, ShouldNotBeNil)
			So(page.Page, ShouldEqual, 2)
			So(page.PageSize, ShouldEqual, 3)
			So(page.TotalPages, ShouldEqual, 4)
			So(page.TotalRecords, ShouldEqual, totalRecords)
			So(len(page.Records), ShouldEqual, 3)
			So(page.Records[0].(User).Name, ShouldEqual, "user3")
			So(page.Records[1].(User).Name, ShouldEqual, "user4")
			So(page.Records[2].(User).Name, ShouldEqual, "user5")
		})
	})
}

type User struct {
	Name string
}

func countUsers() (totalRecords int64, err error) {
	return 10, nil
}

func getUsersFromDB(offset, limit int64) (users []User, err error) {
	all := []User{
		{Name: "user0"},
		{Name: "user1"},
		{Name: "user2"},
		{Name: "user3"},
		{Name: "user4"},
		{Name: "user5"},
		{Name: "user6"},
		{Name: "user7"},
		{Name: "user8"},
		{Name: "user9"},
	}

	return all[offset : offset+limit], nil
}

func TestEmptyPage(t *testing.T) {
	Convey("Empty page", t, func() {
		page := EmptyPage(1, 10)
		So(page, ShouldNotBeNil)
		So(page.Page, ShouldEqual, 1)
		So(page.PageSize, ShouldEqual, 10)
		So(page.TotalPages, ShouldEqual, 0)
		So(page.TotalRecords, ShouldEqual, 0)
		So(page.Records, ShouldBeEmpty)
	})
}
