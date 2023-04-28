// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    runtimeConfig: {
        // Warning this is test data and should be left as such even on prod
        // When deploying please override this using env variables instead
        apiSecret: 'de7d235b14f6dec69f9795e0b6c9d5b8e775919ee6f338d26d623d6a77a94da8',
        // Keys within public are also exposed client-side
        public: {
          apiBase: 'http://localhost:5000/'
        }
    }
})
