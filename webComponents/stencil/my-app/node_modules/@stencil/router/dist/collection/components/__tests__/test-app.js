var TestApp = /** @class */ (function () {
    function TestApp() {
    }
    TestApp.prototype.render = function () {
        return (h("stencil-router", { "a": { "title-suffix": " - Stencil Router Demos" } },
            h("ul", 0,
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/" }, "p": { "exact": true } }, t("Exact Base Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/" } }, t("Base Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo" }, "p": { "urlMatch": ['/demo', '/demox'], "exact": true } }, t("Demo Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo2" } }, t("Demo2 Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo3" } }, t("Demo3 Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo3/page1" } }, t("Demo3 Page1 Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo3/page2" } }, t("Demo3 Page2 Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo4" } }, t("Demo4 Link"))),
                h("li", 0,
                    h("stencil-route-link", { "a": { "url": "/demo6/" } }, t("Demo6 Link")))),
            h("stencil-route", { "a": { "url": "/" }, "p": { "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return h("span", 0, t("rendering /"));
                    } } }),
            h("stencil-route", { "p": { "url": ['/demo', '/demox'], "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return [
                            h("stencil-route-title", { "a": { "title": "DEMO" } }),
                            h("span", 0, t("rendering /demo"))
                        ];
                    } } }),
            h("stencil-route", { "a": { "url": "/demo2" }, "p": { "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return [
                            h("span", 0, t("rendering /demo2")),
                            h("stencil-router-redirect", { "a": { "url": "/demo3" } })
                        ];
                    } } }),
            h("stencil-route", { "a": { "url": "/demo3" }, "p": { "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return [
                            h("stencil-route-title", { "a": { "title": "Demo 3" } }),
                            h("span", 0, t("rendering /demo 3"))
                        ];
                    } } }),
            h("stencil-route", { "a": { "url": "/demo3", "component": "test-demo-three" }, "p": { "componentProps": { "pages": ['intro/index.html'] } } }),
            h("stencil-route", { "a": { "url": "/demo4", "component": "test-demo-four" } }),
            h("stencil-route", { "a": { "url": "/demo5", "component": "async-content" }, "p": { "componentProps": { "location": '/' } } }),
            h("stencil-route", { "a": { "url": "/demo6", "component": "test-demo-six" } })));
    };
    return TestApp;
}());
export { TestApp };
