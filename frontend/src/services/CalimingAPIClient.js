class CalimingAPIClient {

    constructor() {
        this.baseUrl = import.meta.env.VITE_API_BASE_URL || "http://caliming-backend-lb5bdj-b519c8-92-112-177-3.traefik.me";
        console.log(this.baseUrl)
    }
    
  

    async sendRequest({ path, method, body }) {
        const token = localStorage.getItem('token');
        const headers = { 'Content-Type': 'application/json' };

        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }

        try {
            const resp = await fetch(`${this.baseUrl}${path}`, {
                method,
                headers,
                body: body ? JSON.stringify(body) : undefined,
                credentials: 'include',
            });

            if (resp.status === 401) {
                alert("Sessão expirada. Redirecionando para login.");
                localStorage.removeItem('token');
                window.location.href = '/login';
                return;
            }

            return resp;
        } catch (err) {
            console.error("Erro na requisição:", err);
            throw err;
        }
    }
}

const calimingAPI = new CalimingAPIClient()
export default calimingAPI