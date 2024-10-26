<script setup>
import { ref } from 'vue'

const props = defineProps({
    parentDirectory: String,
})

let newDirectoryName = ref('')
let createDirectory = () => {
    console.log('Creating directory: ' + props.parentDirectory + "/" + newDirectoryName.value)
}
</script>

<template>
    <div>
        <input type="text" v-model="newDirectoryName" placeholder="Enter new directory name"
            @keyup.enter="createDirectory">
        <button @click="createDirectory">Create</button>
    </div>
</template>

<style scoped></style>
