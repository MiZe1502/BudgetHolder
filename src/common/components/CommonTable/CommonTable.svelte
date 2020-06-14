<script lang="typescript">
    import { ColumnConfig } from "./utils";

    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'
    import CommonTableTitle from '../CommonTableTitle/CommonTableTitle.svelte'
    import CommonTableHeader from '../CommonTableHeader/CommonTableHeader.svelte'
    import LoadingSpinner from '../LoadingSpinner/LoadingSpinner.svelte'


    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
    export let title: string = "";
</script>

<style>

</style>

<section>
    {#if !data || data.length === 0}
        <LoadingSpinner />
    {:else}
        <CommonTableTitle title={title} dataLength={data.length}/>
        <CommonTableHeader config={config}/>
        {#each data as dataItem}
            <CommonTableRow data={dataItem} config={config} />
        {/each}
    {/if}
</section>