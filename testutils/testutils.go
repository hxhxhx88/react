package testutils

import (
	"fmt"

	"honnef.co/go/js/dom"

	"github.com/hxhxhx88/react/internal/core"
	"myitcv.io/react"

	"github.com/gopherjs/gopherjs/js"

	_ "github.com/hxhxhx88/react/internal/testutils"
)

var (
	testUtilsObj *js.Object
)

func init() {
	testUtilsObj = js.Global.Get("ReactTestUtils")

	if testUtilsObj == nil || testUtilsObj == js.Undefined {
		panic(fmt.Errorf("Could not load React TestUtils - ensure you are using a development build"))
	}
}

func RenderIntoDocument(elem react.Element) *core.ElementHolder {
	v := testUtilsObj.Call("renderIntoDocument", elem)

	return &core.ElementHolder{
		Elem: v,
	}
}

func FindRenderedDOMComponentWithClass(elem react.Element, class string) dom.HTMLElement {
	return dom.WrapHTMLElement(testUtilsObj.Call("findRenderedDOMComponentWithClass", elem, class))
}
