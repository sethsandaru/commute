## Defining Crawler-Rules we should follow

```javascript
crawlerRules = 
{
	timeOut: 3000,
	injectJQuery: Boolean,
	
	// entity is necessary to extract the data
	entity: {
		entityName: {
			key: "int",
			key2: "string",
			key3: "string",
			// type exists: int, string, float
		},
		// more entity here
	},
	
	actions: [
		{
			type: 1, 
				//- 1: wait
				//- 2: extract-single
				//- 3: extract-multiple
				//- 4: custom-js
			value: 1,
				//(for type 1)
				//	- 1: wait for selector
				//	- 2: wait for timer
				//	- 3: wait for xpath
				//(type 2)
				//	- 1: by selector
				//  - 2: by xPath

			extractionRules: {}, // for entity
			containerSelector: ".", // for bulk-extraction (container > children)
		},
	]
}
```