const App2 = {
    data() {
        return {
            user: null
        }
    },
    mounted: function () {
        this.GetUser()
    },
    methods: {
        async GetUser(){
            try{
                const res = await fetch("/api/user")
                if (!res.ok) throw new Error("Ошибка получения пользователя")
                const data = await res.json()
                this.user = data

            }catch(err){
                console.log(err)
            }
        }
        
    }
}

Vue.createApp(App2).mount('#VUE2')
