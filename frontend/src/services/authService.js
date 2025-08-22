import calimingAPI from "./CalimingAPIClient";


export default {
  async login({ email, senha }) {
    const resp = await calimingAPI.sendRequest({
      method: 'POST',
      path: '/login',
      body: { email, senha },
    });

    if (resp.ok) {
      const data = await resp.json();
      localStorage.setItem('token', data.token);
      return true;
    } else {

      if (resp.status === 401) {
        alert("Email ou senha inválidos. Faça login novamente.");
      }

      if (resp.status === 400) {
        alert("Credenciais faltando. Tente novamente.")
      }

      console.error("Erro ao logar:", await resp.text());
      return false;
    }
  },

  async logout() {
    await calimingAPI.sendRequest({
      method: 'POST',
      path: '/logout',
    });

    localStorage.removeItem('token');
  },

  async signUp({ email, senha, nome }) {
    const resp = await calimingAPI.sendRequest({
      method: 'POST',
      path: '/signup',
      body: { email, senha, nome },
    });

    if (resp.ok) {
      const data = await resp.json();
      localStorage.setItem('token', data.token);
      return true;
    } else {
      if (resp.status === 400) {
        alert("Esse email já está cadastrado. Tente outro.")
      }
      const message = await resp.text()
      throw new Error(message);
    }
  },
};
