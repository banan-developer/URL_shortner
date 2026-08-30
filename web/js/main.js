const App = {
    data() {
        return {
            count: 1,
            link: "",
            shortLink: ""
        }
    },
    mounted(){
        
    },
    methods: {
        // async GetLinksByID(){
        //     try{
        //         const res = await fetch("/api/link")
        //          if (!res.ok) throw new Error("Ошибка получения сообщений")
        //         const data = await res.json()
        //         this.link = data
        //     }catch(err){
        //         console.log(err)
        //     }
        // },
        async CreateLink(){
            try{
                const res = await fetch("/api/link", {
                    method: 'POST',
                    credentials: 'same-origin',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({
                    original_url: this.link
                })
                })
                const data = await res.json()
                this.shortLink = data

            }catch(err){
                console.log(err)
            }
        }
    }
}

Vue.createApp(App).mount('#VUE')
