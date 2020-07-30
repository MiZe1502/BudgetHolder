<script lang="ts">
    import {onMount} from 'svelte';
    import {Wrapper} from "./style";
    import {purchases} from "../../../stores/purchases";

    import GoodsItemActionsElement
        from "../GoodsItemActionsElement/GoodsItemActionsElement.svelte";
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
                style: 'flex: 1 0 20%',
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
                style: 'flex: 1 0 25%',
                mapping: (data: GoodsDetails) => data.comment,
            },
            {
                header: "",
                component: GoodsItemActionsElement,
                style: 'flex: 1 0 10%',
                mapping: (data: GoodsDetails): GoodsDetails => data,
            }
        ],
        data: []
    }

    onMount(async () => {
        tableData.status = LoadingStatus.Finished;
        tableData.total = data.length;
    })

    export let purchaseId: number = -1;
</script>

<div class="{Wrapper}">
    <CommonTable inPopup={true} status={tableData.status}
                 total={tableData.total}
                 data={$purchases.find((purchase) => purchase.id === purchaseId).goods}
                 config={tableData.columnsConfig}/>
</div>