package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetSquare(t *testing.T) {
	Convey("Given integers", t, func() {
		Convey("It should return the square", func() {
			So(getSquare(10).Root, ShouldEqual, 10)
			So(getSquare(10).Square, ShouldEqual, 100)
		})
	})
}
