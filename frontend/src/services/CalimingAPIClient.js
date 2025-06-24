class CalimingAPIClient {

    constructor() {
        this.baseUrl = 'http://localhost:8080'
        this.user = {}
        this.token = localStorage.getItem('token') || null
    }

    async sendRequest({ path, method, body }) {

        const token = this.token || localStorage.getItem('token')

        const headers = {
            'Content-Type': 'application/json',
        }
    
        if (token) {
            headers['Authorization'] = `Bearer ${token}`
        }
    
        try {
            const resp = await fetch(`${this.baseUrl}${path}`, {
                method,
                headers,
                body: body ? JSON.stringify(body) : undefined,
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

            localStorage.removeItem('token')
            localStorage.setItem('token', data.token)

            this.token = data.token
            console.log('token recebido:', data.token)
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

        localStorage.removeItem('token')
        this.token = null

    }

    async signUp({ email, senha, nome }) {
        const resp = await this.sendRequest({
            method: 'POST',
            path: '/signup',
            body: { email, senha, nome }
        })
    
        if (resp.ok) {
            const data = await resp.json()

            localStorage.removeItem('token')
            localStorage.setItem('token', data.token)

            this.token = data.token
            console.log('token recebido:', data.token)
            this.user.email = email

            return true 

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

    async gerarCalculo ({ ctc, magnesio, calcio, argila, prnt, potassio, aluminio, sat_desejada, id_metodo }) {
        const resp = await this.sendRequest({
            method: 'POST',
            path:'/analise',
            body: {
                ctc,
                magnesio,
                calcio,
                argila,
                prnt,
                potassio,
                aluminio,
                id_metodo,
                sat_desejada,  
            }
        })

        console.log("aluminio na requisição", aluminio)

        if (resp.ok) {

            const data = await resp.json()
            return data.resultado
        } else {

            const erro = await resp.text()
            console.error("Erro ao fazer analise:", erro)
            throw new Error("Erro ao fazer analise")

        }

    }

    async getListaCalculos () {
        const resp = await this.sendRequest ({
            method: 'GET',
            path:'/listar'
        })
    
        const text = await resp.text()
    
        if (resp.ok) {
            if (!text) {
                throw new Error("Resposta vazia.")
            }
            try {
                return JSON.parse(text)
            } catch (e) {
                console.error("Resposta não é JSON:", text)
                throw new Error("Resposta inválida.")
            }
        } else {
            console.error("Erro ao buscar lista de calculos:", text)
            throw new Error("Erro ao buscar lista de calculos:")
        }
    }

    async gerarExcel (payload, filename = 'Recomendacao.xlsx') {

        console.log("Dados enviados para SaveExcel:", JSON.stringify(payload));
        
        const resp = await this.sendRequest({
            method: 'POST',
            path: '/excel',
            body: payload
        })

        console.log("payload do excel", payload)
    
        if (resp.ok) {
            const blob = await resp.blob()
            const url = window.URL.createObjectURL(blob)
            const link = document.createElement('a')
            link.href = url
            link.setAttribute('download', filename)
            document.body.appendChild(link)
            link.click()
            link.remove()
            console.log('Download realizado com sucesso')
        } else {
            const erro = await resp.text()
            console.error("Erro ao gerar Excel:", erro)
            throw new Error("Erro ao gerar Excel")
        }
    }

    async AlterarDadosUsuario ({ nome, senha }) {
        const body = {}

        if (nome) body.nome = nome
        if (senha) body.senha = senha

        if (Object.keys(body).length === 0) {
            throw new Error("Nenhum dado informado para alterar.")
        }

        const resp = await this.sendRequest({
            method: 'PATCH',
            path: '/alterar',
            body
        })

        if (resp.ok){
            const data = await resp.json()
            console.log("Dados alterados com sucesso", data)
            return true
        } else {
            const erro = await resp.text()
            console.error("Erro ao alterar dados: ", erro)
            throw new Error("Erro ao alterar dados")
        }


    }
}

const calimingAPI = new CalimingAPIClient()
export default calimingAPI