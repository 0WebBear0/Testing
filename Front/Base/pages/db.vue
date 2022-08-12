<template>
  <div>
    <Header/>
    <Container :table-date="dateTable" />
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
      this.dateTable = await this.$axios.$get('https://mocki.io/v1/d4867d8b-b5d5-4a48-a4ab-79131b5809b8')
    }
  },
  created() {
    this.fetchSomething()
    console.log('test')
    console.log(this.dateTable)
  }
}
</script>

<style scoped lang="scss">

</style>
