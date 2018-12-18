		Slim.tag(
			'geocomponents-citation',
			`<div style="font-size:10pt">
			<span bind>{{Authors}}.</span>
			<span bind>{{Year}}.</span>
			<span bind>{{Dataset}}</span>
			<span style="font-style: italic;" bind>{{Title}}.</span>
			<span bind>{{Archive}}</span>
			<span bind>{{Version}}</span>
			<span bind>Retrieved from {{PID}}</span>
			</div>`,
			class CitationTag extends Slim {
				// your code here
				onBeforeCreated() {
					var element = document.getElementById('schemaorg');
					var jsonld = element.innerHTML;
					var obj = JSON.parse(jsonld);

					this.Authors = ""
					this.Year = ""
					this.Dataset = obj.name
					this.Title = obj.description
					this.Archive = ""
					this.Version = ""
					this.PID = obj.url
				}
			})
