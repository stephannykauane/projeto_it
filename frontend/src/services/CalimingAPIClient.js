class CalimingAPIClient {
    baseUrl = 'http://localhost:8080'
    user = {}

    async sendRequest ({ path, method, body }) {
        const token = localStorage.getItem('token')
        console.log("token", token)
        try {
            const headers = {
                'Content-Type': 'application/json',
            }
    
            if (this.token) {
                headers['Authorization'] = `Bearer ${this.token}`
            }

            const resp = await fetch(`${this.baseUrl}${path}`, {
                method,
                headers,
                body: method === 'GET' ? undefined : JSON.stringify(body),
                credentials: 'include',
            });
            return resp;
        } catch (err) {
            console.error('Erro na requisição: ', err)
            throw err; 
        }
    }

    async login({ email, senha }) {
        const resp = await this.sendRequest({
            method: 'POST', 
            path: '/login',
            body: { email, senha }
        })
        
        if (resp.ok) {
            const data = await resp.json()
            localStorage.setItem('token', data.token)
            this.token = data.token
            this.user.email = email

            return true 
        } else { 
            const erro = await resp.text()
            console.error("Erro ao logar: ", erro)

            return false
        }
    }

    async logout() {
        await this.sendRequest({
            method: 'POST',
            path: '/logout',
        });

        this.user = {}
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

    async getDados () {
        const resp = await this.sendRequest ({
            method: 'GET',
            path: '/profile'
        })

        if (resp.ok) {
            return await resp.json()
        } else {
            const erro = await resp.text()
            console.error("Erro ao buscar dados:", erro)
            throw new Error("Erro ao buscar dados")
        }
    }

}

const calimingAPI = new CalimingAPIClient()
export default calimingAPI