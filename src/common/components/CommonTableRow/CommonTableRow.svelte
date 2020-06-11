<script lang="typescript">
    import { ColumnConfig } from "../CommonTable/utils";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any> = {};
</script>

<style>
    .Row {
        display: flex;
        flex-direction: row;
        align-items: center;
    }
</style>

<div class="Row">

    {#each Object.entries(data) as [fieldName, fieldData], index}
            <svelte:component this={config[index].component} data={fieldData}/>
    {/each}

    <!-- {#each config as configItem, index}
        <svelte:component this={configItem.component} data={data}/>
    {/each} -->
</div>
