<script setup>
import { ref } from 'vue'

const props = defineProps({
    panelIndex: Number,
})

const emit = defineEmits(['dialog-accept'])
let newDirectoryName = ref('')
let createDirectory = () => {
    console.log('Creating directory: ' + newDirectoryName.value)
    emit('dialog-accept', newDirectoryName.value, props.panelIndex)
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
