<script lang="ts">
    import {BudgetActionsData} from "./utils";
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import ShowDetailsBudgetActionElement
        from "../ShowDetailsBudgetActionElement/ShowDetailsBudgetActionElement.svelte";

    export let data: BudgetActionsData = {};
</script>

<CommonActionElement>
    <ShowDetailsBudgetActionElement entityId={data.id} data={[...data.goodsData]}/>
</CommonActionElement>