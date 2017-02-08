// Header content dealing with login/signup modals + accurately representing nav state based on user authentication
Vue.component('modal', {
    template: '#modal-template'
})

const config = { headers: { 'Content-Type' : 'application/x-www-form-urlencoded'} };

new Vue({
    el: '#navbar-target',
    data: {
        showLoginModal: false,
        showSignupModal: false,
        showPasswordError: false,
        userAuthenticated: false,
        loginEmail: '',
        loginPass: '',
        signupEmail: '',
        signupPass: '',
        signupConfPass: ''
    },

    // Check for a jwt_token cookie, and if one is found, hit /auth with a get to validate the token
    created: function() {
        var self = this;
        var token = document.cookie.match('(^|;)\\s*' + "jwt_token" + '\\s*=\\s*([^;]+)');
        if (token) {
            axios.get('/auth').then(function(response) {
                if(response.status == 200) {
                    self.userAuthenticated = true;
                }
            })
         
        }
    },

    methods: {
        getLogin: function() {
            var self = this;
            axios.post('/login', {
                email: this.loginEmail,
                password: this.loginPass
            }, config).then(function(response) {
                self.userAuthenticated = true;
            })
            this.showLoginModal = false;
        },

        getSignup: function() {
            var self = this;
            if (this.signupPass !== this.signupConfPass) {
                this.showPasswordError = true;
                return;
            }
            this.showPasswordError = false;
            axios.post('/signup', {
                email: this.signupEmail,
                password: this.signupPass,
                confPassword: this.signupConfPass,
            }, config).then(function(response) {
                console.log("POST to /signup");
            })
            this.showSignupModal = false;
        },
        logout: function() {
            var self = this;
            axios.get('/logout').then(function(response) {
               if(response.status == 200) {
                    self.userAuthenticated = false;
               }
            })
        }
        
    },
    delimiters: ["<%", "%>"]
})

// Body content dealing with anagram input handling + presentation

new Vue({
    el: "#root",

    data: {
        wordInput: '',
        words: [],
    },

    methods: {
        getWords: function() {
            var self = this;
            axios.get('/search/' + this.wordInput.toLowerCase()).then(function(response) {
                self.words = response.data;
            })
        }
    },
    delimiters: ["<%", "%>"]
})
