<template>
  <div class="modal" v-if="open">
      <div class="backdrop" @click="onClose"></div>
      <div class="modal-container">
          <div class="modal-title">
              <h1>{{ title }}</h1>
              <a @click="onClose">X</a>
          </div>
          <div class="modal-content">
              <slot>
              </slot>
          </div>
          <div class="modal-footer">
              <slot name="footer">
                <a @click="onClose">Close</a>
                <a @click="onSubmit">Submit</a>
              </slot>
          </div>
      </div>
  </div>
</template>

<script>
export default {
    name: "Modal",
    props: {
        open: Boolean,
        title: String,
        onClose: Function,
        onSubmit: Function
    }
}
</script>

<style>
    .modal {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .modal > .backdrop {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
        background-color: black;
        opacity: 0.3;
    }
    .modal > .modal-container {
        min-height: 25%;
        max-height: 75%;
        min-width: 40%;
        max-width: 80%;
        background-color: aliceblue;
        display: flex;
        flex-direction: column;
        padding: 10px;
        z-index: 1;
    }
    .modal-title {
        height: 30px;
        width: 100%;
        margin-bottom: 10px;
        border-bottom: 1px solid black;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }
    .modal-content {
        flex: 1;
        width: 100%;
        overflow-x: hidden;
        overflow-y: scroll;
    }
    .modal-footer {
        height: 30px;
        width: 100%;
        margin-top: 10px;
        border-top: 1px solid black;
        display: flex;
        flex-direction: row;
        align-items: center;
    }
</style>