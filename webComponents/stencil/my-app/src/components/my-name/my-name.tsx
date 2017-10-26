import { Component, Prop } from '@stencil/core';
// import jsonld from 'jsonld';


// Notes:
// Look at:
// https://stackoverflow.com/questions/41408471/parse-json-array-in-typescript
// to see about parsing JSON in typescript..  may not work for JSON-LD, but 
// maybe if it is flattened.  Also the JSON from Bleve call will likely 
// just be JSON, 

// declare module namespace {
  
//       export interface IcalDtstart {
//           @type: string;
//       }
  
//       export interface Context {
//           ical: string;
//           xsd: string;
//           ical:dtstart: IcalDtstart;
//       }
  
//       export interface RootObject {
//           @context: Context;
//           ical:summary: string;
//           ical:location: string;
//           ical:dtstart: string;
//       }
  
//   }
   

// class Person {
//   constructor(public name: string, 
//               public surname: string, 
//         public age: number){}
// }

@Component({
  tag: 'my-name',
  styleUrl: 'my-name.scss'
})
export class MyName {


  @Prop() first: string;
  @Prop() last: string;

  render() {

  

    // let mark = new Person('Mark', 'Galea', 30);
    // let mark = new namespace();
    
    // interface Test {
    //   icalSummary: string;
    // }

    // jsonld.compact(doc, context, function(err, compacted) {
      // console.log(JSON.stringify(compacted, null, 2));
      /* Output:
      {
        "@context": {...},
        "name": "Manu Sporny",
        "homepage": "http://manu.sporny.org/",
        "image": "http://manu.sporny.org/images/manu.png"
      }
      */
    // });

    var e = document.getElementById('test');
    var jsonld = e.textContent;
    // need to take JSONLD and try and parse it


    interface MyObj {
      id: string
      type: string
    }
    
    let obj: { string: MyObj[] } = JSON.parse(jsonld.toString());

    // console.log(obj.string[0].id, obj.string[0].type);
    console.log(obj["@graph"]);
    
    
    return (
      <div>
        Hello There, my name is <br/><br/>{this.first} {this.last}  <br/><br/>{jsonld}  <br/><br/>
      </div>
    );
  }
}
