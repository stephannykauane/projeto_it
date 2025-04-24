class CalimingAPIClient {
    baseUrl = 'http://localhost:8080'
    token = null
    user = {}

    async sendRequest ({ path, method, body }) {
        return await fetch(`${this.baseUrl}${path}`, {
            method,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': this.token,
            },
            body: JSON.stringify(body),
            credentials:  'include',
        })
    }

    async login({ email, senha }) {
        const resp = await this.sendRequest({
            method: 'POST', 
            path: '/login',
            body: { email, senha }
        })
        
         if (resp.ok) {
            this.user.email = email 
            return true 
         } else { 
            const erro = await resp.text()
            console.error("Erro ao logar: ", erro)
            return false
         }
    }

    async signUp({ email, senha, nome }) {
        const resp = await this.sendRequest({
            method: 'POST',
            path: '/signup',
            body: { email, senha, nome }
        })
    
        if (resp.ok) {
            return await resp.json()
        } else {
            const erro = await resp.text()
            console.error("Erro ao cadastrar:", erro)
            throw new Error("Erro ao cadastrar")
        }
    }
}

const calimingAPI = new CalimingAPIClient()
export default calimingAPI