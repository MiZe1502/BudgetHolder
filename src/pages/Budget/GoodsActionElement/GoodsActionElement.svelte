<script lang="ts">
    import CommonActionElement
        from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import RemoveGoodsItemElement
        from "../RemoveGoodsItemElement/RemoveGoodsItemElement.svelte";
    import {GoodsData} from "../types";

    export let data: GoodsData = {}
</script>

<CommonActionElement>
    <RemoveGoodsItemElement {data}/>
</CommonActionElement>