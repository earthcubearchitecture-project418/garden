var TestDemoFour = /** @class */ (function () {
    function TestDemoFour() {
    }
    TestDemoFour.prototype.handleClick = function (e, linkLocation) {
        e.preventDefault();
        this.history.push(linkLocation, { 'blue': 'blue' });
    };
    TestDemoFour.prototype.render = function () {
        var _this = this;
        console.log('pages: ', this.pages);
        console.log('match: ', this.match);
        console.log('history: ', this.history.location);
        var linkLocation = '/demo3/page1';
        return (h("div", 0,
            h("a", { "o": { "click": function (e) { return _this.handleClick(e, linkLocation); } }, "p": { "href": linkLocation } }, t("History push to "),
                linkLocation),
            h("pre", 0,
                h("b", 0, t("this.pages")), t(":"),
                h("br", 0),
                JSON.stringify(this.pages, null, 2)),
            h("pre", 0,
                h("b", 0, t("this.match")), t(":"),
                h("br", 0),
                JSON.stringify(this.match, null, 2)),
            h("pre", 0,
                h("b", 0, t("this.history.location")), t(":"),
                h("br", 0),
                JSON.stringify(this.history.location, null, 2))));
    };
    return TestDemoFour;
}());
export { TestDemoFour };
