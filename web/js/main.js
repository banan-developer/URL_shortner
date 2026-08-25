const App = {
    data() {
        return {
            count: 0
        }
    },
    methods: {
        click_button(){
            this.count++
        }
    }
}

Vue.createApp(App).mount('#VUE')
