import {
	html,
	render
} from './node_modules/lit-html/lit-html.js';


// lit-html constant
const greeting = (prefix, name) => {
	return html `
	    <h1>Well, hello there ${prefix} ${name}</h1>
		  `;
};

// lit-html constant
const greeting2 = (barval) => {
	console.log(barval)
	const itemTemplates = [];
	var i;
	for (i = 0; i < barval; i++) {
		itemTemplates.push(html `<li>${i}</li>`);
	}

	return html `
	  <ul>
		    ${itemTemplates}
	  </ul>
		`;
};

// lit-html constant
const providerTemplate = (barval) => {
	console.log("-----------------------------------------------")
	console.log(barval)
	var count = Object.keys(barval).length;
	const itemTemplates = [];
	var i;
	for (i = 0; i < count; i++) {
		itemTemplates.push(html `<li>Name:${barval[i].name}  Description: ${barval[i].description}</li>`);
	}

	return html `
	  <ul>
		    ${itemTemplates}
	  </ul>
		`;
};

// lit-html constant
const searchTemplate = (barval) => {
	console.log("-----------------------------------------------")
	console.log(barval)
	var count = Object.keys(barval).length;
	const itemTemplates = [];
	var i;
	for (i = 0; i < count; i++) {
		itemTemplates.push(html `<li>${barval[i].position} - 
		<a target="_blank" href="${barval[i].URL}">${barval[i].URL}</a> 
		(${barval[i].score})  </li>`);
	}

	return html `
	  <ul>
		    ${itemTemplates}
	  </ul>
		`;
};


// lit-html constant
const threadTemplate = (barval) => {
	console.log("threadtemplate-----------------------------------------------")
	// console.log(barval)
	var count = Object.keys(barval).length;
	const itemTemplates = [];


	var i;
	for (i = 0; i < count; i++) {
		const detailsTemplate = []
		let orset = barval[i].or
		let orlen = orset.length

		var j;
		for (j = 0; j < orlen; j++) {
			// console.log("checkpoint3");
            detailsTemplate.push(html `<div>${orset[j].URL}</div`)
		}
		itemTemplates.push(html `<div>${barval[i].index} -  ${barval[i].highscore} <br> ${detailsTemplate} </div>`);
	}

	return html `
	  <div>
		    ${itemTemplates}
      </div>
		`;
};

// core init code
let params = (new URL(location)).searchParams;
let q = params.get('q');
let n = params.get('n');
let s = params.get('s')
let i = params.get('i')

// Set the values of the query boxes based on URL parameters
let qdo = document.querySelector('#q');
let ndo = document.querySelector('#n');
let sdo = document.querySelector('#s');
let ido = document.querySelector('#i');
qdo.value = q
ndo.value = n
sdo.value = s
ido.value = i

// event listener
document.querySelector('#update').addEventListener('click', searchActions);
document.querySelector('#providers').addEventListener('click', providerList);


// --------  funcs and constants below here   ---------------------
function searchActions() {
	updateURL();
	threadSearch(); // simpleSearch();
	// renderResults();
}

// function renderResults() {
// 	const el = document.querySelector('#container1');
// 	let barval = document.querySelector('#q');
// 	render(greeting2(barval.value), el);
// }

function threadSearch() {
	let params = (new URL(location)).searchParams;
	let q = params.get('q');
	let n = params.get('n');
	let s = params.get('s')
	let i = params.get('i')

	fetch(`https://geodex.org/api/v1/textindex/searchset?q=${q}&n=${n}&s=${s}&i=${i}`)
		.then(function (response) {
			return response.json();
		})
		.then(function (myJson) {
			// console.log(myJson);
			const el = document.querySelector('#container2');
			render(threadTemplate(myJson), el);
		});
}


function simpleSearch() {
	let params = (new URL(location)).searchParams;
	let q = params.get('q');
	let n = params.get('n');
	let s = params.get('s')
	let i = params.get('i')

	fetch(`https://geodex.org/api/v1/textindex/search?q=${q}&n=${n}&s=${s}&i=${i}`)
		.then(function (response) {
			return response.json();
		})
		.then(function (myJson) {
			// console.log(myJson);
			const el = document.querySelector('#container2');
			render(searchTemplate(myJson), el);
		});
}

function providerList() {
	fetch('https://geodex.org/api/v1/typeahead/providers')
		.then(function (response) {
			return response.json();
		})
		.then(function (myJson) {
			// console.log(myJson);
			const el = document.querySelector('#container2');
			render(providerTemplate(myJson), el);
		});
}

function updateURL() {

	let qdo = document.querySelector('#q');
	let ndo = document.querySelector('#n');
	let sdo = document.querySelector('#s');
	let ido = document.querySelector('#i');

	let params = new URLSearchParams(location.search.slice(1));
	params.set('q', qdo.value);
	params.set('n', ndo.value);
	params.set('s', sdo.value);
	params.set('i', ido.value);

	window.history.replaceState({}, '', location.pathname + '?' + params);
}