import { Component, Prop } from '@stencil/core';
import d3 from 'd3';
import jsonldVis from 'jsonld-vis';
jsonldVis(d3);

// Notes:
// Look at:
// https://stackoverflow.com/questions/41408471/parse-json-array-in-typescript
// to see about parsing JSON in typescript..  may not work for JSON-LD, but 
// maybe if it is flattened.  Also the JSON from Bleve call will likely 
// just be JSON, 


@Component({
  tag: 'my-name',
  styleUrl: 'my-name.scss'
})
export class MyName {

  @Prop() first: string;
  @Prop() last: string;



  render() {

  var e = document.getElementById('test');
  var jsonld = e.textContent;
  // need to take JSONLD and try and parse it

  interface MyObj {
    id: string
    type: string
  }

  let obj: { string: MyObj[] } = JSON.parse(jsonld.toString());


  var element = obj["@graph"]
  console.log(element[0]["schema:contentUrl"])
  var url = element[0]["schema:contentUrl"]


    return (
      <div class="earthcube-provinfo">
        Viz test: {this.first} {this.last}
        <div>
        <a href={url} target="_blank">{url}</a>
        </div>
      </div>
    );
  }
}
