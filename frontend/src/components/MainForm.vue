<script setup>
import FilePanel from './FilePanel.vue'
import { ref, reactive } from 'vue'

window.addEventListener('keydown', function (event) {
    if (event.key === 'Tab') {
        event.preventDefault();
        data.currentTab = data.currentTab === 0 ? 1 : 0;
    }
});

const data = reactive(
    {
        currentTab: 0,
        height: '300px',
    }
)

const updateHeight = () => {
    data.height = (window.innerHeight - 200);
}

function onWindowResize() {
    updateHeight();
}

window.addEventListener('resize', onWindowResize);

updateHeight();

setInterval(() => {
    updateHeight();
}, 200);

const styleForFilePanel = () => {
    return {
        flexGrow: 1,
        height: data.height + 'px',
    }
}

const styleForSeparator = () => {
    return {
        //flexGrow: 1,
        width: '3px',
        height: data.height + 'px',
        backgroundColor: '#0F0'
    }
}

</script>

<template>
    <div style="display: flex; flex-direction: column; height: 100%; max-height: 100%;">
        <div style="background-color: #E00; height: 100px;">TOOLBAR</div>
        <div style="flex-grow: 1; display: flex; flex-direction: row; background-color: #777;">
            <FilePanel :style="styleForFilePanel()" :panelHeight="data.height" :panelIndex="0"
                :isActive="data.currentTab == 0" />
            <div :style="styleForSeparator()"></div>
            <FilePanel :style="styleForFilePanel()" :panelHeight="data.height" :panelIndex="1"
                :isActive="data.currentTab == 1" />
        </div>
        <div style="background-color: #A94; height: 100px;">FOOTER</div>
    </div>
</template>
