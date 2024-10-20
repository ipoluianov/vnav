<script setup>
import FilePanel from './FilePanel.vue'
import { ref, reactive } from 'vue'

window.addEventListener('keydown', function (event) {
    if (event.key === 'Tab') {
        event.preventDefault();
        data.currentTab = data.currentTab === 0 ? 1 : 0;
    }
});

const activatePanel = (index) => {
    data.currentTab = index;
}

const data = reactive(
    {
        currentTab: 0,
        height: '300px',
        toolbarHeight: 100,
        footerHeight: 30,
    }
)

const updateSizes = () => {
    data.height = (window.innerHeight - data.toolbarHeight - data.footerHeight);
    data.width = (window.innerWidth);
    data.sepWidth = 3;
}

function onWindowResize() {
    updateSizes();
}

window.addEventListener('resize', onWindowResize);

updateSizes();

setInterval(() => {
    updateSizes();
}, 200);

const panelWidth = () => {
    return (data.width - data.sepWidth) / 2;
}

const styleForFilePanel = () => {
    return {
        flexGrow: 1,
        width: panelWidth() + 'px',
        height: data.height + 'px',
    }
}

const styleForSeparator = () => {
    return {
        //flexGrow: 1,
        width: data.sepWidth + 'px',
        height: data.height + 'px',
        backgroundColor: '#CCC'
    }
}

const styleForToolbar = () => {
    return {
        height: data.toolbarHeight + 'px',
        backgroundColor: '#222'
    }
}

const styleForFooter = () => {
    return {
        height: data.footerHeight + 'px',
        backgroundColor: '#222'
    }
}

</script>

<template>
    <div style="display: flex; flex-direction: column; height: 100%; max-height: 100%;">
        <div :style="styleForToolbar()">TOOLBAR</div>
        <div style="flex-grow: 1; display: flex; flex-direction: row; background-color: #777;">
            <FilePanel @activate-panel="activatePanel(0)" :style="styleForFilePanel()" :panelHeight="data.height" :panelWidth="panelWidth()" :panelIndex="0"
                :isActive="data.currentTab == 0" />
            <div :style="styleForSeparator()"></div>
            <FilePanel @activate-panel="activatePanel(1)" :style="styleForFilePanel()" :panelHeight="data.height"  :panelWidth="panelWidth()" :panelIndex="1"
                :isActive="data.currentTab == 1" />
        </div>
        <div :style="styleForFooter()">FOOTER</div>
    </div>
</template>
