<script lang="ts">
    import {BudgetActionsData} from "./utils";
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import ShowDetailsBudgetActionElement
        from "../ShowDetailsBudgetActionElement/ShowDetailsBudgetActionElement.svelte";
    import RemovePurchaseActionElement
        from "../RemovePurchaseActionElement/RemovePurchaseActionElement.svelte";
    import EditPurchaseActionElement
        from "../EditPurchaseActionElement/EditPurchaseActionElement.svelte";

    export let data: BudgetActionsData = {};
</script>

<CommonActionElement>
    <ShowDetailsBudgetActionElement entityId={data.id}
                                    data={data.goods_data}/>
    <EditPurchaseActionElement data={data}/>
    <RemovePurchaseActionElement
            data={{id: data.id, date: data.date, total_price: data.total_price}}/>
</CommonActionElement>