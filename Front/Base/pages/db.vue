<template>
  <div>
    <Header/>
    <Container :table-date="dateTable?.users" />
  </div>
</template>

<script>

import Header from "~/components/Header.vue";
import Container from "~/components/Container.vue";

export default {
  name: "db",
  components: {Container, Header},

  data() {
    let dateTable = null
    return{
      dateTable
    }
  },

  methods: {
    async fetchSomething() {
      this.dateTable = await this.$axios.$get('http://localhost:8080/users', (res) =>{
        res.set('Access-Control-Allow-Origin', 'http://localhost:8080');
      })
    }
  },
  created() {
    this.fetchSomething()
  }
}
</script>

<style scoped lang="scss">

</style>
