<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="info" fixed="top" class="nav">
      <b-navbar-brand href="#">Truora - Recetas</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
        </b-navbar-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input size="sm" class="mr-sm-2" v-model="Nombre" placeholder="Buscar"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit" @click="buscar">Buscar</b-button>
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>

<script>
export default {
  methods: {
    buscar() {
      const _this = this;
      if(this.Nombre==""){
        fetch("http://localhost:9001/recetas" + this.Nombre, {
        method: "GET"
      })
        .then(function(response) {
          if (response.ok) {
            return response.text();
          }
          throw new Error("Something went wrong.");
        })
        .then(json => {
          const parsed = JSON.parse(json.replace(/'/g, '"'));
          _this.$root.$emit("RecetasPorNombre", parsed);
        })
        .catch(error => console.log("Error"));
      }

      fetch("http://localhost:9001/recetas?nombre=" + this.Nombre, {
        method: "GET"
      })
        .then(function(response) {
          if (response.ok) {
            return response.text();
          }
          throw new Error("Something went wrong.");
        })
        .then(json => {
          const parsed = JSON.parse(json.replace(/'/g, '"'));
          _this.$root.$emit("RecetasPorNombre", parsed);
        })
        .catch(error => console.log(error));
    }
  },
  data() {
    return {
      Nombre: ""
    };
  }
};
</script>

<style scoped>
</style>
