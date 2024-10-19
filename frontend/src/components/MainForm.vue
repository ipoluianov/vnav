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
    }
)
</script>

<template>
    <div style="display: flex; flex-direction: column; height: 100%;">
        <div style="background-color: #E00;">TOOLBAR</div>
        <div style="flex-grow: 1; display: flex; flex-direction: row; background-color: #777; max-height: 80%;">
            <FilePanel style="flex-grow: 1;" :panelIndex="0" :isActive="data.currentTab == 0" />
            <div style="width: 3px; background-color: #0F0;"></div>
            <FilePanel style="flex-grow: 1;" :panelIndex="1" :isActive="data.currentTab == 1" />
        </div>
        <div style="background-color: #A94;">FOOTER</div>
    </div>
</template>
