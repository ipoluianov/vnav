<script setup>
import { ref, reactive } from 'vue'
import { GetFilePanelContentAsJson, SetCurrentItemIndex, UpdateContent } from '../../wailsjs/go/main/App'

const data = reactive(
    {
        content: null,
    }
)

const onClickItem = (index) => {
    setCurrentItemIndex(index);
}

const onDblClickItem = (index) => {
}

const styleForItem = (index) => {
    return {
        backgroundColor: index === data.content.currentItemIndex ? '#444' : '#00000000'
    }
}

const styleForHeader = () => {
    return {
        backgroundColor: props.isActive ? '#F00' : '#000'
    }
}

const props = defineProps({
    isActive: Boolean,
    panelIndex: Number
})

window.addEventListener('keydown', function (event) {
    if (props.isActive === false) {
        return
    }
    if (event.key === 'ArrowDown') {
        event.preventDefault();
        setCurrentItemIndex(data.content.currentItemIndex + 1);
    }
    if (event.key === 'ArrowUp') {
        event.preventDefault();
        setCurrentItemIndex(data.content.currentItemIndex - 1);
    }
    if (event.key == 'F2') {
        console.log("F2");
        event.preventDefault();
        UpdateContent(props.panelIndex);
        loadContent();
    }
});

const setCurrentItemIndex = async (index) => {
    console.log("setCurrentItemIndex", index);
    await SetCurrentItemIndex(props.panelIndex, index);
    loadContent();
}

const loadContent = async () => {
    const content = await GetFilePanelContentAsJson(props.panelIndex);
    data.content = JSON.parse(content);
    scrollToRow(idForRow(data.content.currentItemIndex));
}

const idForRow = (index) => {
    return 'panel_' + props.panelIndex + '_row_' + index;
}

function scrollToRow(rowId) {
    scrollToRowIfNotVisible(rowId);
    return;
    const element = document.getElementById(rowId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }
}

function scrollToRowIfNotVisible(rowId) {
    const element = document.getElementById(rowId);
    if (element) {
        const tableContainer = document.querySelector('.table-container');
        const elementRect = element.getBoundingClientRect();
        const containerRect = tableContainer.getBoundingClientRect();
        const thead = document.querySelector('thead');
        const headerHeight = thead ? thead.offsetHeight : 0; 

        const isElementVisible = (
            elementRect.top >= containerRect.top + headerHeight &&
            elementRect.bottom <= containerRect.bottom
        );

        if (!isElementVisible) {
            const scrollOffset = elementRect.top - containerRect.top - headerHeight;
            tableContainer.scrollBy({ top: scrollOffset, behavior: 'instant' });
        }
    }
}

loadContent();

</script>

<template>
    <div style="display: block;" v-if="data.content != null">
        FILEPANEL
        <div :style="styleForHeader()">
            {{ props.isActive }}
        </div>
        <div class="table-container">
            <table>
                <thead>
                    <tr>
                        <th>FileName</th>
                        <th>Size</th>
                    </tr>
                </thead>
                <tbody>
                    <tr :id="idForRow(index)" v-for="(file, index) in data.content.items" :key="index" @click="onClickItem(index)"
                        @dblclick="onDblClickItem(index)" :style="styleForItem(index)">
                        <td class="fileName">{{ file.name }}</td>
                        <td class="fileSize">{{ file.size }}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<style scoped>
table {
    width: 100%;
    border-collapse: collapse;
}

th,
td {
    border: 1px solid #000;
    padding: 8px;
    text-align: left;
    user-select: none;
}

td {
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
}

.table-container {
    height: 70%;
    overflow-y: scroll;
    position: relative;
    background-color: #111;
}

thead {
    position: sticky;
    top: 0;
    z-index: 1;
    background-color: #000;
}
</style>
