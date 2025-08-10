module.exports = {
  apps: [
    {
      name: "quran-server",
      script: "./quran-server",
      env: {
        PORT: "8080",
        API_SECRET: "my-secret-key",
      },
      watch: false,
    },
  ],
};
