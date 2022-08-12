<template>
  <b-navbar toggleable="lg" type="dark" variant="info">

    <!--  Left  -->
    <b-navbar-brand href="/">Back</b-navbar-brand>
    <b-navbar-brand href="/db">DB</b-navbar-brand>
    <b-navbar-brand href="/api">Api</b-navbar-brand>

    <!--  Right  -->
    <b-collapse id="nav-collapse" is-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-form>
          <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </b-nav-form>

        <b-nav-item-dropdown text="Lang" right>
          <b-dropdown-item href="#">EN</b-dropdown-item>
          <b-dropdown-item href="#">ES</b-dropdown-item>
          <b-dropdown-item href="#">RU</b-dropdown-item>
          <b-dropdown-item href="#">FA</b-dropdown-item>
        </b-nav-item-dropdown>

      </b-navbar-nav>

    </b-collapse>
  </b-navbar>

</template>

<script lang="ts">


export default {
  name: "Header"
}
</script>

<style lang="scss">
.header{
  display: flex;
  justify-content: space-around;
  align-items: center;
  width: 100%;
  height: 7vh;
  background-color: antiquewhite;

  &__item{
    cursor: pointer;

    &:hover{
      color: #444;
    }

    &:active{
      opacity: 0.7;
    }
  }
}
</style>
