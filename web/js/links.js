const App3 = {
    data() {
        return {
            links: {},
            isActive: 1
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
       },

       formatDate(date) {
            return new Date(date).toLocaleString("ru-RU", {
                day: "2-digit",
                month: "2-digit",
                year: "numeric",
                hour: "2-digit",
                minute: "2-digit"
            }).replace(",", "")
        },
       async deleteLinks(LinkID){
        try{
            const res = await fetch(`/api/link?LinkID=${LinkID}`,{
                method: 'DELETE',
                credentials: 'same-origin',
                headers: { 'Content-Type': 'application/json' },
            })
            if (!res.ok) throw new Error("Ошибка удаления ссылки")
        }catch(err){
            console.log(err)
        }
        this.getLinks()
       },
       async updateIsActive(LinkID, Active){
        try{
            const res = await fetch(`/api/link?LinkID=${LinkID}&active=${Active}`,{
                method: 'PUT',
                credentials: 'same-origin',
                headers: { 'Content-Type': 'application/json' },
            })
        }catch(err){
            console.log(err)
        }
        await this.getLinks()
       }
    }
}

Vue.createApp(App3).mount('#VUE')
