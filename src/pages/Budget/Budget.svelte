<script lang="ts">
    import {get, writable} from 'svelte/store'

    import {onMount} from "svelte";
    import {LoadingStatus} from "../../stores/utils";
    import {mockGoods, mockPurchases} from "./data";
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
    onMount(() => {
        goodsStatus.set(LoadingStatus.Loading);
        purchasesStatus.set(LoadingStatus.Loading);

        setTimeout(() => {
            goodsStatus.set(LoadingStatus.Finished);
            purchasesStatus.set(LoadingStatus.Finished);

            goods.set(mockGoods);
            purchases.set(mockPurchases);

            goodsTotal.set(mockGoods.length);
            purchasesTotal.set(mockPurchases.length);
        }, 5000);
    })
</script>

<section>
    BUDGET PAGE
</section>