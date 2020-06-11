<script lang="typescript">
    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
</script>

<style>
</style>

<section>
    {#each data as dataItem}
        <CommonTableRow data={dataItem} config={config} />
    {/each}
</section>