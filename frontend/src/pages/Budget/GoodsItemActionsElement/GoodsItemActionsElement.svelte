<script lang="ts">
    import {GoodsDetails} from "../types";
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import RemoveGoodsItemActionElement
        from "../RemoveGoodsItemActionElement/RemoveGoodsItemActionElement.svelte";
    import EditGoodsItemActionElement
        from "../EditGoodsItemActionElement/EditGoodsItemActionElement.svelte";


    export let data: GoodsDetails = {};
</script>

<CommonActionElement>
    <EditGoodsItemActionElement {data}/>
    <RemoveGoodsItemActionElement {data}/>
</CommonActionElement>