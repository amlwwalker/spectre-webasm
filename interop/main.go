package main

import (
	"syscall/js"
	//"amlwwalker/go-app-tuts/pkg/spectre/layouts"
	//"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var icons = []string{
	"icon-cross",
	"icon-check",
	"icon-stop",
	"icon-shutdown",
	"icon-refresh",
	"icon-search",
	"icon-flag",
	"icon-bookmark",
	"icon-edit",
	"icon-delete",
	"icon-share",
	"icon-download",
	"icon-upload",
	"icon-copy",
	"icon-arrow-up",
	"icon-arrow-right",
	"icon-arrow-down",
	"icon-arrow-left",
	"icon-upward",
	"icon-forward",
	"icon-downward",
	"icon-back",
	"icon-caret",
	"icon-menu",
	"icon-apps",
	"icon-more-horiz",
	"icon-more-vert",
}
//// sideBar
//type sideBar struct {
//	app.Compo
//}
//
//func (h *sideBar) Render() app.UI {
//	sidebar := app.Window().GetElementByID("overlay_sidebar")
//	sidebarContent := layouts.FlexBox("100%",
//		app.Range(icons).Slice(func(i int) app.UI {
//			return app.Div().Class("column").Body(
//				app.I().Class("icon icon-2x").Class(icons[i]),
//			).Style("padding", "0.5rem")
//		}),
//	)
//	sidebar.Call("appendChild", sidebarContent)
//}
//func add(this js.Value, inputs []js.Value) interface{} {
//	return inputs[0].Float() + inputs[1].Float()
//}
func main() {
	////channel prevents program closing
	c := make(chan int)
	//js.Global().Set("go_add", js.FuncOf(add))
	//
	//fmt.Println("hello World")
	//num := js.Global().Call("add", 3, 4) //call a js function and pass any amount of values it expects
	//fmt.Println(num.Float())
	//s := js.Global().Call("hello").String() //call a js function that returns a string
	//fmt.Println(s)
	//env := js.Global().Get("env") //get a globally scoped (var) variable
	//fmt.Println(env)
	//js.Global().Set("env", "DEV") //set a globally scoped variable
	//js.Global().Get("config").Set("key", "1234") //set a nested value on an object
	//
	////add an element to the DOM
	//document := js.Global().Get("document")
	//h1 := document.Call("createElement", "h1")
	//h1.Set("innerText", "an h1 value")

	//
	//someDiv := document.Call("createElement", "div")
	//someDiv.Set("id", "someDiv")
	//someDiv.Set("innerText", "a div injected from web assembly")
	//document.Get("body").Call("appendChild", someDiv)
	//sidebar := app.Window().Get("body")
	//sidebarContent := layouts.FlexBox("100%",
	//	app.Range(icons).Slice(func(i int) app.UI {
	//		return app.Div().Class("column").Body(
	//			app.I().Class("icon icon-2x").Class(icons[i]),
	//		).Style("padding", "0.5rem")
	//	}),
	//).JSValue()
	//document.Get("body").Call("appendChild", sidebarContent)
	//sidebar.Call("appendChild", sidebarContent)
	//.Call("setProperty", "display", "none")
	//.Call("setProperty", "display", "block")
	//https://ian-says.com/articles/golang-in-the-browser-with-web-assembly/
	//https://dev.bitolog.com/go-in-the-browser-using-webassembly/
	<- c
}

/*

let l = document.createElement("div")
l.id = "mouse_overlay"
const style = (node, styles) => Object.keys(styles).forEach(key => node.style[key] = styles[key])
style(l, {
  "position": "fixed",
  "z-index": 999999999999999,
  "left": 0,
  "top": 0,
  "width": 0,
  "height": 0,
  "background": "rgba(0, 100, 255, 0.3)",
	"pointer-events": "none",
  "transition": "0.2s"
})
document.body.appendChild(l)
    function getXPath(el) {
      let nodeElem = el;
      if (nodeElem.id && this.options.shortid) {
        return `//*[@id="${nodeElem.id}"]`;
      }
      const parts = [];
      while (nodeElem && nodeElem.nodeType === Node.ELEMENT_NODE) {
        let nbOfPreviousSiblings = 0;
        let hasNextSiblings = false;
        let sibling = nodeElem.previousSibling;
        while (sibling) {
          if (sibling.nodeType !== Node.DOCUMENT_TYPE_NODE && sibling.nodeName === nodeElem.nodeName) {
            nbOfPreviousSiblings++;
          }
          sibling = sibling.previousSibling;
        }
        sibling = nodeElem.nextSibling;
        while (sibling) {
          if (sibling.nodeName === nodeElem.nodeName) {
            hasNextSiblings = true;
            break;
          }
          sibling = sibling.nextSibling;
        }
        const prefix = nodeElem.prefix ? nodeElem.prefix + ':' : '';
        const nth = nbOfPreviousSiblings || hasNextSiblings ? `[${nbOfPreviousSiblings + 1}]` : '';
        parts.push(prefix + nodeElem.localName + nth);
        nodeElem = nodeElem.parentNode;
      }
      return parts.length ? '/' + parts.reverse().join('/') : '';
    }
document.addEventListener("mouseover", e => {
	let elem = e.target;
	let rect = elem.getBoundingClientRect();
	let overlay = document.querySelector('#mouse_overlay');
	  overlay.style.top = rect.top +'px';
	  overlay.style.left = rect.left +'px';
	  overlay.style.width = rect.width +'px';
	  overlay.style.height = rect.height +'px';
    console.log(elem, rect)
	let XPath = getXPath(elem);
	const element = document.evaluate(XPath, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE,null).singleNodeValue;
	console.log("xpath ", XPath, element)
})
 */
