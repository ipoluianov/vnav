<script setup>
import FilePanel from './FilePanel.vue'
import SelectDriverDialog from './SelectDriverDialog.vue';
import CreateDirectoryDialog from './CreateDirectoryDialog.vue';
import { EventsEmit } from '../../wailsjs/runtime/runtime';
import { GetFilePanelContentAsJson, CreateDirectory } from '../../wailsjs/go/main/App'

import { ref, reactive } from 'vue'

const dialogKeyDown = (event) => {
    if (event.key == 'Escape') {
        event.preventDefault();
        closePanelDialog();
    }
}

window.addEventListener('keydown', function (event) {
    if (dialogIsOpen.value == true) {
        dialogKeyDown(event);
        return
    }

    if (event.altKey && event.key == 'F1') {
        event.preventDefault();
        openPanelDialog("selectDriver0", 0);
    }

    if (event.altKey && event.key == 'F2') {
        event.preventDefault();
        openPanelDialog("selectDriver1", 1);
    }

    if (!event.altKey && event.key == 'F7') {
        event.preventDefault();
        openPanelDialog("createDirectory", -1);
    }

    if (event.key === 'Tab') {
        event.preventDefault();
        data.currentTab = data.currentTab === 0 ? 1 : 0;
    }
});

const onCloseDialog = () => {
}   

const activatePanel = (index) => {
    data.currentTab = index;
}

const data = reactive(
    {
        currentTab: 0,
        height: '300px',
        toolbarHeight: 30,
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

let dialogIsOpen = ref(false);
let dialogCode = '';    
let dialogPosition = -1;
let dialogOperationIsProcessing = false;
let dialogError = '';

const openPanelDialog = (dialogToOpen, position) => {
    dialogCode = dialogToOpen;
    dialogPosition = position;
    dialogIsOpen.value = true;
}

const closePanelDialog = () => {
    dialogIsOpen.value = false;
}

const dialogDivStyle = () => {
    let top = 100;
    let left = 0;
    let dialogWidth = 300;
    let dialogHeight = 200;
    console.log('dialogPosition', dialogPosition);

    if (dialogPosition == 0) {
        left = 0 + panelWidth() / 2 - dialogWidth / 2;
    }

    if (dialogPosition == 1) {
        left = panelWidth() + data.sepWidth + panelWidth() / 2 - dialogWidth / 2;
    }

    if (dialogPosition == -1) {
        left = data.width / 2 - dialogWidth / 2;
        top = data.height / 2 - dialogHeight / 2;
    }

    return {
        top: top + 'px',
        left: left + 'px',
        width: dialogWidth + 'px',
        height: dialogHeight + 'px',
    }
}

const overlayId = () => {
    return 'overlay_';
}

const dialogId = () => {
    return 'dialog_';
}

const createDirectory = async (directoryName, panelIndex) => {
        closePanelDialog();
}

</script>

<template>
    <div style="display: flex; flex-direction: column; height: 100%; max-height: 100%;">
        <div :style="styleForToolbar()">TOOLBAR</div>
        <div style="position: relative;">
            <div v-if="dialogIsOpen">
                <div :id="overlayId()" class="overlay" @click="closePanelDialog"></div>
                <div :id="dialogId()" class="dialog" :style="dialogDivStyle()">
                    <SelectDriverDialog v-if="dialogCode == 'selectDriver0'" />
                    <SelectDriverDialog v-if="dialogCode == 'selectDriver1'" />
                    <CreateDirectoryDialog v-if="dialogCode == 'createDirectory'" :panelIndex="data.currentTab" @dialog-accept="createDirectory" />
                </div>
            </div>

            <div style="flex-grow: 1; display: flex; flex-direction: row; background-color: #777;">
                <FilePanel @activate-panel="activatePanel(0)" :style="styleForFilePanel()" :panelHeight="data.height"
                    :panelWidth="panelWidth()" :panelIndex="0" :isActive="data.currentTab == 0 && !dialogIsOpen" />
                <div :style="styleForSeparator()"></div>
                <FilePanel @activate-panel="activatePanel(1)" :style="styleForFilePanel()" :panelHeight="data.height"
                    :panelWidth="panelWidth()" :panelIndex="1" :isActive="data.currentTab == 1 && !dialogIsOpen" />
            </div>
        </div>
        <div :style="styleForFooter()">FOOTER</div>
    </div>
</template>
<style scoped>
.dialog {
    display: block;
    position: absolute;
    top: 100px;
    left: 100px;
    background-color: #222;
    border: 1px solid #555;
    padding: 20px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    z-index: 1000;
}

.overlay {
    display: block;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    z-index: 999;
}
</style>
