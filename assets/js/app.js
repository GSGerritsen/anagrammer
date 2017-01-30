new Vue({
	el: "#root",

	data: {
		wordInput: '',
		words: 	[],
	},

	methods: {
		getWords: function() {
			var self = this;
			axios.get('/search/' + this.wordInput).then(function(response)  {
				self.words = response.data
			})
		}
	},
	delimiters: ["<%", "%>"]
})
