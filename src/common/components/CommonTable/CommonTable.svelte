<script lang="typescript">
    import { ColumnConfig } from "./utils";

    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'
    import CommonTableHeader from '../CommonTableHeader/CommonTableHeader.svelte'

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
</script>

<style>

</style>

<section>
    <CommonTableHeader config={config}/>
    {#each data as dataItem}
        <CommonTableRow data={dataItem} config={config} />
    {/each}
</section>