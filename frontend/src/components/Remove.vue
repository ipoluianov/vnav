<script setup>
import { ref } from 'vue'
import { CreateDirectory } from '../../wailsjs/go/main/App';
import { Remove } from '../../wailsjs/go/main/App';
import { EventsEmit } from '../../wailsjs/runtime/runtime';

const props = defineProps({
    panelIndex: Number,
})

const emit = defineEmits(['dialog-accept'])
const dialogError = ref('')

let removeDirectory = async () => {
    //console.log('Remove directory: ' + newDirectoryName.value)
    dialogError.value = await Remove(props.panelIndex);
    if (dialogError.value == '') {
        EventsEmit('updateContent', props.panelIndex);
        emit('dialog-accept', "", props.panelIndex)
    }
}

const cancel = () => {
    emit('dialog-accept', '', props.panelIndex)
}

setTimeout(() => {
    document.getElementById('okButton').focus();
}, 100);

</script>

<template>
    <div style="display: block;">
        <div>{{ dialogError }}</div>
        <button id="okButton" @click="removeDirectory">Create</button>
        <button @click="cancel">Cancel</button>
    </div>
</template>

<style scoped></style>
