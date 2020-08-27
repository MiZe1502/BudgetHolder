<script lang="ts">
    import {onMount} from 'svelte';
    import {style} from "./style";
    import {purchases} from "../../../stores/purchases";
    import {buildCategoryList} from "../../../stores/categories";

    import CategoriesListElement
        from '../../../common/components/ElementsAndBlocks/CategoriesListElement/CategoriesListElement.svelte'
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
                component: CategoriesListElement,
                overflowed: true,
                style: 'flex: 1 0 15%',
                mapping: (data: GoodsDetails) => {
                    console.log(buildCategoryList(data.category ? data.category.id : null))
                    return buildCategoryList(data.category ? data.category.id : null);
                }
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
    })

    export let purchaseId: number = -1;
</script>

<div class="{style.Wrapper}">
    <CommonTable maxHeightWithScroll={500} withScroll={true} inPopup={true} status={tableData.status}
                 total={tableData.total}
                 data={$purchases.find((purchase) => purchase.id === purchaseId).goods}
                 config={tableData.columnsConfig}/>
</div>