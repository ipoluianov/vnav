<script setup>
import { ref, reactive } from 'vue'
import { GetFilePanelContentAsJson, GoBack, MainAction, SetCurrentItemIndex, UpdateContent } from '../../wailsjs/go/main/App'

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
        height: '30px',
        backgroundColor: props.isActive ? '#F00' : '#000'
    }
}

const props = defineProps({
    isActive: Boolean,
    panelIndex: Number,
    panelHeight: Number,
    panelWidth: Number,
})

window.addEventListener('keydown',  async (event) => {
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
        event.preventDefault();
        await UpdateContent(props.panelIndex);
        await loadContent();
    }
    if (event.key == 'Enter') {
        event.preventDefault();
        await MainAction(props.panelIndex);
        await loadContent();
    }
    if (event.key == 'Backspace') {
        event.preventDefault();
        await GoBack(props.panelIndex);
        await loadContent();
    }
});

const setCurrentItemIndex = async (index) => {
    console.log("setCurrentItemIndex", '['+props.panelIndex+']', index, " isActive:", props.isActive);
    await SetCurrentItemIndex(props.panelIndex, index);
    await loadContent();
}

const loadContent = async () => {
    const content = await GetFilePanelContentAsJson(props.panelIndex);
    data.content = JSON.parse(content);
    console.log("loadContent", data.content.currentItemIndex);
    setTimeout(() => {
        scrollToRow(idForRow(data.content.currentItemIndex));
    }, 10);
    //scrollToRow(idForRow(data.content.currentItemIndex));
}

const idForRow = (index) => {
    return 'panel_' + props.panelIndex + '_row_' + index;
}

const  scrollToRow = (rowId) => {
    scrollToRowIfNotVisible(rowId);
    return;
    const element = document.getElementById(rowId);
    if (element) {
        element.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }
}

const scrollToRowIfNotVisible = (rowId, behavior = 'instant') => {
    const element = document.getElementById(rowId);
    if (element) {
        const tableContainer = document.getElementById(tableContainerId());
        const elementRect = element.getBoundingClientRect();
        const containerRect = tableContainer.getBoundingClientRect();
        const thead = document.querySelector('thead');
        const headerHeight = thead ? thead.offsetHeight : 0;

        const isElementVisible = (
            elementRect.top >= containerRect.top + headerHeight &&
            elementRect.bottom <= containerRect.bottom
        );

        if (!isElementVisible) {
            let scrollOffset = 0;

            if (elementRect.top < containerRect.top + headerHeight) {
                scrollOffset = elementRect.top - containerRect.top - headerHeight;
            } else if (elementRect.bottom > containerRect.bottom) {
                scrollOffset = elementRect.bottom - containerRect.bottom;
            }

            tableContainer.scrollBy({ top: scrollOffset, behavior });
            console.log("SCROLL to", rowId);
        }
    }
};


loadContent();

const styleForContainer = () => {
    //console.log("styleForContainer", props.panelHeight);

    return {
        height: (props.panelHeight-32) + 'px',
        overflowY: 'scroll',
        position: 'relative',
        backgroundColor: '#111',
    }
}

const tableContainerId = () => {
    return 'table-container-' + props.panelIndex;
}

const styleForColumn = (column, type) => {
    let sizeColumnWidth = 100;

    if (column == 'filename') {
        return {
            width: props.panelWidth - sizeColumnWidth + 'px',
            maxWidth: sizeColumnWidth + 'px',
            textAlign: 'left',
            whiteSpace: 'nowrap',
            overflow: 'hidden',
            textOverflow: 'ellipsis',
        }
    }
    if (column == 'size') {
        return {
            width: sizeColumnWidth + 'px',
            maxWidth: sizeColumnWidth + 'px',
            textAlign: 'right',
            whiteSpace: 'nowrap',
            overflow: 'hidden',
            textOverflow: 'ellipsis',
        }
    }

    return {
    }
}

</script>

<template>
    <div style="display: block;" v-if="data.content != null">
        <div :style="styleForHeader()">
            {{ data.content.currentPath }}
        </div>
        <div class="scrollable-content" :id="tableContainerId()" :style="styleForContainer()">
            <table>
                <thead>
                    <tr>
                        <th :style="styleForColumn('filename', 'header')">FileName</th>
                        <th :style="styleForColumn('size', 'header')">Size</th>
                    </tr>
                </thead>
                <tbody>
                    <tr :id="idForRow(index)" v-for="(file, index) in data.content.items" :key="index"
                        @click="onClickItem(index)" @dblclick="onDblClickItem(index)" :style="styleForItem(index)">
                        <td :style="styleForColumn('filename', 'header')" class="fileName">{{ file.name }}</td>
                        <td :style="styleForColumn('size', 'header')" class="fileSize">{{ file.size }}</td>
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


thead {
    position: sticky;
    top: 0;
    z-index: 1;
    background-color: #000;
}

.scrollable-content {
  scrollbar-width: thin;
  scrollbar-color: #EEE #333333;
}

.scrollable-content::-webkit-scrollbar {
  width: 12px;
}

.scrollable-content::-webkit-scrollbar-track {
  background: #333333;
}

.scrollable-content::-webkit-scrollbar-thumb {
  background-color: #EEE;
  border-radius: 10px;
  border: 3px solid #333333;
}

.scrollable-content::-webkit-scrollbar-thumb:hover {
  background-color: #EEE;
}


</style>
