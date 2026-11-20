const API_CONFIG = window.APP_CONFIG;

const API = {
  url(path) {
    return API_CONFIG.BASE_URL + API_CONFIG.API_PREFIX + path;
  },

  async request(path, options = {}) {
    const response = await fetch(this.url(path), {
      credentials: "include",
      headers: { "Content-Type": "application/json" },
      ...options,
    });

    const data = await response.json().catch(() => ({}));

    if (!response.ok) {
      throw new Error(data.error || data.message || "Ошибка сервера");
    }
    return data;
  },

  login(email) {
    return this.request("/users/login", {
      method: "POST",
      body: JSON.stringify({ email }),
    });
  },

  register(fullName, email) {
    return this.request("/users", {
      method: "POST",
      body: JSON.stringify({ full_name: fullName, email }),
    });
  },

  me() {
    return this.request("/users/me");
  },

  getCount() {
    return this.request("/counter");
  },

  addClick(delta = 1) {
    return this.request("/counter", {
      method: "POST",
      body: JSON.stringify({ delta }),
    });
  },

  logout() {
    return this.request("/logout", { method: "POST" });
  },

  saveClicksOnUnload(delta) {
    return fetch(this.url("/counter"), {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ delta }),
      keepalive: true,
      credentials: "include",
    });
  },
};