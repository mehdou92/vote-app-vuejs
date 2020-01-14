<template>
    <Modal :open="open" :onClose="onClose" title="New Card">
        <template slot="default">
            <div class="grid-row">
                <label for="title">Title</label>
                <input v-model="name" id="title">
            </div>
            <div class="grid-row">
                <label for="description">Description</label>
                <textarea v-model="description" id="description" />
            </div>
            <div class="grid-row">
                <label for="velocity">Velocity</label>
                <input v-model.number="velocity" id="velocity">
            </div>
            <div class="grid-row">
                <RadioGroup 
                    label="Status" 
                    name="status" 
                    v-model="status" 
                    :choices="[{id: 'new', text: 'New'}, {id: 'done', text: 'Done' }]"
                />               
            </div>
        </template>
        <template slot="footer">
            <button @click="handleSubmit">Submit</button>
        </template>
    </Modal>
</template>

<script>
import Modal from './lib/Modal';
import RadioGroup from './lib/RadioGroup';

export default {
    name: "NewCard",
    components: {
        Modal,
        RadioGroup
    },
    data: () => ({
        name: "",
        description: "",
        velocity: 0,
        status: "new"
    }),
    props: {
        open: Boolean,
        onClose: Function
    },
    methods: {
        handleSubmit: function () {
            this.$emit('new-card', {...this.$data});
        }
    }
}
</script>

<style scoped>
    .grid-row {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-content: center;
    }
</style>