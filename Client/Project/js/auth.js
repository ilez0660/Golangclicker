const Auth = {
  mode: "login",
  currentUser: null,
  els: {},

  init() {
    this.els = {
      page:       document.getElementById("authPage"),
      title:      document.getElementById("authTitle"),
      form:       document.getElementById("authForm"),
      fullName:   document.getElementById("fullNameInput"),
      email:      document.getElementById("emailInput"),
      error:      document.getElementById("errorMsg"),
      submit:     document.getElementById("submitBtn"),
      switchText: document.getElementById("switchText"),
      switchBtn:  document.getElementById("switchBtn"),
    };

    this.els.switchBtn.addEventListener("click", () => {
      this.setMode(this.mode === "login" ? "register" : "login");
    });

    this.els.form.addEventListener("submit", (e) => this.handleSubmit(e));

    this.setMode("login");
  },

  setMode(newMode) {
    this.mode = newMode;
    this.els.error.textContent = "";

    if (this.mode === "login") {
      this.els.title.textContent      = "Вход";
      this.els.submit.textContent     = "Войти";
      this.els.fullName.classList.add("hidden");
      this.els.switchText.textContent = "Нет аккаунта?";
      this.els.switchBtn.textContent  = "Зарегистрироваться";
    } else {
      this.els.title.textContent      = "Регистрация";
      this.els.submit.textContent     = "Зарегистрироваться";
      this.els.fullName.classList.remove("hidden");
      this.els.switchText.textContent = "Уже есть аккаунт?";
      this.els.switchBtn.textContent  = "Войти";
    }
  },

  validate(email, fullName) {
    if (!email || !email.includes("@") || !email.includes(".")) {
      return "Введите корректный email";
    }
    if (this.mode === "register" && !fullName) {
      return "Введите полное имя";
    }
    if (this.mode === "register" && fullName.length < 3) {
      return "Имя должно быть не короче 3 символов";
    }
    return null;
  },

  async handleSubmit(e) {
    e.preventDefault();
    this.els.error.textContent = "";

    const email    = this.els.email.value.trim();
    const fullName = this.els.fullName.value.trim();

    const err = this.validate(email, fullName);
    if (err) {
      this.els.error.textContent = err;
      return;
    }

    try {
      this.els.submit.disabled = true;

      let user;
      if (this.mode === "register") {
        user = await API.register(fullName, email);
      } else {
        user = await API.login(email);
      }

      this.currentUser = user;
      Pages.show("counter");
      Counter.open(user);

    } catch (err) {
      this.els.error.textContent = err.message;
    } finally {
      this.els.submit.disabled = false;
    }
  },

  async logout() {
    try {
      await API.logout();
    } catch {}

    this.currentUser = null;
    this.els.email.value = "";
    this.els.fullName.value = "";
    this.setMode("login");
    Pages.show("auth");
  },

  async restore() {
    try {
      const user = await API.me();
      if (user && user.email) {
        this.currentUser = user;
        Pages.show("counter");
        Counter.open(user);
        return true;
      }
    } catch {}
    return false;
  },
};