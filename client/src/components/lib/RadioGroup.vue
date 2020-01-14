<template>
    <div>
        <label>{{label}}</label>
        <label v-for="choice in choices" :key="choice.id">
            {{choice.text}}
            <input 
                type="radio" 
                :value="choice.id"
                :checked="choice.id === value"
                :name="name" 
                @change="$emit('input', $event.target.value)" 
            />
        </label>
    </div>
</template>

<script>
export default {
    name: "RadioGroup",
    props: {
        label: String,
        choices: Array,
        name: String,
        value: String
    }
};
</script>

<style scoped>

</style>