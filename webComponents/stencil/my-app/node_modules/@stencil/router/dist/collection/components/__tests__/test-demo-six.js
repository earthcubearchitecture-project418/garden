var TestDemoSix = /** @class */ (function () {
    function TestDemoSix() {
    }
    TestDemoSix.prototype.render = function () {
        console.log('pages: ', this.pages);
        console.log('match: ', this.match);
        console.log('history: ', this.history.location);
        return [
            h("span", 0, t("Demo 6 Test Page"),
                h("br", 0)),
            h("stencil-route", { "a": { "url": "/demo6/", "group": "main" }, "p": { "exact": true, "routeRender": function (props) {
                        return [
                            h("h1", 0, t("One")),
                            h("stencil-route-link", { "a": { "url": "/demo6/asdf" } }, t("Next"))
                        ];
                    } } }),
            h("stencil-route", { "a": { "url": "/demo6/:any*", "group": "main" }, "p": { "routeRender": function (props) {
                        console.log('Got match props:', props.match);
                        return [
                            h("h1", 0, t("Two: "),
                                props.match),
                            h("stencil-route-link", { "a": { "url": "/demo6/" } }, t("Back"))
                        ];
                    } } })
        ];
    };
    return TestDemoSix;
}());
export { TestDemoSix };
