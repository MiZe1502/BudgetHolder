<script lang="typescript">
    import CommonActionElement from "../../../common/components/ActionElements/CommonActionElement/CommonActionElement.svelte";
    import MapActionElement from "../../../common/components/ActionElements/MapActionElement/MapActionElement.svelte";
    import EditShopActionElement from "../EditShopActionElement/EditShopActionElement.svelte";
    import RemoveShopActionElement from "../RemoveShopActionElement/RemoveShopActionElement.svelte";

    import { ShopActionsData} from "./utils";

    export let data: ShopActionsData = {};
</script>

<CommonActionElement>
    <MapActionElement data={data.mapData} />
    <EditShopActionElement data={data.shopData} />
    <RemoveShopActionElement data={data.shopData} />
</CommonActionElement>

