
greet({
    greeting: "hello world",
    duration: 4000
  });

  
function getGreeting() {
    return "howdy";
}
class MyGreeter extends Greeter { }

greet("hello");
greet(getGreeting);
greet(new MyGreeter());