package page

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPager(t *testing.T) {
	Convey("Business layer pagination", t, func() {
		Convey("Invalid parameters and use default value", func() {
			totalRecords, _ := countUsers()
			pager := NewPager[*User](-1, -1, totalRecords)
			limiter := pager.BuildLimiter()
			users, _ := getUsersFromDB(limiter.Offset, limiter.Limit)
			for i := range users {
				pager.AddRecords(&users[i])
			}
			page := pager.BuildPage()
			So(page, ShouldNotBeNil)
			So(page.PageNo, ShouldEqual, 1)
			So(page.PageSize, ShouldEqual, 10)
			So(page.TotalPages, ShouldEqual, 1)
			So(page.TotalRecords, ShouldEqual, totalRecords)
			So(len(page.Records), ShouldEqual, 10)

		})
		Convey("Valid parameters", func() {
			Convey("Page number less than or equal total pages", func() {
				totalRecords, _ := countUsers()        // 10
				pageNo, pageSize := int64(2), int64(3) // 2 <= 4

				pager := NewPager[User](pageNo, pageSize, totalRecords)
				limiter := pager.BuildLimiter()
				users, _ := getUsersFromDB(limiter.Offset, limiter.Limit)
				pager.AddRecords(users...)

				page := pager.BuildPage()
				So(page, ShouldNotBeNil)
				So(page.PageNo, ShouldEqual, pageNo)
				So(page.PageSize, ShouldEqual, pageSize)
				So(page.TotalPages, ShouldEqual, 4)
				So(page.TotalRecords, ShouldEqual, totalRecords)
				So(len(page.Records), ShouldEqual, pageSize)
				So(page.Records[0].Name, ShouldEqual, "user3")
				So(page.Records[1].Name, ShouldEqual, "user4")
				So(page.Records[2].Name, ShouldEqual, "user5")
			})

			Convey("Page number greater than total pages", func() {
				totalRecords, _ := countUsers()        // 10
				pageNo, pageSize := int64(5), int64(3) // 5 > 4

				pager := NewPager[User](pageNo, pageSize, totalRecords)
				limiter := pager.BuildLimiter()
				users, _ := getUsersFromDB(limiter.Offset, limiter.Limit)
				pager.AddRecords(users...)

				page := pager.BuildPage()
				So(page, ShouldNotBeNil)
				So(page.PageNo, ShouldEqual, pageNo)
				So(page.PageSize, ShouldEqual, pageSize)
				So(page.TotalPages, ShouldEqual, 4)
				So(page.TotalRecords, ShouldEqual, totalRecords)
				So(page.Records, ShouldBeEmpty)
			})
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
	if offset < 0 || (offset+limit) > int64(len(all)) {
		return []User{}, nil
	}
	return all[offset : offset+limit], nil
}

func TestEmptyPage(t *testing.T) {
	Convey("Empty page", t, func() {
		page := EmptyPage[string](1, 10)
		So(page, ShouldNotBeNil)
		So(page.PageNo, ShouldEqual, 1)
		So(page.PageSize, ShouldEqual, 10)
		So(page.TotalPages, ShouldEqual, 0)
		So(page.TotalRecords, ShouldEqual, 0)
		So(page.Records, ShouldBeEmpty)
	})
}
