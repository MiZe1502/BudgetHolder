<script lang="typescript">
    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
</script>

<style>

    .Header {
        position: relative;
        z-index: 1;
        min-height: 40px;
        display: flex;
        flex-direction: row;
        align-items: center;
        background-color: white;
        box-shadow: 0 2px 4px -2px rgba(0, 0, 0, 0.25);
    }
</style>

<div class="Header">
    {#each config as configItem}
        <span style={configItem.style}>{configItem.header}</span>
    {/each}
</div>