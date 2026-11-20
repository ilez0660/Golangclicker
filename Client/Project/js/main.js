const Pages = {
  auth:    null,
  counter: null,

  init() {
    this.auth    = document.getElementById("authPage");
    this.counter = document.getElementById("counterPage");
  },

  show(name) {
    this.auth.classList.toggle("active", name === "auth");
    this.counter.classList.toggle("active", name === "counter");
  },
};

(async function boot() {
  Pages.init();
  Auth.init();
  Counter.init();

  const restored = await Auth.restore();
  if (!restored) {
    Pages.show("auth");
  }
})();