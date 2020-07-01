<script lang="typescript">
    import { fade } from 'svelte/transition';

    import { ColumnConfig } from "./utils";
    import { LoadingStatus } from "../../../stores/utils";

    import { SectionBottomMargin } from "./style";

    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'
    import CommonTitle from '../CommonTitle/CommonTitle.svelte'
    import CommonTableHeader from '../CommonTableHeader/CommonTableHeader.svelte'
    import LoadingSpinner from '../LoadingSpinner/LoadingSpinner.svelte'
    import NoDataFoundBlock from '../NoDataFoundBlock/NoDataFoundBlock.svelte'
    
    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
    export let title: string = "";
    export let total: number = 0;
    export let status: LoadingStatus = LoadingStatus.None;
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
    export let buttonClickHandler = () => {};
</script>

<section class="{SectionBottomMargin}">
    <CommonTitle buttonClickHandler={buttonClickHandler} withButton={withButton} buttonTitle={buttonTitle} title={title} dataLength={total}/>
    {#if status === LoadingStatus.Loading}
        <LoadingSpinner />
    {:else if status === LoadingStatus.Finished}
        {#if !data || data.length === 0}
            <NoDataFoundBlock />
        {:else}
            <CommonTableHeader config={config}/>
            {#each data as dataItem (dataItem.id)}
                <CommonTableRow data={dataItem} config={config} />
            {/each}
        {/if}
    {:else if status === LoadingStatus.Error}
        Error fetching data
    {/if}
</section>