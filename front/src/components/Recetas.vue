<template>
  <div class="hello">
    <h2>Recetas</h2>
    <b-table striped hover :items="recetas" :fields="fields" :sort-by.sync="sortBy">
      <template slot="Acciones" class="mr-2" slot-scope="row">
        <b-button size="sm" variant="warning" @click="row.toggleDetails">
          <font-awesome-icon icon="edit"/>
        </b-button>

        <b-button size="sm" variant="danger" @click="eliminarReceta(row)">
          <font-awesome-icon icon="trash"/>
        </b-button>
      </template>
      <template slot="row-details" slot-scope="row">
        <b-card>
          <b-form inline class="form" @submit="onSubmit(row)">
            <label class="sr-only" for="inline-form-input-name">Name</label>
            <b-input
              id="inline-form-input-name"
              class="mb-2 mr-sm-2 mb-sm-0"
              :placeholder="row.item.Nombre"
              v-model="form.Nombre"
            ></b-input>

            <label class="sr-only" for="inline-form-input-username">Descripcion</label>
            <b-input-group class="mb-2 mr-sm-2 mb-sm-0">
              <b-input
                v-model="form.Descripcion"
                id="inline-form-input-username"
                :placeholder="row.item.Descripcion"
              ></b-input>
            </b-input-group>

            <b-button variant="success" type="submit" class="guardar">Guardar</b-button>

            <b-button variant="success" @click="row.toggleDetails">Cerrar</b-button>
          </b-form>
        </b-card>
      </template>
    </b-table>
    <b-button variant="success" class="pluss" @click="open">
      <font-awesome-icon icon="plus"/>
    </b-button>

    <b-form inline class="form formF" @submit="create" v-bind:style="{visibility: active}">
      <label class="sr-only" for="inline-form-input-name">Nombre</label>
      <b-input
        id="inline-form-input-name"
        class="mb-2 mr-sm-2 mb-sm-0"
        placeholder="Nombre"
        v-model="formNew.Nombre"
      ></b-input>

      <label class="sr-only" for="inline-form-input-username">Descripcion</label>
      <b-input-group class="mb-2 mr-sm-2 mb-sm-0">
        <b-input
          v-model="formNew.Descripcion"
          id="inline-form-input-username"
          placeholder="Descripcion"
        ></b-input>
      </b-input-group>

      <b-button variant="success" type="submit" class="guardar">Guardar</b-button>

      <b-button variant="success" @click="close">Cerrar</b-button>
    </b-form>

    <b-alert
      :show="dismissCountDown"
      variant="success"
      @dismissed="dismissCountDown=0"
      @dismiss-count-down="countDownChanged"
      class="alerta"
    >
      <p>{{mensaje}}</p>
    </b-alert>
  </div>
</template>

<script>
export default {
  name: "Recetas",
  data() {
    return {
      mensaje: "",
      dismissSecs: 5,
      dismissCountDown: 0,
      fields: ["Nombre", "Descripcion", "Acciones"],
      recetas: [],
      form: {
        Nombre: "",
        Descripcion: ""
      },
      formNew: {
        Nombre: "",
        Descripcion: ""
      },
      active: "hidden",
      sortBy: "Nombre"
    };
  },
  mounted: function() {
    const _this = this;
    fetch("http://localhost:9001/recetas", {
      method: "GET"
    })
      .then(function(response) {
        if (response.ok) {
          return response.text();
        }

        _this.showAlert("Error");
        throw new Error("Something went wrong.");
      })
      .then(json => {
        const parsed = JSON.parse(json.replace(/'/g, '"'));
        _this.recetas = parsed;
      })
      .catch(error => _this.showAlert(error));

      _this.$root.$on("RecetasPorNombre", function(data){
        _this.recetas = data
      })
  },

  methods: {
    handleErrors(response) {
      if (!response.ok) {
        throw Error(response.statusText);
      }
      return response;
    },
    eliminarReceta(receta) {
      const _this = this;
      var cuerpo = {
        nombre: receta.item.Nombre
      };
      fetch("http://localhost:9001/deleteReceta", {
        method: "DELETE",
        body: JSON.stringify(cuerpo),
        contentType: "text/plain",
        mode: "cors"
      })
        .then(function(response) {
          if (response.ok) {
            return response.text();
          }
          _this.showAlert("Error");
          throw new Error("Something went wrong.");
        })

        .then(response => {
          _this.showAlert(response);
          var indice = _this.recetas.indexOf(receta.item);
          if (indice > -1) {
            _this.recetas.splice(indice, 1);
          }
        })
        .catch(error => _this.showAlert(error));
    },
    countDownChanged(dismissCountDown) {
      this.dismissCountDown = dismissCountDown;
    },
    showAlert(resp) {
      this.mensaje = resp;
      this.dismissCountDown = this.dismissSecs;
    },
    onSubmit(row) {
      const _this = this;
      var cuerpo = {
        Nombre: this.form.Nombre,
        Descripcion: this.form.Descripcion,
        NombreActualizar: row.item.Nombre
      };
      fetch("http://localhost:9001/updateReceta", {
        method: "PUT",
        body: JSON.stringify(cuerpo),
        contentType: "text/plain",
        mode: "cors"
      })
        .then(function(response) {
          if (response.ok) {
            return response.text();
          } else {
            _this.showAlert("Error");
            throw new Error("Something went wrong.");
          }
        })
        .then(response => {
          _this.showAlert(response);
          var indice = _this.recetas.indexOf(row.item);
          _this.recetas[indice].Nombre = this.form.Nombre;
          _this.recetas[indice].Descripcion = this.form.Descripcion;
          this.form = "";
          this.Descripcion = "";
        })
        .catch(error => console.log(error));
    },
    open() {
      this.active = "visible";
    },
    close() {
      this.active = "hidden";
    },
    create() {
      const _this = this;
      var cuerpo = {
        Nombre: _this.formNew.Nombre,
        Descripcion: _this.formNew.Descripcion
      };
      fetch("http://localhost:9001/createReceta", {
        method: "POST",
        body: JSON.stringify(cuerpo),
        contentType: "text/plain",
        mode: "cors"
      })
        .then(function(response) {
          if (response.ok) {
            return response.text();
          } else {
            _this.showAlert("Error");
            throw new Error("Something went wrong.");
          }
        })
        .then(response => {
          _this.showAlert(response);
          _this.recetas.push(cuerpo)
          this.form = "";
          this.Descripcion = "";
        })
        .catch(error => console.log(error));
    }
  }
};
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

.alerta {
  position: fixed;
  bottom: 0;
  right: 0;
  margin-right: 10px;
}

.form {
  width: 50%;
  margin: 0 auto;
}

.guardar {
  margin-right: 10px;
}

.pluss {
  margin-bottom: 15px;
}

.formF {
  margin-bottom: 15px;
}
</style>
