module.exports = {
  apps: [
    {
      name: "quran-server",
      script: "./quran-server",
      exec_mode: "fork",

      env: {
        PORT: "8080",
        API_SECRET: "my-secret-key",
      },
      watch: false,
    },
  ],
};
