<script lang="typescript">
    import {_} from 'svelte-i18n'
    import {fade} from 'svelte/transition';
    import SimpleBar from '@woden/svelte-simplebar';

    import {ColumnConfig} from "./utils";
    import {LoadingStatus} from "../../../../stores/utils";

    import {style} from "./style";

    import CommonTableRow from '../CommonTableRow/CommonTableRow.svelte'
    import CommonTitle from '../CommonTitle/CommonTitle.svelte'
    import CommonTableHeader
        from '../CommonTableHeader/CommonTableHeader.svelte'
    import LoadingSpinner
        from '../../ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte'
    import NoDataFoundBlock
        from '../../ElementsAndBlocks/NoDataFoundBlock/NoDataFoundBlock.svelte'
    import Pagination
        from '../../Pagination/PaginationBlock/PaginationBlock.svelte'

    import {maxRecordsPerPage} from "../../Pagination/PaginationBlock/utils";
    import EditShopActionElement
        from "../../../../pages/Shops/EditShopActionElement/EditShopActionElement.svelte";

    export let config: ColumnConfig[] = [];
    export let data: Record<any, any>[] = [];
    export let inPopup: boolean = false;
    export let title: string = "";
    export let total: number = 0;
    export let status: LoadingStatus = LoadingStatus.None;
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
    export let withScroll: boolean = false;
    export let maxHeightWithScroll: number = 700;
    export let buttonClickHandler = () => {
    };
    export let onPageChange = () => {
    };
</script>

<section class="{style.SectionBottomMargin}">
    {#if title}
        <CommonTitle buttonClickHandler={buttonClickHandler}
                     withButton={withButton}
                     buttonTitle={buttonTitle} title={title} dataLength={total}>
            <slot name="titleButton" />
        </CommonTitle>
    {/if}
    {#if status === LoadingStatus.Loading}
        <LoadingSpinner/>
    {:else if status === LoadingStatus.Finished}
        {#if !data || data.length === 0}
            <NoDataFoundBlock/>
        {:else}
            <CommonTableHeader {inPopup} config={config}/>
            {#if withScroll}
                <SimpleBar
                        style="max-height: {maxHeightWithScroll}px; width: 100%">
                    {#each data as dataItem (dataItem.id || dataItem.name)}
                        <CommonTableRow {inPopup} data={dataItem}
                                        config={config}/>
                    {/each}
                </SimpleBar>
            {:else}
                {#each data as dataItem (dataItem.id)}
                    <CommonTableRow {inPopup} data={dataItem}
                                    config={config}/>
                {/each}
                {#if total > maxRecordsPerPage}
                    <Pagination onPageChange={onPageChange} totalCount={total}/>
                {/if}
            {/if}
        {/if}
    {:else if status === LoadingStatus.Error}
        {$_("common.messages.errors.fetching")}
    {/if}
</section>