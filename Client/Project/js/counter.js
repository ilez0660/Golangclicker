const Counter = {
  count: 0,
  pending: 0,
  FLUSH_EVERY: 10,
  flushing: false,
  els: {},

  init() {
    this.els = {
      page:   document.getElementById("counterPage"),
      user:   document.getElementById("userInfo"),
      value:  document.getElementById("counterValue"),
      button: document.getElementById("counterBtn"),
      logout: document.getElementById("logoutBtn"),
    };

    this.els.button.addEventListener("click", () => this.handleClick());
    this.els.logout.addEventListener("click", () => this.handleLogout());

    window.addEventListener("beforeunload", () => this.flushOnUnload());

    document.addEventListener("visibilitychange", () => {
      if (document.visibilityState === "hidden") {
        this.flushOnUnload();
      }
    });

    window.addEventListener("pagehide", () => this.flushOnUnload());
  },

  async open(user) {
    this.els.user.textContent = user.full_name || user.email;

    try {
      const data = await API.getCount();
      this.count = data.count || 0;
    } catch {
      this.count = 0;
    }
    this.pending = 0;
    this.render();
  },

  render() {
    this.els.value.textContent = String(this.count);
  },

  handleClick() {
    this.count += 1;
    this.pending += 1;
    this.render();

    if (this.pending >= this.FLUSH_EVERY) {
      this.flush();
    }
  },

  async flush() {
    if (this.pending === 0 || this.flushing) return;

    this.flushing = true;
    const delta = this.pending;
    this.pending = 0;

    try {
      const data = await API.addClick(delta);
      this.count = data.count;
      this.render();
    } catch (err) {
      this.pending += delta;
      console.error("Не удалось сохранить клики:", err);
    } finally {
      this.flushing = false;
    }
  },

  flushOnUnload() {
    if (this.pending === 0) return;

    const delta = this.pending;
    this.pending = 0;

    try {
      API.saveClicksOnUnload(delta);
    } catch (err) {
      this.pending += delta;
    }
  },

  async handleLogout() {
    await this.flush();
    await Auth.logout();
  },
};