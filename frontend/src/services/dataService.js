import calimingAPI from "./CalimingAPIClient";

export default {
  async getDados() {
    const resp = await calimingAPI.sendRequest({ method: 'GET', path: '/profile' });
    if (resp.ok) return await resp.json();
    throw new Error(await resp.text());
  },

  async gerarCalculo(payload) {
    const resp = await calimingAPI.sendRequest({
      method: 'POST',
      path: '/analise',
      body: payload,
    });

    if (!resp.ok) throw new Error(await resp.text());

    const text = await resp.text();
    if (!text) throw new Error("Resposta vazia.");
    return JSON.parse(text);
  },

  async getListaCalculos(page = 1, limit = 10) {
    const resp = await calimingAPI.sendRequest({
      method: 'GET',
      path: `/listar?page=${page}&limit=${limit}`,
    });

    const text = await resp.text();
    if (!resp.ok) throw new Error(text);
    return JSON.parse(text);
  },

  async gerarExcel(payload, filename = 'Recomendacao.xlsx') {
    const resp = await calimingAPI.sendRequest({
      method: 'POST',
      path: '/excel',
      body: payload,
    });

    if (!resp.ok) throw new Error(await resp.text());

    const blob = await resp.blob();
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    link.remove();
  },

  async alterarDadosUsuario({ nome, senha }) {
    const body = {};
    if (nome) body.nome = nome;
    if (senha) body.senha = senha;

    const resp = await calimingAPI.sendRequest({
      method: 'PATCH',
      path: '/alterar',
      body,
    });

    if (!resp.ok) throw new Error(await resp.text());

    return await resp.json();
  },
};
