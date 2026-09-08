const App3 = {
    data() {
        return {
            links: {}
        }
    },
    mounted: function () {
        this.getLinks()
    },
    methods: {
       async getLinks(){
        try{
            const res = await fetch("/api/link")
            if (!res.ok) throw new Error("Ошибка получения ссылки")
            const data = await res.json()
            this.links = data
        }catch(err){
            console.log(err)
        }
       }
    }
}

Vue.createApp(App3).mount('#VUE3')
