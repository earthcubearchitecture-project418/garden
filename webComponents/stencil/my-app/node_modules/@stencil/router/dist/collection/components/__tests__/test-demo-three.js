var TestDemoThree = /** @class */ (function () {
    function TestDemoThree() {
    }
    TestDemoThree.prototype.render = function () {
        var _this = this;
        console.log('pages: ', this.pages);
        console.log('match: ', this.match);
        console.log('history: ', this.history.location);
        return [
            h("span", 0, t("Demo 3 Test Page"),
                h("br", 0)),
            h("stencil-route", { "a": { "url": "/demo3/page1" }, "p": { "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return [
                            h("a", { "o": { "click": function (e) {
                                        e.preventDefault();
                                        _this.history.push('/demo3/page2', { 'blue': 'blue' });
                                    } }, "a": { "href": "#" } }, t("History push to /demo3/page2")),
                            h("pre", 0,
                                h("b", 0, t("props.pages")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.pages, null, 2)),
                            h("pre", 0,
                                h("b", 0, t("props.match")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.match, null, 2)),
                            h("pre", 0,
                                h("b", 0, t("props.history.location")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.history.location, null, 2))
                        ];
                    } } }),
            h("stencil-route", { "a": { "url": "/demo3/page2" }, "p": { "exact": true, "routeRender": function (props) {
                        console.log(props);
                        return [
                            h("a", { "o": { "click": function (e) {
                                        e.preventDefault();
                                        _this.history.push('/demo3/page1', { 'red': 'red' });
                                    } }, "a": { "href": "#" } }, t("History push to /demo3/page1")),
                            h("pre", 0,
                                h("b", 0, t("props.pages")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.pages, null, 2)),
                            h("pre", 0,
                                h("b", 0, t("props.match")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.match, null, 2)),
                            h("pre", 0,
                                h("b", 0, t("props.history.location")), t(":"),
                                h("br", 0),
                                JSON.stringify(_this.history.location, null, 2))
                        ];
                    } } })
        ];
    };
    return TestDemoThree;
}());
export { TestDemoThree };
