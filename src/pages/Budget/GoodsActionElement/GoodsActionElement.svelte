<script lang="ts">
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import RemoveGoodsItemElement
        from "../RemoveGoodsItemElement/RemoveGoodsItemElement.svelte";
    import {GoodsData} from "../types";
    import EditGoodsActionElement
        from "../EditGoodsActionElement/EditGoodsActionElement.svelte";

    export let data: GoodsData = {}
</script>

<CommonActionElement>
    <EditGoodsActionElement {data} />
    <RemoveGoodsItemElement {data} />
</CommonActionElement>