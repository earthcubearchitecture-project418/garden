import {
	html,
	render
} from '../node_modules/lit-html/lit-html.js';


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
		itemTemplates.push(html `<div style="margin-top:30px"> 
		<img style="height:50px" src="${barval[i].logo}"><br>  ${barval[i].description}   (${barval[i].name} )  </div>`);
	}

	return html `
	  <div style="margin-top:30px">
		    ${itemTemplates}
      </div>
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


// lit-html constant
const nusearch = (barval, q) => {
	console.log("nusearchtemplate-----------------------------------------------")
	var count = Object.keys(barval.search_result.hits).length;
	const itemTemplates = [];

	function nametest(t) {
		if (t == "undefined") {
			return 'Un-named facility data set';
		}
		return t;
	}

	function desctest(t) {
		if (t == "undefined") {
			return 'No description for this data set is provided by the facility';
		}
		return t;
	}

	// obviously pointless function..   do this test and a template push below....
	function curltest(t) {
		if (t == "undefined") {
			return 'undefined';
		}
		return t;
	}

	function UpdateQueryString(key, value, url) {
		if (!url) url = window.location.href;
		var re = new RegExp("([?&])" + key + "=.*?(&|#|$)(.*)", "gi"),
			hash;
	
		if (re.test(url)) {
			if (typeof value !== 'undefined' && value !== null)
				return url.replace(re, '$1' + key + "=" + value + '$2$3');
			else {
				hash = url.split('#');
				url = hash[0].replace(re, '$1$3').replace(/(&|\?)$/, '');
				if (typeof hash[1] !== 'undefined' && hash[1] !== null) 
					url += '#' + hash[1];
				return url;
			}
		}
		else {
			if (typeof value !== 'undefined' && value !== null) {
				var separator = url.indexOf('?') !== -1 ? '&' : '?';
				hash = url.split('#');
				url = hash[0] + separator + key + '=' + value;
				if (typeof hash[1] !== 'undefined' && hash[1] !== null) 
					url += '#' + hash[1];
				return url;
			}
			else
				return url;
		}
	}

	var i;
	for (i = 0; i < count; i++) {
		var desc = `${barval.search_result.hits[i].fields.description}`;
		var shortdesc = desc.slice(0, 500);
		shortdesc = desctest(shortdesc)

		var name = `${barval.search_result.hits[i].fields.name}`;
		name = nametest(name)

		// need to check for distribution..  then contentUrl
		// .search_result.hits["0"].fields["distribution.contentUrl"]
		var curl = null
		if (barval.search_result.hits[i].fields["distribution.contentUrl"] != null) {
			var curl = `${barval.search_result.hits[i].fields["distribution.contentUrl"]}`; //
			curl = curltest(curl)
		}

		// Set up the datadownload template
		var dataDownloadTemplates = ""
		if (curl != null) {
			dataDownloadTemplates = (html `<a target="_blank" href="${curl}"><img style="margin-left:40px;height:20px" src="./download.svg"></a>`)
		} else {
			dataDownloadTemplates = (html `<span> </span>`)
		}

		// Set up the filter on source section
		var filterTemplates = ""
		var newq = `${q}  p418source:${barval.search_result.hits[i].fields.p418source}`
		var urlrewrite = UpdateQueryString("q", newq, null)
		filterTemplates = (html `<a href="${urlrewrite}">
		<img style="margin-left:20px;height:20px" src="./filter.svg"></a>`)


		// Main Item template
		itemTemplates.push(html `<div style="margin-top:15px">
		<a target="_blank" href="${barval.search_result.hits[i].fields.p418url}">${name}</a>
		  <br/>
			<img style="height:20px" src="${barval.search_result.hits[i].fields.p418logo}">
			
			
			${filterTemplates}
		 
			${dataDownloadTemplates}

			
			<br/>
	<span> ${shortdesc}... </span>
	<br/>
		<span style="font-size: smaller;" >(${barval.search_result.hits[i].score}) <span> </div>`);
	}

	return html `
	  <div>
		   ${itemTemplates}
      </div>
		`;
};

const query1 = (q, n) => {

	return `
	{
		"search_request": {
		  "query": {
			"query": "${q}"
		  },
		  "size": ${n},
		  "from": 0,
		  "fields": [
			"*"
		  ],
		  "sort": [
			"-_score"
		  ],
		  "highlight": {
			"style": "html",
			"fields": [
			  "name",
			  "description"
			]
		  }
		}
	  }
	  `

}


// popstate for history button
window.onpopstate = event => {
	console.log("opnpopstate seen")
	console.log(event.state)
  }


// core init code
let params = (new URL(location)).searchParams;
let q = params.get('q');
let n = params.get('n');
let s = params.get('s');
let i = params.get('i');

// trap n = null to prime the number return do
if (n == null) {
	n = 20
}

// Set the values of the query boxes based on URL parameters
let qdo = document.querySelector('#q');
let ndo = document.querySelector('#nn');
let sdo = document.querySelector('#s');
let ido = document.querySelector('#i');
qdo.value = q;
ndo.value = n;
sdo.value = s;
ido.value = i;

// if q is not null..   fire off a search, 
if (q != null) {
		searchActions();
}




// event listeners
document.querySelector('#q').addEventListener('keyup', function (e) {
	if (e.keyCode === 13) {
		searchActions();
	}
});

document.querySelector('#update').addEventListener('click', searchActions);
document.querySelector('#providers').addEventListener('click', providerList);

// --------  funcs and constants below here   ---------------------
function searchActions() {
	// let params = (new URL(location)).searchParams;
	let q = document.querySelector('#q').value
	let n = document.querySelector('#nn').value
	// let s = params.get('s');
	// let i = params.get('i');

	updateURL();

	// Different search options
	blastsearchsimple(q, n);
	// threadSearch(q, n, s, i); 
	// simpleSearch();
}

function blastsearchsimple(q, n) {
	// var formData = new FormData();
	var data = query1(q, n);
	console.log(data)

	//fetch(`http://localhost:6789/api/v1/textindex/getnusearch?q=${data}`)
	fetch(`http://192.168.2.50:6789/api/v1/textindex/getnusearch?q=${data}`)
		.then(function (response) {
			return response.json();
		})
		.then(function (myJson) {
			console.log(myJson);
			const el = document.querySelector('#container2');
			render(nusearch(myJson, q), el);
		});
}

function threadSearch(q, n, s, i) {
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

function simpleSearch(q, n, s, i) {
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
	fetch('http://192.168.2.50:6789/api/v1/typeahead/providers')
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
	let ndo = document.querySelector('#nn');
	let sdo = document.querySelector('#s');
	let ido = document.querySelector('#i');

	let params = new URLSearchParams(location.search.slice(1));
	params.set('q', qdo.value);
	params.set('n', ndo.value);
	params.set('s', sdo.value);
	params.set('i', ido.value);

	//window.history.replaceState({}, '', location.pathname + '?' + params);
	const state = { geodexsearch: q }
	window.history.pushState({}, '', location.pathname + '?' + params);
}


