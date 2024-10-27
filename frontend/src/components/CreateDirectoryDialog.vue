<script setup>
import { ref } from 'vue'
import { CreateDirectory } from '../../wailsjs/go/main/App';
import { EventsEmit } from '../../wailsjs/runtime/runtime';

const props = defineProps({
    panelIndex: Number,
})

const emit = defineEmits(['dialog-accept'])
let newDirectoryName = ref('')
const dialogError = ref('')

let createDirectory = async () => {
    console.log('Creating directory: ' + newDirectoryName.value)
    dialogError.value = await CreateDirectory(newDirectoryName.value, props.panelIndex);
    EventsEmit('updateContent', props.panelIndex);
    emit('dialog-accept', newDirectoryName.value, props.panelIndex)
}

const cancel = () => {
    emit('dialog-accept', '', props.panelIndex)
}

setTimeout(() => {
    document.getElementById('createDirectoryNameField').focus();
}, 100);

</script>

<template>
    <div style="display: block;">
        <input id="createDirectoryNameField" type="text" v-model="newDirectoryName"
            placeholder="Enter new directory name" @keyup.enter="createDirectory">
        <div>{{ dialogError }}</div>
        <button @click="createDirectory">Create</button>
        <button @click="cancel">Cancel</button>
    </div>
</template>

<style scoped></style>
