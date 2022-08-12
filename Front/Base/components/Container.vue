<template>
  <div class="container1">
    <b-table class="container1__box" striped hover :items="TableDate"></b-table>
  </div>
</template>


<script lang="js">

export default {
  name: "Container",
  props: {
    TableDate: {
      required: false,
      type: Object | undefined,
    }
  },


}
</script>

<style scoped lang="scss">
.container1 {
  width: 100%;
  height: 92vh;
  background-color: lavender;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  &__title{
    padding-bottom: 15px;
    font-size: 24px;
  }

  &__box {
    width: 50vw;
    cursor: pointer;
  }

}
</style>
