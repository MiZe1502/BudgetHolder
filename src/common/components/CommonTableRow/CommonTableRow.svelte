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

    {#each Object.entries(data) as [cat_name, cat_number], index}
        {cat_name}: {cat_number}
    {/each}

    {#each config as configItem, index}
        <svelte:component this={configItem.component} data={data}/>
    {/each}
</div>
