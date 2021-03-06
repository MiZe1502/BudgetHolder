<script lang="ts">
    import {_} from 'svelte-i18n'
    import {get} from 'svelte/store';
    import {onMount} from "svelte";
    import {Purchase, GoodsData} from "./types";
    import {BudgetActionsData} from "./BudgetActionsElement/utils";
    import {LoadingStatus} from "../../stores/utils";
    import CategoriesListElement
        from '../../common/components/ElementsAndBlocks/CategoriesListElement/CategoriesListElement.svelte';
    import {
        mockData as mockCategoriesData,
        mockCategories
    } from "../Categorization/data";
    import {mockData as mockShops} from "../Shops/data";
    import GoodsActionElement
        from "./GoodsActionElement/GoodsActionElement.svelte";
    import BudgetActionsElement
        from "./BudgetActionsElement/BudgetActionsElement.svelte";
    import CommonTable
        from '../../common/components/Tables/CommonTable/CommonTable.svelte'
    import UrlElement
        from '../../common/components/ElementsAndBlocks/UrlElement/UrlElement.svelte'
    import SimpleDateElement
        from '../../common/components/ElementsAndBlocks/SimpleDateElement/SimpleDateElement.svelte'
    import SimpleTextElement
        from '../../common/components/ElementsAndBlocks/SimpleTextElement/SimpleTextElement.svelte'
    import {CommonTable as CommonTableInterface} from '../../common/components/Tables/CommonTable/utils'
    import {
        categoriesStatus,
        categories,
        simpleCategories,
        simpleCategoriesStatus,
        categoriesTotal,
        findParentCategory,
        buildCategoryList,
        loadSimpleCategoriesList,
        loadCategoriesTree
    } from "../Categorization/categories";
    import {
        goodsStatus,
        goods,
        goodsTotal,
        allGoods, updateCurrentGoodsSlice, getGoodsItemsForSuggestions,
    } from "./goods";
    import {allShops, getSimpleShopsData} from "../Shops/shops";
    import {
        purchasesStatus,
        purchases,
        purchasesTotal, updatePurchasesWithGoodsSlice,
    } from "./purchases";
    import BudgetFormContainer
        from "./BudgetFormContainer/BudgetFormContainer.svelte";
    import BudgetPurchaseForm
        from "./BudgetPurchaseForm/BudgetPurchaseForm.svelte";
    import LoadingSpinner
        from "../../common/components/ElementsAndBlocks/LoadingSpinner/LoadingSpinner.svelte";
    import {maxRecordsPerPage} from "../../common/components/Pagination/PaginationBlock/utils";
    import EditGoodsActionElement
        from "./EditGoodsActionElement/EditGoodsActionElement.svelte";

    const goodsTableData: CommonTableInterface = {
        title: "budget.titles.goods",
        total: 0,
        withButton: true,
        buttonTitle: "goods.titles.new",
        buttonClickHandler: () => {
        },
        status: LoadingStatus.Loading,
        columnsConfig: [
            {
                header: 'common.labels.name',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 30%',
                mapping: (data: GoodsData) => data.name,
            },
            {
                header: 'budget.labels.category',
                component: CategoriesListElement,
                overflowed: true,
                style: 'flex: 1 0 20%',
                mapping: (data: GoodsData) => {
                    return buildCategoryList(data.category ? data.category.id : null);
                }
            },
            {
                header: 'common.labels.comment',
                component: SimpleTextElement,
                style: 'flex: 1 0 30%',
                mapping: (data: GoodsData) => data.comment,
            },
            {
                header: "",
                component: GoodsActionElement,
                style: 'flex: 1 0 10%',
                mapping: (data: GoodsData) => data,
            }
        ]
    }

    const tableData: CommonTableInterface = {
        title: "budget.titles.purchases",
        total: 0,
        buttonClickHandler: () => {
        },
        status: LoadingStatus.Loading,
        columnsConfig: [
            {
                header: 'budget.labels.date',
                component: SimpleDateElement,
                overflowed: true,
                style: 'flex: 1 0 10%',
                mapping: (data: Purchase) => data.date,
            },
            {
                header: 'budget.labels.price',
                component: SimpleTextElement,
                overflowed: true,
                style: 'flex: 1 0 10%',
                mapping: (data: Purchase) => data.total_price,
            },
            {
                header: 'budget.labels.shop',
                component: UrlElement,
                overflowed: true,
                style: 'flex: 1 0 10%',
                mapping: (data: Purchase) => {
                    return {
                        name: data.shop.name,
                        url: data.shop.url,
                    }
                }
            },
            {
                header: 'common.labels.comment',
                component: SimpleTextElement,
                style: 'flex: 1 0 40%',
                mapping: (data: Purchase) => data.comment,
            },
            {
                header: "",
                component: BudgetActionsElement,
                style: 'flex: 1 0 10%',
                mapping: (data: Purchase): BudgetActionsData => {
                    return {
                        id: data.id,
                        goodsData: data.goods_data,
                        date: data.date,
                        total_price: data.total_price,
                        shop: data.shop,
                        shop_id: data.shop_id,
                        comment: data.comment,
                    }
                },
            }
        ]
    }

    const onPurchasesPageChange = async (currentPage: number) => {
        await updatePurchasesWithGoodsSlice((currentPage - 1) * maxRecordsPerPage, maxRecordsPerPage)
    }

    const onGoodsPageChange = async (currentPage: number) => {
        await updateCurrentGoodsSlice((currentPage - 1) * maxRecordsPerPage, maxRecordsPerPage)
    }

    onMount(async () => {
        await loadCategoriesTree();
        await getGoodsItemsForSuggestions();
        await getSimpleShopsData();
        await loadSimpleCategoriesList();
        await onGoodsPageChange(1);
        await onPurchasesPageChange(1);

        goodsTableData.data = get(goods);
        goodsTableData.status = get(goodsStatus);

        tableData.data = get(purchases);
        tableData.status = get(purchasesStatus);

        setTimeout(() => {
            allShops.set(mockShops);

            // goodsStatus.set(LoadingStatus.Finished);
            // purchasesStatus.set(LoadingStatus.Finished);
            //
            // purchases.set(mockPurchases);
            // purchasesTotal.set(mockPurchases.length);

            // tableData.status = get(purchasesStatus);
            // tableData.total = get(purchasesTotal);
            // tableData.data = get(purchases);
        }, 5000);
    })
</script>

<section>
    <BudgetFormContainer>
        {#if tableData.status === LoadingStatus.Loading}
            <LoadingSpinner/>
        {:else}
            <BudgetPurchaseForm/>
        {/if}
    </BudgetFormContainer>

    <CommonTable onPageChange={onPurchasesPageChange} withButton={tableData.withButton}
                 buttonTitle={tableData.buttonTitle} status={tableData.status}
                 total={$purchasesTotal} data={$purchases}
                 config={tableData.columnsConfig} title={tableData.title}/>

    <CommonTable onPageChange={onGoodsPageChange}
                 withButton={goodsTableData.withButton}
                 buttonTitle={goodsTableData.buttonTitle}
                 status={goodsTableData.status} total={$goodsTotal}
                 data={$goods} config={goodsTableData.columnsConfig}
                 title={goodsTableData.title}>
        <div slot="titleButton">
            <EditGoodsActionElement withButton={true}
                                    buttonTitle={$_("goods.titles.new")}/>
        </div>
    </CommonTable>
</section>