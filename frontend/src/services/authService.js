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
      console.error("Erro ao logar:", await resp.text());
      return false;
    }
  },

  async logout() {
    await apiClient.sendRequest({
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
      console.error("Erro ao cadastrar:", await resp.text());
      throw new Error("Erro ao cadastrar");
    }
  },
};
