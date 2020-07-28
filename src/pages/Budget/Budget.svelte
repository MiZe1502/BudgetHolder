<script lang="ts">
    import { get } from 'svelte/store';
    import {onMount} from "svelte";
    import {Purchase} from "./types";
    import {LoadingStatus} from "../../stores/utils";
    import {mockGoods, mockPurchases} from "./data";

    import CommonTable from '../../common/components/Tables/CommonTable/CommonTable.svelte'
    import UrlElement
        from '../../common/components/ElementsAndBlocks/UrlElement/UrlElement.svelte'
    import SimpleTextElement
        from '../../common/components/ElementsAndBlocks/SimpleTextElement/SimpleTextElement.svelte'
    import {CommonTable as CommonTableInterface} from '../../common/components/Tables/CommonTable/utils'
    import {
        goodsStatus,
        goods,
        goodsTotal,
    } from "../../stores/goods";
    import {
        purchasesStatus,
        purchases,
        purchasesTotal,
    } from "../../stores/purchases";

    const tableData: CommonTableInterface = {
        title: "budget.titles.purchases",
        total: 0,
        withButton: true,
        buttonTitle: "budget.buttons.new",
        buttonClickHandler: () => {
        },
        status: LoadingStatus.Loading,
        columnsConfig: [
            {
                header: 'budget.labels.price',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 20%',
                mapping: (data: Purchase) => data.totalPrice,
            },
            {
                header: 'common.labels.comment',
                component: SimpleTextElement,
                style: 'flex: 1 0 40%',
                mapping: (data: Purchase) => data.comment,
            },
            {
                header: 'budget.labels.shop',
                component: UrlElement,
                overflowed: true,
                style: 'flex: 1 0 20%',
                mapping: (data: Purchase) => {
                    return {
                        name: data.shop.name,
                        url: data.shop.url,
                    }
                }
            },
        ]
    }

    const onPageChange = () => {
        //TODO: implement change page
    }


    onMount(async () => {
        goodsStatus.set(LoadingStatus.Loading);
        purchasesStatus.set(LoadingStatus.Loading);

        setTimeout(() => {
            goodsStatus.set(LoadingStatus.Finished);
            purchasesStatus.set(LoadingStatus.Finished);

            goods.set(mockGoods);
            purchases.set(mockPurchases);

            goodsTotal.set(mockGoods.length);
            purchasesTotal.set(mockPurchases.length);

            tableData.status = get(purchasesStatus);
            tableData.total = get(purchasesTotal);
            tableData.data = get(purchases);
        }, 5000);
    })
</script>

<section>
    <CommonTable onPageChange={onPageChange} withButton={tableData.withButton}
                 buttonTitle={tableData.buttonTitle} status={tableData.status}
                 total={tableData.total} data={$purchases}
                 config={tableData.columnsConfig} title={tableData.title}/>
</section>