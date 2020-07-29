<script lang="ts">
    import {onMount} from 'svelte';
    import {Wrapper} from "./style";

    import SimpleTextElement
        from '../../../common/components/ElementsAndBlocks/SimpleTextElement/SimpleTextElement.svelte'
    import CommonTable
        from '../../../common/components/Tables/CommonTable/CommonTable.svelte'
    import {CommonTable as CommonTableInterface} from '../../../common/components/Tables/CommonTable/utils'
    import {LoadingStatus} from "../../../stores/utils";
    import {GoodsDetails} from "../types";

    const tableData: CommonTableInterface = {
        total: 0,
        status: LoadingStatus.Finished,
        columnsConfig: [
            {
                header: 'common.labels.name',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 25%',
                mapping: (data: GoodsDetails) => data.name,
            },
            {
                header: 'budget.labels.category',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 15%',
                mapping: (data: GoodsDetails) => data.category ? data.category.value : null,
            },
            {
                header: 'budget.labels.amount',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 10%',
                mapping: (data: GoodsDetails) => data.amount,
            },
            {
                header: 'budget.labels.price',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 10%',
                mapping: (data: GoodsDetails) => data.price,
            },
            {
                header: 'common.labels.comment',
                component: SimpleTextElement,
                style: 'flex: 1 0 30%',
                mapping: (data: GoodsDetails) => data.comment,
            },
        ],
        data: []
    }

    onMount(async () => {
        tableData.status = LoadingStatus.Finished;
        tableData.total = data.length;
        tableData.data = data;
    })

    export let data: GoodsDetails[] = [];
</script>

<div class="{Wrapper}">
    <CommonTable inPopup={true} status={tableData.status}
                 total={tableData.total} data={tableData.data}
                 config={tableData.columnsConfig}/>
</div>