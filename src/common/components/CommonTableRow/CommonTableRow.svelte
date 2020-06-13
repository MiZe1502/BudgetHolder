<script lang="typescript">
    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
</script>

<style>
    .Row {
        min-height: 40px;
        display: flex;
        flex-direction: row;
        align-items: center;
    }

    .Row:nth-of-type(2n) {
        background-color: #faf8f6;
    }
</style>

<div class="GlobalPadding Row">
    {#each Object.entries(data) as [fieldName, fieldData], index}
            <svelte:component style={config[index].style} this={config[index].component} data={fieldData}/>
    {/each}
</div>
