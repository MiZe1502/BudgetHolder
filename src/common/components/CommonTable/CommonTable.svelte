<script lang="typescript">
    import { ColumnConfig } from "./utils";

    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'
    import CommonTableHeader from '../CommonTableHeader/CommonTableHeader.svelte'
    import LoadingSpinner from '../LoadingSpinner/LoadingSpinner.svelte'


    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
</script>

<style>

</style>

<section>
    {#if !data || data.length === 0}
        <LoadingSpinner />
    {:else}
        <CommonTableHeader config={config}/>
        {#each data as dataItem}
            <CommonTableRow data={dataItem} config={config} />
        {/each}
    {/if}
</section>