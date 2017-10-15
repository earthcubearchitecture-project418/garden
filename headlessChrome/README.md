### About
At some point in P418 it would be interesting to use headless chrome to make screen shots 
of the pages and use in the display.   The hidden reason for this though is to look at headless
chrome as the page render in order to potentially get at data in .js or web components which dont
render on read for simple library driven GET.

* 
* https://chromeless.netlify.com/#src=const%20chromeless%20=%20new%20Chromeless(%7B%20remote:%20true%20%7D)%0A%0Aconst%20screenshot%20=%20await%20chromeless%0A%20%20.goto('http://opencoredata.org')%0A%20%20.scrollTo(0,%202000)%0A%20%20.screenshot()%0A%0Aconsole.log(screenshot)%0A%0Aawait%20chromeless.end()

https://github.com/knq/chromedp

https://github.com/knq/chromedp/tree/master/examples/headless
```
# retrieve docker image
docker pull knqz/chrome-headless

# start chrome-headless
docker run -d -p 9222:9222 --rm --name chrome-headless knqz/chrome-headless

# run chromedp headless example
go build && ./headless
```